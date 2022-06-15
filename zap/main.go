package main

import (
	"encoding/json"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func sugger_logger() {

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "www.baidu.com"

	/**
	*输出日志到控制台
	*打印日志错误结构体数据
	*自定义输出参数， 如 url,attempt, ceshi, file: "./1.txt""
	 */
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"ceshi", "cs",
		"file", "./1.txt",
		"backoff", time.Second,
	)
	/*
	*输出数据到控制台
	*  等价于： msg := fmt.Sprintf("Failed to fetch URL: %s", fmtArgs...)
	*
	 */
	sugar.Infof("Failed to fetch URL: %s", url)
}

func log_logger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	url := "www.aiqiyi.com"
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

/*
*日志写入文件
 */
func zap_file_logger() {

	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "./logs.txt"],
		"errorOutputPaths": ["stderr"],
		"initialFields": {"foo": "bar"},
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	url := "www.aiqiyi.com"
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

/**
* 改写 zap日志文件记录
* zapcore:核心配置内容
 */
func zap_file_logger1() {

	cfg := zap.NewProductionConfig()

	cfg.OutputPaths = []string{
		"./logs.txt",
	}

	//重构时间格式配置：EncodeTime:     zapcore.ISO8601TimeEncoder,
	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "www.baiduafaf.com"

	/**
	*输出日志到控制台
	*打印日志错误结构体数据
	*自定义输出参数， 如 url,attempt, ceshi, file: "./1.txt""
	 */
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"ceshi", "cs",
		"file", "./1.txt",
		"backoff", time.Second,
	)
	/*
	*输出数据到控制台
	*  等价于： msg := fmt.Sprintf("Failed to fetch URL: %s", fmtArgs...)
	*
	 */
	// sugar.Infof("Failed to fetch URL: %s", url)
}

type LoggerInfo struct {
	url  string
	msg  string
	time string
}

func file_logger() {

	// s := LoggerInfo{
	// 	url:  "www.baidu",
	// 	msg:  "2323323",
	// 	time: "2020-02-11",
	// }

	// os.OpenFile("./aa.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)

}

func main() {

	// sugger_logger()
	// log_logger()
	zap_file_logger1()
}
