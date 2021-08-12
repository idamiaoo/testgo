package proto

import (
	"encoding/binary"
	"io"
)

type File struct {
	Name string
	Data []byte
}

func (f *File) ReadFile(reader io.Reader) error {
	headLenBuf := make([]byte, 4)
	_, err := io.ReadFull(reader, headLenBuf)
	if err != nil {
		return err
	}
	headLen := binary.BigEndian.Uint32(headLenBuf)
	headBuf := make([]byte, headLen-4)
	_, err = io.ReadFull(reader, headBuf)
	if err != nil {
		return err
	}
	f.Name = string(headBuf[0 : headLen-8])
	totalLen := binary.BigEndian.Uint32(headBuf[headLen-8:])
	data := make([]byte, totalLen-headLen)
	_, err = io.ReadFull(reader, data)
	if err != nil {
		return err
	}
	f.Data = data
	return nil
}

func (f *File) WriteFile(writer io.Writer) error {
	headLen := 4 + len(f.Name) + 4
	totalLen := headLen + len(f.Data)
	headBuf := make([]byte, headLen)
	binary.BigEndian.PutUint32(headBuf[0:4], uint32(headLen))
	copy(headBuf[4:4+len(f.Name)], f.Name)
	binary.BigEndian.PutUint32(headBuf[4+len(f.Name):], uint32(totalLen))
	_, err := writer.Write(headBuf)
	if err != nil {
		return err
	}
	_, err = writer.Write(f.Data)
	if err != nil {
		return err
	}
	return nil
}
