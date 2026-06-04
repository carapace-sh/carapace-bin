package cmd

import (
	"fmt"
	"go/format"
	"os"

	"github.com/carapace-sh/carapace"
	genCompleter "github.com/carapace-sh/carapace-bin/cmd/carapace-generate/pkg/completer"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
	"github.com/carapace-sh/carapace-bridge/pkg/choices"
	"github.com/spf13/cobra"
)

var generateAllCmd = &cobra.Command{
	Use:   "generate-all DIR",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]
		if err := genCompleter.Optimize(dir); err != nil {
			return err
		}

		allCompleters, err := genCompleter.ReadCompleters(dir, "force_all")
		if err != nil {
			return err
		}

		allReleaseCompleters, err := genCompleter.ReadCompleters(dir+"_release", "force_all")
		if err != nil {
			return err
		}

		type target struct {
			tag     string
			goos    string
			release bool
			output  string
		}

		targets := []target{
			{"!release && !force_all &&  android", "android", false, "cmd/completers/completers_termux.go"},
			{"!release && !force_all && !android", "linux", false, "cmd/completers/completers_linux.go"},
			{"!release && !force_all", "darwin", false, "cmd/completers/completers_darwin.go"},
			{"!release && !force_all", "freebsd", false, "cmd/completers/completers_freebsd.go"},
			{"!release && !force_all", "netbsd", false, "cmd/completers/completers_netbsd.go"},
			{"!release && !force_all", "openbsd", false, "cmd/completers/completers_openbsd.go"},
			{"!release && !force_all", "windows", false, "cmd/completers/completers_windows.go"},
			{"!release &&  force_all", "force_all", false, "cmd/completers/completers_all.go"},
			{"release && !force_all &&  android", "android", true, "cmd/completers/completers_release_termux.go"},
			{"release && !force_all && !android", "freebsd", true, "cmd/completers/completers_release_freebsd.go"},
			{"release && !force_all && !android", "linux", true, "cmd/completers/completers_release_linux.go"},
			{"release && !force_all && !android", "netbsd", true, "cmd/completers/completers_release_netbsd.go"},
			{"release && !force_all && !android", "openbsd", true, "cmd/completers/completers_release_openbsd.go"},
			{"release && !force_all", "darwin", true, "cmd/completers/completers_release_darwin.go"},
			{"release && !force_all", "windows", true, "cmd/completers/completers_release_windows.go"},
			{"release &&  force_all", "force_all", true, "cmd/completers/completers_release_all.go"},
		}

		for _, t := range targets {
			var base completer.CompleterMap
			if t.release {
				base = allReleaseCompleters
			} else {
				base = allCompleters
			}

			filtered := filterByGoos(base, t.goos)
			s := filtered.Format("completers", t.tag)
			formatted, err := format.Source([]byte(s))
			if err != nil {
				return fmt.Errorf("formatting %s: %w", t.output, err)
			}
			if err := os.WriteFile(t.output, formatted, 0644); err != nil {
				return err
			}
		}
		return nil
	},
}

func filterByGoos(all completer.CompleterMap, goos string) completer.CompleterMap {
	groups := map[string][]string{
		"android":   {"common", "unix", "linux", "android", "bash"},
		"darwin":    {"common", "unix", "bsd", "darwin", "bash"},
		"freebsd":   {"common", "unix", "bsd", "freebsd", "bash"},
		"linux":     {"common", "unix", "linux", "bash"},
		"netbsd":    {"common", "unix", "bsd", "netbsd", "bash"},
		"openbsd":   {"common", "unix", "bsd", "openbsd", "bash"},
		"windows":   {"common", "windows"},
		"force_all": {"common", "unix", "linux", "bsd", "darwin", "android", "windows", "freebsd", "netbsd", "openbsd", "bash"},
	}

	filtered := make(completer.CompleterMap)
	for _, group := range groups[goos] {
		filtered.Merge(all.Filter(choices.Choice{Group: group}))
	}
	return filtered
}

func init() {
	carapace.Gen(generateAllCmd).Standalone()

	rootCmd.AddCommand(generateAllCmd)

	carapace.Gen(generateAllCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
