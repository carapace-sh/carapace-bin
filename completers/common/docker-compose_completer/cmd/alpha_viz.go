package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alpha_vizCmd = &cobra.Command{
	Use:   "viz [OPTIONS]",
	Short: "EXPERIMENTAL - Generate a graphviz graph from your compose file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alpha_vizCmd).Standalone()

	alpha_vizCmd.Flags().Bool("image", false, "Include service's image name in output graph")
	alpha_vizCmd.Flags().String("indentation-size", "", "Number of tabs or spaces to use for indentation")
	alpha_vizCmd.Flags().Bool("networks", false, "Include service's attached networks in output graph")
	alpha_vizCmd.Flags().Bool("ports", false, "Include service's exposed ports in output graph")
	alpha_vizCmd.Flags().Bool("spaces", false, "If given, space character ' ' will be used to indent")
	alphaCmd.AddCommand(alpha_vizCmd)
}
