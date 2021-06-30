// Package ggin
// Create by guoxin 2020.12.15
package ggin

import (
    "bytes"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "io/ioutil"
    "net"
    "net/http"
    "net/http/httputil"
    "os"
    "runtime/debug"
    "strings"
    "time"
)

func HandlerNotFound(c *gin.Context) {
    c.JSON(http.StatusNotFound, NotFound)
    return
}

// Logger GinLogger 接收gin框架默认的日志
func Logger(logger *zap.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        gin.Default()
        start := time.Now()
        path := c.Request.URL.Path
        query := c.Request.URL.RawQuery
        c.Next()

        cost := time.Since(start)
        logger.Info(path,
            zap.Int("status", c.Writer.Status()),
            zap.String("method", c.Request.Method),
            zap.String("path", path),
            zap.String("query", query),
            zap.String("ip", c.ClientIP()),
            zap.String("user-agent", c.Request.UserAgent()),
            zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
            zap.Duration("cost", cost),
        )
    }
}

type bodyWriterWrapper struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w bodyWriterWrapper) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}

func Ginzap(logger *zap.Logger, timeFormat string, utc bool, enableRequestId bool) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        writer := &bodyWriterWrapper{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
        c.Writer = writer

        // some evil middlewares modify this values
        path := c.Request.URL.Path
        query := c.Request.URL.RawQuery

        // request id
        var requestIdField zap.Field
        if enableRequestId {
            requestId := c.GetString(requestIdKey)
            if len(requestId) != 0 {
                requestIdField = zap.String(requestIdKey, requestId)
            }
        }
        // There is no need to close after reading the request body
        var bodyBytes []byte
        if c.Request.Body != nil {
            bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
        }
        // 把刚刚读出来的再写进去
        c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

        logger.Info("START",
            requestIdField,
            zap.Int("status", c.Writer.Status()),
            zap.String("method", c.Request.Method),
            zap.ByteString("body", bodyBytes),
            zap.String("path", path),
            zap.String("query", query),
            zap.String("ip", c.ClientIP()),
            //zap.String("user-agent", c.Request.UserAgent()),
        )

        c.Next()

        end := time.Now()
        latency := end.Sub(start)
        if utc {
            end = end.UTC()
        }

        if len(c.Errors) > 0 {
            // Append error field if this is an erroneous request.
            for _, e := range c.Errors.Errors() {
                logger.Error(e)
            }
        }
        logger.Info("END",
            requestIdField,
            zap.String("time", end.Format(timeFormat)),
            zap.Duration("latency", latency),
            zap.String("body", writer.body.String()),
        )
    }
}

// Recovery GinRecovery recover掉项目可能出现的panic
func Recovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                // Check for a broken connection, as it is not really a
                // condition that warrants a panic stack trace.
                var brokenPipe bool
                if ne, ok := err.(*net.OpError); ok {
                    if se, ok := ne.Err.(*os.SyscallError); ok {
                        if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
                            brokenPipe = true
                        }
                    }
                }

                httpRequest, _ := httputil.DumpRequest(c.Request, false)
                if brokenPipe {
                    logger.Error(c.Request.URL.Path,
                        zap.Any("error", err),
                        zap.String("request", string(httpRequest)),
                    )
                    // If the connection is dead, we can't write a status to it.
                    _ = c.Error(err.(error)) // nolint: errcheck
                    c.Abort()
                    return
                }

                if stack {
                    logger.Error("[Recovery from panic]",
                        zap.Any("error", err),
                        zap.String("request", string(httpRequest)),
                        zap.String("stack", string(debug.Stack())),
                    )
                } else {
                    logger.Error("[Recovery from panic]",
                        zap.Any("error", err),
                        zap.String("request", string(httpRequest)),
                    )
                }
                c.AbortWithStatus(http.StatusInternalServerError)
            }
        }()
        c.Next()
    }
}
