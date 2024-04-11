package logging

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Logger interface {
	Init(LoggingLevel, string) error
	// SetLoggingLevel()
	LogRequest(*gin.Context)
	LogResponse(*gin.Context)
	Debugw(str, useCase, requestId string)
	Infow(str, useCase, requestId string)
	Warnw(str, useCase, requestId string)
	Errorw(str, useCase, requestId string)
}

type LoggingLevel string

const (
	LOG_DEBUG_LEVEL LoggingLevel = "debug"
	LOG_INFO_LEVEL  LoggingLevel = "info"
	LOG_WARN_LEVEL  LoggingLevel = "warn"
	LOG_ERROR_LEVEL LoggingLevel = "error"
)

func (loggingLevel *LoggingLevel) Sprint() string {
	return fmt.Sprintf("%s", loggingLevel)
}

type logFormatLocal struct {
	Timestamp    time.Time
	StatusCode   int
	Latency      time.Duration
	ClientIP     string
	Method       string
	Path         string
	ErrorMessage string
	RequestProto string
}

func FormatLogsJson(param gin.LogFormatterParams) string {
	params := &logFormatLocal{
		Timestamp:    param.TimeStamp,
		StatusCode:   param.StatusCode,
		Latency:      param.Latency,
		ClientIP:     param.ClientIP,
		Method:       param.Method,
		Path:         param.Path,
		ErrorMessage: param.ErrorMessage,
		RequestProto: param.Request.Proto,
	}
	jsonData, _ := json.Marshal(params)
	return string(jsonData)
}
