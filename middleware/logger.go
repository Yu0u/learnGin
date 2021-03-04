package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rta "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func Log() gin.HandlerFunc{
	logFilePath := "./log/log"

	fileName := path.Join(logFilePath)

	src, err := os.OpenFile(fileName,os.O_APPEND|os.O_WRONLY,os.ModeAppend)
	if err != nil {
		fmt.Println("err:",err)
	}

	logger := logrus.New()

	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)

	logWriter,err := rta.New(
		fileName+".%Y%m%d.log",
		rta.WithLinkName(fileName),
		rta.WithMaxAge(7*24*time.Hour),
		rta.WithRotationTime(24*time.Hour),
		)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI

		statusCode := c.Writer.Status()

		clientIP := c.ClientIP()

		logger.WithFields(logrus.Fields{
			"status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
		}).Info()
	}
}
