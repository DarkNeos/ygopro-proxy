package darkneos

import (
	"errors"
	"unicode/utf16"

	"github.com/sktt1ryze/ygopro-proxy/DarkNeos/ygopropb"
	"google.golang.org/protobuf/proto"
)

const FILLING_TOKEN uint16 = 0xcccc
const UTF16_BUFFER_MAX_LEN int = 20

const (
	ProtobufToRawBuf = 1
	RawBufToProtobuf = 2
)

func Transform(src []byte, tranformType int) ([]byte, error) {
	if tranformType == ProtobufToRawBuf {
		message := &ygopropb.YgoCtosMsg{}
		err := proto.Unmarshal(src, message)
		if err != nil {
			return nil, err
		}

		switch message.Msg.(type) {
		case *(ygopropb.YgoCtosMsg_CtosPlayerInfo):
			return transformPlayerInfo(message.GetCtosPlayerInfo()), nil
		case *(ygopropb.YgoCtosMsg_CtosJoinGame):
			return transformJoinGame(message.GetCtosJoinGame()), nil
		default:
			return nil, errors.New("Unhandled YgoCtosMsg type")
		}
	} else if tranformType == RawBufToProtobuf {
		return nil, errors.New("Unhandled tranformType")
	} else {
		return nil, errors.New("Unknown tranformType")
	}
}

func transformPlayerInfo(pb *ygopropb.CtosPlayerInfo) []byte {
	buf := strToUtf16Buffer(pb.Name)

	return uint16BufToByteBuf(buf)
}

func transformJoinGame(pb *ygopropb.CtosJoinGame) []byte {
	buf := make([]byte, 0)

	version := uint16(pb.Version)
	buf = append(buf, byte(version), byte(pb.Version>>8), 0, 0)

	buf = append(buf, byte(pb.Gameid), byte(pb.Gameid>>8), byte(pb.Gameid>>16), byte(pb.Gameid>>24))

	for _, v := range uint16BufToByteBuf(strToUtf16Buffer(pb.Passwd)) {
		buf = append(buf, v)
	}

	return buf
}

func strToUtf16Buffer(s string) []uint16 {
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

func uint16BufToByteBuf(u16_b []uint16) []byte {
	b := make([]byte, 0, len(u16_b)*2)

	for _, v := range u16_b {
		// little endian
		b = append(b, byte(v), byte(v>>8))
	}

	return b
}
