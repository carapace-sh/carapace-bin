package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var localCmd = &cobra.Command{
	Use:     "local",
	Short:   "Set or show local tool versions (deprecated; use `mise use`)",
	Aliases: []string{"l"},
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(localCmd).Standalone()

	localCmd.Flags().Bool("fuzzy", false, "Save the version as written, such as `24`")
	localCmd.Flags().BoolP("help", "h", false, "Print help")
	localCmd.Flags().BoolP("parent", "p", false, "Use the nearest `mise.toml` in this or a parent directory")
	localCmd.Flags().Bool("path", false, "Print the path of the config file")
	localCmd.Flags().Bool("pin", false, "Save the exact version, such as `24.11.0`")
	localCmd.Flags().StringSlice("remove", nil, "Remove these tools from the local config")
	localCmd.Flags().StringSlice("rm", nil, "Remove these tools from the local config")
	localCmd.Flags().StringSlice("unset", nil, "Remove these tools from the local config")
	localCmd.Flag("rm").Hidden = true
	localCmd.Flag("unset").Hidden = true
	rootCmd.AddCommand(localCmd)

	carapace.Gen(localCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)
}
