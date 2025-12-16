package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dagger",
	Short: "The Dagger CLI provides a command-line interface to Dagger.",
	Long:  "https://dagger.io/",
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
	rootCmd.AddGroup(
		&cobra.Group{ID: "module", Title: ""},
		&cobra.Group{ID: "cloud", Title: ""},
		&cobra.Group{ID: "exec", Title: ""},
	)

	rootCmd.PersistentFlags().String("api", "", "Dagger Cloud API URL")
	rootCmd.PersistentFlags().Bool("debug", false, "Show more information for debugging")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	rootCmd.PersistentFlags().String("progress", "", "progress output format (auto, plain, tty)")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "disable terminal UI and progress output")
	rootCmd.PersistentFlags().String("workdir", "", "The host workdir loaded into dagger")
	rootCmd.Flag("api").Hidden = true
	rootCmd.Flag("help").Hidden = true
	rootCmd.Flag("workdir").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"progress": carapace.ActionValues("auto", "plain", "tty"),
		"workdir":  carapace.ActionDirectories(),
	})
}
