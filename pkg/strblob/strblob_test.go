package strblob

import (
	"encoding/binary"
	"strings"
	"testing"
)

func mustEncode(t *testing.T, vals []string) string {
	t.Helper()
	encoded, err := Encode(vals)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	return string(encoded)
}

func TestBlobRoundtrip(t *testing.T) {
	vals := []string{"first", "second value with some length to it", "third"}
	blob := Blob(mustEncode(t, vals))

	for i, expected := range vals {
		if got := blob.Get(i, len(expected)); got != expected {
			t.Errorf("blob.Get(%d) = %q, want %q", i, got, expected)
		}
	}
}

func TestBlobMultiByteUTF8(t *testing.T) {
	vals := []string{"héllo wörld", "日本語のテキスト", "emoji 👋🏽 string", "δκξ"}
	blob := Blob(mustEncode(t, vals))

	for i, expected := range vals {
		if got := blob.Get(i, len(expected)); got != expected {
			t.Errorf("blob.Get(%d) = %q, want %q", i, got, expected)
		}
	}
}

func TestBlobBoundaries(t *testing.T) {
	vals := []string{"a", "", "bc", "", "def"}
	blob := Blob(mustEncode(t, vals))

	// empty strings at boundaries
	if got := blob.Get(1, 0); got != "" {
		t.Errorf("blob.Get(1, 0) = %q, want empty", got)
	}
	if got := blob.Get(3, 0); got != "" {
		t.Errorf("blob.Get(3, 0) = %q, want empty", got)
	}
	if got := blob.Get(2, 2); got != "bc" {
		t.Errorf("blob.Get(2, 2) = %q, want %q", got, "bc")
	}
	if got := blob.Get(4, 3); got != "def" {
		t.Errorf("blob.Get(4, 3) = %q, want %q", got, "def")
	}
}

func TestBlobEmpty(t *testing.T) {
	blob := Blob(mustEncode(t, []string{}))
	if got := blob.Get(0, 0); got != "" {
		t.Errorf("blob.Get(0, 0) = %q, want empty", got)
	}
}

func TestBlobOutOfBounds(t *testing.T) {
	vals := []string{"one", "two"}
	blob := Blob(mustEncode(t, vals))

	cases := []struct{ off, length int }{
		{-1, 3}, {3, 0}, {4, 0}, {100, 1}, {1, -1}, {2, 5},
	}
	for _, c := range cases {
		if got := blob.Get(c.off, c.length); got != "" {
			t.Errorf("blob.Get(%d, %d) = %q, want empty", c.off, c.length, got)
		}
	}
}

func TestBlobCorruptData(t *testing.T) {
	blob := Blob("\x00not a zstd frame")
	if got := blob.Get(0, 3); got != "" {
		t.Errorf("blob.Get(0, 3) = %q, want empty", got)
	}

	// valid zstd frame with truncated header
	encoded, err := Encode([]string{"hello"})
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	blob = Blob(string(encoded[:len(encoded)/2]))
	if got := blob.Get(0, 5); got != "" {
		t.Errorf("blob.Get(0, 5) = %q, want empty", got)
	}
}

func TestBlobLargeValues(t *testing.T) {
	vals := []string{
		strings.Repeat("long-value-", 1000),
		strings.Repeat("ümlaut-", 500),
		"tail",
	}
	blob := Blob(mustEncode(t, vals))

	for i, expected := range vals {
		if got := blob.Get(i, len(expected)); got != expected {
			t.Errorf("blob.Get(%d) mismatch (len %d, want %d)", i, len(got), len(expected))
		}
	}
}

func TestBlobDedupedOffsets(t *testing.T) {
	// offsets table must be strictly monotonic for consecutive entries
	vals := []string{"x", "y", "z"}
	encoded, err := Encode(vals)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded, err := zstdDecoder().DecodeAll([]byte(encoded), nil)
	if err != nil {
		t.Fatalf("DecodeAll failed: %v", err)
	}
	count, pos := binary.Uvarint(decoded)
	if count != uint64(len(vals)) {
		t.Fatalf("count = %d, want %d", count, len(vals))
	}
	previous := uint64(0)
	for i := 0; i <= int(count); i++ {
		offset, n := binary.Uvarint(decoded[pos:])
		if n <= 0 {
			t.Fatalf("failed to read offset %d", i)
		}
		if offset < previous {
			t.Errorf("offset %d = %d, before previous %d", i, offset, previous)
		}
		previous = offset
		pos += n
	}
}
