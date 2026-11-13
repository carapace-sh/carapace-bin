package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generatePipelineCmd = &cobra.Command{
	Use:    "generate-pipeline",
	Short:  "[ALPHA] Generate tekton pipeline from skaffold.yaml",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generatePipelineCmd).Standalone()

	generatePipelineCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	generatePipelineCmd.Flags().StringSlice("config-files", nil, "Select additional files whose artifacts to use when generating pipeline.")
	generatePipelineCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	generatePipelineCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	generatePipelineCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	generatePipelineCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	generatePipelineCmd.Flags().StringP("user", "u", "", "")
	generatePipelineCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(generatePipelineCmd)
	rootCmd.AddCommand(generatePipelineCmd)
}
