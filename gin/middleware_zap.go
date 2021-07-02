// Package ggin Create by guoxin 2021.07.02
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

type bodyWriterWrapper struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriterWrapper) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func Ginzap(logger *zap.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		// There is no need to close after reading the request body
		var requestBodyBytes []byte
		if c.Request.Body != nil {
			requestBodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		// 把刚刚读出来的再写进去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))

		logger.Info("GComponent [gin] REQUEST",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.ByteString("body", requestBodyBytes),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			//zap.Duration("latency", latency),
			zap.String("time", start.Format(timeFormat)),
			zap.String("user-agent", c.Request.UserAgent()),
		)

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}
		// 响应
		responseWriter := &bodyWriterWrapper{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = responseWriter

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			logger.Error("GComponent [gin] RESPONSE",
				zap.String("time", end.Format(timeFormat)),
				zap.Duration("latency", latency),
				zap.Strings("errors", c.Errors.Errors()))
		} else {
			logger.Info("GComponent [gin] RESPONSE",
				zap.String("time", end.Format(timeFormat)),
				zap.Duration("latency", latency),
				zap.Int("size", responseWriter.body.Len()),
				zap.String("body", responseWriter.body.String()),
			)
		}
	}
}

// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
func RecoveryWithZap(logger *zap.Logger, stack bool) gin.HandlerFunc {
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
					logger.Error("GComponent [gin] RESPONSE Broken pipe",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.

					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("GComponent [gin] RESPONSE Recovery from panic",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("GComponent [gin] RESPONSE Recovery from panic",
						zap.Time("time", time.Now()),
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
