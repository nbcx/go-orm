package buffers

import "github.com/valyala/bytebufferpool"

var _ Buffer = &bytebufferpool.ByteBuffer{}

type Buffer interface {
	Write(p []byte) (int, error)
	WriteString(s string) (int, error)
	WriteByte(c byte) error
	String() string
}

func Get() Buffer {
	return bytebufferpool.Get()
}

func Put(bf Buffer) {
	bytebufferpool.Put(bf.(*bytebufferpool.ByteBuffer))
}
