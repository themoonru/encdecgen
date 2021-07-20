//CODE GENERATED AUTOMATICALLY
//THIS FILE SHOULD NOT BE EDITED BY HAND
package main

import (
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func (u *User) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(3); err != nil {
		return err
	}
	
	if err := e.EncodeString(u.Name); err != nil { // 1
		return err
	}
	
	if err := e.EncodeString(u.Password); err != nil { // 2
		return err
	}
	
	if err := e.EncodeInt(u.Id); err != nil { // 3
		return err
	}
	
	return nil
}

func (u *User) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 3 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.DecodeString(); err != nil { // 1
		return err
	} else {
		u.Name = v
	}
	
	if v, err := d.DecodeString(); err != nil { // 2
		return err
	} else {
		u.Password = v
	}
	
	if v, err := d.DecodeInt(); err != nil { // 3
		return err
	} else {
		u.Id = v
	}
	
	return nil
}

func (p *Pswd) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(2); err != nil {
		return err
	}
	
	if err := e.EncodeString(p.Value); err != nil { // 1
		return err
	}
	
	if err := e.EncodeUint32(p.Salt); err != nil { // 2
		return err
	}
	
	return nil
}

func (p *Pswd) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < 2 {
		return fmt.Errorf("incorrect len: %d", length)
	}
	
	if v, err := d.DecodeString(); err != nil { // 1
		return err
	} else {
		p.Value = v
	}
	
	if v, err := d.DecodeUint32(); err != nil { // 2
		return err
	} else {
		p.Salt = v
	}
	
	return nil
}
