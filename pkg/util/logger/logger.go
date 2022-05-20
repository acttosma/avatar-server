package logger

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"acttos.com/avatar/pkg/setting"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

var Monitor *zap.SugaredLogger

const TIME_FORMAT = "2006-01-02 15:04:05.000"

func InitLoggers() {
	initMonitorLogger()
}

func initMonitorLogger() {
	writeSyncer := getMonitorLogWriter()
	encoder := getMonitorEncoder()
	var logLevel zapcore.LevelEnabler
	if setting.GetInstance().ServerMode == "debug" {
		logLevel = zapcore.DebugLevel
	} else {
		logLevel = zapcore.InfoLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger := zap.New(core, zap.AddCaller())
	Monitor = logger.Sugar()
}

func getMonitorEncoder() zapcore.Encoder {
	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format(TIME_FORMAT) + "]")
	}
	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}

	// 自定义文件:行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder   //zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = customLevelEncoder //zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = customCallerEncoder //zapcore.ShortCallerEncoder
	encoderConfig.ConsoleSeparator = " - "
	encoderConfig.TimeKey = "ts"

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getMonitorLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   setting.GetInstance().MonitorLoggerFileName, // 日志文件位置
		MaxSize:    setting.GetInstance().LoggerMaxSize,         // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxBackups: setting.GetInstance().LoggerMaxBackups,      // 保留旧文件的最大个数
		MaxAge:     setting.GetInstance().LoggerMaxAge,          // 保留旧文件的最大天数
		Compress:   true,                                        // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		Monitor.Info("path:", path,
			"  status:", c.Writer.Status(),
			"  method:", c.Request.Method,
			"  query:", query,
			"  ip:", c.ClientIP(),
			"  user-agent:", c.Request.UserAgent(),
			"  content-type:", c.ContentType(),
			"  errors:", c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"  cost:", cost,
		)
	}
}

// GinRecovery recover掉项目可能出现的panic,并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
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
					Monitor.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					Monitor.Error("[Recovery from panic]",
						"  error:", err,
						"  request:", string(httpRequest),
						"  stack:", string(debug.Stack()),
					)
				} else {
					Monitor.Error("[Recovery from panic]",
						"  error", err,
						"  request", string(httpRequest),
					)
				}
				c.AbortWithStatus(http.StatusOK)
			}
		}()
		c.Next()
	}
}
