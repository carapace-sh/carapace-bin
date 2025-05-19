package common

import "github.com/spf13/cobra"

func ActionInstallFlags(cmd *cobra.Command) {
	// TODO documentation is lacking and descriptions are likely wrong
	cmd.Flags().Bool("auto-remove", false, "automatically remove packages")
	cmd.Flags().Bool("dry-run", false, "perform a simulation of events taken")
	cmd.Flags().Bool("force-yes", false, "continue without prompting if it is changing held packages")
	cmd.Flags().Bool("no-list-columns", false, "display package lists without arranging them in columns")
	cmd.Flags().Bool("no-strict-pinning", false, "consider all versions of a package")
	cmd.Flags().Bool("purge", false, "remove and purge packages")
	cmd.Flags().Bool("reinstall", false, "reinstall package")
	cmd.Flags().Bool("show-progress", false, "show progress")
	cmd.Flags().BoolP("simulate", "s", false, "perform a simulation of events taken")
	cmd.Flags().String("solver", "", "set solver")
	cmd.Flags().Bool("verbose-versions", false, "show full versions")
}
