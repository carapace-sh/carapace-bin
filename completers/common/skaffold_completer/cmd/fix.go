package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Update old configuration to a newer schema version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fixCmd).Standalone()

	fixCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	fixCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	fixCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	fixCmd.Flags().StringP("output", "o", "", "File to write the changed config (instead of standard output)")
	fixCmd.Flags().Bool("overwrite", false, "Overwrite original config with fixed config")
	fixCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	fixCmd.Flags().String("sync-remote-cache", "missing", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	fixCmd.Flags().StringP("user", "u", "", "")
	fixCmd.Flags().String("version", "skaffold/v4beta14", "Target schema version to upgrade to")
	fixCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(fixCmd)
	rootCmd.AddCommand(fixCmd)
}
