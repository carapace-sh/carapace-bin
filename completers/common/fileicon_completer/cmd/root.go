package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fileicon",
	Short: "macOS CLI for managing custom icons for files and folders",
	Long:  "https://github.com/mklement0/fileicon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "print usage information")
	rootCmd.Flags().Bool("home", false, "open fileicon's home page")
	rootCmd.Flags().Bool("man", false, "display the manual page")
	rootCmd.Flags().Bool("man-source", false, "print raw, embedded Markdown-formatted man-page source")
	rootCmd.Flags().Bool("version", false, "output version number")
	rootCmd.Flag("man-source").Hidden = true
}
