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
	"fmt"
	"sync"
	"unsafe"

	"github.com/klauspost/compress/zstd"
)

// Table serves substrings of a compressed blob.
type Table struct {
	once    sync.Once
	data    string
	table   []uint32
	payload []byte
}

// Blob returns a table over the strings encoded in data.
//
// data is kept as a string so that package initialization does not copy
// the compressed bytes; decompression happens once, on first access.
func Blob(data string) *Table {
	return &Table{data: data}
}

// Get returns the string at byte offset off with byte length length.
// Malformed blobs and out-of-bounds access yield an empty string.
func (t *Table) Get(off, length int) string {
	t.once.Do(func() {
		t.table, t.payload, _ = decodeBlob(t.data)
	})

	switch {
	case length <= 0:
		return ""
	case off < 0 || off >= len(t.table):
		return ""
	default:
		begin := t.table[off]
		if begin+uint32(length) > uint32(len(t.payload)) {
			return ""
		}
		return unsafe.String(&t.payload[begin], length)
	}
}

// Decode recovers the strings encoded in a blob produced by Encode.
func Decode(data string) ([]string, error) {
	table, payload, err := decodeBlob(data)
	if err != nil {
		return nil, err
	}
	if len(table) == 0 {
		return nil, fmt.Errorf("strblob: malformed blob")
	}
	vals := make([]string, 0, len(table)-1)
	for i := 0; i+1 < len(table); i++ {
		begin, end := table[i], table[i+1]
		if end < begin || int(end) > len(payload) {
			return nil, fmt.Errorf("strblob: malformed blob")
		}
		vals = append(vals, string(payload[begin:end]))
	}
	return vals, nil
}

func decodeBlob(data string) (table []uint32, payload []byte, err error) {
	decoded, err := zstdDecoder().DecodeAll([]byte(data), nil)
	if err != nil {
		return nil, nil, err
	}
	count, pos := binary.Uvarint(decoded)
	if pos <= 0 {
		return nil, nil, fmt.Errorf("strblob: malformed blob")
	}
	table = make([]uint32, count+1)
	for i := range table {
		v, n := binary.Uvarint(decoded[pos:])
		if n <= 0 {
			return nil, nil, fmt.Errorf("strblob: malformed blob")
		}
		table[i] = uint32(v)
		pos += n
	}
	if pos > len(decoded) {
		return nil, nil, fmt.Errorf("strblob: malformed blob")
	}
	return table, decoded[pos:], nil
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
