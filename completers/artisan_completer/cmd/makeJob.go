package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeJobCmd = &cobra.Command{
	Use:   "make:job [-f|--force] [--sync] [--test] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new job class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeJobCmd).Standalone()

	makeJobCmd.Flags().Bool("force", false, "Create the class even if the job already exists")
	makeJobCmd.Flags().Bool("pest", false, "Generate an accompanying Pest test for the Job")
	makeJobCmd.Flags().Bool("phpunit", false, "Generate an accompanying PHPUnit test for the Job")
	makeJobCmd.Flags().Bool("sync", false, "Indicates that job should be synchronous")
	makeJobCmd.Flags().Bool("test", false, "Generate an accompanying Test test for the Job")
	rootCmd.AddCommand(makeJobCmd)
}
