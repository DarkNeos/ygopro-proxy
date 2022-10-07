// todo: use interface
package darkneos

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"

	"github.com/sktt1ryze/ygopro-proxy/DarkNeos/ygopropb"
	util "github.com/sktt1ryze/ygopro-proxy/util"
	"google.golang.org/protobuf/proto"
)

const PACKET_MIN_LEN int = 3
const COMPONENT = "[transform]"
const (
	ProtobufToRawBuf = 1
	RawBufToProtobuf = 2
)

const (
	CtosUpdateDeck      = 2
	CtosProtoPlayerInfo = 16
	CtosProtoJoinGame   = 18
	CtosHsReady         = 34
	CtosHsStart         = 37

	StocJoinGame       = 18
	StocTypeChange     = 19
	StocChat           = 25
	StocHsPlayerEnter  = 32
	StocHsPlayerChange = 33
	StocHsWatchChange  = 34
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

func bufferToPacket(p []byte, ctx *util.Context) (YgoPacket, error) {
	if len(p) < PACKET_MIN_LEN {
		return YgoPacket{}, errors.New(fmt.Sprintf("Packet len too short, len=%d", len(p)))
	}

	// todo: impl Reader/Writer for buffer
	packet_len := binary.LittleEndian.Uint16(p)
	proto := p[2]

	if len(p) < int(packet_len+2) {
		log.Printf(COMPONENT+
			`Unmatched packet size, proto=%d, buffer_length=%d, packet_len=%d, infa_read_len=%d
Use the buffer length.\n`,
			proto, len(p), packet_len, ctx.InfaReadBufferLen,
		)

		packet_len = uint16(len(p) - 2)
	}

	exdata := p[3 : packet_len+2]

	return YgoPacket{
		PacketLen: packet_len,
		Proto:     proto,
		Exdata:    exdata,
	}, nil
}

func Transform(src []byte, tranformType int, ctx *util.Context) ([]byte, error) {
	if tranformType == ProtobufToRawBuf {
		message := &ygopropb.YgoCtosMsg{}
		err := proto.Unmarshal(src, message)
		if err != nil {
			return nil, err
		}

		var packet YgoPacket
		switch message.Msg.(type) {
		case *(ygopropb.YgoCtosMsg_CtosPlayerInfo):
			packet = (*pCtosPlayerInfo)(message.GetCtosPlayerInfo()).Pb2Packet()
		case *(ygopropb.YgoCtosMsg_CtosJoinGame):
			packet = (*pCtosJoinGame)(message.GetCtosJoinGame()).Pb2Packet()
		case *(ygopropb.YgoCtosMsg_CtosUpdateDeck):
			packet = (*pCtosUpdateDeck)(message.GetCtosUpdateDeck()).Pb2Packet()
		case *(ygopropb.YgoCtosMsg_CtosHsReady):
			packet = (*pCtosHsReady)(message.GetCtosHsReady()).Pb2Packet()
		case *(ygopropb.YgoCtosMsg_CtosHsStart):
			packet = (*pCtosHsStart)(message.GetCtosHsStart()).Packet2Pb()
		default:
			return nil, errors.New(COMPONENT + "Unhandled YgoCtosMsg type")
		}

		return packetToBuffer(packet), nil

	} else if tranformType == RawBufToProtobuf {
		packet, err := bufferToPacket(src, ctx)
		if err != nil {
			return nil, err
		}

		var pb ygopropb.YgoStocMsg
		switch packet.Proto {
		case StocChat:
			pb = pStocChat{}.Packet2Pb(packet)
		case StocJoinGame:
			pb = pStocJoinGame{}.Packet2Pb(packet)
		case StocHsPlayerEnter:
			pb = pStocHsPlayerEnter{}.Packet2Pb(packet)
		case StocTypeChange:
			pb = pStocTypeChage{}.Packet2Pb(packet)
		case StocHsPlayerChange:
			pb = pStocHsPlayerChange{}.Packet2Pb(packet)
		case StocHsWatchChange:
			pb = pStocHsWatchChange{}.Packet2Pb(packet)
		default:
			return nil, errors.New(fmt.Sprintf(COMPONENT+"Unhandled YgoStocMsg type, proto=%d", packet.Proto))
		}

		return proto.Marshal(&pb)
	} else {
		return nil, errors.New(COMPONENT + "Unknown tranformType")
	}
}

// +++++ Client To Server +++++

type client2Server interface {
	Pb2Packet() YgoPacket
}

type pCtosPlayerInfo ygopropb.CtosPlayerInfo

// @Name: [20]uint16
func (pb *pCtosPlayerInfo) Pb2Packet() YgoPacket {
	buf := util.StrToUtf16Buffer(pb.Name)
	exdata := util.Uint16BufToByteBuf(buf)

	return YgoPacket{
		PacketLen: uint16(len(exdata)) + 1,
		Proto:     CtosProtoPlayerInfo,
		Exdata:    exdata,
	}
}

type pCtosJoinGame ygopropb.CtosJoinGame

// @Version: uint16
// @Gameid: uint32
// @Passwd: [20]uint16
func (pb *pCtosJoinGame) Pb2Packet() YgoPacket {
	exdata := make([]byte, 0)

	version := uint16(pb.Version)
	exdata = append(exdata, byte(version), byte(version>>8), 0, 0)

	exdata = append(exdata, byte(pb.Gameid), byte(pb.Gameid>>8), byte(pb.Gameid>>16), byte(pb.Gameid>>24))

	for _, v := range util.Uint16BufToByteBuf(util.StrToUtf16Buffer(pb.Passwd)) {
		exdata = append(exdata, v)
	}

	return YgoPacket{
		PacketLen: uint16(len(exdata)) + 1,
		Proto:     CtosProtoJoinGame,
		Exdata:    exdata,
	}
}

type pCtosUpdateDeck ygopropb.CtosUpdateDeck

// @main: []int32
// @extra: []int32
// @size: []int32
func (pb *pCtosUpdateDeck) Pb2Packet() YgoPacket {
	v := make([]int32, 0)

	v = append(v, int32(len(pb.Main)+len(pb.Extra)))
	v = append(v, int32(len(pb.Side)))
	v = append(v, pb.Main...)
	v = append(v, pb.Extra...)
	v = append(v, pb.Side...)

	exdata := util.Int32ArrayToByteArray(v)

	return YgoPacket{
		PacketLen: uint16(len(exdata)) + 1,
		Proto:     CtosUpdateDeck,
		Exdata:    exdata,
	}
}

type pCtosHsReady ygopropb.CtosHsReady

// empty message
func (_ *pCtosHsReady) Pb2Packet() YgoPacket {
	return YgoPacket{
		PacketLen: 1,
		Proto:     CtosHsReady,
		Exdata:    make([]byte, 0),
	}
}

type pCtosHsStart ygopropb.CtosHsStart

// empty message
func (_ *pCtosHsStart) Packet2Pb() YgoPacket {
	return YgoPacket{
		PacketLen: 1,
		Proto:     CtosHsStart,
		Exdata:    make([]byte, 0),
	}
}

// +++++ Server To Client +++++

type server2Client interface {
	Packet2Pb(pkt YgoPacket) ygopropb.YgoStocMsg
}

type pStocChat struct{}

// @player: uint16
// @message: []uint16
func (_ pStocChat) Packet2Pb(pkt YgoPacket) ygopropb.YgoStocMsg {
	player := int32(binary.LittleEndian.Uint16(pkt.Exdata))
	message := util.Utf16BufferToStr(pkt.Exdata[2:])

	msg := ygopropb.YgoStocMsg_StocChat{
		StocChat: &ygopropb.StocChat{
			Player: player,
			Msg:    message,
		},
	}

	return ygopropb.YgoStocMsg{
		Msg: &msg,
	}
}

type pStocJoinGame struct{}

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
func (_ pStocJoinGame) Packet2Pb(pkt YgoPacket) ygopropb.YgoStocMsg {
	hostInfo := HostInfo{}
	exData := bytes.NewBuffer(pkt.Exdata)

	if err := binary.Read(exData, binary.LittleEndian, &hostInfo); err != nil {
		log.Fatal(err)
	}

	msg := ygopropb.YgoStocMsg_StocJoinGame{
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

	return ygopropb.YgoStocMsg{
		Msg: &msg,
	}
}

type pStocHsPlayerEnter struct{}

func (_ pStocHsPlayerEnter) Packet2Pb(pkt YgoPacket) ygopropb.YgoStocMsg {
	name_max := util.UTF16_BUFFER_MAX_LEN * 2
	name := util.Utf16BufferToStr(pkt.Exdata[:name_max])
	pos := pkt.Exdata[name_max] & 0x3 // todo: make sure

	msg := ygopropb.YgoStocMsg_StocHsPlayerEnter{
		StocHsPlayerEnter: &ygopropb.StocHsPlayerEnter{
			Name: name,
			Pos:  int32(pos),
		},
	}

	return ygopropb.YgoStocMsg{
		Msg: &msg,
	}
}

type pStocTypeChage struct{}

func (_ pStocTypeChage) Packet2Pb(pkt YgoPacket) ygopropb.YgoStocMsg {
	type_ := pkt.Exdata[0]
	isHost := ((type_ >> 4) & 0xf) != 0

	var selfType ygopropb.StocTypeChange_SelfType
	switch type_ & 0xf {
	case 0:
		{
			selfType = ygopropb.StocTypeChange_PLAYER1
		}
	case 1:
		{
			selfType = ygopropb.StocTypeChange_PLAYER2
		}
	case 2:
		{
			selfType = ygopropb.StocTypeChange_PLAYER3
		}
	case 3:
		{
			selfType = ygopropb.StocTypeChange_PLAYER4
		}
	case 4:
		{
			selfType = ygopropb.StocTypeChange_PLAYER5
		}
	case 5:
		{
			selfType = ygopropb.StocTypeChange_PLAYER6
		}
	default:
		{
			selfType = ygopropb.StocTypeChange_UNKNOWN
		}
	}

	msg := ygopropb.YgoStocMsg_StocTypeChange{
		StocTypeChange: &ygopropb.StocTypeChange{
			SelfType: selfType,
			IsHost:   isHost,
		},
	}

	return ygopropb.YgoStocMsg{
		Msg: &msg,
	}
}

type pStocHsPlayerChange struct{}

func (_ pStocHsPlayerChange) Packet2Pb(pkt YgoPacket) ygopropb.YgoStocMsg {
	pb := ygopropb.StocHsPlayerChange{}
	pb.State = ygopropb.StocHsPlayerChange_UNKNOWN

	status := pkt.Exdata[0]
	pos := (status >> 4) & 0xf
	state := status & 0xf

	pb.Pos = int32(pos)

	if pos < 4 {
		if state < 8 {
			pb.State = ygopropb.StocHsPlayerChange_MOVE
			pb.MovedPos = int32(state)
		} else if state == 0x9 {
			pb.State = ygopropb.StocHsPlayerChange_READY
		} else if state == 0xa {
			pb.State = ygopropb.StocHsPlayerChange_NO_READY
		} else if state == 0xb {
			pb.State = ygopropb.StocHsPlayerChange_LEAVE
		} else if state == 0x8 {
			pb.State = ygopropb.StocHsPlayerChange_TO_OBSERVER
		}
	}

	msg := ygopropb.YgoStocMsg_StocHsPlayerChange{
		StocHsPlayerChange: &pb,
	}

	return ygopropb.YgoStocMsg{
		Msg: &msg,
	}
}

type pStocHsWatchChange struct{}

func (_ pStocHsWatchChange) Packet2Pb(pkt YgoPacket) ygopropb.YgoStocMsg {
	count := binary.LittleEndian.Uint16(pkt.Exdata)

	msg := ygopropb.YgoStocMsg_StocHsWatchChange{
		StocHsWatchChange: &ygopropb.StocHsWatchChange{
			Count: int32(count),
		},
	}

	return ygopropb.YgoStocMsg{
		Msg: &msg,
	}
}
