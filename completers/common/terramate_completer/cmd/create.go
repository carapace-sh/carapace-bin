package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/terramate"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a stack on the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().StringSlice("after", nil, "Add a stack as after")
	createCmd.Flags().Bool("all-terraform", false, "initialize all Terraform directories containing terraform.backend blocks defined")
	createCmd.Flags().StringSlice("before", nil, "Add a stack as before")
	createCmd.Flags().String("description", "", "Description of the stack, defaults to the stack name")
	createCmd.Flags().Bool("ensure-stack-ids", false, "generate an UUID for the stack.id of all stacks which does not define it")
	createCmd.Flags().String("id", "", "ID of the stack, defaults to UUID")
	createCmd.Flags().Bool("ignore-existing", false, "If the stack already exists do nothing and don't fail")
	createCmd.Flags().StringSlice("import", nil, "Add import block for the given path on the stack")
	createCmd.Flags().String("name", "", "Name of the stack, defaults to stack dir base name")
	createCmd.Flags().Bool("no-generate", false, "Disable code generation for the newly created stacks")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"after":  terramate.ActionStacks().UniqueList(","),
		"before": terramate.ActionStacks().UniqueList(","),
	})

	carapace.Gen(createCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
