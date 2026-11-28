package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Run a diagnostic on Skaffold",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnoseCmd).Standalone()

	diagnoseCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	diagnoseCmd.Flags().StringP("config", "c", "", "File for global configurations (defaults to $HOME/.skaffold/config)")
	diagnoseCmd.Flags().Bool("enable-templating", false, "Render supported templated fields with golang template engine")
	diagnoseCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	diagnoseCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	diagnoseCmd.Flags().StringP("output", "o", "", "File to write diagnose result")
	diagnoseCmd.Flags().StringSliceP("profile", "p", nil, "Activate profiles by name (prefixed with `-` to disable a profile)")
	diagnoseCmd.Flags().Bool("profile-auto-activation", true, "Set to false to disable profile auto activation")
	diagnoseCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	diagnoseCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	diagnoseCmd.Flags().String("sync-remote-cache", "missing", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	diagnoseCmd.Flags().StringP("user", "u", "", "")
	diagnoseCmd.Flags().Bool("yaml-only", false, "Only prints the effective skaffold.yaml configuration")
	diagnoseCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(diagnoseCmd)
	rootCmd.AddCommand(diagnoseCmd)
}
