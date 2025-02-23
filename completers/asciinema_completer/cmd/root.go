package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "asciinema",
	Short: "Record and share your terminal sessions, the right way.",
	Long:  "https://asciinema.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Quiet mode, i.e. suppress diagnostic messages")
	rootCmd.PersistentFlags().String("server-url", "", "asciinema server URL")
	rootCmd.Flags().Bool("version", false, "show program's version number and exit")
}
