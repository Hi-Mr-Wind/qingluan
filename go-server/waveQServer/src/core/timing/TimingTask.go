package timing

import (
	"github.com/robfig/cron/v3"
	"time"
	"waveQServer/src/config"
	"waveQServer/src/core/database"
	"waveQServer/src/core/message"
	"waveQServer/src/utils/logutil"
)

func init() {
	go clear()
}

// 清理持久化文件
func clear() {
	cronJob := cron.New()
	// 定义一个每分钟一次的任务
	_, err := cronJob.AddFunc("0 0 23 * * *", func() {
		now := time.Now()
		date := now.AddDate(0, 0, -int(config.GetConfig().PossessTime)).UnixNano() // 获取配置的消息持有天数
		deleteSubMessage(date)
		deleteDelayedMessage(date)
		deleteWeightMessage(date)
		deleteExclusiveMessage(date)
		deleteRandomMessage(date)
	})
	if err != nil {
		logutil.LogError(err.Error())
		return
	}
	cronJob.Start()
	// 程序不会退出，除非手动停止
	select {}
}

// 删除指定天数之前的订阅消息
func deleteSubMessage(tim int64) {
	subMessage := new(message.SubMessage)
	database.GetDb().Where("timestamp <= ?", tim).Delete(subMessage)
	logutil.LogInfo("delete past due subMessage accomplish ")
}

// 删除指定天数之前的延迟消息
func deleteDelayedMessage(tim int64) {
	m := new(message.DelayedMessage)
	database.GetDb().Where("timestamp <= ?", tim).Delete(m)
	logutil.LogInfo("delete past due DelayedMessage accomplish ")
}

// 删除指定天数之前的权重消息
func deleteWeightMessage(tim int64) {
	m := new(message.WeightMessage)
	database.GetDb().Where("timestamp <= ?", tim).Delete(m)
	logutil.LogInfo("delete past due WeightMessage accomplish ")
}

// 删除指定天数之前的独享消息
func deleteExclusiveMessage(tim int64) {
	m := new(message.ExclusiveMessage)
	database.GetDb().Where("timestamp <= ?", tim).Delete(m)
	logutil.LogInfo("delete past due ExclusiveMessage accomplish ")
}

// 删除指定天数之前的权重随机消息
func deleteRandomMessage(tim int64) {
	m := new(message.RandomMessage)
	database.GetDb().Where("timestamp <= ?", tim).Delete(m)
	logutil.LogInfo("delete past due RandomMessage accomplish ")
}
