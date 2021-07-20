//CODE GENERATED AUTOMATICALLY
//THIS FILE SHOULD NOT BE EDITED BY HAND
package test

import (
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func (s *StructPointer) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(1); err != nil {
		return fmt.Errorf("error encode struct StructPointer to MsgPack: %w", err)
	}
	
	if err := e.Encode(s.Pointer); err != nil { // 1
		return fmt.Errorf("error encode field s.Pointer to MsgPack: %w", err)
	}
	
	return nil
}

func (s *StructPointer) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 1 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.Decode(); err != nil { // 1
		return err
	} else {
		s.Pointer = v
	}
	
	return nil
}

func (s *SlicePointer) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(1); err != nil {
		return fmt.Errorf("error encode struct SlicePointer to MsgPack: %w", err)
	}
	
	if err := e.Encode(s.Slice); err != nil { // 1
		return fmt.Errorf("error encode field s.Slice to MsgPack: %w", err)
	}
	
	return nil
}

func (s *SlicePointer) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 1 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.Decode(); err != nil { // 1
		return err
	} else {
		s.Slice = v
	}
	
	return nil
}

func (s *Slice) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(3); err != nil {
		return fmt.Errorf("error encode struct Slice to MsgPack: %w", err)
	}
	
	if err := e.Encode(s.Int); err != nil { // 1
		return fmt.Errorf("error encode field s.Int to MsgPack: %w", err)
	}
	
	if err := e.Encode(s.String); err != nil { // 2
		return fmt.Errorf("error encode field s.String to MsgPack: %w", err)
	}
	
	if err := e.Encode(s.Float); err != nil { // 3
		return fmt.Errorf("error encode field s.Float to MsgPack: %w", err)
	}
	
	return nil
}

func (s *Slice) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 3 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.Decode(); err != nil { // 1
		return err
	} else {
		s.Int = v
	}
	
	if v, err := d.Decode(); err != nil { // 2
		return err
	} else {
		s.String = v
	}
	
	if v, err := d.Decode(); err != nil { // 3
		return err
	} else {
		s.Float = v
	}
	
	return nil
}

func (s *SliceIntStruct) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(1); err != nil {
		return fmt.Errorf("error encode struct SliceIntStruct to MsgPack: %w", err)
	}
	
	if err := e.Encode(s.Slice); err != nil { // 1
		return fmt.Errorf("error encode field s.Slice to MsgPack: %w", err)
	}
	
	return nil
}

func (s *SliceIntStruct) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 1 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.Decode(); err != nil { // 1
		return err
	} else {
		s.Slice = v
	}
	
	return nil
}

func (s *SliceFloatStruct) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(1); err != nil {
		return fmt.Errorf("error encode struct SliceFloatStruct to MsgPack: %w", err)
	}
	
	if err := e.Encode(s.Slice); err != nil { // 1
		return fmt.Errorf("error encode field s.Slice to MsgPack: %w", err)
	}
	
	return nil
}

func (s *SliceFloatStruct) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 1 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.Decode(); err != nil { // 1
		return err
	} else {
		s.Slice = v
	}
	
	return nil
}
