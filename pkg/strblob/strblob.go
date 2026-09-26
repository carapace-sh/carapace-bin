// Package strblob provides lazy access to a compressed string table
// embedded at build time.
//
// A blob is a deflate-compressed payload holding a lengths table followed
// by the raw string bytes:
//
//	uvarint(n)          number of strings
//	uvarint(n) times    byte length of each string
//	raw bytes           concatenated string data
//
// Blob is intended for generated code: the build-time rewriter replaces
// long string literals with lookups into a per-package blob so the literal
// data is stored compressed and decompressed only when actually accessed.
package strblob

import (
	"bytes"
	"compress/flate"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"sync"
	"unsafe"
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
	decoded, err := io.ReadAll(flate.NewReader(strings.NewReader(data)))
	if err != nil {
		return nil, nil, err
	}
	count, pos := binary.Uvarint(decoded)
	if pos <= 0 {
		return nil, nil, fmt.Errorf("strblob: malformed blob")
	}
	lengths := make([]uint32, count)
	table = make([]uint32, count+1)
	for i := range lengths {
		v, n := binary.Uvarint(decoded[pos:])
		if n <= 0 {
			return nil, nil, fmt.Errorf("strblob: malformed blob")
		}
		lengths[i] = uint32(v)
		pos += n
	}
	if pos > len(decoded) {
		return nil, nil, fmt.Errorf("strblob: malformed blob")
	}
	payload = decoded[pos:]
	offset := uint32(0)
	for i, length := range lengths {
		table[i] = offset
		offset += length
	}
	table[count] = offset
	if int(offset) != len(payload) {
		return nil, nil, fmt.Errorf("strblob: malformed blob")
	}
	return table, payload, nil
}

// Encode packs vals into a blob accepted by Blob.
func Encode(vals []string) ([]byte, error) {
	var header []byte
	header = binary.AppendUvarint(header, uint64(len(vals)))
	for _, val := range vals {
		header = binary.AppendUvarint(header, uint64(len(val)))
	}

	var buf bytes.Buffer
	writer, err := flate.NewWriter(&buf, flate.BestCompression)
	if err != nil {
		return nil, err
	}
	if _, err := writer.Write(header); err != nil {
		return nil, err
	}
	for _, val := range vals {
		if _, err := writer.Write([]byte(val)); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
