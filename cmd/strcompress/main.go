// strcompress rewrites long string literals in completer packages into
// lookups into a per-package compressed blob (see pkg/strblob).
//
// For every package under the given roots it collects string literals
// (including "a" + "b" concat chains) whose decoded length reaches the
// threshold, deduplicates them, and replaces each occurrence with
// <var>.Get(off, length) where <var> is declared in a generated
// strblob_generated.go holding the zstd-compressed blob.
//
// Literals in constant contexts (const declarations, switch cases, array
// lengths, struct tags, import paths, unsafe builtins) are left as-is.
// Run with -verify afterwards: go build fails loudly on anything the
// skip list missed.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	threshold  = flag.Int("threshold", 100, "minimum decoded length of a string literal to compress")
	writeFiles = flag.Bool("write", true, "write rewritten files and blob (false for a dry run)")
	statsMode  = flag.Bool("stats", false, "print a literal census and exit")
	verifyMode = flag.Bool("verify", false, "run go build on the rewritten tree afterwards")
)

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"completers/common"}
	}

	if *statsMode {
		if err := runStats(roots, *threshold); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	dirs, err := discover(roots)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var total report
	for _, dir := range dirs {
		r, err := compressPackage(dir, *threshold, *writeFiles)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", dir, err)
			os.Exit(1)
		}
		if !r.empty() {
			fmt.Print(r.string())
		}
		total.add(r)
	}
	if !total.empty() {
		total.pkg = "TOTAL"
		fmt.Print(total.string())
	}

	if *verifyMode {
		if err := verifyBuild(roots); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("verify: build ok")
	}
}

// verifyBuild compiles the rewritten tree so any literal rewritten in a
// context that requires a compile-time constant fails loudly.
func verifyBuild(roots []string) error {
	cmd := exec.Command("go", "build", "./...")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("build failed after rewrite:\n%s", out)
	}
	return nil
}

func discover(roots []string) ([]string, error) {
	var dirs []string
	for _, root := range roots {
		info, err := os.Stat(root)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			return nil, fmt.Errorf("not a directory: %s", root)
		}
		err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				return nil
			}
			if hasGoFiles(path) {
				dirs = append(dirs, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return dirs, nil
}

func hasGoFiles(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		name := e.Name()
		if strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go") && name != blobFileName {
			return true
		}
	}
	return false
}
