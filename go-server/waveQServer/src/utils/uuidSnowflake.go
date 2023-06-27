package utils

import (
	"errors"
	"strconv"
	"sync"
	"time"
	"waveQServer/src/utils/logutil"
)

// 雪花算法ID，用于生成各个唯一性ID使用，同时为后续分布式架构做准备

const (
	nodeBits        = 10                                  // 节点位数
	stepBits        = 12                                  // 序列号位数
	nodeMax         = int64(-1) ^ (int64(-1) << nodeBits) // 节点数最大值
	stepMask        = int64(-1) ^ (int64(-1) << stepBits) // 序列号掩码
	timeShift       = nodeBits + stepBits                 // 时间戳左偏移量
	nodeShift       = stepBits                            // 节点左偏移量
	DefaultNodeBits = 6                                   // 默认的节点位数
	DefaultStepBits = 6                                   // 默认的序列号位数
)

// 序列开始的时间
var starTime = time.Date(2023, 6, 27, 0, 0, 0, 0, time.UTC)

var snow *snowflake

type snowflake struct {
	mu          sync.Mutex // 互斥锁
	lastTime    int64      // 上次生成的时间戳
	node        int64      // 节点ID
	step        int64      // 序列号
	nodeBits    uint8      // 节点位数
	stepBits    uint8      // 序列号位数
	maxNode     int64      // 最大节点数
	maxStep     int64      // 最大序列号数
	startTime   int64      // 开始时间
	lastErrTime int64      // 上次出错的时间
}

// NewSnowflake node为节点编号
func NewSnowflake(node int64) (*snowflake, error) {
	return newEx(node, DefaultNodeBits, DefaultStepBits)
}

func newEx(node int64, nodeBits uint8, stepBits uint8) (*snowflake, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("node number must be between 0 and 1023")
	}
	s := &snowflake{
		node:     node,
		nodeBits: nodeBits,
		stepBits: stepBits,
		maxNode:  nodeMax,
		maxStep:  stepMask,
	}
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *snowflake) init() error {
	if s.node < 0 || s.node > s.maxNode {
		return errors.New("Node number must be between 0 and " + strconv.FormatInt(s.maxNode, 10))
	}
	if s.nodeBits <= 0 || s.nodeBits > 22 {
		return errors.New("node bits must be between 1 and 22")
	}
	if s.stepBits <= 0 || s.stepBits > 20 {
		return errors.New("step bits must be between 1 and 20")
	}

	s.startTime = s.epoch()
	if s.startTime >= time.Now().UnixNano() {
		return errors.New("clock moved backwards")
	}

	s.lastTime = s.startTime
	return nil
}

func (s *snowflake) epoch() int64 {
	return starTime.UnixNano() / 1000000
}

func (s *snowflake) gen() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UnixNano() / 1000000
	if now < s.lastTime {
		// 时间回拨，则等待到最后一次生成ID时间之后再生成
		time.Sleep(time.Duration(s.lastTime - now))
		now = time.Now().UnixNano() / 1000000
		if now < s.lastTime {
			s.lastErrTime = now
			return 0
		}
	}
	if s.lastTime == now {
		s.step = (s.step + 1) & s.maxStep
		if s.step == 0 {
			now = s.waitNextMillisecond(now)
		}
	} else {
		s.step = 0
	}
	s.lastTime = now
	return (now-s.startTime)<<timeShift | (s.node << nodeShift) | s.step
}

// 等待下一毫秒
func (s *snowflake) waitNextMillisecond(timestamp int64) int64 {
	next := timestamp
	for {
		now := time.Now().UnixNano() / 1000000
		if now > next {
			return now
		}
		time.Sleep(time.Duration(next - now))
		next = time.Now().UnixNano() / 1000000
	}
}

func (s *snowflake) GetId() int64 {
	id := s.gen()
	if id == 0 {
		panic(errors.New("Error generating ID,lastTime: " + strconv.FormatInt(s.lastTime, 10) + ", lastErrTime: " + strconv.FormatInt(s.lastErrTime, 10)))
	}
	return id
}

// 初始化
func init() {
	s, err := NewSnowflake(1)
	if err != nil {
		logutil.LogError(err.Error())
	}
	snow = s
}

// GetSnowflakeId 使用默认节点1 获取一个雪花ID
func GetSnowflakeId() int64 {
	if snow == nil {
		s, err := NewSnowflake(1)
		if err != nil {
			logutil.LogError(err.Error())
		}
		snow = s
	}
	return snow.GetId()
}

// GetSnowflakeIdStr 使用默认节点1 获取一个雪花ID
func GetSnowflakeIdStr() string {
	return strconv.FormatInt(GetSnowflakeId(), 10)
}
