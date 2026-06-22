package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "List all defined routes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(routesCmd).Standalone()

	routesCmd.Flags().StringP("controller", "c", "", "Filter by controller")
	routesCmd.Flags().BoolP("expanded", "E", false, "Print expanded vertically")
	routesCmd.Flags().StringP("grep", "g", "", "Grep by pattern")
	routesCmd.Flags().BoolP("help", "h", false, "Show help")
	routesCmd.Flags().BoolP("unused", "u", false, "Show unused routes")

	common.AddEnvironmentFlag(routesCmd, "development")
}
