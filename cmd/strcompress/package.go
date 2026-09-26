package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace-bin/pkg/strblob"
)

const blobFileName = "strblob_generated.go"

type report struct {
	pkg           string
	files         int
	literals      int
	unique        int
	originalBytes int
	blobBytes     int
	skippedConst  int
	skippedSmall  int
}

func (r report) empty() bool { return r.literals == 0 && r.files == 0 }

func (r report) string() string {
	return fmt.Sprintf("%s: %d literals (%d unique, %d bytes) -> blob %d bytes, %d files, skipped %d const / %d small\n",
		r.pkg, r.literals, r.unique, r.originalBytes, r.blobBytes, r.files, r.skippedConst, r.skippedSmall)
}

func (r *report) add(other report) {
	r.files += other.files
	r.literals += other.literals
	r.unique += other.unique
	r.originalBytes += other.originalBytes
	r.blobBytes += other.blobBytes
	r.skippedConst += other.skippedConst
	r.skippedSmall += other.skippedSmall
}

func compressPackage(dir string, threshold int, write bool) (report, error) {
	r := report{pkg: dir}

	files, err := parsePackage(dir)
	if err != nil {
		return r, err
	}

	existingVals, existingVar, err := readExistingBlob(dir)
	if err != nil {
		return r, fmt.Errorf("reading existing blob: %w", err)
	}

	type fileCands struct {
		parsed parsedFile
		c      *collector
	}
	fcs := make([]fileCands, 0, len(files))
	for _, f := range files {
		c := &collector{threshold: threshold}
		c.collect(f.ast)
		fcs = append(fcs, fileCands{parsed: f, c: c})
		r.skippedConst += c.skippedConst
		r.skippedSmall += c.skippedSmall
	}

	candidates := 0
	for _, fc := range fcs {
		candidates += len(fc.c.cands)
	}
	if candidates == 0 && existingVals == nil {
		return report{}, nil
	}

	varName := existingVar
	if varName == "" {
		varName = pickVarName(files)
	}

	vals := []string{}
	if existingVals != nil {
		vals = existingVals
	}
	index := make(map[string]int, len(vals))
	for i, v := range vals {
		if _, exists := index[v]; !exists {
			index[v] = i
		}
	}
	for _, fc := range fcs {
		for _, cand := range fc.c.cands {
			r.literals++
			r.originalBytes += len(cand.value)
			if _, exists := index[cand.value]; !exists {
				index[cand.value] = len(vals)
				vals = append(vals, cand.value)
			}
		}
	}
	r.unique = len(vals)

	type replacement struct {
		file string
		rep  textReplacement
	}
	var replacements []replacement
	for _, fc := range fcs {
		cands := append([]candidate(nil), fc.c.cands...)
		sort.Slice(cands, func(i, j int) bool { return cands[i].start < cands[j].start })
		for _, cand := range cands {
			idx := index[cand.value]
			replacements = append(replacements, replacement{
				file: fc.parsed.path,
				rep: textReplacement{
					start: fc.parsed.tokenFile.Offset(cand.start),
					end:   fc.parsed.tokenFile.Offset(cand.end),
					text:  fmt.Sprintf("%s.Get(%d, %d)", varName, idx, len(cand.value)),
				},
			})
		}
	}

	blob := []byte(nil)
	if len(vals) > 0 {
		blob, err = strblob.Encode(vals)
		if err != nil {
			return r, fmt.Errorf("encoding blob: %w", err)
		}
		r.blobBytes = len(blob)
	}

	if write {
		byFile := make(map[string][]textReplacement)
		for _, r := range replacements {
			byFile[r.file] = append(byFile[r.file], r.rep)
		}
		for path, reps := range byFile {
			var src []byte
			for _, fc := range fcs {
				if fc.parsed.path == path {
					src = fc.parsed.src
					break
				}
			}
			out, err := applyReplacements(src, reps)
			if err != nil {
				return r, fmt.Errorf("%s: %w", path, err)
			}
			if err := writeFile(path, out); err != nil {
				return r, err
			}
			r.files++
		}
		if len(vals) > 0 {
			content, err := generateBlobFile(files[0].ast.Name.Name, varName, blob)
			if err != nil {
				return r, fmt.Errorf("generating %s: %w", blobFileName, err)
			}
			if err := writeFile(filepath.Join(dir, blobFileName), content); err != nil {
				return r, err
			}
		}
	}

	return r, nil
}

type parsedFile struct {
	path      string
	src       []byte
	ast       *ast.File
	tokenFile *token.File
}

func parsePackage(dir string) ([]parsedFile, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	var files []parsedFile
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") || name == blobFileName {
			continue
		}
		path := filepath.Join(dir, name)
		src, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		file, err := parser.ParseFile(fset, path, src, parser.SkipObjectResolution)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		files = append(files, parsedFile{path: path, src: src, ast: file, tokenFile: fset.File(file.Pos())})
	}
	return files, nil
}

// pickVarName returns an identifier that does not collide with any name
// appearing in the package (declarations, imports, locals).
func pickVarName(files []parsedFile) string {
	taken := make(map[string]bool)
	for _, f := range files {
		ast.Inspect(f.ast, func(n ast.Node) bool {
			if ident, ok := n.(*ast.Ident); ok {
				taken[ident.Name] = true
			}
			return true
		})
	}
	for i := 0; ; i++ {
		name := "_strblob"
		if i > 0 {
			name = fmt.Sprintf("_strblob%d", i)
		}
		if !taken[name] {
			return name
		}
	}
}

type textReplacement struct {
	start int
	end   int
	text  string
}

// applyReplacements splices replacements into src. Reps must not overlap.
func applyReplacements(src []byte, reps []textReplacement) ([]byte, error) {
	sorted := append([]textReplacement(nil), reps...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].start < sorted[j].start })
	for i, rep := range sorted {
		if rep.start < 0 || rep.end > len(src) || rep.start >= rep.end {
			return nil, fmt.Errorf("invalid replacement range %d:%d", rep.start, rep.end)
		}
		if i > 0 && sorted[i-1].end > rep.start {
			return nil, fmt.Errorf("overlapping replacements at offset %d", rep.start)
		}
	}
	out := make([]byte, 0, len(src))
	last := 0
	for _, rep := range sorted {
		out = append(out, src[last:rep.start]...)
		out = append(out, rep.text...)
		last = rep.end
	}
	out = append(out, src[last:]...)
	return out, nil
}

func writeFile(path string, content []byte) error {
	mode := os.FileMode(0o644)
	if info, err := os.Stat(path); err == nil {
		mode = info.Mode()
	}
	return os.WriteFile(path, content, mode)
}
