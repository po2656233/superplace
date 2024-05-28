package config

import (
	"fmt"
	"time"

	face "github.com/po2656233/superplace/facade"
	"go.uber.org/zap/zapcore"
)

type (
	LogConfig struct {
		LogLevel        string     `json:"level"`             // 输出日志等级
		StackLevel      string     `json:"stack_level"`       // 堆栈输出日志等级
		EnableConsole   bool       `json:"enable_console"`    // 是否控制台输出
		EnableWriteFile bool       `json:"enable_write_file"` // 是否输出文件(必需配置FilePath)
		MaxAge          int        `json:"max_age"`           // 最大保留天数(达到限制，则会被清理)
		TimeFormat      string     `json:"time_format"`       // 打印时间输出格式
		PrintCaller     bool       `json:"print_caller"`      // 是否打印调用函数
		RotationTime    int        `json:"rotation_time"`     // 日期分割时间(秒)
		FileLinkPath    string     `json:"file_link_path"`    // 日志文件连接路径
		FilePathFormat  string     `json:"file_path_format"`  // 日志文件路径格式
		IncludeStdout   bool       `json:"include_stdout"`    // 是否包含os.stdout输出
		IncludeStderr   bool       `json:"include_stderr"`    // 是否包含os.stderr输出
		User            string     `json:"user"`              // 使用者名称
		EmailHook       *EmailHook `json:"email"`
		RedisHook       *RedisHook `json:"redis"`
		MongoHook       *MongoHook `json:"mongo"`
	}
	EmailHook struct {
		AppName  string `json:"appname"` // 主题
		Host     string `json:"host"`
		Port     int    `json:"port"`
		From     string `json:"from"`
		To       string `json:"to"`
		Message  string `json:"message"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	RedisHook struct {
		Key      string `json:"key"`
		Format   string `json:"format"`
		App      string `json:"app"`
		Host     string `json:"host"`
		Password string `json:"password"`
		Username string `json:"username"`
		Port     int    `json:"port"`
		DB       int    `json:"db"`
		TTL      int    `json:"ttl"`
	}
	MongoHook struct {
		Addrs    []string `json:"addrs"`
		Timeout  int      `json:"port"`
		Database string   `json:"database"`
		Username string   `json:"username"`
		Password string   `json:"password"`
	}
)

func DefaultLogConfig() *LogConfig {
	config := &LogConfig{
		LogLevel:        "debug",
		StackLevel:      "panic",
		EnableConsole:   true,
		EnableWriteFile: false,
		MaxAge:          7,
		TimeFormat:      "15:04:05.000", //2006-01-02 15:04:05.000
		PrintCaller:     true,
		RotationTime:    86400,
		FileLinkPath:    "logs/debug.log",
		FilePathFormat:  "logs/debug_%Y%m%d%H%M.log",
		IncludeStdout:   false,
		IncludeStderr:   false,
	}
	return config
}

func NewConfig(jsonConfig face.ProfileJSON) (*LogConfig, error) {
	config := &LogConfig{
		LogLevel:        jsonConfig.GetString("level", "debug"),
		StackLevel:      jsonConfig.GetString("stack_level", "panic"),
		EnableConsole:   jsonConfig.GetBool("enable_console", true),
		EnableWriteFile: jsonConfig.GetBool("enable_write_file", false),
		MaxAge:          jsonConfig.GetInt("max_age", 7),
		TimeFormat:      jsonConfig.GetString("time_format", "15:04:05.000"),
		PrintCaller:     jsonConfig.GetBool("print_caller", true),
		RotationTime:    jsonConfig.GetInt("rotation_time", 86400),
		FileLinkPath:    jsonConfig.GetString("file_link_path", ""),
		FilePathFormat:  jsonConfig.GetString("file_path_format", ""),
		IncludeStdout:   jsonConfig.GetBool("include_stdout", false),
		IncludeStderr:   jsonConfig.GetBool("include_stderr", false),
	}

	if config.EnableWriteFile {
		if config.FileLinkPath == "" {
			defaultValue := fmt.Sprintf("logs/%s.log", config.LogLevel)
			config.FileLinkPath = jsonConfig.GetString("file_link_path", defaultValue)
		}

		if config.FilePathFormat == "" {
			defaultValue := fmt.Sprintf("logs/%s_%s.log", config.LogLevel, "%Y%m%d%H%M")
			config.FilePathFormat = jsonConfig.GetString("file_path_format", defaultValue)
		}
	}

	return config, nil
}

func NewConfigWithName(refLoggerName string) (*LogConfig, error) {
	loggerConfig := GetConfig("logger")
	if loggerConfig.LastError() != nil {
		return nil, loggerConfig.LastError()
	}

	jsonConfig := loggerConfig.GetConfig(refLoggerName)
	if jsonConfig.LastError() != nil {
		return nil, jsonConfig.LastError()
	}

	return NewConfig(jsonConfig)
}

func (c *LogConfig) TimeEncoder() zapcore.TimeEncoder {
	return func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format(c.TimeFormat))
	}
}
