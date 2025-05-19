package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [--without-tty] [--compact] [--coverage] [--min [MIN]] [-p|--parallel] [--profile] [--recreate-databases] [--drop-databases] [--without-databases]",
	Short: "Run the application tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().Bool("compact", false, "Indicates whether the compact printer should be used")
	testCmd.Flags().Bool("coverage", false, "Indicates whether code coverage information should be collected")
	testCmd.Flags().Bool("drop-databases", false, "Indicates if the test databases should be dropped")
	testCmd.Flags().String("min", "", "Indicates the minimum threshold enforcement for code coverage")
	testCmd.Flags().Bool("parallel", false, "Indicates if the tests should run in parallel")
	testCmd.Flags().Bool("profile", false, "Lists top 10 slowest tests")
	testCmd.Flags().Bool("recreate-databases", false, "Indicates if the test databases should be re-created")
	testCmd.Flags().Bool("without-databases", false, "Indicates if database configuration should be performed")
	testCmd.Flags().Bool("without-tty", false, "Disable output to TTY")
	rootCmd.AddCommand(testCmd)
}
