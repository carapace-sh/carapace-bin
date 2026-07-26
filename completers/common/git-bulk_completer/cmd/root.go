package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-bulk",
	Short: "Run git commands on multiple repositories across registered workspaces",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-bulk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func actionBulkWorkspaces() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommandE("git", "config", "--global", "--get-regexp", "^bulkworkspaces\\.")(func(output []byte, err error) carapace.Action {
			if err != nil {
				return carapace.ActionValues()
			}
			vals := make([]string, 0)
			for line := range strings.SplitSeq(string(output), "\n") {
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}
				parts := strings.SplitN(line, " ", 2)
				if len(parts) == 2 && strings.HasPrefix(parts[0], "bulkworkspaces.") {
					name := strings.TrimPrefix(parts[0], "bulkworkspaces.")
					vals = append(vals, name, parts[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...).Tag("workspaces")
		})
	})
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("addcurrent", "", "Register the current directory as a workspace")
	rootCmd.Flags().StringArray("addworkspace", nil, "Register a workspace")
	rootCmd.Flags().BoolP("all", "a", false, "Run command on all registered workspaces")
	rootCmd.Flags().Bool("listall", false, "List all registered workspaces")
	rootCmd.Flags().Bool("no-follow-hidden", false, "Don't follow hidden directories")
	rootCmd.Flags().Bool("no-follow-symlinks", false, "Don't follow symlinks")
	rootCmd.Flags().Bool("purge", false, "Remove all registered workspaces")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	rootCmd.Flags().String("removeworkspace", "", "Remove a registered workspace")
	rootCmd.Flags().StringP("workspace", "w", "", "Run command on a specific workspace")

	rootCmd.MarkFlagsMutuallyExclusive("addworkspace", "addcurrent", "listall", "purge", "removeworkspace")
	rootCmd.MarkFlagsMutuallyExclusive("all", "workspace")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"addcurrent":      carapace.ActionValues(),
		"addworkspace":    carapace.ActionValues(),
		"removeworkspace": actionBulkWorkspaces(),
		"workspace":       actionBulkWorkspaces(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("git"),
	)
}
