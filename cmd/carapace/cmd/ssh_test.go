package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/sandbox"
)

func TestSSHMatchHosts(t *testing.T) {
	tests := []struct {
		name, match string
		hosts       []string
	}{
		{"plain", "", nil},
		{"exec", `Match exec "echo host fake"`, nil},
		{"host", `Match host desk,server,!excluded,*.example user *,!admin`, []string{"desk", "server"}},
		{"originalhost", `Match originalhost dev,prod user *,!admin`, []string{"dev", "prod"}},
		{"combined", `Match !final user host exec "echo originalhost fake" originalhost box host desk`, []string{"box", "desk"}},
		{"negated", `Match !host excluded !originalhost excluded-too host desk`, []string{"desk"}},
		{"other", `Match localnetwork 10.0.0.0/8 tagged work localuser host command "echo hi" version OpenSSH_* host desk`, []string{"desk"}},
		{"all", "Match all", nil},
		{"canonical", "Match canonical all", nil},
	}
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/cmd/carapace")(func(s *sandbox.Sandbox) {
		for _, test := range tests {
			s.Files(test.name, "Host before\n"+test.match+"\n  User somebody\nHost after\n")
			config := filepath.ToSlash(filepath.Join(s.Dir(), test.name))
			s.Run("--macro", "net.ssh.Hosts("+config+")", "").
				Expect(carapace.ActionValues(append([]string{"after", "before"}, test.hosts...)...).
					Tag("ssh hosts").Usage("--macro [macro] [...args]"))
		}
	})
}

func TestSSHHostIncludes(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/cmd/carapace")(func(s *sandbox.Sandbox) {
		s.Env("HOME", s.Dir())
		s.Env("USERPROFILE", s.Dir())
		s.Files(
			"custom/config", "Include conf.d/* \"~/.ssh/with space\" missing\nInclude = =extra\nHoSt = custom\n",
			".ssh/conf.d/first", "Include nested\nHost first\n",
			".ssh/conf.d/second", "Match exec \"true\"\nHost second\n",
			".ssh/nested", "Include conf.d/*\nHost nested\n",
			".ssh/with space", "hOsT\t\"quoted\" *.example !excluded # ignored\n",
			".ssh/=extra", "Host equals\n",
			".ssh/config", "Host default-host\n",
		)
		s.Run("--macro", "net.ssh.Hosts(custom/config)", "").
			Expect(carapace.ActionValues("custom", "equals", "first", "nested", "quoted", "second").
				Tag("ssh hosts").Usage("--macro [macro] [...args]"))
	})
}

func TestSSHHostsWithMatchExec(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/cmd/carapace")(func(s *sandbox.Sandbox) {
		s.Env("HOME", s.Dir())
		s.Env("USERPROFILE", s.Dir())
		marker := filepath.ToSlash(filepath.Join(s.Dir(), "executed"))
		s.Files(
			".ssh/config", "Include rex.conf\nMatch originalhost desk,server user *,!admin\nHost work\n",
			".ssh/rex.conf", "Match exec \"echo > \\\""+marker+"\\\"\"\n  SetEnv TERM=xterm-256color\n",
		)
		s.Run("--macro", "net.ssh.Hosts", "de").
			Expect(carapace.ActionValues("desk").Tag("ssh hosts").Usage("--macro [macro] [...args]"))
		if _, err := os.Stat(marker); !os.IsNotExist(err) {
			t.Fatalf("Match exec must not run during completion: %v", err)
		}
	})
}
