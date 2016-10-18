package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func InitLog() {
	// 使用JSON格式错误输出，而非常规错误输出
	log.SetFormatter(&log.JSONFormatter{})
	// stdout输出错误，也可以是一个文件
	log.SetOutput(os.Stderr)
	// 仅仅输出warning之上的Log
	log.SetLevel(log.DebugLevel)
}

func main() {
	InitLog()

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	//  这里会直接宕掉
	/*
		log.WithFields(log.Fields{
			"omg":    true,
			"number": 100,
		}).Fatal("The ice breaks!")
	*/

	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// 这里会调出堆栈
	log.Panic("I'm bailing.")
	// 调用Fatal会直接AppExit(1)
	log.Fatal("Bye.")
}
