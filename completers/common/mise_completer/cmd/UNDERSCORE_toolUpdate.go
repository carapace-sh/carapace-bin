// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var __toolUpdateCmd = &cobra.Command{
	Use:    "__tool-update",
	Short:  "Update globally configured tools for `auto_update`.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(__toolUpdateCmd).Standalone()

	__toolUpdateCmd.Flags().Bool("background", false, "Started detached from a shell prompt: stop the update if it runs longer than a service pass may")
	__toolUpdateCmd.Flags().Bool("due", false, "Update every opted-in tool that is due, once (one `--watch` pass)")
	__toolUpdateCmd.Flags().BoolP("help", "h", false, "Print help")
	__toolUpdateCmd.Flags().String("id", "", "The tool id the launch claimed the check under; the result is recorded there, so `mise doctor` finds it under the same id")
	__toolUpdateCmd.Flags().Bool("watch", false, "Run until stopped, updating due tools once an hour")
	__toolUpdateCmd.Flag("background").Hidden = true
	__toolUpdateCmd.Flag("due").Hidden = true
	__toolUpdateCmd.Flag("id").Hidden = true
	rootCmd.AddCommand(__toolUpdateCmd)
}
