package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get and set settings and shortcuts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	rootCmd.AddCommand(configCmd)
}

func actionConfigKeys() carapace.Action {
	return carapace.ActionExecCommand("terminal-browser", "config", "list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		for _, line := range strings.Split(string(output), "\n") {
			if key, value, ok := strings.Cut(line, "="); ok {
				vals = append(vals, key, value)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
