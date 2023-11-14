package utils

import "sync"

// code from shadowsocks-go

type Buffer struct {
	Data []byte
	Next *Buffer
}

type LeakyBuf struct {
	bufSize  int // size of each buffer
	maxCount int
	count    int
	head     *Buffer
	tail     *Buffer
	mu       sync.Mutex
}

const leakyBigBufSize = 4200 // data.len(2) + hmacsha1(10) + data(4096)
const leakyBufSize = 4096 // data.len(2) + hmacsha1(10) + data(4096)=4108
const LeakyBufSize = 4096 // data.len(2) + hmacsha1(10) + data(4096)=4108
const ReadsBufSize = leakyBufSize - 96
const UDPBufSize = 30 * 1024 // supported max udp packet is 30k
const maxNBuf = 120000  //30000 * 4

var leakyBuf = NewLeakyBuf(maxNBuf, leakyBufSize)

// NewLeakyBuf creates a leaky buffer which can hold at most n buffer, each
// with bufSize bytes.
func NewLeakyBuf(n, bufSize int) *LeakyBuf {
	return &LeakyBuf{
		bufSize:  bufSize,
		maxCount: n,
	}
}

// Get returns a buffer from the leaky buffer or create a new buffer.
func (lb *LeakyBuf) GetBigBuf() (*Buffer) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	if lb.head == nil {
		return &Buffer{
			Data: make([]byte, leakyBigBufSize),
		}
	} else {
		b := lb.head
		lb.head = b.Next
		if lb.head == nil {
			lb.tail = nil
		}
		b.Next = nil
		lb.count--
		return b
	}
}

// Get returns a buffer from the leaky buffer or create a new buffer.
func (lb *LeakyBuf) Get() (*Buffer) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	if lb.head == nil {
		return &Buffer{
			Data: make([]byte, lb.bufSize),
		}
	} else {
		b := lb.head
		lb.head = b.Next
		if lb.head == nil {
			lb.tail = nil
		}
		b.Next = nil
		lb.count--
//		b.Data = make([]byte,lb.bufSize)
		return b
	}
}


// Put add the buffer into the free buffer pool for reuse. Panic if the buffer
// size is not the same with the leaky buffer's. This is intended to expose
// error usage of leaky buffer.
func (lb *LeakyBuf) Put(b *Buffer) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	if lb.count >= lb.maxCount {
		return
	}
	b.Next = nil
	if lb.tail == nil {
		lb.head = b
		lb.tail = b
	} else {
		lb.tail.Next = b
		lb.tail = b
	}
	lb.count++
}

func GetBuffer() *Buffer {
	return leakyBuf.Get()
}

func GetBigBuffer() *Buffer {
	return leakyBuf.GetBigBuf()
}

func PutBuffer(b *Buffer) {
	leakyBuf.Put(b)
}

func GetUdpBuffer() *Buffer {
	return &Buffer{
		Data: make([]byte, UDPBufSize),
	}
}

func PutUdpBuffer(buf *Buffer) {

}