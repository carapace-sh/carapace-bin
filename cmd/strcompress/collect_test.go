package main

import (
	"go/parser"
	"go/token"
	"testing"
)

func collectSource(t *testing.T, src string, threshold int) *collector {
	t.Helper()
	file, err := parser.ParseFile(token.NewFileSet(), "test.go", src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	c := &collector{threshold: threshold}
	c.collect(file)
	return c
}

func TestCollectSelectsLongLiterals(t *testing.T) {
	c := collectSource(t, `package p
var _ = "short"
var _ = "a reasonably long literal that easily exceeds the hundred byte threshold used by the string literal compressor"
`, 100)
	if len(c.cands) != 1 {
		t.Fatalf("got %d candidates, want 1", len(c.cands))
	}
	if c.cands[0].value != "a reasonably long literal that easily exceeds the hundred byte threshold used by the string literal compressor" {
		t.Errorf("unexpected candidate value %q", c.cands[0].value)
	}
}

func TestCollectSkipsConstantContexts(t *testing.T) {
	src := `package p
const A = "a constant string that is definitely longer than one hundred bytes to trip the compressor threshold"
const B = A + " and a concatenated tail that pushes this constant expression well past the threshold too"
var arr [len("an array length that is also a constant context expression and far beyond the threshold")]int
type s struct {
	Field string ` + "`tag:\"a struct tag that is way beyond the one hundred byte threshold on purpose\"`" + `
}
var _ = s{Field: "a runtime value that should be collected as a candidate because it is comfortably long enough to pass the threshold"}
func f() {
	switch "x" {
	case "a switch case value that is a constant context and way longer than the threshold":
	}
}
`
	c := collectSource(t, src, 100)
	if len(c.cands) != 1 {
		t.Fatalf("got %d candidates, want 1 (struct field value)", len(c.cands))
	}
	if c.skippedConst != 5 {
		t.Errorf("skippedConst = %d, want 5", c.skippedConst)
	}
}

func TestCollectFoldsConcatChains(t *testing.T) {
	c := collectSource(t, `package p
var _ = "first part of a very long literal that goes on and on " +
	"second part of a very long literal that keeps going " +
	"third part of a very long literal"
var _ = runtime() + "a literal following a non literal expression that is on its own long enough to become a candidate value"
func runtime() string { return "" }
`, 100)
	if len(c.cands) != 2 {
		t.Fatalf("got %d candidates, want 2", len(c.cands))
	}
	if len(c.cands[0].value) <= len(c.cands[1].value) {
		t.Errorf("expected folded chain first, got %q", c.cands[0].value)
	}
}

func TestApplyReplacements(t *testing.T) {
	src := []byte(`x = "AAAAAAAAAA"; y = "BBBBBBBBBB"`)
	out, err := applyReplacements(src, []textReplacement{
		{start: 4, end: 16, text: "a()"},
		{start: 22, end: 34, text: "b()"},
	})
	if err != nil {
		t.Fatalf("applyReplacements failed: %v", err)
	}
	if string(out) != `x = a(); y = b()` {
		t.Errorf("got %q", out)
	}
}

func TestApplyReplacementsOverlap(t *testing.T) {
	_, err := applyReplacements([]byte(`x = "AAAA"`), []textReplacement{
		{start: 0, end: 5, text: "a"},
		{start: 3, end: 8, text: "b"},
	})
	if err == nil {
		t.Error("expected error for overlapping replacements")
	}
}
