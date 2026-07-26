package ui

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/dustin/go-humanize"
)

func EncryptReport(in, out string, size int64, elapsed time.Duration) {
	fmt.Printf(`
✓ Encryption successful

┌────────────────────────────────────┐
│ Input   %-25s│
│ Output  %-25s│
│ Size    %-25s│
│ Time    %-25s│
└────────────────────────────────────┘
`,
		filepath.Base(in),
		filepath.Base(out),
		humanize.Bytes(uint64(size)),
		elapsed,
	)
}
