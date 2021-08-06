package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

func init() {
	// 如果设置此处，则日志输出为json格式，如下：
	// {"level":"info","msg":"I am info","name":"test","namespace":"kube-system","time":"2021-08-06T16:51:11+08:00"}
	// logrus.SetFormatter(&logrus.JSONFormatter{})

	// 开启 Debug 模式
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	flag.Parse()
	// 添加的字段WithField或WithFields, 字段会自动添加到所有日志记录事件中
	log := logrus.WithFields(logrus.Fields{
		"name":      "test",
		"namespace": "kube-system",
	})

	log.WriterLevel(1)

	log.Debug("I am debug")
	log.Info("I am info")
	log.Error("I am error")
	log.Warn("I am warn")
	log.Fatal("I am fatal,i will exit now")
}
