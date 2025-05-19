package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeControllerCmd = &cobra.Command{
	Use:   "make:controller [--api] [--type TYPE] [--force] [-i|--invokable] [-m|--model [MODEL]] [-p|--parent [PARENT]] [-r|--resource] [-R|--requests] [-s|--singleton] [--creatable] [--test] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new controller class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeControllerCmd).Standalone()

	makeControllerCmd.Flags().Bool("api", false, "Exclude the create and edit methods from the controller")
	makeControllerCmd.Flags().Bool("creatable", false, "Indicate that a singleton resource should be creatable")
	makeControllerCmd.Flags().Bool("force", false, "Create the class even if the controller already exists")
	makeControllerCmd.Flags().Bool("invokable", false, "Generate a single method, invokable controller class")
	makeControllerCmd.Flags().String("model", "", "Generate a resource controller for the given model")
	makeControllerCmd.Flags().String("parent", "", "Generate a nested resource controller class")
	makeControllerCmd.Flags().Bool("pest", false, "Generate an accompanying Pest test for the Controller")
	makeControllerCmd.Flags().Bool("phpunit", false, "Generate an accompanying PHPUnit test for the Controller")
	makeControllerCmd.Flags().Bool("requests", false, "Generate FormRequest classes for store and update")
	makeControllerCmd.Flags().Bool("resource", false, "Generate a resource controller class")
	makeControllerCmd.Flags().Bool("singleton", false, "Generate a singleton resource controller class")
	makeControllerCmd.Flags().Bool("test", false, "Generate an accompanying Test test for the Controller")
	makeControllerCmd.Flags().String("type", "", "Manually specify the controller stub file to use")
	rootCmd.AddCommand(makeControllerCmd)
}
