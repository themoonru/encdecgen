//CODE GENERATED AUTOMATICALLY
//THIS FILE SHOULD NOT BE EDITED BY HAND
package test

import (
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func (s *Struct01) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(1); err != nil {
		return fmt.Errorf("error encode struct Struct01 to MsgPack: %w", err)
	}
	
	if err := e.EncodeInt(int(s.count)); err != nil { // 1
		return fmt.Errorf("error encode field s.count to MsgPack: %w", err)
	}
	
	return nil
}

func (s *Struct01) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 1 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.DecodeInt(); err != nil { // 1
		return err
	} else {
		s.count = v
	}
	
	return nil
}
