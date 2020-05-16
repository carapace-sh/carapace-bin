package cmd

import (
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show various types of objects",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	showCmd.Flags().String("decorate", "", "decorate options")
	showCmd.Flags().String("decorate-refs-exclude", "", "do not decorate refs that match <pattern>")
	showCmd.Flags().String("decorate-refs", "", "only decorate refs that match <pattern>")
	showCmd.Flags().StringP("L", "L", "", "Process line range n,m in file, counting from 1")
	showCmd.Flags().BoolP("quiet", "q", false, "suppress diff output")
	showCmd.Flags().Bool("source", false, "show source")
	showCmd.Flags().Bool("use-mailmap", false, "Use mail map file")
	rootCmd.AddCommand(showCmd)
}
