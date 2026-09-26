package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tool_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List installed tools",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tool_listCmd).Standalone()

	tool_listCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	tool_listCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	tool_listCmd.Flags().Bool("no-outdated", false, "")
	tool_listCmd.Flags().Bool("no-python-downloads", false, "")
	tool_listCmd.Flags().Bool("outdated", false, "List outdated tools")
	tool_listCmd.Flags().String("python-preference", "", "")
	tool_listCmd.Flags().Bool("show-extras", false, "Whether to display the extra requirements installed with each tool")
	tool_listCmd.Flags().Bool("show-paths", false, "Whether to display the path to each tool environment and installed executable")
	tool_listCmd.Flags().Bool("show-python", false, "Whether to display the Python version associated with each tool")
	tool_listCmd.Flags().Bool("show-version-specifiers", false, "Whether to display the version specifier(s) used to install each tool")
	tool_listCmd.Flags().Bool("show-with", false, "Whether to display the additional requirements installed with each tool")
	tool_listCmd.Flag("no-outdated").Hidden = true
	tool_listCmd.Flag("no-python-downloads").Hidden = true
	tool_listCmd.Flag("python-preference").Hidden = true
	toolCmd.AddCommand(tool_listCmd)
}
