package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "run0",
	Short: "Elevate privileges interactively",
	Long:  "https://www.man7.org/linux/man-pages/man1/run0.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().String("background", "", "Set ANSI color for background")
	rootCmd.Flags().StringP("chdir", "D", "", "Set working directory")
	rootCmd.Flags().String("description", "", "Description for unit")
	rootCmd.Flags().StringP("group", "g", "", "Run as system group")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().String("machine", "", "Operate on local container")
	rootCmd.Flags().String("nice", "", "Nice level")
	rootCmd.Flags().Bool("no-ask-password", false, "Do not prompt for password")
	rootCmd.Flags().String("property", "", "Set service or scope unit property")
	rootCmd.Flags().String("setenv", "", "Set environment variable")
	rootCmd.Flags().String("slice", "", "Run in the specified slice")
	rootCmd.Flags().Bool("slice-inherit", false, "Inherit the slice")
	rootCmd.Flags().String("unit", "", "Run under the specified unit name")
	rootCmd.Flags().StringP("user", "u", "", "Run as system user")
	rootCmd.Flags().BoolP("version", "V", false, "Show package version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"chdir":  carapace.ActionDirectories(),
		"group":  os.ActionGroups(),
		"setenv": os.ActionEnvironmentVariables().UniqueList(","),
		"user":   os.ActionUsers(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
