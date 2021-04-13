package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "carafmt",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exitCode := 0
		for _, arg := range args {
			if tmpfile, err := format(arg); err != nil {
				println(err.Error())
				exitCode = 1
			} else {
				if err := os.Rename(tmpfile, arg); err != nil {
					println(err.Error())
				}
			}
		}
		os.Exit(exitCode)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".go"),
	)
}

func checkGitStatus(path string) error {
	if output, err := exec.Command("git", "status", "--porcelain", path).Output(); err != nil {
		return err
	} else {
		lines := strings.Split(string(output), "\n")
		if len(lines[0]) > 0 {
			return fmt.Errorf("git status of '%v' not clean", path)
		}
	}
	return nil
}

func format(path string) (string, error) {
	if !strings.HasSuffix(path, ".go") {
		return "", fmt.Errorf("file '%v' has no '.go' suffix", path)
	}
	if err := checkGitStatus(path); err != nil {
		return "", err
	}
	if err := exec.Command("go", "fmt", path).Run(); err != nil {
		return "", fmt.Errorf("failed execute go fmt on '%v': %v", path, err.Error())
	}

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := os.Stat(path + "_carafmt"); err == nil {
		return "", fmt.Errorf("tmpfile '%v' already exists", path+"_carafmt")
	}
	tmpfile, err := os.Create(path + "_carafmt")
	if err != nil {
		return "", err
	}
	defer tmpfile.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(tmpfile)

	r := regexp.MustCompile(`\.Flags\(\)\.(Bool|String)[^(]*\("(?P<name>[^"]+)"`)

	toSort := make(map[string]string)
	for scanner.Scan() {
		if !r.MatchString(scanner.Text()) {
			if len(toSort) > 0 {
				keys := make([]string, 0)
				for key := range toSort {
					keys = append(keys, key)
				}
				sort.Strings(keys)

				for _, sortedKey := range keys {
					writer.WriteString(toSort[sortedKey] + "\n")
				}
				toSort = make(map[string]string)
			}
			writer.WriteString(scanner.Text() + "\n")
		} else {
			flagName := r.FindStringSubmatch(scanner.Text())[2]
			toSort[flagName] = scanner.Text()
		}
	}
	writer.Flush()
	return tmpfile.Name(), nil
}
