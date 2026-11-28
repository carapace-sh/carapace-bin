package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate configuration for deploying an application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().String("XXdefaultBuildpacksBuilder", "gcr.io/buildpacks/builder:v1", "")
	initCmd.Flags().Bool("XXenableBuildpacksInit", true, "")
	initCmd.Flags().Bool("XXenableJibGradleInit", false, "")
	initCmd.Flags().Bool("XXenableJibInit", true, "")
	initCmd.Flags().Bool("XXenableKoInit", false, "")
	initCmd.Flags().Bool("XXenableNewInitFormat", false, "")
	initCmd.Flags().Bool("analyze", false, "Print all discoverable Dockerfiles and images in JSON format to stdout")
	initCmd.Flags().StringSliceP("artifact", "a", nil, "'='-delimited Dockerfile/image pair, or JSON string, to generate build artifact")
	initCmd.Flags().Bool("assume-yes", false, "If true, skaffold will skip yes/no confirmation from the user and default to yes")
	initCmd.Flags().String("compose-file", "", "Initialize from a docker-compose file")
	initCmd.Flags().String("default-kustomization", "", "Default Kustomization overlay path (others will be added as profiles)")
	initCmd.Flags().StringP("filename", "f", "skaffold.yaml", "Path or URL to the Skaffold config file")
	initCmd.Flags().Bool("force", false, "Force the generation of the Skaffold config")
	initCmd.Flags().Bool("generate-manifests", false, "Allows skaffold to try and generate basic kubernetes resources to get your project started")
	initCmd.Flags().StringSliceP("kubernetes-manifest", "k", nil, "A path or a glob pattern to kubernetes manifests (can be non-existent) to be added to the kubectl deployer (overrides detection of kubernetes manifests). Repeat the flag for multiple entries. E.g.: skaffold init -k pod.yaml -k k8s/*.yml")
	initCmd.Flags().StringSliceP("module", "m", nil, "Filter Skaffold configs to only the provided named modules")
	initCmd.Flags().String("remote-cache-dir", "", "Specify the location of the remote cache (default $HOME/.skaffold/remote-cache)")
	initCmd.Flags().Bool("skip-build", false, "Skip generating build artifacts in Skaffold config")
	initCmd.Flags().Bool("skip-deploy", false, "Skip generating deploy stanza in Skaffold config")
	initCmd.Flags().Bool("skip-unreachable-dirs", false, "Instead of erroring, it will skip the directories that cannot be accessed due to permissions")
	initCmd.Flags().String("sync-remote-cache", "always", "Controls how Skaffold manages the remote config cache (see `remote-cache-dir`). One of `always` (default), `missing`, or `never`. `always` syncs remote repositories to latest on access. `missing` only clones remote repositories if they do not exist locally. `never` means the user takes responsibility for updating remote repositories.")
	initCmd.Flags().StringP("user", "u", "", "")
	initCmd.Flag("XXdefaultBuildpacksBuilder").Hidden = true
	initCmd.Flag("XXenableBuildpacksInit").Hidden = true
	initCmd.Flag("XXenableJibGradleInit").Hidden = true
	initCmd.Flag("XXenableJibInit").Hidden = true
	initCmd.Flag("XXenableKoInit").Hidden = true
	initCmd.Flag("XXenableNewInitFormat").Hidden = true
	initCmd.Flag("skip-deploy").Hidden = true
	initCmd.Flag("user").Hidden = true
	applyCommonFlagCompletion(initCmd)
	rootCmd.AddCommand(initCmd)
}
