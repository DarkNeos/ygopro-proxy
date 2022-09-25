// todo: use interface
package darkneos

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"unicode/utf16"

	"github.com/sktt1ryze/ygopro-proxy/DarkNeos/ygopropb"
	"google.golang.org/protobuf/proto"
)

const FILLING_TOKEN uint16 = 0xcccc
const UTF16_BUFFER_MAX_LEN int = 20
const PACKET_MIN_LEN int = 3

const (
	ProtobufToRawBuf = 1
	RawBufToProtobuf = 2
)

const (
	CtosProtoPlayerInfo = 16
	CtosProtoJoinGame   = 18

	StocJoinGame = 18
	StocChat     = 25
)

type YgoPacket struct {
	PacketLen uint16
	Proto     uint8
	Exdata    []byte
}

type HostInfo struct {
	Lflist        uint32
	Rule          uint8
	Mode          uint8
	DuelRule      uint8
	NoCheckDeck   bool
	NoShuffleDeck bool
	StartLp       uint32
	StartHand     uint8
	DrawCount     uint8
	TimeLimit     uint16
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

func bufferToPacket(p []byte) (YgoPacket, error) {
	if len(p) < PACKET_MIN_LEN {
		return YgoPacket{}, errors.New(fmt.Sprintf("Packet len too short, len=%d", len(p)))
	}

	// todo: impl Reader/Writer for buffer
	packet_len := binary.LittleEndian.Uint16(p)
	proto := p[2]
	exdata := p[3 : packet_len+2]

	return YgoPacket{
		PacketLen: packet_len,
		Proto:     proto,
		Exdata:    exdata,
	}, nil
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
		packet, err := bufferToPacket(src)
		if err != nil {
			return nil, err
		}

		var pb ygopropb.YgoStocMsg
		switch packet.Proto {
		case StocChat:
			msg := transformChat(packet)
			pb = ygopropb.YgoStocMsg{
				Msg: &msg,
			}
		case StocJoinGame:
			msg := transformJoinGame_(packet)
			pb = ygopropb.YgoStocMsg{
				Msg: &msg,
			}
		default:
			return nil, errors.New("Unhandled YgoStocMsg type")
		}

		return proto.Marshal(&pb)
	} else {
		return nil, errors.New("Unknown tranformType")
	}
}

// +++++ Client To Server +++++

// @Name: [20]uint16
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
// @Passwd: [20]uint16
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

// +++++ Server To Client +++++

// @player: uint16
// @message: []uint16
func transformChat(pkt YgoPacket) ygopropb.YgoStocMsg_StocChat {
	player := int32(binary.LittleEndian.Uint16(pkt.Exdata))
	message := utf16BufferToStr(pkt.Exdata[2:])

	return ygopropb.YgoStocMsg_StocChat{
		StocChat: &ygopropb.StocChat{
			Player: player,
			Msg:    message,
		},
	}
}

// @lflist: uint32
// @rule: uint8
// @mode: uint8
// @duel_rule: uint8
// @no_check_deck: bool
// @no_shuffle_deck: bool
// @start_lp: uint32
// @start_hand: uint8
// @draw_count: uint8
// @time_limit: uint16
func transformJoinGame_(pkt YgoPacket) ygopropb.YgoStocMsg_StocJoinGame {
	hostInfo := HostInfo{}
	exData := bytes.NewBuffer(pkt.Exdata)

	if err := binary.Read(exData, binary.LittleEndian, &hostInfo); err != nil {
		log.Fatal(err)
	}

	return ygopropb.YgoStocMsg_StocJoinGame{
		StocJoinGame: &ygopropb.StocJoinGame{
			Lflist:        int32(hostInfo.Lflist),
			Rule:          int32(hostInfo.Rule),
			Mode:          int32(hostInfo.Mode),
			DuelRule:      int32(hostInfo.DuelRule),
			NoCheckDeck:   hostInfo.NoCheckDeck,
			NoShuffleDeck: hostInfo.NoShuffleDeck,
			StartLp:       int32(hostInfo.StartLp),
			StartHand:     int32(hostInfo.StartHand),
			DrawCount:     int32(hostInfo.DrawCount),
			TimeLimit:     int32(hostInfo.TimeLimit),
		},
	}
}

// +++++ Util Functions +++++

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

func utf16BufferToStr(p []byte) string {
	v := chunkBytesToUint16s(p)

	return string(utf16.Decode(v))
}

func uint16BufToByteBuf(u16_b []uint16) []byte {
	b := make([]byte, 0, len(u16_b)*2)

	for _, v := range u16_b {
		// little endian
		b = append(b, byte(v), byte(v>>8))
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
