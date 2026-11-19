package ssh

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/assert"
)

func TestActionHosts(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "config")

	content := `# comment
Host alpha *.wild !negated
Match exec "sh -c 'echo marker > '` + filepath.Join(dir, "marker") + `' '"
Host from-exec

Match originalhost jump,*.example host prod user deploy
Host prod-db

Match command "echo hi" localnetwork 10.0.0.0/8
Host from-command

Match all
Host all-host
`
	if err := os.WriteFile(configPath, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	actual := ActionHosts(configPath).Invoke(carapace.Context{})
	expected := carapace.ActionValues(
		"all-host",
		"alpha",
		"from-command",
		"from-exec",
		"jump",
		"prod",
		"prod-db",
	).Tag("ssh hosts").Invoke(carapace.Context{})

	assert.Equal(t, expected, actual)

	if _, err := os.Stat(filepath.Join(dir, "marker")); !os.IsNotExist(err) {
		t.Fatalf("expected Match exec to not run, marker stat err: %v", err)
	}
}

func TestActionHostsExplicitConfigIsolation(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "explicit.conf")
	includedPath := filepath.Join(dir, "included.conf")

	if err := os.WriteFile(includedPath, []byte("Host included-only\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(configPath, []byte("Host explicit-only\nInclude included.conf\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	actual := ActionHosts(configPath).Invoke(carapace.Context{})
	expected := carapace.ActionValues("explicit-only", "included-only").Tag("ssh hosts").Invoke(carapace.Context{})
	assert.Equal(t, expected, actual)
}
