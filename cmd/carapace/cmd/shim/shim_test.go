package shim

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"sort"
	"strings"
	"sync"
	"testing"

	shlex "github.com/carapace-sh/carapace-shlex/v2"
	"github.com/carapace-sh/carapace/pkg/xdg"
)

func TestIsRunnable(t *testing.T) {
	tests := map[string]struct {
		content string
		want    bool
	}{
		"plain root key":       {content: "name: demo\nrun: echo ok\n", want: true},
		"quoted root key":      {content: "'run':\n  - echo ok\n", want: true},
		"mapping value":        {content: "run:\n  shell: echo ok\n", want: true},
		"literal text":         {content: "description: |\n  run: echo no\n", want: false},
		"folded text":          {content: "description: >\n  run: echo no\n", want: false},
		"nested key":           {content: "command:\n  run: echo no\n", want: false},
		"comment":              {content: "name: demo\n# run: echo no\n", want: false},
		"malformed":            {content: "run: [\n", want: false},
		"empty":                {content: "", want: false},
		"scalar root":          {content: "run echo no\n", want: false},
		"sequence root":        {content: "- run: echo no\n", want: false},
		"missing root run key": {content: "name: demo\n", want: false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "spec.yaml")
			if err := os.WriteFile(path, []byte(test.content), 0o600); err != nil {
				t.Fatal(err)
			}

			candidate := shim{Target: path}
			for attempt := 0; attempt < 2; attempt++ {
				if got := candidate.IsRunnable(); got != test.want {
					t.Fatalf("IsRunnable() = %v, want %v", got, test.want)
				}
			}

			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if string(content) != test.content {
				t.Fatal("IsRunnable modified the spec")
			}
		})
	}

	if (shim{Target: filepath.Join(t.TempDir(), "missing.yaml")}).IsRunnable() {
		t.Fatal("IsRunnable() = true for a missing file")
	}
}

func TestBodyEscaping(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		t.Fatal(err)
	}

	names := []string{
		"cal`touch PROOF`.yaml",
		"d$(touch PROOF)z.yaml",
		"a\" b.yaml",
		`back\slash.yaml`,
		"spa ce.yaml",
		`test$(a)b,c'd.yaml`,
	}

	targets := make([]string, 0, len(names)+1)
	subtests := make([]string, 0, len(names)+1)
	for _, name := range names {
		targets = append(targets, fmt.Sprintf("%v/carapace/specs/%v", configDir, name))
		subtests = append(subtests, name)
	}
	// environment-dependent config dir
	targets = append(targets, "$HOME/`back`/tick$(x)s/pipe|and$var.yaml")
	subtests = append(subtests, "env substituted path")

	for index, target := range targets {
		t.Run(subtests[index], func(t *testing.T) {
			full := strings.TrimPrefix(string(body(target)), "#!/bin/sh\n")
			tokens, err := shlex.Split(full, shlex.Bash)
			if err != nil {
				t.Fatalf("Split(%q): %v", full, err)
			}
			want := []string{"exec", "carapace", "--run", target, "$@"}
			if len(tokens) != len(want) {
				t.Fatalf("token count = %v, want %v: %v", len(tokens), len(want), tokens)
			}
			for index, token := range tokens {
				if token.Value != want[index] {
					t.Fatalf("token %v = %q, want %q", index, token.Value, want[index])
				}
			}
		})
	}
}

func TestBodyStaysRunnable(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	target := fmt.Sprintf("%v/carapace/specs/demo.yaml", configDir)

	if got, want := string(body(target)), fmt.Sprintf("#!/bin/sh\nexec carapace --run \"%v\" \"$@\"", target); got != want {
		t.Fatalf("body() = %q, want %q", got, want)
	}
}

func TestValidSpecName(t *testing.T) {
	tests := map[string]bool{
		"demo.yaml":          true,
		"my.tool-2.yaml":     true,
		"with space.yaml":    true,
		"tick`name`.yaml":    true,
		".yaml":              false,
		"..yaml":             false,
		"-lead.yaml":         false,
		"new\nline.yaml":     false,
		"carriage\ract.yaml": false,
	}

	for name, want := range tests {
		if got := validSpecName(name); got != want {
			t.Errorf("validSpecName(%q) = %v, want %v", name, got, want)
		}
	}
}

