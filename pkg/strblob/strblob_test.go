package strblob

import (
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
	vals := []string{"x", "y", "z"}
	decoded, err := Decode(mustEncode(t, vals))
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if len(decoded) != len(vals) {
		t.Fatalf("Decode returned %d strings, want %d", len(decoded), len(vals))
	}
	for i, expected := range vals {
		if decoded[i] != expected {
			t.Errorf("Decode()[%d] = %q, want %q", i, decoded[i], expected)
		}
	}
}

func TestDecodeMalformed(t *testing.T) {
	if _, err := Decode("not a deflate stream"); err == nil {
		t.Error("Decode of malformed data should fail")
	}
	if _, err := Decode(""); err == nil {
		t.Error("Decode of empty data should fail")
	}
}
