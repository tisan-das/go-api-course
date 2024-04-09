package logging

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

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
