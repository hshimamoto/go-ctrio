// MIT License Copyright (C) 2023 Hiroshi Shimamoto
package ctrio

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type Reader struct {
	r cipher.StreamReader
}

func NewReader(key []byte, in io.Reader) (*Reader, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var iv [aes.BlockSize]byte
	if _, err := io.ReadFull(in, iv[:]); err != nil {
		return nil, err
	}
	s := cipher.NewCTR(block, iv[:])
	return &Reader{cipher.StreamReader{S: s, R: in}}, nil
}

func (r *Reader) Read(dst []byte) (int, error) {
	return r.r.Read(dst)
}

type Writer struct {
	w cipher.StreamWriter
}

func NewWriter(key []byte, out io.Writer) (*Writer, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var iv [aes.BlockSize]byte
	if _, err := io.ReadFull(rand.Reader, iv[:]); err != nil {
		return nil, err
	}
	s := cipher.NewCTR(block, iv[:])
	out.Write(iv[:])
	return &Writer{cipher.StreamWriter{S: s, W: out}}, nil
}

func (w *Writer) Close() error {
	return w.w.Close()
}

func (w *Writer) Write(src []byte) (int, error) {
	return w.w.Write(src)
}
