package just

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

type attributes []any

func (a attributes) Description() string {
	for _, attr := range a {
		switch attr := attr.(type) {
		case map[string]interface{}:
			if doc, ok := attr["doc"]; ok {
				if doc == nil {
					return ""
				}
				return doc.(string)
			}
		}
	}
	return ""
}

type parameter struct {
	Name  string `json:"name"`
	Kind  string `json:"kind"`
	Long  string `json:"long"`
	Short string `json:"short"`
	Help  string `json:"help"`
	Value string `json:"value"`
}

type recipe struct {
	Attributes attributes  `json:"attributes"`
	Parameters []parameter `json:"parameters"`
	Name       string      `json:"name"`
	Doc        string      `json:"doc"`
	Private    bool        `json:"private"`
}

func (r recipe) IsVariadic() bool {
	for _, p := range r.Parameters {
		if p.Kind == "star" {
			return true
		}
	}
	return false
}

func (r recipe) Description() string {
	if doc := r.Doc; doc != "" {
		return doc
	}
	return r.Attributes.Description()
}

type justfile struct {
	Aliases map[string]struct {
		Attributes attributes `json:"attributes"`
		Name       string     `json:"name"`
		Target     string     `json:"target"`
	} `json:"aliases"`
	Recipes map[string]recipe   `json:"recipes"`
	Modules map[string]justfile `json:"modules"`
}

func (jf justfile) AllRecipesWithPrefix(prefix string) []string {
	vals := make([]string, 0, 4)
	for _, recipe := range jf.Recipes {
		if !recipe.Private {
			vals = append(vals, prefix+recipe.Name, recipe.Description())
		}
	}

	for _, alias := range jf.Aliases {
		if recipe, ok := jf.Recipes[alias.Target]; ok {
			vals = append(vals, prefix+alias.Name, recipe.Description())
		}
	}

	for mod, jf := range jf.Modules {
		vals = append(vals, jf.AllRecipesWithPrefix(prefix+mod+"::")...)
	}
	return vals
}

func (jf justfile) AllRecipes() []string {
	return jf.AllRecipesWithPrefix("")
}

// ActionRecipes completes recipes
// default
// build (build project)
func ActionRecipes(path string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--dump", "--dump-format", "json"}
		if path != "" {
			args = append(args, "--justfile", path)
		}

		return carapace.ActionExecCommand("just", args...)(func(output []byte) carapace.Action {
			var jf justfile
			if err := json.Unmarshal(output, &jf); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := jf.AllRecipes()
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("recipes")
}

// ActionRecipesWithArgs completes recipes with their arguments
func ActionRecipesWithArgs(path string) carapace.Action {
	return actionRecipesWithArgs(path)
}

func actionRecipesWithArgs(path string, filter ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) == 0 {
			return ActionRecipes(path).Filter(filter...)
		}

		args := []string{"--json"}
		if path != "" {
			args = append(args, "--justfile", path)
		}
		return carapace.ActionExecCommand("just", args...)(func(output []byte) carapace.Action {
			var jf justfile
			if err := json.Unmarshal(output, &jf); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			recipe, ok := jf.Recipes[c.Args[0]]
			if !ok {
				return carapace.ActionMessage("unknown recipe: " + c.Args[0])
			}

			flags := make(carapace.ActionMap)
			positionals := make([]carapace.Action, 0)

			command := &cobra.Command{Use: recipe.Name}
			command.Flags().SetInterspersed(false)
			carapace.Gen(command).Standalone()
			for _, p := range recipe.Parameters {
				switch {
				case p.Long != "" && p.Short != "":
					switch p.Value {
					case "":
						command.Flags().StringP(p.Long, p.Short, "", p.Help)
					default:
						command.Flags().BoolP(p.Long, p.Short, false, p.Help)
					}
					flags[p.Long] = carapace.ActionFiles()
				case p.Long != "" && p.Short == "":
					switch p.Value {
					case "":
						command.Flags().String(p.Long, "", p.Help)
					default:
						command.Flags().Bool(p.Long, false, p.Help)
					}
					flags[p.Long] = carapace.ActionFiles()
				case p.Long == "" && p.Short != "":
					switch p.Value {
					case "":
						command.Flags().StringS(p.Name, p.Short, "", p.Help)
					default:
						command.Flags().BoolS(p.Name, p.Short, false, p.Help)
					}
					flags[p.Name] = carapace.ActionFiles()
				default:
					positionals = append(positionals, carapace.ActionFiles())
				}
			}

			carapace.Gen(command).FlagCompletion(flags)
			carapace.Gen(command).PositionalCompletion(positionals...)

			switch recipe.IsVariadic() {
			case true:
				carapace.Gen(command).PositionalAnyCompletion(
					carapace.ActionFiles(),
				)
			default:
				carapace.Gen(command).PositionalAnyCompletion(
					actionRecipesWithArgs(path, append(filter, recipe.Name)...).Shift(len(positionals)),
				)
			}
			return carapace.ActionExecute(command).Shift(1)
		})
	})
}
