/**
  Create by guoxin 2020.12.15
*/
package zap

import (
    "go.uber.org/zap/zapcore"
    "time"
)

func ISOTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format("2006-01-02T15:04:05.000"))
}
