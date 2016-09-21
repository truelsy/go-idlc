//////////////////////////////////////////////////////////////////
// Automatically-generated file. Do not edit!
//////////////////////////////////////////////////////////////////

package msg

import "encoding/binary"
import "unsafe"
import "fmt"

type Messager interface {
	GetId() uint16
	Encode([]byte) (int, error)
	Decode([]byte) error
}

type BOError struct {
	size   int
	offset int
	bufcap int
}

func (e BOError) Error() string {
	return fmt.Sprintf("ERROR - Buffer Overflow! size(%v) offset(%v) bufcap(%v)", e.size, e.offset, e.bufcap)
}

///////////////////////////////////////////////////////////////////////
// declare message id
var (
	TestMsg_ID = uint16(7000)
)

///////////////////////////////////////////////////////////////////////
// struct Item
type Item struct {
	Item_seq  uint64
	Item_name string
}

func (p *Item) Encode(offset *int, buf []byte) error {
	var bufcap int = cap(buf)
	var size int = 0

	{
		size = int(unsafe.Sizeof(p.Item_seq))
		if (size + *offset) > bufcap {
			return BOError{size, *offset, bufcap}
		}
		binary.LittleEndian.PutUint64(buf[*offset:], p.Item_seq)
		*offset += size
	}
	{
		strlen := uint16(len(p.Item_name))
		size = int(unsafe.Sizeof(strlen))
		if (size + *offset) > bufcap {
			return BOError{size, *offset, bufcap}
		}
		binary.LittleEndian.PutUint16(buf[*offset:], strlen)
		*offset += size

		size = int(strlen)
		if (size + *offset) > bufcap {
			return BOError{size, *offset, bufcap}
		}
		copy(buf[*offset:], p.Item_name)
		*offset += size
	}
	return nil
}

func (p *Item) Decode(offset *int, buf []byte) error {
	var bufcap int = cap(buf)
	var size int = 0

	{
		size = int(unsafe.Sizeof(p.Item_seq))
		if (size + *offset) > bufcap {
			return BOError{size, *offset, bufcap}
		}
		p.Item_seq = binary.LittleEndian.Uint64(buf[*offset:])
		*offset += size
	}
	{
		size = int(unsafe.Sizeof(uint16(0)))
		if (size + *offset) > bufcap {
			return BOError{size, *offset, bufcap}
		}
		strlen := binary.LittleEndian.Uint16(buf[*offset:])
		*offset += size

		size = int(strlen)
		if (size + *offset) > bufcap {
			return BOError{size, *offset, bufcap}
		}
		p.Item_name = string(buf[*offset : *offset+int(strlen)])
		*offset += size
	}
	return nil
}

///////////////////////////////////////////////////////////////////////
// message TestMsg
type TestMsg struct {
	M_s8    int8
	M_u8    uint8
	M_s16   int16
	M_u16   uint16
	M_s32   int32
	M_u32   uint32
	M_s64   int64
	M_u64   uint64
	M_str   string
	M_items []Item
}

func (p TestMsg) GetId() uint16 {
	return TestMsg_ID
}

func (p *TestMsg) Encode(buf []byte) (int, error) {
	var offset int = 0
	var bufcap int = cap(buf)
	var size int = 0

	{
		size = int(unsafe.Sizeof(p.M_s8))
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		buf[offset] = byte(p.M_s8)
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_u8))
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		buf[offset] = uint8(p.M_u8)
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_s16))
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		binary.LittleEndian.PutUint16(buf[offset:], uint16(p.M_s16))
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_u16))
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		binary.LittleEndian.PutUint16(buf[offset:], p.M_u16)
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_s32))
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		binary.LittleEndian.PutUint32(buf[offset:], uint32(p.M_s32))
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_u32))
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		binary.LittleEndian.PutUint32(buf[offset:], p.M_u32)
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_s64))
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		binary.LittleEndian.PutUint64(buf[offset:], uint64(p.M_s64))
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_u64))
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		binary.LittleEndian.PutUint64(buf[offset:], p.M_u64)
		offset += size
	}
	{
		strlen := uint16(len(p.M_str))
		size = int(unsafe.Sizeof(strlen))
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		binary.LittleEndian.PutUint16(buf[offset:], strlen)
		offset += size

		size = int(strlen)
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		copy(buf[offset:], p.M_str)
		offset += size
	}
	{
		arrlen := uint16(len(p.M_items))
		size = int(unsafe.Sizeof(arrlen))
		if (size + offset) > bufcap {
			return 0, BOError{size, offset, bufcap}
		}
		binary.LittleEndian.PutUint16(buf[offset:], arrlen)
		offset += size
		for _, v := range p.M_items {
			if err := v.Encode(&offset, buf); err != nil {
				return 0, err
			}
		}
	}

	return offset, nil
}

func (p *TestMsg) Decode(buf []byte) error {
	var offset int = 0
	var bufcap int = cap(buf)
	var size int = 0

	{
		size = int(unsafe.Sizeof(p.M_s8))
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		p.M_s8 = int8(buf[offset])
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_u8))
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		p.M_u8 = uint8(buf[offset])
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_s16))
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		p.M_s16 = int16(binary.LittleEndian.Uint16(buf[offset:]))
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_u16))
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		p.M_u16 = binary.LittleEndian.Uint16(buf[offset:])
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_s32))
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		p.M_s32 = int32(binary.LittleEndian.Uint32(buf[offset:]))
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_u32))
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		p.M_u32 = binary.LittleEndian.Uint32(buf[offset:])
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_s64))
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		p.M_s64 = int64(binary.LittleEndian.Uint64(buf[offset:]))
		offset += size
	}
	{
		size = int(unsafe.Sizeof(p.M_u64))
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		p.M_u64 = binary.LittleEndian.Uint64(buf[offset:])
		offset += size
	}
	{
		strlen := binary.LittleEndian.Uint16(buf[offset:])
		size = int(unsafe.Sizeof(strlen))
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		offset += size

		size = int(strlen)
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		p.M_str = string(buf[offset : offset+int(strlen)])
		offset += size
	}
	{
		size = int(unsafe.Sizeof(uint16(0)))
		if (size + offset) > bufcap {
			return BOError{size, offset, bufcap}
		}
		arrlen := binary.LittleEndian.Uint16(buf[offset:])
		offset += size
		for i := uint16(0); i < arrlen; i++ {
			v := Item{}
			if err := v.Decode(&offset, buf); err != nil {
				return err
			}
			p.M_items = append(p.M_items, v)
		}
	}

	return nil
}
