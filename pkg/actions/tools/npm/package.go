package npm

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/util"
)

// ActionPackageSearch completes packages@version for given registry
//
//	express
//	lodash
func ActionPackageSearch(registry string) carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionPackageNames(registry).NoSpace()
		case 1:
			return ActionPackageVersions(PackageOpts{Registry: registry, Package: c.Parts[0]})
		default:
			return carapace.ActionValues()
		}
	})
}

// ActionPackageNames completes package names for given registry
//
//	express
//	lodash
func ActionPackageNames(registry string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"search", "--parseable", "--searchlimit", "250", fmt.Sprintf(`/^%v`, c.Value)}
		if registry != "" {
			args = append(args, "--registry", registry)
		}

		return carapace.ActionExecCommand("npm", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				fields := strings.Split(line, "\t")
				vals = append(vals, fields[0], fields[1])
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

type PackageOpts struct {
	Registry string
	Package  string
}

// ActionPackageVersions completes versions for given package
//
//	4.18.0
//	4.17.1
func ActionPackageVersions(opts PackageOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"view", opts.Package, "versions", "--json"}
		if opts.Registry != "" {
			args = append(args, "--registry", opts.Registry)
		}

		return carapace.ActionExecCommand("npm", args...)(func(output []byte) carapace.Action {
			var versions []string
			if err := json.Unmarshal(output, &versions); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionValues(versions...)
		})
	})
}

// ActionPackageTags completes tags for given package
//
//	latest
//	next
func ActionPackageTags(opts PackageOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"view", opts.Package, "dist-tags", "--json"}
		if opts.Registry != "" {
			args = append(args, "--registry", opts.Registry)
		}

		return carapace.ActionExecCommand("npm", args...)(func(output []byte) carapace.Action {
			var tags map[string]string
			if err := json.Unmarshal(output, &tags); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0, len(tags)*2)
			for tag, version := range tags {
				vals = append(vals, tag, version)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

type packageJson struct {
	Scripts    map[string]string
	Workspaces []string
}

// ActionPackumentFields completes field names for a published package
//
//	name
//	version
//	description
func ActionPackumentFields(opts PackageOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"view", opts.Package, "--json"}
		if opts.Registry != "" {
			args = append(args, "--registry", opts.Registry)
		}

		return carapace.ActionExecCommand("npm", args...)(func(output []byte) carapace.Action {
			var pkg map[string]any
			if err := json.Unmarshal(output, &pkg); err == nil {
				fields := getCompletionFields(pkg, nil)
				return carapace.ActionValues(fields...)
			}

			var pkgs []map[string]any
			if err := json.Unmarshal(output, &pkgs); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			merged := make(map[string]any)
			for _, p := range pkgs {
				for k, v := range p {
					merged[k] = v
				}
			}
			fields := getCompletionFields(merged, nil)
			return carapace.ActionValues(fields...)
		})
	})
}

func getCompletionFields(d map[string]any, pref []string) []string {
	var fields []string
	for key, val := range d {
		if len(key) > 0 && (key[0] == '_' || strings.Contains(key, ".")) {
			continue
		}
		path := strings.Join(append(pref, key), ".")
		fields = append(fields, path)
		switch v := val.(type) {
		case map[string]any:
			fields = append(fields, getCompletionFields(v, append(pref, key))...)
		case []any:
			for i, item := range v {
				idxPath := fmt.Sprintf("%s[%d]", path, i)
				fields = append(fields, idxPath)
				if m, ok := item.(map[string]any); ok {
					fields = append(fields, getCompletionFields(m, append(pref, key))...)
				}
			}
		}
	}
	return fields
}

func loadPackageJson(c carapace.Context) (pj packageJson, err error) {
	var packageFile string
	if packageFile, err = util.FindReverse(c.Dir, "package.json"); err == nil {
		var content []byte
		if content, err = os.ReadFile(packageFile); err == nil {
			err = json.Unmarshal(content, &pj)
		}
	}
	return
}

// ActionPackageJsonKeys completes keys of package.json using dot notation
//
//	name
//	scripts.test
func ActionPackageJsonKeys() carapace.Action {
	return carapace.ActionMultiParts(".", func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			packageFile, err := util.FindReverse(c.Dir, "package.json")
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			content, err := os.ReadFile(packageFile)
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			var pkg map[string]any
			if err := json.Unmarshal(content, &pkg); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			current := any(pkg)
			for _, part := range c.Parts {
				obj, ok := current.(map[string]any)
				if !ok {
					return carapace.ActionValues()
				}
				next, ok := obj[part]
				if !ok {
					return carapace.ActionValues()
				}
				current = next
			}

			obj, ok := current.(map[string]any)
			if !ok {
				return carapace.ActionValues()
			}

			vals := make([]string, 0, len(obj)*2)
			hasNested := false
			for key, value := range obj {
				vals = append(vals, key, jsonType(value))
				if _, ok := value.(map[string]any); ok {
					hasNested = true
				}
			}
			action := carapace.ActionValuesDescribed(vals...)
			if hasNested {
				action = action.NoSpace('.')
			}
			return action
		})
	}).Tag("package.json keys")
}

func jsonType(value any) string {
	switch value.(type) {
	case map[string]any:
		return "object"
	case []any:
		return "array"
	case string:
		return "string"
	case float64:
		return "number"
	case bool:
		return "boolean"
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%T", value)
	}
}
