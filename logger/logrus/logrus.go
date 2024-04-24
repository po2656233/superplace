package logrus

import (
	"crypto/tls"
	cprofile "github.com/po2656233/superplace/config"
	"github.com/po2656233/superplace/logger/rotatelogs"
	"github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
	"github.com/zbindenren/logrus_mail"
	"gopkg.in/mgo.v2"
	"net"
	"strings"
	"time"
)

// logrus在记录Levels()返回的日志级别的消息时会触发HOOK,
// 按照Fire方法定义的内容修改logrus.Entry.
type Hook interface {
	Levels() []logrus.Level
	Fire(*logrus.Entry) error
}
type DefaultFieldHook struct {
}

func (hook *DefaultFieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["appName"] = "MyAppName"
	return nil
}
func (hook *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

var defaultLogger *logrus.Entry
var singleLog *logrus.Logger

func init() {
	// 设置日志格式为json格式
	singleLog = logrus.New()
	conf := cprofile.DefaultLogConfig()
	if conf.EnableWriteFile {
		singleLog.SetFormatter(&logrus.JSONFormatter{})
		writer, _ := rotatelogs.New(
			conf.FilePathFormat,
			rotatelogs.WithLinkName(conf.FileLinkPath),
			rotatelogs.WithMaxAge(time.Hour*24*time.Duration(conf.MaxAge)),
			rotatelogs.WithRotationTime(time.Second*time.Duration(conf.RotationTime)),
		)
		singleLog.SetOutput(writer)
	}
	singleLog.SetLevel(toLevel(conf.LogLevel))
	if conf.EmailHook != nil {
		AddEmailHook(conf.EmailHook)
	}
	if conf.RedisHook != nil {
		AddRedisHook(conf.RedisHook)
	}
	if conf.MongoHook != nil {
		AddMongoHook(conf.MongoHook)
	}
	defaultLogger = singleLog.WithFields(logrus.Fields{"user": conf.User})

	//hook, err := logrus_mail.NewMailAuthHook("testapp", "smtp.163.com", 25, "username@163.com", "username@163.com", "smtp_name", "smtp_password")
	//if err == nil {
	//	logrus.Hooks.Add(hook)
	//}

}
func AddEmailHook(eConf *cprofile.EmailHook) error {
	//parameter"APPLICATION_NAME", "HOST", PORT, "FROM", "TO"
	//首先开启smtp服务,最后两个参数是smtp的用户名和密码
	hook, err := logrus_mail.NewMailAuthHook(eConf.AppName, eConf.Host, eConf.Port, eConf.From, eConf.To, eConf.Username, eConf.Password)
	if err != nil {
		return err
	}

	//设置时间戳和message
	defaultLogger.Time = time.Now()
	defaultLogger.Message = eConf.Message
	singleLog.Hooks.Add(hook)
	//使用Fire发送,包含时间戳,message
	err = hook.Fire(defaultLogger)
	if err != nil {
		return err
	}
	return nil
}
func AddRedisHook(rConf *cprofile.RedisHook) error {
	hookConfig := logredis.HookConfig{
		Host:     rConf.Host,
		Key:      rConf.Key,
		Format:   rConf.Format,
		App:      rConf.App,
		Hostname: rConf.Username,
		Password: rConf.Password,
		Port:     rConf.Port,
		DB:       rConf.DB,
		TTL:      rConf.TTL,
	}

	hook, err := logredis.NewHook(hookConfig)
	if err != nil {
		return err
	}
	singleLog.AddHook(hook)
	err = hook.Fire(defaultLogger)
	if err != nil {
		return err
	}
	return nil
}

func AddMongoHook(mConf *cprofile.MongoHook) error {
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    mConf.Addrs,
		Timeout:  time.Duration(mConf.Timeout) * time.Second,
		Database: mConf.Database,
		Username: mConf.Username,
		Password: mConf.Password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), &tls.Config{InsecureSkipVerify: true})
			return conn, err
		},
	})
	if err != nil {
		singleLog.Fatalf("can't create session: %s\n", err)
		return err
	}

	c := s.DB("db").C("collection")
	hooker := mgorus.NewHookerFromCollection(c)
	singleLog.Hooks.Add(hooker)
	err = hooker.Fire(defaultLogger)
	if err != nil {
		return err
	}
	return nil
}
func toLevel(level string) logrus.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "release":
		return logrus.InfoLevel
	case "official":
		return logrus.WarnLevel
	case "trace":
		return logrus.TraceLevel
	case "strict":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	defaultLogger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}
