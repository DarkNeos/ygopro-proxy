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

const (
	CtosProtoPlayerInfo = 16
	CtosProtoJoinGame   = 18
)

type YgoPacket struct {
	PacketLen uint16
	Proto     uint8
	Exdata    []byte
}

func packetToBuffer(pkt YgoPacket) []byte {
	buf := make([]byte, 0)

	// packet len
	buf = append(buf, byte(pkt.PacketLen), byte(pkt.PacketLen>>8))
	// proto
	buf = append(buf, byte(pkt.Proto))
	// exdata
	for _, v := range pkt.Exdata {
		buf = append(buf, v)
	}

	return buf
}

func Transform(src []byte, tranformType int) ([]byte, error) {
	if tranformType == ProtobufToRawBuf {
		message := &ygopropb.YgoCtosMsg{}
		err := proto.Unmarshal(src, message)
		if err != nil {
			return nil, err
		}

		var packet YgoPacket
		switch message.Msg.(type) {
		case *(ygopropb.YgoCtosMsg_CtosPlayerInfo):
			packet = transformPlayerInfo(message.GetCtosPlayerInfo())
		case *(ygopropb.YgoCtosMsg_CtosJoinGame):
			packet = transformJoinGame(message.GetCtosJoinGame())
		default:
			return nil, errors.New("Unhandled YgoCtosMsg type")
		}

		return packetToBuffer(packet), nil

	} else if tranformType == RawBufToProtobuf {
		return nil, errors.New("Unhandled tranformType")
	} else {
		return nil, errors.New("Unknown tranformType")
	}
}

// todo: use interface

// @Name: [uint16, 20]
func transformPlayerInfo(pb *ygopropb.CtosPlayerInfo) YgoPacket {
	buf := strToUtf16Buffer(pb.Name)
	exdata := uint16BufToByteBuf(buf)

	return YgoPacket{
		PacketLen: uint16(len(exdata)) + 1,
		Proto:     CtosProtoPlayerInfo,
		Exdata:    exdata,
	}
}

// @Version: uint16
// @Gameid: uint32
// @Passwd: [uint16, 20]
func transformJoinGame(pb *ygopropb.CtosJoinGame) YgoPacket {
	exdata := make([]byte, 0)

	version := uint16(pb.Version)
	exdata = append(exdata, byte(version), byte(version>>8), 0, 0)

	exdata = append(exdata, byte(pb.Gameid), byte(pb.Gameid>>8), byte(pb.Gameid>>16), byte(pb.Gameid>>24))

	for _, v := range uint16BufToByteBuf(strToUtf16Buffer(pb.Passwd)) {
		exdata = append(exdata, v)
	}

	return YgoPacket{
		PacketLen: uint16(len(exdata)) + 1,
		Proto:     CtosProtoJoinGame,
		Exdata:    exdata,
	}
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
