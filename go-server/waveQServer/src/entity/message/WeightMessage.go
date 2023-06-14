package message

type WeightMessage struct {
	//消息id
	Id string `json:"id"`
	//生产者ID
	ProducerID []byte `json:"producerID"`
	// 消息生成的时间戳
	Timestamp int64 `json:"timestamp"`
	//可消费次数
	Number int32
	//权重
	Weight uint32
	//消息内容
	Data []byte
}
