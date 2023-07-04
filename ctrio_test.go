// MIT License Copyright (C) 2023 Hiroshi Shimamoto
package ctrio

import "testing"
import (
	"bytes"
	"io"
)

func TestEncodeDecode(t *testing.T) {
	plain := "Plain Text for Test"
	orgbuf := bytes.NewBufferString(plain)
	key := []byte("0123456789ABCDEF")
	encbuf := new(bytes.Buffer)
	w, err := NewWriter(key, encbuf)
	if err != nil {
		t.Errorf("NewWriter: %v", err)
		return
	}
	io.Copy(w, orgbuf)
	cipher := bytes.NewBuffer(encbuf.Bytes())
	r, err := NewReader(key, cipher)
	if err != nil {
		t.Errorf("NewReader: %v", err)
		return
	}
	decbuf := new(bytes.Buffer)
	io.Copy(decbuf, r)
	decoded := string(decbuf.Bytes())
	if plain != decoded {
		t.Errorf("Decoded text != Plain text")
		return
	}
}
