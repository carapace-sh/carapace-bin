package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_initCmd = &cobra.Command{
	Use:   "init [options] CONTAINER [CONTAINER...]",
	Short: "Initialize one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_initCmd).Standalone()

	container_initCmd.Flags().BoolP("all", "a", false, "Initialize all containers")
	container_initCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	containerCmd.AddCommand(container_initCmd)
}
