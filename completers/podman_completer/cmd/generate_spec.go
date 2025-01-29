package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_specCmd = &cobra.Command{
	Use:   "spec [options] {CONTAINER|POD}",
	Short: "Generate Specgen JSON based on containers or pods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_specCmd).Standalone()

	generate_specCmd.Flags().BoolP("compact", "c", false, "Print the json in a compact format for consumption")
	generate_specCmd.Flags().StringP("filename", "f", "", "Write output to the specified path")
	generate_specCmd.Flags().BoolP("name", "n", false, "Specify a new name for the generated spec")
	generateCmd.AddCommand(generate_specCmd)
}
