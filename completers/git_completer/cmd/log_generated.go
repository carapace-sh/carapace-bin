package cmd

import (
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit logs",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	logCmd.Flags().String("decorate", "", "decorate options")
	logCmd.Flags().String("decorate-refs-exclude", "", "do not decorate refs that match <pattern>")
	logCmd.Flags().String("decorate-refs", "", "only decorate refs that match <pattern>")
	logCmd.Flags().StringS("L", "L", "", "Process line range n,m in file, counting from 1")
	logCmd.Flags().BoolP("quiet", "q", false, "suppress diff output")
	logCmd.Flags().Bool("source", false, "show source")
	logCmd.Flags().Bool("use-mailmap", false, "Use mail map file")
	rootCmd.AddCommand(logCmd)
}
