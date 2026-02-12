package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scripts_runCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scripts_runCmd).Standalone()

	scripts_runCmd.Flags().BoolP("help", "h", false, "Print help")
	scripts_runCmd.Flags().StringSliceP("var", "v", nil, "Specify template variables in the format KEY=VALUE Example: -v name=John -v greeting=\"Hello there\"")
	scriptsCmd.AddCommand(scripts_runCmd)
}
