package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:     "shell",
	Short:   "Set a tool version for the current shell session",
	Aliases: []string{"sh"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()

	shellCmd.Flags().BoolP("help", "h", false, "Print help")
	shellCmd.Flags().StringP("jobs", "j", "", "Number of jobs to run in parallel (default: the `jobs` setting)")
	shellCmd.Flags().Bool("raw", false, "Connect the install commands' stdin, stdout, and stderr to the terminal")
	shellCmd.Flags().BoolP("unset", "u", false, "Remove a version set earlier with `mise shell`")
	rootCmd.AddCommand(shellCmd)

	carapace.Gen(shellCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)
}
