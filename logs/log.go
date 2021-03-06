package logs

import (
	"io"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//var logger *zap.Logger
var (
	ZapLogger   *zap.Logger
	SugarLogger *zap.SugaredLogger
)

// 初始化日志 logger
func InitLog(logPath, errPath string, level string) {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	config := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder, //将级别转换成大写
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}
	encoder := zapcore.NewConsoleEncoder(config)
	// 设置级别
	logLevel := zap.DebugLevel
	switch level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	case "panic":
		logLevel = zap.PanicLevel
	case "fatal":
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}
	// 实现三个判断日志等级的interface  可以自定义级别展示
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.InfoLevel && lvl >= logLevel
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter(logPath)
	warnWriter := getWriter(errPath)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		// 将info及以下写入logPath,  warn及以上写入errPath
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
		//日志都会在console中展示
		zapcore.NewCore(encoder,
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), debugLevel),
	)
	//logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel)) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	ZapLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel)) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	SugarLogger = ZapLogger.Sugar()
}

// func getWriter(filename string) io.Writer {
// 	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
// 	// demo.log是指向最新日志的链接
// 	hook, err := rotatelogs.New(
// 		filename+".%Y%m%d%H", // 没有使用go风格反人类的format格式
// 		rotatelogs.WithLinkName(filename),
// 		rotatelogs.WithMaxAge(time.Hour*24*30),    // 保存30天
// 		rotatelogs.WithRotationTime(time.Hour*24), //切割频率 24小时
// 	)
// 	if err != nil {
// 		log.Println("日志启动异常")
// 		panic(err)
// 	}
// 	return hook
// }

func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    100,   //最大M数，超过则切割
		MaxBackups: 100,   //最大文件保留数，超过就删除最老的日志文件
		MaxAge:     30,    //保存30天
		Compress:   false, //是否压缩
	}
}

// func Sync() {
// 	logger.Sync()
// }

// // logs.Debug(...)
// func Debug(format string, v ...interface{}) {
// 	logger.Sugar().Debugf(format, v...)
// }

// func Info(format string, v ...interface{}) {
// 	logger.Sugar().Infof(format, v...)
// }

// func Warn(format string, v ...interface{}) {
// 	logger.Sugar().Warnf(format, v...)
// }

// func Error(format string, v ...interface{}) {
// 	logger.Sugar().Errorf(format, v...)
// }

// func Panic(format string, v ...interface{}) {
// 	logger.Sugar().Panicf(format, v...)
// }
