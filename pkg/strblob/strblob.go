// Package strblob provides lazy access to a compressed string table
// embedded at build time.
//
// A blob is a zstd-compressed payload holding a offsets table followed by
// the raw string bytes:
//
//	uvarint(n)              number of strings
//	uvarint(n+1) times      offsets into the string data
//	raw bytes               concatenated string data
//
// Blob is intended for generated code: the build-time rewriter replaces
// long string literals with lookups into a per-package blob so the literal
// data is stored compressed and decompressed only when actually accessed.
package strblob

import (
	"encoding/binary"
	"sync"
	"unsafe"

	"github.com/klauspost/compress/zstd"
)

// Blob returns a lookup function over the strings encoded in data.
//
// data is kept as a string so that package initialization does not copy
// the compressed bytes; decompression happens once, on first access.
func Blob(data string) func(off, length int) string {
	var once sync.Once
	var table []uint32
	var payload []byte

	lookup := func(off, length int) string {
		switch {
		case length <= 0:
			return ""
		case off < 0 || off >= len(table):
			return ""
		default:
			begin := table[off]
			if begin+uint32(length) > uint32(len(payload)) {
				return ""
			}
			return unsafe.String(&payload[begin], length)
		}
	}

	decode := func() {
		decoded, err := zstdDecoder().DecodeAll([]byte(data), nil)
		if err != nil {
			return
		}
		count, pos := binary.Uvarint(decoded)
		if pos <= 0 {
			return
		}
		table = make([]uint32, count+1)
		for i := range table {
			v, n := binary.Uvarint(decoded[pos:])
			if n <= 0 {
				table = nil
				return
			}
			table[i] = uint32(v)
			pos += n
		}
		if pos > len(decoded) {
			table = nil
			return
		}
		payload = decoded[pos:]
	}

	return func(off, length int) string {
		once.Do(decode)
		return lookup(off, length)
	}
}

var (
	decoder     *zstd.Decoder
	decoderOnce sync.Once
)

func zstdDecoder() *zstd.Decoder {
	decoderOnce.Do(func() {
		decoder, _ = zstd.NewReader(nil)
	})
	return decoder
}

// Encode packs vals into a blob accepted by Blob.
func Encode(vals []string) ([]byte, error) {
	offsets := make([]uint64, len(vals)+1)
	for i, val := range vals {
		offsets[i+1] = offsets[i] + uint64(len(val))
	}

	header := make([]byte, 0, binary.MaxVarintLen64*(len(offsets)+1))
	header = binary.AppendUvarint(header, uint64(len(vals)))
	for _, offset := range offsets {
		header = binary.AppendUvarint(header, offset)
	}

	payload := make([]byte, 0, len(header)+int(offsets[len(vals)]))
	payload = append(payload, header...)
	for _, val := range vals {
		payload = append(payload, val...)
	}

	encoder, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
	if err != nil {
		return nil, err
	}
	defer encoder.Close()

	return encoder.EncodeAll(payload, nil), nil
}
