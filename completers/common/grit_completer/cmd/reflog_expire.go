package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reflog_expireCmd = &cobra.Command{
	Use:   "expire",
	Short: "Drop reflog entries older than configured or explicit cutoffs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reflog_expireCmd).Standalone()

	reflog_expireCmd.Flags().Bool("all", false, "Process every ref under `refs/`, not only refs that already have log files")
	reflog_expireCmd.Flags().String("expire", "", "Expire reachable entries older than this (Git `--expire` syntax)")
	reflog_expireCmd.Flags().String("expire-unreachable", "", "Expire unreachable entries older than this")
	reflog_expireCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	reflogCmd.AddCommand(reflog_expireCmd)
}
