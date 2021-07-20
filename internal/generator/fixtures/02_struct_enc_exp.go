//CODE GENERATED AUTOMATICALLY
//THIS FILE SHOULD NOT BE EDITED BY HAND
package test

import (
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func (s *StructWithCustomTypes) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(2); err != nil {
		return fmt.Errorf("error encode struct StructWithCustomTypes to MsgPack: %w", err)
	}
	
	if err := e.EncodeInt(int(s.customInt)); err != nil { // 1
		return fmt.Errorf("error encode field s.customInt to MsgPack: %w", err)
	}
	
	if err := e.EncodeString(string(s.customString)); err != nil { // 2
		return fmt.Errorf("error encode field s.customString to MsgPack: %w", err)
	}
	
	return nil
}

func (s *StructWithCustomTypes) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 2 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.DecodeInt(); err != nil { // 1
		return err
	} else {
		s.customInt = v
	}
	
	if v, err := d.DecodeString(); err != nil { // 2
		return err
	} else {
		s.customString = v
	}
	
	return nil
}
