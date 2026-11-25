package ssh

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
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

	actual := ActionHosts(configPath).Invoke(carapace.Context{Env: []string{"HOME=" + dir}})
	b, err := json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	for _, value := range []string{"all-host", "alpha", "from-command", "from-exec", "jump", "prod", "prod-db"} {
		if !strings.Contains(s, `"`+value+`"`) {
			t.Fatalf("expected value %q in %s", value, s)
		}
	}
	if !strings.Contains(s, `"ssh hosts"`) {
		t.Fatalf("expected tag in %s", s)
	}

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

	actual := ActionHosts(configPath).Invoke(carapace.Context{Env: []string{"HOME=" + dir}})
	expected := carapace.ActionValues("explicit-only", "included-only").Tag("ssh hosts").Invoke(carapace.Context{})
	assert.Equal(t, expected, actual)
}

func TestActionHostsDefaultUsesContextHome(t *testing.T) {
	dir := t.TempDir()
	sshDir := filepath.Join(dir, ".ssh")
	if err := os.MkdirAll(sshDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sshDir, "config"), []byte("Host home-host\nInclude conf.d/*.conf\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(sshDir, "conf.d"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sshDir, "conf.d", "extra.conf"), []byte("Host extra-host\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	ctx := carapace.Context{Env: []string{"HOME=" + dir}}
	actual := ActionHosts("").Invoke(ctx)
	b, err := json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	for _, value := range []string{"extra-host", "home-host"} {
		if !strings.Contains(s, `"`+value+`"`) {
			t.Fatalf("expected value %q in %s", value, s)
		}
	}
}

func TestActionHostsMergesKnownHosts(t *testing.T) {
	dir := t.TempDir()
	sshDir := filepath.Join(dir, ".ssh")
	if err := os.MkdirAll(sshDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sshDir, "config"), []byte("Host config-host\nUserKnownHostsFile ~/.ssh/known_hosts.custom\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sshDir, "known_hosts"), []byte("default.example ssh-ed25519 AAAA\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sshDir, "known_hosts.custom"), []byte("custom.example ssh-ed25519 AAAA\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	actual := ActionHosts("").Invoke(carapace.Context{Env: []string{"HOME=" + dir}})
	b, err := json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	for _, value := range []string{"config-host", "custom.example", "default.example"} {
		if !strings.Contains(s, `"`+value+`"`) {
			t.Fatalf("expected value %q in %s", value, s)
		}
	}
}
