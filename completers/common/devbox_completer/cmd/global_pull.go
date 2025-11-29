package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_pullCmd = &cobra.Command{
	Use:   "pull <file> | <url>",
	Short: "Pull a config from a file or URL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_pullCmd).Standalone()

	global_pullCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	global_pullCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	global_pullCmd.Flags().BoolP("force", "f", false, "Force overwrite of existing [global] config files")
	global_pullCmd.Flag("config").Hidden = true
	globalCmd.AddCommand(global_pullCmd)

	// TODO environment
	carapace.Gen(global_pullCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
