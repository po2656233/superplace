package logger

import (
	"os"
	"strings"
	"sync"
	"time"

	cprofile "github.com/po2656233/superplace/config"
	cfacade "github.com/po2656233/superplace/facade"
	"github.com/po2656233/superplace/logger/rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	rw             sync.RWMutex            // mutex
	DefaultLogger  *SuperLogger            // 默认日志对象(控制台输出)
	loggers        map[string]*SuperLogger // 日志实例存储map(key:日志名称,value:日志实例)
	printLevel     zapcore.Level           // super log print level
	fileNameVarMap = map[string]string{}   // 日志输出文件名自定义变量
)

func init() {
	DefaultLogger = NewConfigLogger(cprofile.DefaultLogConfig(), zap.AddCallerSkip(1))
	loggers = make(map[string]*SuperLogger)
}

// 定义颜色
const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPink   = "\033[35m" // 粉色
	colorCyan   = "\033[36m" // 青色
	colorWhite  = "\033[37m"
	colorReset  = "\033[0m"
)

// 自定义颜色编码器
func colorLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	color := colorWhite
	switch level {
	case zapcore.DebugLevel:
		color = colorPink
	case zapcore.InfoLevel:
		color = colorCyan
	case zapcore.WarnLevel:
		color = colorYellow
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		color = colorRed
	}
	enc.AppendString(color + level.CapitalString() + colorReset)
}

type SuperLogger struct {
	*zap.SugaredLogger
	*cprofile.LogConfig
}

func (c *SuperLogger) Print(v ...interface{}) {
	c.Warn(v)
}

func SetNodeLogger(node cfacade.INode) {
	refLoggerName := node.Settings().Get("ref_logger").ToString()
	if refLoggerName == "" {
		DefaultLogger.Infof("RefLoggerName not found, used default console logger.")
		return
	}

	SetFileNameVar("nodeId", node.NodeId())     // %nodeId
	SetFileNameVar("nodeType", node.NodeType()) // %nodeTyp

	DefaultLogger = NewLogger(refLoggerName, zap.AddCallerSkip(1))
	printLevel = GetLevel(cprofile.PrintLevel())
}

func SetFileNameVar(key, value string) {
	fileNameVarMap[key] = value
}

func Flush() {
	_ = DefaultLogger.Sync()

	for _, logger := range loggers {
		_ = logger.Sync()
	}
}

func NewLogger(refLoggerName string, opts ...zap.Option) *SuperLogger {
	if refLoggerName == "" {
		return nil
	}

	defer rw.Unlock()
	rw.Lock()

	if logger, found := loggers[refLoggerName]; found {
		return logger
	}

	config, err := cprofile.NewConfigWithName(refLoggerName)
	if err != nil {
		Panicf("New Config fail. err = %v", err)
	}

	logger := NewConfigLogger(config, opts...)
	loggers[refLoggerName] = logger
	return logger
}

func NewConfigLogger(config *cprofile.LogConfig, opts ...zap.Option) *SuperLogger {
	if config.EnableWriteFile {
		for key, value := range fileNameVarMap {
			config.FileLinkPath = strings.ReplaceAll(config.FileLinkPath, "%"+key, value)
			config.FilePathFormat = strings.ReplaceAll(config.FilePathFormat, "%"+key, value)
		}
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		NameKey:        "name",
		StacktraceKey:  "stack",
		EncodeLevel:    colorLevelEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	//encoderConfig.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
	//	if nodeId != "" {
	//		encoder.AppendString(fmt.Sprintf("%s", nodeId))
	//	} else {
	//		encoder.AppendString(level.CapitalString())
	//	}
	//	colorLevelEncoder(level, encoder)
	//}

	if config.PrintCaller {
		encoderConfig.EncodeTime = config.TimeEncoder()
		encoderConfig.EncodeName = zapcore.FullNameEncoder
		encoderConfig.FunctionKey = zapcore.OmitKey
		opts = append(opts, zap.AddCaller())
	}

	opts = append(opts, zap.AddStacktrace(GetLevel(config.StackLevel)))

	var writers []zapcore.WriteSyncer

	if config.EnableWriteFile {
		hook, err := rotatelogs.New(
			config.FilePathFormat, //filename+"_%Y%m%d%H%M.log",
			rotatelogs.WithLinkName(config.FileLinkPath),
			rotatelogs.WithMaxAge(time.Hour*24*time.Duration(config.MaxAge)),
			rotatelogs.WithRotationTime(time.Second*time.Duration(config.RotationTime)),
		)

		if err != nil {
			panic(err)
		}

		writers = append(writers, zapcore.AddSync(hook))
	}

	if config.EnableConsole {
		writers = append(writers, zapcore.AddSync(os.Stderr))
	}

	if config.IncludeStdout {
		writers = append(writers, zapcore.Lock(os.Stdout))
	}

	if config.IncludeStderr {
		writers = append(writers, zapcore.Lock(os.Stderr))
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(zapcore.NewMultiWriteSyncer(writers...)),
		zap.NewAtomicLevelAt(GetLevel(config.LogLevel)),
	)

	superLogger := &SuperLogger{
		SugaredLogger: NewSugaredLogger(core, opts...),
		LogConfig:     config,
	}

	return superLogger
}

func NewSugaredLogger(core zapcore.Core, opts ...zap.Option) *zap.SugaredLogger {
	zapLogger := zap.New(core, opts...)
	return zapLogger.Sugar()
}

func Enable(level zapcore.Level) bool {
	return DefaultLogger.Desugar().Core().Enabled(level)
}

func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	DefaultLogger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	DefaultLogger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	DefaultLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	DefaultLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	DefaultLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	DefaultLogger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	DefaultLogger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	DefaultLogger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	DefaultLogger.Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	DefaultLogger.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	DefaultLogger.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	DefaultLogger.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	DefaultLogger.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	DefaultLogger.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	DefaultLogger.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	DefaultLogger.Fatalw(msg, keysAndValues...)
}

func PrintLevel(level zapcore.Level) bool {
	return level >= printLevel
}

func GetLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
