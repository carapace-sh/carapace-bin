package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec",
	Short:   "Run a command with mise's tools and environment",
	Aliases: []string{"x"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().StringSlice("allow-env", nil, "Pass only VAR through from the environment (implies --deny-env)")
	execCmd.Flags().StringSlice("allow-net", nil, "Allow network access only to HOST (not supported on any platform)")
	execCmd.Flags().StringSlice("allow-read", nil, "Allow reads from PATH (implies --deny-read for everything else)")
	execCmd.Flags().StringSlice("allow-write", nil, "Allow writes to PATH (implies --deny-write for everything else)")
	execCmd.Flags().StringP("command", "c", "", "Command string to run through a shell (supports pipes and redirection)")
	execCmd.Flags().Bool("deny-all", false, "Block reads, writes, network, and env vars")
	execCmd.Flags().Bool("deny-env", false, "Block env var inheritance except PATH, HOME, USER, SHELL, TERM, COLORTERM, LANG")
	execCmd.Flags().Bool("deny-net", false, "Block all network access")
	execCmd.Flags().Bool("deny-read", false, "Block filesystem reads (system libs and tool dirs still accessible)")
	execCmd.Flags().Bool("deny-write", false, "Block filesystem writes except temporary and device paths (/tmp, /dev)")
	execCmd.Flags().Bool("fresh-env", false, "Bypass the environment cache and recompute the environment")
	execCmd.Flags().BoolP("help", "h", false, "Print help")
	execCmd.Flags().StringP("jobs", "j", "", "How many tools to install in parallel; defaults to the `jobs` setting")
	execCmd.Flags().Bool("no-deps", false, "Skip automatic dependency preparation")
	execCmd.Flags().Bool("raw", false, "Connect install commands directly to the terminal; implies --jobs=1")
	execCmd.Flags().StringSlice("secrets", nil, "[experimental] Give the command these secrets (comma-separated); by default it gets none")
	execCmd.Flags().Bool("secrets-all", false, "[experimental] Give the command every secret the project can inject, except file secrets")
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"command": bridge.ActionCarapaceBin().Split(),
		"secrets": mise.ActionSecrets(),
	})

	carapace.Gen(execCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)

	carapace.Gen(execCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
