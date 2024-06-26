package utils

import (
	"sync"
)

var BytePool1k = NewBytePool(1 << 10)
var BytePool4k = NewBytePool(4 << 10)
var BytePool16k = NewBytePool(16 << 10)
var BytePool32k = NewBytePool(32 << 10)

type BytesBuf struct {
	Bytes []byte
}

// BytePool pool for get byte slice
type BytePool struct {
	length  int
	rawPool *sync.Pool
}

// NewBytePool 创建新对象
func NewBytePool(length int) *BytePool {
	if length < 0 {
		length = 1024
	}
	return &BytePool{
		length: length,
		rawPool: &sync.Pool{
			New: func() any {
				return &BytesBuf{
					Bytes: make([]byte, length),
				}
			},
		},
	}
}

// Get 获取一个新的byte slice
func (this *BytePool) Get() *BytesBuf {
	return this.rawPool.Get().(*BytesBuf)
}

// Put 放回一个使用过的byte slice
func (this *BytePool) Put(ptr *BytesBuf) {
	this.rawPool.Put(ptr)
}

// Length 单个字节slice长度
func (this *BytePool) Length() int {
	return this.length
}
