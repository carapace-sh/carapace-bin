package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rails"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Run a code generator",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()

	generateCmd.Flags().BoolP("help", "h", false, "Show help")

	common.AddGeneratorFlags(generateCmd)

	generateCmd.Flags().String("actions", "", "Controller actions (comma-separated)")
	generateCmd.Flags().String("database", "", "Database config name for multi-db")
	generateCmd.Flags().String("db", "", "Alias for --database")
	generateCmd.Flags().Bool("helper", false, "Generate helper")
	generateCmd.Flags().Bool("indexes", true, "Generate indexes for references")
	generateCmd.Flags().Bool("migration", true, "Create migration")
	generateCmd.Flags().String("orm", "", "ORM for the controller")
	generateCmd.Flags().String("parent", "", "Parent class name")
	generateCmd.Flags().String("primary-key-type", "", "Primary key type (e.g. uuid)")
	generateCmd.Flags().String("queue", "default", "Queue name for jobs")
	generateCmd.Flags().Bool("resource-route", true, "Add resource route")
	generateCmd.Flags().Bool("skip-routes", false, "Don't add routes")
	generateCmd.Flags().Bool("timestamps", false, "Add timestamps")

	carapace.Gen(generateCmd).FlagCompletion(carapace.ActionMap{
		"database": action.ActionDatabaseConfigs(),
	})

	carapace.Gen(generateCmd).PositionalCompletion(
		rails.ActionGenerators(),
	)

	carapace.Gen(generateCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			switch c.Args[0] {
			case "model", "scaffold", "resource":
				return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValues().Usage("model name")
					case 1:
						return rails.ActionFieldTypes().Usage("field type")
					case 2:
						return carapace.ActionValues("index", "uniq").Usage("index")
					default:
						return carapace.ActionValues()
					}
				})
			case "controller":
				return carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValues().Usage("controller name")
					default:
						return carapace.ActionValues().Usage("action name")
					}
				})
			case "migration":
				return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValues().Usage("migration name")
					case 1:
						return rails.ActionFieldTypes().Usage("field type")
					case 2:
						return carapace.ActionValues("index", "uniq").Usage("index")
					default:
						return carapace.ActionValues()
					}
				})
			case "job", "mailbox", "helper", "system_test", "integration_test", "jbuilder", "generator", "benchmark":
				return carapace.ActionValues().Usage("name")
			case "mailer":
				return carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValues().Usage("mailer name")
					default:
						return carapace.ActionValues().Usage("action name")
					}
				})
			case "channel":
				return carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValues().Usage("channel name")
					default:
						return carapace.ActionValues().Usage("action name")
					}
				})
			case "task":
				return carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValues().Usage("task name")
					default:
						return carapace.ActionValues().Usage("action name")
					}
				})
			case "scaffold_controller":
				return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValues().Usage("controller name")
					default:
						return rails.ActionFieldTypes().Usage("field type")
					}
				})
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
