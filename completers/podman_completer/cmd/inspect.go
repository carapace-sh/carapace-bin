package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect [options] {CONTAINER|IMAGE|POD|NETWORK|VOLUME} [...]",
	Short: "Display the configuration of object denoted by ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()

	inspectCmd.Flags().StringP("format", "f", "", "Format the output to a Go template or json")
	inspectCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	inspectCmd.Flags().BoolP("size", "s", false, "Display total file size")
	inspectCmd.Flags().StringP("type", "t", "", "Specify inspect-object type")
	rootCmd.AddCommand(inspectCmd)
}
