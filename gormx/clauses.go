package gormx

import (
	"bytes"
	"encoding/hex"
	"reflect"
)

const (
	tagKeySelect = "gormx_select"
)

func structType(v any) reflect.Type {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Struct:
		return t
	case reflect.Ptr:
		return t.Elem()
	}
	return nil
}

func BytesToSql(bs []byte) string {
	buf := bytes.NewBufferString("0x")
	bsh := make([]byte, hex.EncodedLen(len(bs)))
	hex.Encode(bsh, bs)
	buf.Write(bsh)
	return buf.String()
}
