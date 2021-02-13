package conn

import (
	"fmt"
	"github.com/viamAhmadi/gReceiver2/pkg/util"
	"strconv"
	"strings"
)

type Send struct {
	ConnId string
}

type Message struct {
	Id          int    // 5 bytes
	ConnId      string // 20 char
	Content     string // 8 kilo
}

type Factor struct {
	ConnId     string
	Successful int
	List       *[]string
}

func ConvertToSend(b []byte) (Send, error) {
	if len(b) < 20 {
		return Send{}, ErrConvertToModel
	}
	return Send{ConnId: string(b[1:21])}, nil
}

func SerializeSend(connId string) []byte {
	return []byte(fmt.Sprintf("s%s", connId))
}

func ConvertToMessage(b *[]byte) (*Message, error) {
	if cap(*b) < 26 {
		return nil, ErrConvertToModel
	}
	i, err := strconv.Atoi(util.RemoveAdditionalCharacters((*b)[1:6]))
	if err != nil {
		return nil, ErrConvertToModel
	}
	return &Message{
		Id: i,
		//Destination: util.RemoveAdditionalCharacters((*b)[6:34]),
		ConnId:  string((*b)[6:26]),
		Content: string((*b)[26:]),
	}, nil
}

func SerializeMessage(m *Message) *[]byte {
	//v := []byte(fmt.Sprintf("m%s%s%s%s", util.ConvertIntToBytes(m.Id), util.ConvertDesToBytes(destination), connId, *content))
	v := []byte(fmt.Sprintf("m%s%s%s", util.ConvertIntToBytes(m.Id), m.ConnId, m.Content))
	return &v
}

func ConvertToReceiveConn(from []byte, b []byte) (*ReceiveConn, error) {
	if cap(b) < 52 {
		return nil, ErrConvertToModel
	}
	count, err := strconv.Atoi(util.RemoveAdditionalCharacters(b[48:53]))
	if err != nil {
		return nil, err
	}
	return NewReceiveConn(util.RemoveAdditionalCharacters(b[1:28]), string(b[28:48]), count, from), nil
}

func SerializeReceiveConn(destination string, count int, id string) []byte {
	return []byte(fmt.Sprintf("c%s%s%s", util.ConvertDesToBytes(destination), id, util.ConvertIntToBytes(count)))
}

func ConvertToFactor(b *[]byte) (*Factor, error) {
	if len(*b) < 21 {
		return nil, ErrConvertToModel
	}
	successful, err := strconv.Atoi(string((*b)[21]))
	if err != nil {
		return nil, ErrConvertToModel
	}
	var list []string
	if successful != YES {
		nums := strings.Split(string((*b)[22:]), ".")
		for i := 0; i < len(nums); i++ {
			val := nums[i]
			if val == "" {
				continue
			}
			list = append(list, val)
		}
	}
	return &Factor{
		Successful: successful,
		ConnId:     string((*b)[1:21]),
		List:       &list,
	}, nil
}

func SerializeFactor(f *Factor) *[]byte {
	tmp := ""
	if f.Successful != YES {
		if f.List != nil {
			tmp = strings.Join(*f.List, ".")
		}
	}
	b := []byte(fmt.Sprintf("f%s%s%s", f.ConnId, strconv.Itoa(f.Successful), tmp))
	return &b
}
