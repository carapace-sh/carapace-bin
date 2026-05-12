package nix

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/sandbox"
)

func TestActionInstallablesUsesFlakeAttributes(t *testing.T) {
	binDir := t.TempDir()
	nix := filepath.Join(binDir, "nix")
	err := os.WriteFile(nix, []byte(`#!/bin/sh
case "$*" in
  "registry list")
    printf 'global flake:nixpkgs github:NixOS/nixpkgs/nixpkgs-unstable\n'
    ;;
  "build nixpkgs")
    printf 'attrs\nnixpkgs\n'
    ;;
  "build nixpkgs#he")
    printf 'attrs\nnixpkgs#hello\nnixpkgs#helix\n'
    ;;
  *)
    printf 'unexpected nix arguments: %s\n' "$*" >&2
    exit 1
    ;;
esac
`), 0o755)
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))

	sandbox.Action(t, ActionInstallables)(func(s *sandbox.Sandbox) {
		s.Run("").
			Expect(carapace.ActionValuesDescribed(
				"nixpkgs", "github:NixOS/nixpkgs/nixpkgs-unstable",
			).Tag("flakes").Suffix("#").NoSpace('#'))

		s.Run("nixpkgs#he").
			Expect(carapace.ActionValues("helix", "hello").Prefix("nixpkgs#").NoSpace('#'))
	})
}
