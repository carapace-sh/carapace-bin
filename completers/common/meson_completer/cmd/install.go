package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().StringS("C", "C", "", "directory to cd into before running")
	installCmd.Flags().String("destdir", "", "Sets or overrides DESTDIR environment.")
	installCmd.Flags().BoolP("dry-run", "n", false, "Doesn't actually install, but print logs.")
	installCmd.Flags().Bool("no-rebuild", false, "Do not rebuild before installing.")
	installCmd.Flags().Bool("only-changed", false, "Only overwrite files that are older than the copied file.")
	installCmd.Flags().Bool("profile-self", false, "==SUPPRESS==")
	installCmd.Flags().BoolP("quiet", "q", false, "Do not print every file that was installed.")
	installCmd.Flags().String("skip-subprojects", "", "Do not install files from given subprojects.")
	installCmd.Flags().Bool("strip", false, "Strip targets even if strip option was not set during configure.")
	installCmd.Flags().String("tags", "", "Install only targets having one of the given tags.")
	applyCommonFlagCompletion(installCmd)
	rootCmd.AddCommand(installCmd)
}
