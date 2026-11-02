package man

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

// TestCmdYaml ensures that every command man page is valid YAML.
//
// These files are not only read by carapace itself but also by external
// tooling (carapace-man, editor language servers, downstream catalogs),
// which fails on the first file that does not parse. An unquoted plain
// scalar containing ": " is the most common mistake and results in
// `yaml: line N: mapping values are not allowed in this context`.
func TestCmdYaml(t *testing.T) {
	err := filepath.WalkDir("cmd", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".yaml" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if err := validateCmdYAML(content); err != nil {
			t.Errorf("%s: %v", path, err)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func validateCmdYAML(content []byte) error {
	var node yaml.Node
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	if err := decoder.Decode(&node); err != nil {
		return err
	}
	if node.Kind != yaml.DocumentNode || len(node.Content) != 1 || node.Content[0].Kind != yaml.MappingNode {
		return fmt.Errorf("expected a YAML mapping")
	}
	if err := decoder.Decode(&node); !errors.Is(err, io.EOF) {
		if err != nil {
			return fmt.Errorf("after the command document: %w", err)
		}
		return fmt.Errorf("expected exactly one YAML document")
	}
	return nil
}

func TestValidateCmdYAMLRejectsTrailingDocuments(t *testing.T) {
	for _, content := range []string{
		"name: tool\n---\nname: other\n",
		"name: tool\n---\nflags: [unterminated\n",
		"name: tool\n---\n",
	} {
		if err := validateCmdYAML([]byte(content)); err == nil {
			t.Errorf("accepted trailing YAML document: %q", content)
		}
	}
}
