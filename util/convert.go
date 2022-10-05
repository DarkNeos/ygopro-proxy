package util

import (
	"encoding/binary"
	"unicode/utf16"
)

const UTF16_BUFFER_MAX_LEN int = 20
const FILLING_TOKEN uint16 = 0xcccc

func StrToUtf16Buffer(s string) []uint16 {
	b := make([]uint16, UTF16_BUFFER_MAX_LEN, UTF16_BUFFER_MAX_LEN)
	for i := range b {
		b[i] = FILLING_TOKEN
	}

	s_utf16 := utf16.Encode([]rune(s))

	// todo: optimize
	for i, v := range s_utf16 {
		if i < UTF16_BUFFER_MAX_LEN {
			b[i] = v

			if i == len(s_utf16)-1 && i < len(b)-1 {
				b[i+1] = 0
			}
		} else {
			break
		}
	}

	return b
}

func Utf16BufferToStr(p []byte) string {
	v := chunkBytesToUint16s(p)

	return string(utf16.Decode(v))
}

func Uint16BufToByteBuf(u16_b []uint16) []byte {
	b := make([]byte, 0, len(u16_b)*2)

	for _, v := range u16_b {
		// little endian
		b = append(b, byte(v), byte(v>>8))
	}

	return b
}

func Int32ArrayToByteArray(v []int32) []byte {
	b := make([]byte, 0, len(v)*4)

	for _, i := range v {
		b = append(b, byte(i), byte(i>>8), byte(i>>16), byte(i>>24))
	}

	return b
}

func chunkBytesToUint16s(items []byte) []uint16 {
	const chunkSize = 2
	var chunks []uint16

	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, binary.LittleEndian.Uint16(items))
	}

	return chunks
}
