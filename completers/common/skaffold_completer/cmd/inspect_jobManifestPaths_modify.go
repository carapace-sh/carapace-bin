package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_jobManifestPaths_modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Print the list of jobManifestPaths that would be run for a given configuration (default skaffold configuration, specific module, specific profile, etc).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_jobManifestPaths_modifyCmd).Standalone()

	inspect_jobManifestPaths_modifyCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	inspect_jobManifestPaths_modifyCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	inspect_jobManifestPaths_modifyCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	inspect_jobManifestPaths_modifyCmd.Flags().String("output", "", "File to write `inspect jobManifestPath modify` result")
	inspect_jobManifestPaths_modifyCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	inspect_jobManifestPaths_modifyCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	inspect_jobManifestPaths_modifyCmd.Flags().StringP("user", "u", "", "")
	inspect_jobManifestPaths_modifyCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(inspect_jobManifestPaths_modifyCmd)
	inspect_jobManifestPathsCmd.AddCommand(inspect_jobManifestPaths_modifyCmd)
}
