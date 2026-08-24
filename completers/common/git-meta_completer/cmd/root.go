package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-meta",
	Short: "Structured metadata for Git data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
}

func ActionTarget() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValuesDescribed(
				"branch", "branch",
				"change-id", "change-id",
				"commit", "commit",
				"path", "path",
				"project", "project",
			).Suffix(":")
		case 1:
			switch c.Parts[0] {
			case "commit":
				return git.ActionRefs(git.RefOption{}.Default())
			case "branch":
				return git.ActionLocalBranches()
			case "path":
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		default:
			return carapace.ActionValues()
		}
	})
}
