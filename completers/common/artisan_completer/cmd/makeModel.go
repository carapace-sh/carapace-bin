package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeModelCmd = &cobra.Command{
	Use:   "make:model [-a|--all] [-c|--controller] [-f|--factory] [--force] [-m|--migration] [--morph-pivot] [--policy] [-s|--seed] [-p|--pivot] [-r|--resource] [--api] [-R|--requests] [--test] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new Eloquent model class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeModelCmd).Standalone()

	makeModelCmd.Flags().Bool("all", false, "Generate a migration, seeder, factory, policy, resource controller, and form request classes for the model")
	makeModelCmd.Flags().Bool("api", false, "Indicates if the generated controller should be an API resource controller")
	makeModelCmd.Flags().Bool("controller", false, "Create a new controller for the model")
	makeModelCmd.Flags().Bool("factory", false, "Create a new factory for the model")
	makeModelCmd.Flags().Bool("force", false, "Create the class even if the model already exists")
	makeModelCmd.Flags().Bool("migration", false, "Create a new migration file for the model")
	makeModelCmd.Flags().Bool("morph-pivot", false, "Indicates if the generated model should be a custom polymorphic intermediate table model")
	makeModelCmd.Flags().Bool("pest", false, "Generate an accompanying Pest test for the Model")
	makeModelCmd.Flags().Bool("phpunit", false, "Generate an accompanying PHPUnit test for the Model")
	makeModelCmd.Flags().Bool("pivot", false, "Indicates if the generated model should be a custom intermediate table model")
	makeModelCmd.Flags().Bool("policy", false, "Create a new policy for the model")
	makeModelCmd.Flags().Bool("requests", false, "Create new form request classes and use them in the resource controller")
	makeModelCmd.Flags().Bool("resource", false, "Indicates if the generated controller should be a resource controller")
	makeModelCmd.Flags().Bool("seed", false, "Create a new seeder for the model")
	makeModelCmd.Flags().Bool("test", false, "Generate an accompanying Test test for the Model")
	rootCmd.AddCommand(makeModelCmd)
}
