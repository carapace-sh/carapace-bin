package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bench_webCmd = &cobra.Command{
	Use:    "web",
	Short:  "Run Web benchmark tests",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bench_webCmd).Standalone()

	benchCmd.AddCommand(bench_webCmd)
}
