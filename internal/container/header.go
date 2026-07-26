package container

import (
	"encoding/binary"
	"io"
)

const Magic = "EPST"

type Header struct {
	Version byte

	Salt  [16]byte
	Nonce [24]byte

	Filename string
}

func WriteHeader(w io.Writer, h Header) error {
	if _, err := w.Write([]byte(Magic)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Salt); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Nonce); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, uint16(len(h.Filename))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(h.Filename)); err != nil {
		return err
	}

	return nil
}

func ReadHeader(r io.Reader) (Header, error) {
	var magicBuf [4]byte
	var version byte
	var saltBuf [16]byte
	var nonceBuf [24]byte
	var nameLen uint16

	readFull := func(data any) error {
		err := binary.Read(r, binary.LittleEndian, data)
		return err
	}

	_, err := io.ReadFull(r, magicBuf[:])
	if err != nil {
		return Header{}, err
	}

	if string(magicBuf[:]) != Magic {
		return Header{}, ErrInvalidMagic
	}

	if err := readFull(&version); err != nil {
		return Header{}, err
	}

	if err := readFull(saltBuf); err != nil {
		return Header{}, err
	}

	if err := readFull(nonceBuf); err != nil {
		return Header{}, err
	}

	if err := readFull(nameLen); err != nil {
		return Header{}, err
	}

	nameBuf := make([]byte, nameLen)
	_, err = io.ReadFull(r, nameBuf)
	if err != nil {
		return Header{}, err
	}

	return Header{
		Version:  version,
		Salt:     saltBuf,
		Nonce:    nonceBuf,
		Filename: string(nameBuf),
	}, nil
}
