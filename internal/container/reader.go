package container

import (
	"encoding/binary"
	"io"
)

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

	if err := readFull(&saltBuf); err != nil {
		return Header{}, err
	}

	if err := readFull(&nonceBuf); err != nil {
		return Header{}, err
	}

	if err := readFull(&nameLen); err != nil {
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
