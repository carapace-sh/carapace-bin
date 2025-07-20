package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "kitty",
	Short:              "The fast, feature rich terminal emulator",
	Long:               "https://sw.kovidgoyal.net/kitty/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			request := []string{"kitty"}
			request = append(request, c.Args...)
			request = append(request, c.Value)

			matchRequests := [][]string{request}
			if c.Value == "-" { // some longhand flags are missing when called with `-`
				longhandRequest := []string{"kitty"}
				longhandRequest = append(longhandRequest, c.Args...)
				longhandRequest = append(longhandRequest, "--")
				matchRequests = append(matchRequests, longhandRequest)
			}

			// partial directories aren't always completed (e.g. tmp directory with `kitty /tm`)
			filepathRequest := []string{"kitty"}
			filepathRequest = append(filepathRequest, c.Args...)
			filepathRequest = append(filepathRequest, filepath.Dir(c.Value))
			matchRequests = append(matchRequests, filepathRequest)

			m, err := json.Marshal(matchRequests)
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			var stdout, stderr bytes.Buffer
			cmd := c.Command("kitten", "__complete__", "json")
			cmd.Stdin = bytes.NewReader(m)
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			if err := cmd.Run(); err != nil {
				if exitErr, ok := err.(*exec.ExitError); ok {
					if firstLine := strings.SplitN(string(exitErr.Stderr), "\n", 2)[0]; strings.TrimSpace(firstLine) != "" {
						err = errors.New(firstLine)
					}
				}
				return carapace.ActionMessage(err.Error())
			}

			var matchResults []matchResult
			if err := json.Unmarshal(stdout.Bytes(), &matchResults); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			if len(matchResults) == 0 {
				return carapace.ActionValues()
			}

			var nospace bool
			batch := carapace.Batch()
			for _, matchResult := range matchResults {
				for _, group := range matchResult.Groups {
					switch group.Title {
					case "Directories":
						batch = append(batch, carapace.ActionDirectories())
					case "Executables":
						dir := filepath.ToSlash(filepath.Dir(c.Value))
						batch = append(batch,
							carapace.ActionDirectories(),
							carapace.ActionExecutables(dir).Prefix(dir+"/"),
						)
					case "Executables in PATH":
						batch = append(batch, carapace.ActionExecutables())
					case "Options":
						longhandMatches := make(matches, 0)
						shorthandMatches := make(matches, 0)
						for _, match := range group.Matches {
							switch {
							case strings.HasPrefix(match.Word, "--"):
								longhandMatches = append(longhandMatches, match)
							default:
								shorthandMatches = append(shorthandMatches, match)
							}
						}
						batch = append(batch,
							longhandMatches.Action().Tag("longhand flags"),
							shorthandMatches.Action().Tag("shorthand flags"),
						)
					case "Keywords":
						batch = append(batch, group.Matches.Action().Tag("keywords").StyleF(style.ForKeyword))
					case "Sub-commands":
						batch = append(batch, group.Matches.Action().Tag("commands"))
					default:
						switch {
						case group.IsFiles:
							// TODO this tries to use carapace file completion via retain (switch to provided values if there is an issue)
							retain := make([]string, 0)
							for _, match := range group.Matches {
								retain = append(retain, filepath.ToSlash(match.Word))
							}
							batch = append(batch, carapace.ActionFiles().Retain(retain...))
						default:
							batch = append(batch, group.Matches.Action().Tag(strings.ToLower(group.Title)))
						}
					}

					nospace = nospace || group.NoTrailingSpace
				}
			}

			if nospace {
				return batch.ToA().NoSpace()
			}
			return batch.ToA()
		}),
	)
}

type matchResult struct {
	Groups []struct {
		Title           string  `json:"title"`
		NoTrailingSpace bool    `json:"no_trailing_space"`
		IsFiles         bool    `json:"is_files"`
		Matches         matches `json:"matches"`
	} `json:"groups"`
	Delegate struct {
	} `json:"delegate"`
}

type matches []struct {
	Word        string `json:"word"`
	Description string `json:"description"`
}

func (m matches) Action() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		for _, match := range m {
			vals = append(vals, match.Word, match.Description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
