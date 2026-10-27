package action

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionStashIndexes completes stash entries as recorded by `grit stash list`
//
//	stash@{0} (WIP on branchA: ABCDEF1 fix bug)
//	stash@{1} (WIP on branchB: ABCDEF2 work)
func ActionStashIndexes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("grit", "stash", "list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				splitted := strings.SplitN(line, ": ", 2)
				fields := strings.SplitN(splitted[0], " ", 2)
				if len(splitted) == 2 && len(fields) == 2 {
					vals = append(vals, fields[0], splitted[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionHookNames completes hook scripts known to `grit hooks list`
//
//	pre-commit
//	commit-msg
func ActionHookNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("grit", "hooks", "list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				if name := strings.SplitN(line, "  ", 2)[0]; name != "" {
					vals = append(vals, name)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}

// ActionConfigKeys completes configuration keys of `grit config` from `grit config --list`
//
//	commit.autoStage (false)
//	user.name (J. Doe)
//
// ActionConfigKeys completes configuration keys of `grit config` from `grit config --list`
//
// ActionConfigKeys completes configuration keys of `grit config` from `grit config --list`
//
//	commit.autoStage (false)
//	user.name (J. Doe)
func ActionConfigKeys() carapace.Action {
	return carapace.ActionMultiParts(".", func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("grit", "config", "--list")(func(output []byte) carapace.Action {
				vals := make([]string, 0)
				seen := map[string]struct{}{}

				prefix := strings.Join(c.Parts, ".")
				for _, line := range strings.Split(string(output), "\n")[:len(strings.Split(string(output), "\n"))-1] {
					splitted := strings.SplitN(line, "=", 2)
					key := splitted[0]
					if prefix != "" && !strings.HasPrefix(key, prefix+".") {
						continue
					}
					rest := strings.TrimPrefix(key, prefix+".")
					if prefix == "" {
						rest = key
					}
					description := ""
					if index := strings.Index(rest, "."); index >= 0 {
						rest = rest[:index]
					} else if len(splitted) == 2 {
						description = splitted[1]
					}
					if _, exists := seen[rest]; !exists {
						seen[rest] = struct{}{}
						vals = append(vals, rest, description)
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		})
	})
}

// diffSelection mirrors the `grit diff --json` shape used for patch selection.
type diffSelection struct {
	Files []diffFileSelection `json:"files"`
}

type diffFileSelection struct {
	Path  string          `json:"path"`
	Hunks []selectionHunk `json:"hunks"`
}

type selectionHunk struct {
	ID       string `json:"id"`
	NewStart int    `json:"new_start"`
	NewLines int    `json:"new_lines"`
}

// ActionHunkIds completes hunk ids of `grit diff --{unstaged,staged} --json`
// (`--hunks` selections, comma-separated or repeated)
//
//	4a1f0d1e83ab (doc.md)
//	9c02b7f1214c (main.rs)
func ActionHunkIds(side string) carapace.Action {
	return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("grit", "diff", "--"+side, "--json")(func(output []byte) carapace.Action {
				vals := make([]string, 0)
				for _, file := range unmarshalDiffSelection(output) {
					for _, hunk := range file.Hunks {
						if hunk.ID != "" {
							vals = append(vals, hunk.ID, file.Path)
						}
					}
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		})
	})
}

// ActionLineSelections completes new-side line ranges of
// `grit diff --{unstaged,staged} --json` (`path:start-end`)
//
//	doc.md (changed)
//	doc.md:3-5
func ActionLineSelections(side string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("grit", "diff", "--"+side, "--json")(func(output []byte) carapace.Action {
			files := unmarshalDiffSelection(output)
			paths := make([]string, 0)
			for _, file := range files {
				paths = append(paths, file.Path)
			}

			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValues(paths...).Suffix(":")
					default:
						vals := make([]string, 0)
						for _, file := range files {
							if file.Path != c.Parts[0] {
								continue
							}
							for _, hunk := range file.Hunks {
								if hunk.NewLines > 0 {
									vals = append(vals,
										fmt.Sprintf("%v-%v", hunk.NewStart, hunk.NewStart+hunk.NewLines-1),
									)
								}
							}
						}
						return carapace.ActionValues(vals...)
					}
				})
			}).UniqueList(",").NoSpace()
		})
	})
}

// ActionReflogSpecs completes reflog selectors `ref@{id}` from `grit reflog <ref>`
//
//	HEAD@{0} (61bbffc 2 minutes ago commit (initial): first)
//	main@{3} (6f9cd2f 5 hours ago commit message)
func ActionReflogSpecs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		ref := c.Value
		if index := strings.IndexAny(ref, "@{"); index > 0 {
			ref = ref[:index]
		}
		if ref == "" {
			ref = "HEAD"
		}

		return carapace.ActionExecCommand("grit", "reflog", ref)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				entry := strings.SplitN(line, " ", 2)
				if len(entry) == 2 {
					vals = append(vals, entry[0], entry[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

// ActionWorktreePaths completes paths of registered worktrees from `grit worktree list`
//
//	../docs (docs-site)
//	../release (release) (locked)
func ActionWorktreePaths() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("grit", "worktree", "list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				fields := strings.TrimPrefix(line, "*")
				if fst := strings.IndexAny(fields, "[("); fst > 0 {
					fields = fields[:fst]
				} else if fst == 0 {
					continue
				}
				path := strings.TrimSpace(fields)
				if path != "" {
					vals = append(vals, path)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}

func unmarshalDiffSelection(output []byte) []diffFileSelection {
	var selection diffSelection
	if err := json.Unmarshal(output, &selection); err != nil {
		return nil
	}
	return selection.Files
}
