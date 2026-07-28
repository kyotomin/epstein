package container

import (
	"encoding/binary"
	"io"
)

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
