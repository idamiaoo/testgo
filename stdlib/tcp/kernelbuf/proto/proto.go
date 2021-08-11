package proto

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

type Proto struct {
	Seq  int
	Body string
}

func (p *Proto) String() string {
	return fmt.Sprintf("version:%d, body:%s", p.Seq, p.Body)
}

func (p *Proto) ReadFrom(reader io.Reader) error {
	headBuf := make([]byte, 4)
	_, err := io.ReadFull(reader, headBuf)
	if err != nil {
		return err
	}
	bodyBuf := make([]byte, binary.BigEndian.Uint32(headBuf))
	_, err = io.ReadFull(reader, bodyBuf)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bodyBuf, p); err != nil {
		return err
	}
	return nil
}

func (p *Proto) WriteTo(writer io.Writer) error {
	bodyBuf, err := json.Marshal(p)
	if err != nil {
		return err
	}
	headBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(headBuf, uint32(len(bodyBuf)))
	_, err = writer.Write(headBuf)
	if err != nil {
		return err
	}
	_, err = writer.Write(bodyBuf)
	if err != nil {
		return err
	}
	return nil
}
