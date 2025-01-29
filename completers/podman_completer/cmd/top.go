package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top [options] CONTAINER [FORMAT-DESCRIPTORS|ARGS...]",
	Short: "Display the running processes of a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(topCmd).Standalone()

	topCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	topCmd.Flags().Bool("list-descriptors", false, "")
	topCmd.Flag("list-descriptors").Hidden = true
	rootCmd.AddCommand(topCmd)
}
