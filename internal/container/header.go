package container

import (
	"encoding/binary"
	"fmt"
	"io"
)

const Magic = "EPST"

type Header struct {
	Filename string
	Version  byte
	Salt     [16]byte
	Nonce    [24]byte
}

func WriteHeader(w io.Writer, h Header) {
	w.Write([]byte(Magic))

	binary.Write(w, binary.LittleEndian, &h.Version)
	binary.Write(w, binary.LittleEndian, &h.Salt)
	binary.Write(w, binary.LittleEndian, &h.Nonce)

	binary.Write(w, binary.LittleEndian, uint16(len(h.Filename)))

	w.Write([]byte(h.Filename))
}
