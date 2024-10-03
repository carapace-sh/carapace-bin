package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var locateProjectCmd = &cobra.Command{
	Use:   "locate-project",
	Short: "Print a JSON representation of a Cargo.toml file's location",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(locateProjectCmd).Standalone()

	locateProjectCmd.Flags().BoolP("help", "h", false, "Print help")
	locateProjectCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	locateProjectCmd.Flags().String("message-format", "", "Output representation [possible values: json, plain]")
	locateProjectCmd.Flags().Bool("workspace", false, "Locate Cargo.toml of the workspace root")
	rootCmd.AddCommand(locateProjectCmd)

	carapace.Gen(locateProjectCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path":  carapace.ActionFiles(),
		"message-format": carapace.ActionValues("json", "plain"),
	})
}
