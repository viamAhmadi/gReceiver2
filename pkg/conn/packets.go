package conn

import (
	"fmt"
	"github.com/viamAhmadi/gReceiver2/pkg/util"
	"strconv"
)

type Message struct {
	Destination string // 27 byte
	Id          int    // 5 bytes
	ConnId      string // 20 char
	Content     string // 8 kilo
}

type Factor struct {
	ConnId     string
	Successful byte
	List       []string
}

func ConvertToMessage(b *[]byte) (*Message, error) {
	if cap(*b) < 54 {
		return nil, ErrConvertToModel
	}
	i, err := strconv.Atoi(util.RemoveAdditionalCharacters((*b)[1:6]))
	if err != nil {
		return nil, ErrConvertToModel
	}
	return &Message{
		Id:          i,
		Destination: util.RemoveAdditionalCharacters((*b)[6:34]),
		ConnId:      string((*b)[34:55]),
		Content:     string((*b)[32:]),
	}, nil
}

func SerializeMessage(id int, destination, connId string, content *string) *[]byte {
	v := []byte(fmt.Sprintf("m%s%s%s%s", util.ConvertIntToBytes(id), util.ConvertDesToBytes(destination), connId, *content))
	return &v
}

func ConvertToReceiveConn(from []byte, b []byte) (*ReceiveConn, error) {
	if cap(b) < 52 {
		return nil, ErrConvertToModel
	}
	count, err := strconv.Atoi(util.RemoveAdditionalCharacters(b[49:53]))
	if err != nil {
		return nil, err
	}
	return NewReceiveConn(util.RemoveAdditionalCharacters(b[1:28]), string(b[28:49]), count, from), nil
}

func SerializeReceiveConn(destination string, count int, id string) []byte {
	return []byte(fmt.Sprintf("c%s%s%s%s", util.ConvertDesToBytes(destination), id, util.ConvertIntToBytes(count), util.ConvertIntToBytes(count)))
}
