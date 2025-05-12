package just

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
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

type recipe struct {
	Attributes attributes `json:"attributes"`
	Name       string     `json:"name"`
	Doc        string     `json:"doc"`
	Private    bool       `json:"private"`
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
