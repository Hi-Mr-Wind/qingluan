package fileUtil

import (
	"os"
)

type BufferWrite struct {
	file  *os.File
	data  []byte
	index int
}

// NewBufferWrite 构造缓冲写出文件对象
func NewBufferWrite(file *os.File, bufferSize ...int) *BufferWrite {
	buf := new(BufferWrite)
	buf.index = 0
	buf.file = file
	buf.data = make([]byte, bufferSize[0])
	return buf
}

// WriteBytes 写出到缓冲
func (b *BufferWrite) WriteBytes(data []byte) error {
	if len(data) > len(b.data) {
		_, err := b.file.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
	if len(data)+b.index > len(b.data) {
		b.Flush()
		copy(b.data[b.index:], data)
		b.index += len(data)
		return nil
	}
	return nil
}

// WriteString 写出字符串到缓冲
func (b *BufferWrite) WriteString(data string) error {
	return b.WriteBytes([]byte(data))
}

// Flush 将缓冲数据写入磁盘
func (b *BufferWrite) Flush() {
	_, err := b.file.Write(b.data[0:b.index])
	if err != nil {
		return
	}
	b.index = 0
}
