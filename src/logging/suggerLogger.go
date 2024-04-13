package logging

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"runtime"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ZapSugarLogger struct {
	logger *zap.SugaredLogger
}

func NewZapSugarLogger(logLevel LoggingLevel, logFile string) (Logger, error) {
	zapSugarLogger := ZapSugarLogger{}
	err := zapSugarLogger.Init(logLevel, logFile)
	if err != nil {
		return nil, err
	}
	return &zapSugarLogger, nil
}

func (log *ZapSugarLogger) Init(logLevel LoggingLevel, logFile string) error {

	rawJSON := []byte(fmt.Sprint(`{
		"level": "`, logLevel, `",
		"encoding": "json",
		"outputPaths": ["stdout", "`, logFile, `"],
		"errorOutputPaths": ["stderr", "`, logFile, `"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`))
	var cfg zap.Config
	err := json.Unmarshal(rawJSON, &cfg)
	if err != nil {
		return err
	}
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}
	defer logger.Sync()
	log.logger = logger.Sugar()
	return nil
}

func (log *ZapSugarLogger) Debugw(str, useCase, requestId string) {
	_, file, line, _ := runtime.Caller(1)
	log.logger.Debugw(str, "useCase", useCase, "requestId", requestId, "timestamp", time.Now().UTC().Format(time.RFC3339),
		"code", fmt.Sprintf("%s %s", file, line))
}
func (log *ZapSugarLogger) Infow(str, useCase, requestId string) {
	_, file, line, _ := runtime.Caller(1)
	log.logger.Infow(str, "useCase", useCase, "requestId", requestId, "timestamp", time.Now().UTC().Format(time.RFC3339),
		"code", fmt.Sprintf("%s %s", file, line))
}
func (log *ZapSugarLogger) Warnw(str, useCase, requestId string) {
	_, file, line, _ := runtime.Caller(1)
	log.logger.Warnw(str, "use-case", useCase, "requestId", requestId, "timestamp", time.Now().UTC().Format(time.RFC3339),
		"code", fmt.Sprintf("%s %s", file, line))
}
func (log *ZapSugarLogger) Errorw(str, useCase, requestId string) {
	_, file, line, _ := runtime.Caller(1)
	log.logger.Errorw(str, "useCase", useCase, "requestId", requestId, "timestamp", time.Now().UTC().Format(time.RFC3339),
		"code", fmt.Sprintf("%s %s", file, line))
}

func (log *ZapSugarLogger) LogRequest(ctx *gin.Context) {
	requestId, _ := ctx.Get("requestId")
	log.logger.Infow("Incomging request ", "method:", ctx.Request.Method, "url", ctx.Request.URL,
		"header", ctx.Request.Header, "body", ctx.Request.Body, "requestId", requestId, "timestamp", time.Now().UTC().Format(time.RFC3339))
}

func (log *ZapSugarLogger) LogResponse(ctx *gin.Context) {
	requestId, _ := ctx.Get("requestId")
	responseBody, _ := ctx.Copy().GetRawData()
	log.logger.Infow("Outgoing response", "method:", ctx.Request.Method, "url", ctx.Request.URL,
		"header", ctx.Writer.Header(), "status", strconv.Itoa(ctx.Writer.Status()), "body", string(responseBody), "requestId", requestId, "timestamp", time.Now().UTC().Format(time.RFC3339))
}
