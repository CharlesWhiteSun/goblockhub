package logger

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugar *zap.SugaredLogger

func InitLogger(logDir string, level LogLevel) {
    if _, err := os.Stat(logDir); os.IsNotExist(err) {
        os.MkdirAll(logDir, 0755)
    }

    today := time.Now().Format("2006-01-02")
    logFile := filepath.Join(logDir, today+".log")

    // lumberjack：自動切割 log
    lumberjackLogger := &lumberjack.Logger{
        Filename:   logFile,
        MaxSize:    10,   // 每個檔案最大 10MB
        MaxBackups: 7,    // 保留 7 個舊檔案
        MaxAge:     30,   // 保留 30 天
        Compress:   true, // 是否壓縮
    }

    writeSyncer := zapcore.AddSync(lumberjackLogger)

    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    encoderConfig.TimeKey = "time"
    encoderConfig.LevelKey = "level"
    encoderConfig.CallerKey = "caller"
    encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
    encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),
        zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writeSyncer),
        toZapLevel(level),
    )

    logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
    sugar = logger.Sugar()
}

func toZapLevel(level LogLevel) zapcore.Level {
    switch level {
    case DebugLevel:
        return zapcore.DebugLevel
    case InfoLevel:
        return zapcore.InfoLevel
    case WarnLevel:
        return zapcore.WarnLevel
    case ErrorLevel:
        return zapcore.ErrorLevel
    default:
        return zapcore.InfoLevel
    }
}

// Exported methods
func Debug(args ...interface{}) { sugar.Debug(args...) }
func Info(args ...interface{})  { sugar.Info(args...) }
func Warn(args ...interface{})  { sugar.Warn(args...) }
func Error(args ...interface{}) { sugar.Error(args...) }

// 格式化
func Debugf(template string, args ...interface{}) { sugar.Debugf(template, args...) }
func Infof(template string, args ...interface{})  { sugar.Infof(template, args...) }
func Warnf(template string, args ...interface{})  { sugar.Warnf(template, args...) }
func Errorf(template string, args ...interface{}) { sugar.Errorf(template, args...) }
