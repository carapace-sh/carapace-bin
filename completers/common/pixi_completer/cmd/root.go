package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pixi",
	Short: "Developer Workflow and Environment Management",
	Long:  "https://pixi.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("color", "", "Whether the log needs to be colored")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Display help information")
	rootCmd.Flags().Bool("list", false, "List all installed commands (built-in and extensions)")
	rootCmd.PersistentFlags().Bool("no-progress", false, "Hide all progress bars, always turned on if stderr is not a terminal")
	rootCmd.PersistentFlags().CountP("quiet", "q", "Decrease logging verbosity (quiet mode)")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Increase logging verbosity (-v for warnings, -vv for info, -vvv for debug, -vvvv for trace)")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("always", "never", "auto"),
	})
}
