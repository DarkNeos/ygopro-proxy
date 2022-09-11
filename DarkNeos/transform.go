package darkneos

import (
	"errors"

	"github.com/sktt1ryze/ygopro-proxy/DarkNeos/ygoprobuffer"
	"github.com/sktt1ryze/ygopro-proxy/DarkNeos/ygopropb"
)

const FILLING_TOKEN uint16 = 0xcccc;

const (
  ProtobufToRawBuf = 1
  RawBufToProtobuf = 2
)

func Transform(src []byte, tranformType int) ([]byte, error) {
  if tranformType == ProtobufToRawBuf {
    return make([]byte, 0), nil
  } else if tranformType == RawBufToProtobuf {
    // todo
    return make([]byte, 0), nil
  } else {
    return nil, errors.New("Unknown tranformType")
  }
}


func TransformPlayerInfo(pb ygopropb.CtosPlayerInfo) ygoprobuffer.CtosPlayerInfo {
  name := [ygoprobuffer.PLAYER_NAME_MAX_LEN]uint16 {}
  for i := range name {
    name[i] = FILLING_TOKEN
  }

  info_buf := ygoprobuffer.CtosPlayerInfo {
    Name: name,
  }
  
  return info_buf
}
