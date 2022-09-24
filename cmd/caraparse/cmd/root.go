package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "caraparse",
	Short: "GNU help parser.\n(e.g. 'ln --help | ./caraparse -n ln')",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parent, _ := cmd.Flags().GetString("parent")
		short, _ := cmd.Flags().GetString("short")
		convertX(name, parent, short)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringP("name", "n", "", "name of command")
	rootCmd.Flags().StringP("parent", "p", "", "name of parent command")
	rootCmd.Flags().StringP("short", "s", "", "short description of command")

	carapace.Gen(rootCmd)
}

type Command struct {
	Name   string
	Parent string
	Short  string
	Flags  []*Flag
}

func (c *Command) Format() string {
	pattern := `package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var %v = &cobra.Command{
	Use:   "%v",
	Short: "%v",
	Run:   func(cmd *cobra.Command, args []string) {},
}
%v
func init() {
	carapace.Gen(%v).Standalone()

%v%v
}`
	cmdName := c.Name + "Cmd"
	executeFunction := ""
	parent := ""
	if c.Parent != "" {
		if c.Parent != "root" {
			cmdName = c.Parent + "_" + cmdName
		}
		parent = fmt.Sprintf("\n	%vCmd.AddCommand(%v)", c.Parent, cmdName)
	} else {
		cmdName = "rootCmd"
		executeFunction = `
func Execute() error {
	return rootCmd.Execute()
}`
	}

	flags := make([]string, len(c.Flags))
	for i, flag := range c.Flags {
		flags[i] = "\t" + flag.Format(cmdName)
	}

	return fmt.Sprintf(pattern, cmdName, c.Name, c.Short, executeFunction, cmdName, strings.Join(flags, "\n"), parent)
}

type Flag struct {
	Name        string
	Shorthand   string
	Value       string
	Description string
}

func (flag *Flag) Format(cmdName string) string {
	if flag.Shorthand == "" {
		if flag.Value == "" {
			return fmt.Sprintf(`%v.Flags().Bool("%v", false, "%v")`, cmdName, flag.Name, strings.Replace(flag.Description, `"`, `\"`, -1))
		} else {
			return fmt.Sprintf(`%v.Flags().String("%v", "", "%v")`, cmdName, flag.Name, strings.Replace(flag.Description, `"`, `\"`, -1))
		}
	} else {
		if flag.Name == "" {
			if flag.Value == "" {
				return fmt.Sprintf(`%v.Flags().BoolS("%v", "%v", false, "%v")`, cmdName, flag.Shorthand, flag.Shorthand, strings.Replace(flag.Description, `"`, `\"`, -1))
			} else {
				return fmt.Sprintf(`%v.Flags().StringS("%v", "%v", "", "%v")`, cmdName, flag.Shorthand, flag.Shorthand, strings.Replace(flag.Description, `"`, `\"`, -1))
			}
		} else {
			if flag.Value == "" {
				return fmt.Sprintf(`%v.Flags().BoolP("%v", "%v", false, "%v")`, cmdName, flag.Name, flag.Shorthand, strings.Replace(flag.Description, `"`, `\"`, -1))
			} else {
				return fmt.Sprintf(`%v.Flags().StringP("%v", "%v", "", "%v")`, cmdName, flag.Name, flag.Shorthand, strings.Replace(flag.Description, `"`, `\"`, -1))
			}
		}
	}
}

func convertX(name string, parent string, short string) {
	flags := make([]*Flag, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if flag, err := parse(scanner.Text()); err == nil {
			flags = append(flags, flag)
		}
	}
	sort.Slice(flags, func(i, j int) bool { return flags[i].Name < flags[j].Name })

	if err := scanner.Err(); err != nil {
		log.Println(err)
	} else {
		cmd := Command{
			Name:   name,
			Parent: parent,
			Short:  short,
			Flags:  flags,
		}
		fmt.Println(cmd.Format())
	}
}

func parse(line string) (*Flag, error) {
	// git
	pattern := regexp.MustCompile(`^ *(-(?P<shorthand>[^-][^,[< ]*))?([, ]*(--(?P<name>[^ <=]+[^ [<=])))?(?P<value>[ [<=][^ ]+)?(  +(?P<description>[^ ].*)?)?$`)
	if !pattern.MatchString(line) {
		return nil, errors.New("no match")
	}
	m := pattern.FindStringSubmatch(line)

	flag := &Flag{}
	for i, name := range pattern.SubexpNames() {
		switch name {
		case "name":
			flag.Name = m[i]
		case "shorthand":
			flag.Shorthand = m[i]
		case "value":
			flag.Value = m[i]
		case "description":
			flag.Description = m[i]
		}
	}

	if flag.Name == "" {
		if flag.Shorthand == "" {
			return nil, errors.New("neither name nor shorthand matched")
		}
	}
	return flag, nil
}
