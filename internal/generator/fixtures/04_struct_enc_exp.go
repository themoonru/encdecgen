//CODE GENERATED AUTOMATICALLY
//THIS FILE SHOULD NOT BE EDITED BY HAND
package test

import (
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func (m *MapCustomIntStruct) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(5); err != nil {
		return fmt.Errorf("error encode struct MapCustomIntStruct to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapInt); err != nil { // 1
		return fmt.Errorf("error encode field m.MapInt to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapInt8); err != nil { // 2
		return fmt.Errorf("error encode field m.MapInt8 to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapInt16); err != nil { // 3
		return fmt.Errorf("error encode field m.MapInt16 to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapInt32); err != nil { // 4
		return fmt.Errorf("error encode field m.MapInt32 to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapInt64); err != nil { // 5
		return fmt.Errorf("error encode field m.MapInt64 to MsgPack: %w", err)
	}
	
	return nil
}

func (m *MapCustomIntStruct) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 5 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.Decode(); err != nil { // 1
		return err
	} else {
		m.MapInt = v
	}
	
	if v, err := d.Decode(); err != nil { // 2
		return err
	} else {
		m.MapInt8 = v
	}
	
	if v, err := d.Decode(); err != nil { // 3
		return err
	} else {
		m.MapInt16 = v
	}
	
	if v, err := d.Decode(); err != nil { // 4
		return err
	} else {
		m.MapInt32 = v
	}
	
	if v, err := d.Decode(); err != nil { // 5
		return err
	} else {
		m.MapInt64 = v
	}
	
	return nil
}

func (m *MapCustomUIntStruct) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(5); err != nil {
		return fmt.Errorf("error encode struct MapCustomUIntStruct to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapUInt); err != nil { // 1
		return fmt.Errorf("error encode field m.MapUInt to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapUInt8); err != nil { // 2
		return fmt.Errorf("error encode field m.MapUInt8 to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapUInt16); err != nil { // 3
		return fmt.Errorf("error encode field m.MapUInt16 to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapUInt32); err != nil { // 4
		return fmt.Errorf("error encode field m.MapUInt32 to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapUInt64); err != nil { // 5
		return fmt.Errorf("error encode field m.MapUInt64 to MsgPack: %w", err)
	}
	
	return nil
}

func (m *MapCustomUIntStruct) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 5 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.Decode(); err != nil { // 1
		return err
	} else {
		m.MapUInt = v
	}
	
	if v, err := d.Decode(); err != nil { // 2
		return err
	} else {
		m.MapUInt8 = v
	}
	
	if v, err := d.Decode(); err != nil { // 3
		return err
	} else {
		m.MapUInt16 = v
	}
	
	if v, err := d.Decode(); err != nil { // 4
		return err
	} else {
		m.MapUInt32 = v
	}
	
	if v, err := d.Decode(); err != nil { // 5
		return err
	} else {
		m.MapUInt64 = v
	}
	
	return nil
}

func (m *MapCustomFloatStruct) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(2); err != nil {
		return fmt.Errorf("error encode struct MapCustomFloatStruct to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapFloat32); err != nil { // 1
		return fmt.Errorf("error encode field m.MapFloat32 to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapFloat64); err != nil { // 2
		return fmt.Errorf("error encode field m.MapFloat64 to MsgPack: %w", err)
	}
	
	return nil
}

func (m *MapCustomFloatStruct) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 2 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.Decode(); err != nil { // 1
		return err
	} else {
		m.MapFloat32 = v
	}
	
	if v, err := d.Decode(); err != nil { // 2
		return err
	} else {
		m.MapFloat64 = v
	}
	
	return nil
}

func (m *MapCustomComplexStruct) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(2); err != nil {
		return fmt.Errorf("error encode struct MapCustomComplexStruct to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapComplex64); err != nil { // 1
		return fmt.Errorf("error encode field m.MapComplex64 to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapComplex128); err != nil { // 2
		return fmt.Errorf("error encode field m.MapComplex128 to MsgPack: %w", err)
	}
	
	return nil
}

func (m *MapCustomComplexStruct) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 2 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.Decode(); err != nil { // 1
		return err
	} else {
		m.MapComplex64 = v
	}
	
	if v, err := d.Decode(); err != nil { // 2
		return err
	} else {
		m.MapComplex128 = v
	}
	
	return nil
}

func (m *MapCustomOtherStruct) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(4); err != nil {
		return fmt.Errorf("error encode struct MapCustomOtherStruct to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapByte); err != nil { // 1
		return fmt.Errorf("error encode field m.MapByte to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapRune); err != nil { // 2
		return fmt.Errorf("error encode field m.MapRune to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapString); err != nil { // 3
		return fmt.Errorf("error encode field m.MapString to MsgPack: %w", err)
	}
	
	if err := e.Encode(m.MapBool); err != nil { // 4
		return fmt.Errorf("error encode field m.MapBool to MsgPack: %w", err)
	}
	
	return nil
}

func (m *MapCustomOtherStruct) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 4 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.Decode(); err != nil { // 1
		return err
	} else {
		m.MapByte = v
	}
	
	if v, err := d.Decode(); err != nil { // 2
		return err
	} else {
		m.MapRune = v
	}
	
	if v, err := d.Decode(); err != nil { // 3
		return err
	} else {
		m.MapString = v
	}
	
	if v, err := d.Decode(); err != nil { // 4
		return err
	} else {
		m.MapBool = v
	}
	
	return nil
}
