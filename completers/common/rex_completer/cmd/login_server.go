package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginServerCmd = &cobra.Command{
	Use:   "login-server",
	Short: "Run the login server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginServerCmd).Standalone()

	loginServerCmd.Flags().String("accounts", "", "load accounts from FILE")
	loginServerCmd.Flags().String("extra-listen", "", "also listen on ENDPOINT (repeatable)")
	loginServerCmd.Flags().String("listen", "", "listen on ENDPOINT")
	loginServerCmd.Flags().String("log-format", "text", "write server logs as FORMAT: text or json")
	loginServerCmd.Flags().Bool("pam", true, "authenticate with PAM (--pam=false checks /etc/shadow directly)")
	loginServerCmd.Flags().String("pam-service", "rex-login-server", "use PAM service NAME")
	loginServerCmd.Flags().String("runtime-root", "", "store runtime files below DIR")

	rootCmd.AddCommand(loginServerCmd)

	carapace.Gen(loginServerCmd).FlagCompletion(carapace.ActionMap{
		"accounts":     carapace.ActionFiles(),
		"log-format":   carapace.ActionValues("json", "text"),
		"runtime-root": carapace.ActionDirectories(),
	})
}
