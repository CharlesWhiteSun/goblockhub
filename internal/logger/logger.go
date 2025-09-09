package logger

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var sugar *zap.SugaredLogger

func InitLogger(level LogLevel) {
    if goTest() {
		sugar = zap.NewNop().Sugar()
		return
	}

    // 固定寫入根目錄 logs 資料夾
	rootDir := projectRoot()
	logDir := filepath.Join(rootDir, "logs")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)
	}

    today := time.Now().Format("2006-01-02")
    logFile := filepath.Join(logDir, today+".log")

    lumberjackLogger := &lumberjack.Logger{
        Filename:   logFile,
        MaxSize:    10,   // 每個檔案最大 10MB
        MaxBackups: 31,    // 保留 n 個舊檔案
        MaxAge:     31,   // 保留 n 天
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

func goTest() bool {
	// 依據 runtime 調用 stack 判斷是否為測試
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.") {
			return true
		}
	}
	return false
}

func projectRoot() string {
	// 固定為 go module 根目錄
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir { // 到達根目錄
			break
		}
		dir = parent
	}
	return "."
}