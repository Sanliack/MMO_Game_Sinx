package simodel

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"mmm/sinx/config"
	"mmm/sinx/siface"
)

// Len,ID,data
type PackMsgModel struct {
}

func (p *PackMsgModel) GetHeadLen() uint32 {
	// 返回id 和 type （4+4）
	return 8
}

func (p *PackMsgModel) PackMsg(msg siface.MessageFace) ([]byte, error) {
	databuf := bytes.NewBuffer([]byte{})
	err := binary.Write(databuf, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		fmt.Println("pack len error", err)
		return nil, err
	}
	err = binary.Write(databuf, binary.LittleEndian, msg.GetMsgID())
	if err != nil {
		fmt.Println("pack ID error", err)
		return nil, err
	}
	err = binary.Write(databuf, binary.LittleEndian, msg.GetData())
	if err != nil {
		fmt.Println("pack Data error", err)
		return nil, err
	}
	return databuf.Bytes(), nil
}

func (p *PackMsgModel) UnPackMsg(buf []byte) (siface.MessageFace, error) {
	msg := &MessageModel{}
	ioreader := bytes.NewReader(buf)
	err := binary.Read(ioreader, binary.LittleEndian, &msg.Len)
	if err != nil {
		fmt.Println("Unpack len err", err)
		return nil, err
	}
	err = binary.Read(ioreader, binary.LittleEndian, &msg.Id)
	if err != nil {
		fmt.Println("Unpack ID err", err)
		return nil, err
	}

	if uint32(config.SinxConfig.MaxTranSize) < msg.Len {
		return nil, errors.New("msg长度超出限制")
	}

	msg.Content = make([]byte, msg.GetMsgLen())
	_, err = io.ReadFull(ioreader, msg.Content)
	if err != nil {
		fmt.Println("Unpack content err", err)
		return nil, err
	}

	return msg, nil
}

func (p *PackMsgModel) PackMsgByOther(id uint32, content []byte) ([]byte, error) {
	mes := NewMessageModel(id, uint32(len(content)), content)
	msg, err := p.PackMsg(mes)
	if err != nil {
		fmt.Println("packmsgbyother error", err)
		return nil, err
	}
	return msg, err
}

func NewPackMsgModel() *PackMsgModel {
	return &PackMsgModel{}
}
