package shim

import (
	"os"
	"path/filepath"
	"testing"
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
