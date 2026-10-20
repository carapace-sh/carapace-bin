// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secrets_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "[experimental] List secret names from the project's secrets source",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_lsCmd).Standalone()

	secrets_lsCmd.Flags().Bool("complete", false, "Render completions: the injectable key names, one per line")
	secrets_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	secrets_lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	secrets_lsCmd.Flags().Bool("no-header", false, "Do not print the table header")
	secrets_lsCmd.Flag("complete").Hidden = true
	secretsCmd.AddCommand(secrets_lsCmd)
}
