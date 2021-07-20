//CODE GENERATED AUTOMATICALLY
//THIS FILE SHOULD NOT BE EDITED BY HAND
package test

import (
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func (o *OtherPackageScalarTypes) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(3); err != nil {
		return fmt.Errorf("error encode struct OtherPackageScalarTypes to MsgPack: %w", err)
	}
	
	if err := e.Encode(o.Age); err != nil { // 1
		return fmt.Errorf("error encode field o.Age to MsgPack: %w", err)
	}
	
	if err := e.Encode(o.Name); err != nil { // 2
		return fmt.Errorf("error encode field o.Name to MsgPack: %w", err)
	}
	
	if err := e.Encode(o.Flag); err != nil { // 3
		return fmt.Errorf("error encode field o.Flag to MsgPack: %w", err)
	}
	
	return nil
}

func (o *OtherPackageScalarTypes) DecodeMsgpack(d *msgpack.Decoder) error {
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
		o.Age = v
	}
	
	if v, err := d.Decode(); err != nil { // 2
		return err
	} else {
		o.Name = v
	}
	
	if v, err := d.Decode(); err != nil { // 3
		return err
	} else {
		o.Flag = v
	}
	
	return nil
}

func (o *OtherPackagePersonStruct) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(1); err != nil {
		return fmt.Errorf("error encode struct OtherPackagePersonStruct to MsgPack: %w", err)
	}
	
	if err := e.Encode(o.Person); err != nil { // 1
		return fmt.Errorf("error encode field o.Person to MsgPack: %w", err)
	}
	
	return nil
}

func (o *OtherPackagePersonStruct) DecodeMsgpack(d *msgpack.Decoder) error {
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
		o.Person = v
	}
	
	return nil
}

func (o *OtherPackageCarStruct) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(1); err != nil {
		return fmt.Errorf("error encode struct OtherPackageCarStruct to MsgPack: %w", err)
	}
	
	if err := e.Encode(o.Auto); err != nil { // 1
		return fmt.Errorf("error encode field o.Auto to MsgPack: %w", err)
	}
	
	return nil
}

func (o *OtherPackageCarStruct) DecodeMsgpack(d *msgpack.Decoder) error {
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
		o.Auto = v
	}
	
	return nil
}
