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
			args := []string{"kitty"}
			args = append(args, c.Args...)
			args = append(args, c.Value)

			matchRequests := [][]string{args}
			if c.Value == "-" { // some longhand flags are missing when called with `-`
				longhandRequest := []string{"kitty"}
				longhandRequest = append(longhandRequest, c.Args...)
				longhandRequest = append(longhandRequest, "--")
				matchRequests = append(matchRequests, longhandRequest)
			}

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

			batch := carapace.Batch()
			for _, matchResult := range matchResults {
				for _, group := range matchResult.Groups {
					includeFiles := true // no explicit file completion and some positions need this
					switch group.Title {
					case "Directories":
						includeFiles = false
						batch = append(batch, carapace.ActionDirectories())
					case "Executables":
						includeFiles = false
						dir := filepath.ToSlash(filepath.Dir(c.Value))
						batch = append(batch,
							carapace.ActionDirectories(),
							carapace.ActionExecutables(dir).Prefix(dir+"/"),
						)
					case "Executables in PATH":
						includeFiles = false
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
						batch = append(batch, group.Matches.Action().Tag(strings.ToLower(group.Title)))
					}

					if includeFiles {
						batch = append(batch, carapace.ActionFiles())
					}
				}
			}
			return batch.ToA()
		}),
	)
}

type matchResult struct {
	Groups []struct {
		Title   string  `json:"title"`
		Matches matches `json:"matches"`
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
