package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/carapace-sh/carapace"
	shlex "github.com/carapace-sh/carapace-shlex"
	"github.com/carapace-sh/carapace/pkg/execlog"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "--edit FILE...",
	Short: "edit files",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := traverse.UserConfigDir(carapace.NewContext())
		if err != nil {
			return err
		}
		dir = filepath.Join(dir, "carapace")

		e, err := editor()
		if err != nil {
			return err
		}

		for _, arg := range args {
			switch filepath.Dir(filepath.Clean(arg)) {
			case "overlays", "specs":
				if filepath.Ext(arg) != ".yaml" {
					return errors.New("specs should have `.yaml` extension")
				}
			default: // not a spec directory
				continue
			}

			f := filepath.Join(dir, arg)
			if _, err := os.Stat(f); !os.IsNotExist(err) {
				continue // already exists
			}
			if err := os.WriteFile(f, []byte("# yaml-language-server: $schema=https://carapace.sh/schemas/command.json\n"), 0644); err != nil {
				return err
			}
			t := time.Now()
			defer func(f string, t time.Time) {
				if info, err := os.Stat(f); err == nil && !info.ModTime().After(t) {
					os.Remove(f) //  not modified
				}
			}(f, t)
		}

		args = append(e, args...)
		command := execlog.Command(args[0], args[1:]...)
		command.Stdout = cmd.OutOrStdout()
		command.Stderr = cmd.ErrOrStderr()
		command.Stdin = cmd.InOrStdin()
		command.Dir = dir
		return command.Run()
	},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	carapace.Gen(editCmd).PositionalAnyCompletion(
		carapace.ActionFiles().
			Chdir("carapace").
			ChdirF(traverse.UserConfigDir).
			FilterArgs(),
	)
}

func editor() ([]string, error) {
	var e string
	for _, env := range []string{
		"CARAPACE_EDITOR",
		"GIT_EDITOR",
		"VISUAL",
		"EDITOR",
	} {
		if e == "" {
			e = os.Getenv(env)
		}
	}
	if e == "" && runtime.GOOS == "windows" {
		e = "notepad"
	}

	tokens, err := shlex.Split(e)
	if err != nil {
		return nil, err
	}
	words := tokens.Words().Strings()
	if len(words) == 0 || words[0] == "" {
		return nil, errors.New("unable to determine editor")
	}
	return words, nil
}
