package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:     "bootstrap",
	Short:   "Set up this machine from your mise config",
	Aliases: []string{"bs"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrapCmd).Standalone()

	bootstrapCmd.Flags().String("adopt", "", "Clone a repository of global mise config or dotfile history, then apply it")
	bootstrapCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	bootstrapCmd.Flags().Bool("force-dotfiles", false, "Overwrite existing files that conflict with whole-file dotfile entries")
	bootstrapCmd.Flags().String("from", "", "Clone a bootstrap project (a Git repository with a mise.toml) and apply it")
	bootstrapCmd.Flags().String("from-dir", "", "Directory for the `--from` checkout (default: `$MISE_DATA_DIR/bootstrap-repo`)")
	bootstrapCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.Flags().StringSlice("only", nil, "Run only these parts")
	bootstrapCmd.Flags().Bool("prompt-secrets", false, "Prompt securely for missing bootstrap secret inputs")
	bootstrapCmd.Flags().Bool("prompt-vars", false, "Ask for `[vars]` entries that declare a `prompt` and have no saved answer")
	bootstrapCmd.Flags().Bool("prune-dotfiles", false, "Remove files that dotfile groups deployed but no active entry deploys now")
	bootstrapCmd.Flags().Bool("replace-history", false, "Discard this machine's dotfile history and adopt the repository's")
	bootstrapCmd.Flags().StringSlice("skip", nil, "Skip these parts")
	bootstrapCmd.Flags().Bool("skip-dirty", false, "Skip configured repositories with local changes instead of failing")
	bootstrapCmd.Flags().Bool("take-remote-all", false, "Take the setup repository's version of every existing file that differs while adopting")
	bootstrapCmd.Flags().Bool("update", false, "Refresh package metadata and update repositories before applying")
	bootstrapCmd.Flags().Bool("upgrade", false, "Upgrade installed packages and tools as well as installing missing ones")
	bootstrapCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts")
	rootCmd.AddCommand(bootstrapCmd)

	carapace.Gen(bootstrapCmd).FlagCompletion(carapace.ActionMap{
		"from-dir": carapace.ActionDirectories(),
		"only":     action.BootstrapParts(),
		"skip":     action.BootstrapParts(),
	})
}