func TestUpdateSkipsInvalidSpecNames(t *testing.T) {
	configDir := t.TempDir()
	cacheDir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", configDir)
	t.Setenv("XDG_CACHE_HOME", cacheDir)

	specDir := filepath.Join(configDir, "carapace", "specs")
	if err := os.MkdirAll(specDir, 0o755); err != nil {
		t.Fatal(err)
	}

	runnable := []byte("name: valid\ndescription: test\nrun: \"[echo]\"\n")
	invalid := []string{"..yaml", "-lead.yaml"}
	if runtime.GOOS != "windows" {
		invalid = append(invalid, "new\nline.yaml")
	}
	for _, name := range append([]string{"valid.yaml"}, invalid...) {
		if err := os.WriteFile(filepath.Join(specDir, name), runnable, 0o644); err != nil {
			t.Fatalf("write spec %q: %v", name, err)
		}
	}

	binDir := filepath.Join(configDir, "carapace", "bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"-lead", "stale"} {
		if err := os.WriteFile(filepath.Join(binDir, name), []byte("#!/bin/sh\n"), 0o755); err != nil {
			t.Fatalf("write shim %q: %v", name, err)
		}
	}

	if err := Update(); err != nil {
		t.Fatal(err)
	}

	entries, err := os.ReadDir(binDir)
	if err != nil {
		t.Fatal(err)
	}
	shims := make([]string, 0, len(entries))
	for _, entry := range entries {
		shims = append(shims, entry.Name())
	}
	if want := []string{"valid"}; !slices.Equal(shims, want) {
		t.Fatalf("shims = %v, want %v", shims, want)
	}

	if _, err := os.Stat(filepath.Join(cacheDir, "carapace", ".shims")); !os.IsNotExist(err) {
		t.Fatalf("dummy shims written for invalid spec names: %v", err)
	}
}

func lineScanRunnable(content string) bool {
	if !strings.Contains(content, "run") {
		return false
	}
	runnable := regexp.MustCompile(`^ *run:`)
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		if runnable.MatchString(scanner.Text()) {
			return true
		}
	}
	return false
}

func realSpecs(t testing.TB) []string {
	t.Helper()

	root := filepath.Join("..", "..", "..", "..", "man")
	paths := make([]string, 0)
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return filepath.SkipAll
		}
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".yaml") {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil || len(paths) == 0 {
		t.Skip("no spec corpus available")
	}

	sort.Slice(paths, func(i, j int) bool {
		iInfo, _ := os.Stat(paths[i])
		jInfo, _ := os.Stat(paths[j])
		return iInfo.Size() < jInfo.Size()
	})
	return paths
}

type batchMode int

const (
	batchMixed batchMode = iota
	batchNoRun
	batchAllRun
)

func realSpecsBatch(t testing.TB, count int, mode batchMode) []shim {
	t.Helper()

	paths := realSpecs(t)
	if mode == batchNoRun {
		filtered := make([]string, 0, len(paths))
		for _, path := range paths {
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Contains(content, []byte("run")) {
				filtered = append(filtered, path)
			}
		}
		if len(filtered) < count {
			t.Fatalf("corpus of %v specs without run is too small", len(filtered))
		}
		paths = filtered
	}

	dir := t.TempDir()
	targets := make([]shim, 0, count)

	stride := len(paths) / count
	for index := 0; index < count; index++ {
		content, err := os.ReadFile(paths[index*stride])
		if err != nil {
			t.Fatal(err)
		}
		if mode == batchAllRun {
			content = append(append([]byte{}, content...), []byte("run: [echo]\n")...)
		}
		path := filepath.Join(dir, fmt.Sprintf("%03d.yaml", index))
		if err := os.WriteFile(path, content, 0o600); err != nil {
			t.Fatal(err)
		}
		targets = append(targets, shim{Target: path})
	}
	return targets
}

func parseParallel(targets []shim, runnable func(shim) bool) {
	workers := runtime.NumCPU()

	next := make(chan shim)
	go func() {
		defer close(next)
		for _, target := range targets {
			next <- target
		}
	}()

	var group sync.WaitGroup
	group.Add(workers)
	for range workers {
		go func() {
			defer group.Done()
			for target := range next {
				runnable(target)
			}
		}()
	}
	group.Wait()
}

func BenchmarkIsRunnableReal(b *testing.B) {
	for _, quantile := range []struct {
		name     string
		position float64
	}{
		{"median", 0.5},
		{"p95", 0.95},
		{"largest", 1},
	} {
		paths := realSpecs(b)
		path := paths[int(quantile.position*float64(len(paths)-1))]

		b.Run(fmt.Sprintf("%v/yaml", quantile.name), func(b *testing.B) {
			candidate := shim{Target: path}
			for i := 0; i < b.N; i++ {
				candidate.IsRunnable()
			}
		})
		b.Run(fmt.Sprintf("%v/scan", quantile.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				content, err := os.ReadFile(path)
				if err != nil {
					b.Fatal(err)
				}
				lineScanRunnable(string(content))
			}
		})
	}
}

func BenchmarkSpecBatchParallel(b *testing.B) {
	for _, count := range []int{50, 200} {
		for _, mode := range []struct {
			name string
			mode batchMode
		}{
			{"mixed", batchMixed},
			{"no-run", batchNoRun},
			{"all-run", batchAllRun},
		} {
			targets := realSpecsBatch(b, count, mode.mode)
			if len(targets) < count {
				b.Fatalf("got %v of %v targets", len(targets), count)
			}

			b.Run(fmt.Sprintf("%v/%v/yaml", count, mode.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					parseParallel(targets, shim.IsRunnable)
				}
			})
			b.Run(fmt.Sprintf("%v/%v/scan", count, mode.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					parseParallel(targets, func(s shim) bool {
						content, err := os.ReadFile(s.Target)
						if err != nil {
							return false
						}
						return lineScanRunnable(string(content))
					})
				}
			})
		}
	}
}
