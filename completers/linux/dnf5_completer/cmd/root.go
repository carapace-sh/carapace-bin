package cmd

import (
	dnf "github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnf5",
	Short: "DNF5 is a package manager for RPM-based Linux distributions",
	Long:  "https://github.com/rpm-software-management/dnf5",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return dnf.ExecuteWithUse("dnf5")
}
