package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fdesetup",
	Short: "fdesetup utility",
	Long:  "https://keith.github.io/xcode-manpages/fdesetup.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(enableCmd)
	rootCmd.AddCommand(disableCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(changerecoveryCmd)
	rootCmd.AddCommand(removerecoveryCmd)
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(haspersonalrecoverykeyCmd)
	rootCmd.AddCommand(hasinstitutionalrecoverykeyCmd)
	rootCmd.AddCommand(usingrecoverykeyCmd)
	rootCmd.AddCommand(supportsauthrestartCmd)
	rootCmd.AddCommand(authrestartCmd)
	rootCmd.AddCommand(validaterecoveryCmd)
	rootCmd.AddCommand(isactiveCmd)
	rootCmd.AddCommand(showdeferralinfoCmd)
	rootCmd.AddCommand(versionCmd)
}
