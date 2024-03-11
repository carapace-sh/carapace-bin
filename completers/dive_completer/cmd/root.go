package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dive",
	Short: "Docker Image Visualizer & Explorer",
	Long:  "https://github.com/wagoodman/dive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().Bool("ci", false, "Skip the interactive TUI and validate against CI rules (same as env var CI=true)")
	rootCmd.Flags().String("ci-config", ".dive-ci", "If CI=true in the environment, use the given yaml to drive validation rules.")
	rootCmd.PersistentFlags().String("config", "", "config file (default is $HOME/.dive.yaml, ~/.config/dive/*.yaml, or $XDG_CONFIG_HOME/dive.yaml)")
	rootCmd.Flags().BoolP("help", "h", false, "help for dive")
	rootCmd.Flags().String("highestUserWastedPercent", "0.1", "(only valid with --ci given) highest allowable percentage of bytes wasted (as a ratio between 0-1), otherwise CI validation will fail.")
	rootCmd.Flags().String("highestWastedBytes", "disabled", "(only valid with --ci given) highest allowable bytes wasted, otherwise CI validation will fail.")
	rootCmd.PersistentFlags().BoolP("ignore-errors", "i", false, "ignore image parsing errors and run the analysis anyway")
	rootCmd.Flags().StringP("json", "j", "", "Skip the interactive TUI and write the layer analysis statistics to a given file.")
	rootCmd.Flags().String("lowestEfficiency", "0.9", "(only valid with --ci given) lowest allowable image efficiency (as a ratio between 0-1), otherwise CI validation will fail.")
	rootCmd.PersistentFlags().String("source", "docker", "The container engine to fetch the image from. Allowed values: docker, podman, docker-archive")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "display version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ci-config": carapace.ActionFiles(),
		"json":      carapace.ActionFiles(),
		"source":    carapace.ActionValues("docker", "podman", "docker-archive"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
