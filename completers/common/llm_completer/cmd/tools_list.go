package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tools_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available tools, optionally including tools supported by a model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tools_listCmd).Standalone()

	tools_listCmd.Flags().String("functions", "", "Python code block or file path defining functions to register as tools")
	tools_listCmd.Flags().Bool("json", false, "Output as JSON")
	tools_listCmd.Flags().StringP("model", "m", "", "List tools supported by this model")
	toolsCmd.AddCommand(tools_listCmd)
}
