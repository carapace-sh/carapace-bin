package cmd

import (
	"github.com/spf13/cobra"
)

var whatchangedCmd = &cobra.Command{
	Use:   "whatchanged",
	Short: "Show logs with difference each commit introduces",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	whatchangedCmd.Flags().StringS("L", "L", "", "Process line range n,m in file, counting from 1")
	whatchangedCmd.Flags().String("decorate", "", "decorate options")
	whatchangedCmd.Flags().String("decorate-refs", "", "only decorate refs that match <pattern>")
	whatchangedCmd.Flags().String("decorate-refs-exclude", "", "do not decorate refs that match <pattern>")
	whatchangedCmd.Flags().BoolP("quiet", "q", false, "suppress diff output")
	whatchangedCmd.Flags().Bool("source", false, "show source")
	whatchangedCmd.Flags().Bool("use-mailmap", false, "Use mail map file")
	rootCmd.AddCommand(whatchangedCmd)
}
