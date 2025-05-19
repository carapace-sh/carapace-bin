package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:     "audit",
	Short:   "Audit system for changes",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "miscellaneous",
}

func init() {
	carapace.Gen(auditCmd).Standalone()

	auditCmd.Flags().Bool("backup", false, "Audit configuration files only (default)")
	auditCmd.Flags().Bool("check-permissions", false, "Check file permissions too")
	auditCmd.Flags().Bool("details", false, "Enable reporting of detail records")
	auditCmd.Flags().Bool("full", false, "Audit all system files")
	auditCmd.Flags().Bool("ignore-busybox-symlinks", false, "Ignore symlinks whose target is the busybox binary")
	auditCmd.Flags().Bool("packages", false, "Print only the packages with changed files")
	auditCmd.Flags().String("protected-paths", "", "Use given FILE for protected paths listings")
	auditCmd.Flags().BoolP("recursive", "r", false, "Descend into directories and audit them as well")
	auditCmd.Flags().Bool("system", false, "Audit all system files")
	common.AddGlobalFlags(auditCmd)
	rootCmd.AddCommand(auditCmd)

	carapace.Gen(auditCmd).FlagCompletion(carapace.ActionMap{
		"protected-paths": carapace.ActionFiles(),
	})

	carapace.Gen(auditCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
