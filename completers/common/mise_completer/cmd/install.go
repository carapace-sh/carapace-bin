package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install a tool version",
	Aliases: []string{"i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolP("dry-run", "n", false, "Show what would be installed without actually installing")
	installCmd.Flags().Bool("dry-run-code", false, "Like --dry-run but exits with code 1 if there are tools to install")
	installCmd.Flags().BoolP("force", "f", false, "Force reinstall even if already installed")
	installCmd.Flags().Bool("include-lazy", false, "Also install tools configured with lazy = true")
	installCmd.Flags().Bool("include-task-tools", false, "Also install tools required by tasks in current scope")
	installCmd.Flags().StringP("jobs", "j", "", "Number of jobs to run in parallel")
	installCmd.Flags().String("minimum-release-age", "", "Only install versions released before this date")
	installCmd.Flags().Bool("monorepo", false, "Install tools from every config root")
	installCmd.Flags().Bool("raw", false, "Connect backend install command directly to terminal")
	installCmd.Flags().String("shared", "", "Install tool(s) to a shared directory")
	installCmd.Flags().Bool("system", false, "Install tool(s) to the system-wide shared directory")
	installCmd.Flags().BoolP("verbose", "v", false, "Show installation output")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"shared": carapace.ActionDirectories(),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		action.ActionTools(),
	)
}
