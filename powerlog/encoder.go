package powerlog

import "io"

var nullBytes = []byte("null")
var objKeyStr = []byte(`":"`)
var objKeyObj = []byte(`":{`)
var objKeyArr = []byte(`":[`)
var objKey = []byte(`":`)

type Encoder struct {
	buf      []byte
	isPooled byte
	w        io.Writer
	err      error
	hasKeys  bool
	keys     []string
}

const hex = "0123456789abcdef"

// grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, grow panics.
func (enc *Encoder) grow(n int) {
	if cap(enc.buf)-len(enc.buf) < n {
		Buf := make([]byte, len(enc.buf), 2*cap(enc.buf)+n)
		copy(Buf, enc.buf)
		enc.buf = Buf
	}
}

// Write appends the contents of p to b's Buffer.
// Write always returns len(p), nil.
func (enc *Encoder) writeBytes(p []byte) {
	enc.buf = append(enc.buf, p...)
}

func (enc *Encoder) writeTwoBytes(b1 byte, b2 byte) {
	enc.buf = append(enc.buf, b1, b2)
}

// WriteByte appends the byte c to b's Buffer.
// The returned error is always nil.
func (enc *Encoder) writeByte(c byte) {
	enc.buf = append(enc.buf, c)
}

// WriteString appends the contents of s to b's Buffer.
// It returns the length of s and a nil error.
func (enc *Encoder) writeString(s string) {
	enc.buf = append(enc.buf, s...)
}

func (enc *Encoder) writeStringEscape(s string) {
	l := len(s)
	for i := 0; i < l; i++ {
		c := s[i]
		if c >= 0x20 && c != '\\' && c != '"' {
			enc.writeByte(c)
			continue
		}
		switch c {
		case '\\', '"':
			enc.writeTwoBytes('\\', c)
		case '\n':
			enc.writeTwoBytes('\\', 'n')
		case '\f':
			enc.writeTwoBytes('\\', 'f')
		case '\b':
			enc.writeTwoBytes('\\', 'b')
		case '\r':
			enc.writeTwoBytes('\\', 'r')
		case '\t':
			enc.writeTwoBytes('\\', 't')
		default:
			enc.writeString(`\u00`)
			enc.writeTwoBytes(hex[c>>4], hex[c&0xF])
		}
		continue
	}
}

// AppendBytes allows a modular usage by appending bytes manually to the current state of the buffer.
func (enc *Encoder) AppendBytes(b []byte) {
	enc.writeBytes(b)
}

// AppendByte allows a modular usage by appending a single byte manually to the current state of the buffer.
func (enc *Encoder) AppendByte(b byte) {
	enc.writeByte(b)
}

// Buf returns the Encoder's buffer.
func (enc *Encoder) Buf() []byte {
	return enc.buf
}

// Write writes to the io.Writer and resets the buffer.
func (enc *Encoder) Write() (int, error) {
	i, err := enc.w.Write(enc.buf)
	enc.buf = enc.buf[:0]
	return i, err
}

func (enc *Encoder) getPreviousRune() byte {
	last := len(enc.buf) - 1
	return enc.buf[last]
}

func (enc *Encoder) keyExists(k string) bool {
	if enc.keys == nil {
		return false
	}
	for _, key := range enc.keys {
		if key == k {
			return true
		}
	}
	return false
}

// Release sends back a Encoder to the pool.
func (enc *Encoder) Release() {
	enc.isPooled = 1
	encPool.Put(enc)
}
