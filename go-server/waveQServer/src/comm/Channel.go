package comm

import "waveQServer/src/core/message"

// SubMessageChan 订阅消息通道
var SubMessageChan = make(chan *message.SubMessage, 0)
