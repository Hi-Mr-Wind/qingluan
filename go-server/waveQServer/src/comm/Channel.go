package comm

import (
	"sync"
	"waveQServer/src/core/message"
)

var Play = sync.WaitGroup{}

// SubMessageChan 订阅消息通道
var SubMessageChan = make(chan *message.SubMessage, 0)
