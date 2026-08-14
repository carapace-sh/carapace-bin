package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zellij",
	Short: "A terminal workspace with batteries included",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config", "c", "", "Change where zellij looks for the configuration file")
	rootCmd.Flags().String("config-dir", "", "Change where zellij looks for the configuration directory")
	rootCmd.Flags().String("data-dir", "", "Change where zellij looks for plugins")
	rootCmd.Flags().BoolP("debug", "d", false, "Specify emitting additional debug information")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().StringP("layout", "l", "", "Name of a predefined layout inside the layout directory or the path to a layout file if inside a session (or using the --session flag) will be added to the session as a new tab or tabs, otherwise will start a new session")
	rootCmd.Flags().String("layout-string", "", "Raw KDL layout string to use directly (instead of a file path) if inside a session (or using the --session flag) will be added to the session as a new tab or tabs, otherwise will start a new session")
	rootCmd.Flags().String("max-panes", "", "Maximum panes on screen, caution: opening more panes will close old ones")
	rootCmd.Flags().StringP("new-session-with-layout", "n", "", "Name of a predefined layout inside the layout directory or the path to a layout file Will always start a new session, even if inside an existing session")
	rootCmd.Flags().String("server", "", "Run server listening at the specified socket path")
	rootCmd.Flags().StringP("session", "s", "", "Specify name of a new session")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flag("server").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":                  carapace.ActionFiles(),
		"config-dir":              carapace.ActionFiles(),
		"data-dir":                carapace.ActionFiles(),
		"layout":                  carapace.ActionFiles(),
		"new-session-with-layout": carapace.ActionFiles(),
		"server":                  carapace.ActionFiles(),
	})
}
