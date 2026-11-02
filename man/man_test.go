package man

import (
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

		var node yaml.Node
		if err := yaml.Unmarshal(content, &node); err != nil {
			t.Errorf("%s: %v", path, err)
			return nil
		}
		if node.Kind != yaml.DocumentNode || len(node.Content) != 1 || node.Content[0].Kind != yaml.MappingNode {
			t.Errorf("%s: expected a YAML mapping", path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
