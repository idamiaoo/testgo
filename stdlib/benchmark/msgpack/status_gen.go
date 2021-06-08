package msgpack

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Status) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Age":
			z.Age, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Money":
			z.Money, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "Areas":
			var zajw uint32
			zajw, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Areas) >= int(zajw) {
				z.Areas = (z.Areas)[:zajw]
			} else {
				z.Areas = make([]string, zajw)
			}
			for zxvk := range z.Areas {
				z.Areas[zxvk], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "QinFu":
			var zwht uint32
			zwht, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.QinFu == nil && zwht > 0 {
				z.QinFu = make(map[string]int, zwht)
			} else if len(z.QinFu) > 0 {
				for key, _ := range z.QinFu {
					delete(z.QinFu, key)
				}
			}
			for zwht > 0 {
				zwht--
				var zbzg string
				var zbai int
				zbzg, err = dc.ReadString()
				if err != nil {
					return
				}
				zbai, err = dc.ReadInt()
				if err != nil {
					return
				}
				z.QinFu[zbzg] = zbai
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Status) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "Age"
	err = en.Append(0x85, 0xa3, 0x41, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Age)
	if err != nil {
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Money"
	err = en.Append(0xa5, 0x4d, 0x6f, 0x6e, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteFloat64(z.Money)
	if err != nil {
		return
	}
	// write "Areas"
	err = en.Append(0xa5, 0x41, 0x72, 0x65, 0x61, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Areas)))
	if err != nil {
		return
	}
	for zxvk := range z.Areas {
		err = en.WriteString(z.Areas[zxvk])
		if err != nil {
			return
		}
	}
	// write "QinFu"
	err = en.Append(0xa5, 0x51, 0x69, 0x6e, 0x46, 0x75)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.QinFu)))
	if err != nil {
		return
	}
	for zbzg, zbai := range z.QinFu {
		err = en.WriteString(zbzg)
		if err != nil {
			return
		}
		err = en.WriteInt(zbai)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Status) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Age"
	o = append(o, 0x85, 0xa3, 0x41, 0x67, 0x65)
	o = msgp.AppendInt(o, z.Age)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Money"
	o = append(o, 0xa5, 0x4d, 0x6f, 0x6e, 0x65, 0x79)
	o = msgp.AppendFloat64(o, z.Money)
	// string "Areas"
	o = append(o, 0xa5, 0x41, 0x72, 0x65, 0x61, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Areas)))
	for zxvk := range z.Areas {
		o = msgp.AppendString(o, z.Areas[zxvk])
	}
	// string "QinFu"
	o = append(o, 0xa5, 0x51, 0x69, 0x6e, 0x46, 0x75)
	o = msgp.AppendMapHeader(o, uint32(len(z.QinFu)))
	for zbzg, zbai := range z.QinFu {
		o = msgp.AppendString(o, zbzg)
		o = msgp.AppendInt(o, zbai)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Status) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zhct uint32
	zhct, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zhct > 0 {
		zhct--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Age":
			z.Age, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Money":
			z.Money, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "Areas":
			var zcua uint32
			zcua, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Areas) >= int(zcua) {
				z.Areas = (z.Areas)[:zcua]
			} else {
				z.Areas = make([]string, zcua)
			}
			for zxvk := range z.Areas {
				z.Areas[zxvk], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "QinFu":
			var zxhx uint32
			zxhx, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.QinFu == nil && zxhx > 0 {
				z.QinFu = make(map[string]int, zxhx)
			} else if len(z.QinFu) > 0 {
				for key, _ := range z.QinFu {
					delete(z.QinFu, key)
				}
			}
			for zxhx > 0 {
				var zbzg string
				var zbai int
				zxhx--
				zbzg, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zbai, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					return
				}
				z.QinFu[zbzg] = zbai
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Status) Msgsize() (s int) {
	s = 1 + 4 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name) + 6 + msgp.Float64Size + 6 + msgp.ArrayHeaderSize
	for zxvk := range z.Areas {
		s += msgp.StringPrefixSize + len(z.Areas[zxvk])
	}
	s += 6 + msgp.MapHeaderSize
	if z.QinFu != nil {
		for zbzg, zbai := range z.QinFu {
			_ = zbai
			s += msgp.StringPrefixSize + len(zbzg) + msgp.IntSize
		}
	}
	return
}
