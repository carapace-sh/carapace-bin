package action

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

type repository struct {
	name        string
	description string
}

// repositories returns the named repositories of the given composer.json.
func repositories(p project) []repository {
	raw, ok := p["repositories"]
	if !ok {
		return nil
	}

	describe := func(config any) string {
		switch config := config.(type) {
		case map[string]any:
			repositoryType, _ := config["type"].(string)
			url, _ := config["url"].(string)
			return strings.TrimSpace(repositoryType + " " + url)
		case bool:
			if !config {
				return "disabled"
			}
		}
		return ""
	}

	result := make([]repository, 0)

	var list []map[string]any
	if err := json.Unmarshal(raw, &list); err == nil {
		for _, config := range list {
			if name, ok := config["name"].(string); ok && name != "" {
				result = append(result, repository{name: name, description: describe(config)})
			}
		}
		return result
	}

	var object map[string]any
	if err := json.Unmarshal(raw, &object); err == nil {
		for name, config := range object {
			result = append(result, repository{name: name, description: describe(config)})
		}
	}
	return result
}

// ActionRepositoryTypes completes the repository types
func ActionRepositoryTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"artifact", "directory containing archives",
		"bitbucket", "Bitbucket repository",
		"composer", "composer repository",
		"fossil", "Fossil repository",
		"git", "Git repository",
		"git-bitbucket", "Bitbucket Git repository",
		"github", "GitHub repository",
		"gitlab", "GitLab repository",
		"hg", "Mercurial repository",
		"package", "inline package definition",
		"path", "local directory",
		"perforce", "Perforce repository",
		"svn", "Subversion repository",
		"vcs", "version control system repository",
	).Tag("repository types")
}

// ActionRepositories completes the names of the repositories defined in the composer.json (or config.json)
func ActionRepositories(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path, err := configFile(c, cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		p, err := readProject(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, repository := range repositories(p) {
			vals = append(vals, repository.name, repository.description)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("repositories")
}

// ActionRepositoryNames completes the name argument of `composer repository`
func ActionRepositoryNames(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) == 0 {
			return carapace.ActionValues()
		}

		switch c.Args[0] {
		case "disable", "enable":
			return carapace.ActionValues("packagist.org")
		case "get-url", "remove", "set-url":
			return ActionRepositories(cmd)
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionRepositoryTypeArgument completes the type argument of `composer repository add`
func ActionRepositoryTypeArgument() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) > 0 && c.Args[0] == "add" {
			return ActionRepositoryTypes()
		}
		return carapace.ActionValues()
	})
}

// ActionPolicyNames completes the custom dependency policies defined in the composer.json (or config.json)
func ActionPolicyNames(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) == 0 || c.Args[0] != "add-source" {
			return carapace.ActionValues()
		}

		path, err := configFile(c, cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		p, err := readProject(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		policy, ok := p.config()["policy"].(map[string]any)
		if !ok {
			return carapace.ActionValues()
		}

		vals := make([]string, 0, len(policy))
		for name, value := range policy {
			switch name {
			case "abandoned", "advisories", "ignore-unreachable", "malware":
				continue // builtin policies and settings
			}
			if _, ok := value.(map[string]any); !ok {
				continue
			}
			vals = append(vals, name)
		}
		return carapace.ActionValues(vals...)
	}).Tag("policies")
}

// ActionPolicySourceTypes completes the source type argument of `composer policy add-source`
func ActionPolicySourceTypes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) > 0 && c.Args[0] == "add-source" {
			return carapace.ActionValues("url")
		}
		return carapace.ActionValues()
	})
}
