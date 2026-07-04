package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createDbCmd = &cobra.Command{
	Use:   "create-db",
	Short: "Create a db using the DL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createDbCmd).Standalone()
	createDbCmd.Flags().BoolP("autocommit-off", "a", false, "Turn off autocommit")
	createDbCmd.Flags().StringP("dl", "g", "", "Use AppleDL (default) or AppleCspDL")
	createDbCmd.Flags().StringP("mode", "m", "", "Set the file permissions to mode")
	createDbCmd.Flags().BoolP("openparams", "o", false, "Force using openparams argument")
	rootCmd.AddCommand(createDbCmd)
}
