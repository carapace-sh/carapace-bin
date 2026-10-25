package git

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/carapace-sh/carapace"
)

func gitRun(t *testing.T, dir string, args ...string) {
	t.Helper()

	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(),
		"GIT_AUTHOR_NAME=test", "GIT_AUTHOR_EMAIL=test@example.com",
		"GIT_COMMITTER_NAME=test", "GIT_COMMITTER_EMAIL=test@example.com",
		"GIT_CONFIG_GLOBAL=/dev/null", "GIT_CONFIG_SYSTEM=/dev/null", "GIT_CONFIG_NOSYSTEM=1")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatal(strings.TrimSpace(err.Error() + ": " + string(output)))
	}
}

func gitTestRepo(t *testing.T, commands ...string) string {
	t.Helper()

	dir := t.TempDir()
	gitRun(t, dir, "init", "--initial-branch", "main", ".")

	for _, command := range commands {
		arg, value, _ := strings.Cut(command, "=")
		switch arg {
		case "write": // write=filename
			if err := os.MkdirAll(filepath.Dir(filepath.Join(dir, value)), 0755); err != nil {
				t.Fatal(err.Error())
			}
			if err := os.WriteFile(filepath.Join(dir, value), []byte(value+" content\n"), 0644); err != nil {
				t.Fatal(err.Error())
			}
		case "modify": // modify=filename
			f, err := os.OpenFile(filepath.Join(dir, value), os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				t.Fatal(err.Error())
			}
			if _, err := f.WriteString("modified\n"); err != nil {
				t.Fatal(err.Error())
			}
			if err := f.Close(); err != nil {
				t.Fatal(err.Error())
			}
		case "rename": // rename=old->new
			old, new, _ := strings.Cut(value, "->")
			if err := os.Rename(filepath.Join(dir, old), filepath.Join(dir, new)); err != nil {
				t.Fatal(err.Error())
			}
		case "remove": // remove=filename
			if err := os.Remove(filepath.Join(dir, value)); err != nil {
				t.Fatal(err.Error())
			}
		default: // git command
			gitRun(t, dir, strings.Fields(command)...)
		}
	}
	return dir
}

func dumpValues(t *testing.T, action carapace.Action, c carapace.Context) [][4]string {
	t.Helper()

	output, err := action.Invoke(c).MarshalJSON()
	if err != nil {
		t.Fatal(err.Error())
	}

	export := struct {
		Values []struct {
			Value       string
			Display     string
			Description string
			Tag         string
		}
	}{}
	if err := json.Unmarshal(output, &export); err != nil {
		t.Fatal(err.Error() + ": " + string(output))
	}

	vals := make([][4]string, 0)
	for _, value := range export.Values {
		vals = append(vals, [4]string{value.Value, value.Display, value.Description, value.Tag})
	}
	slices.SortFunc(vals, func(a, b [4]string) int { return strings.Compare(strings.Join(a[:], ":"), strings.Join(b[:], ":")) })
	return vals
}

func TestActionChanges(t *testing.T) {
	dir := gitTestRepo(t,
		"write=a.txt", "write=c.txt", "write=e.txt",
		"add -A", "commit -m initial",
		"write=b.go", "write=sub/g.md",
		"add -A",
		"rename=e.txt->f.txt", "add e.txt f.txt",
		"modify=a.txt", "remove=c.txt", "write=d.md",
	)

	tests := []struct {
		name  string
		opts  ChangeOpts
		value string
		wd    string
		want  [][4]string
	}{
		{name: "none-selected", opts: ChangeOpts{}, want: [][4]string{}},
		{name: "default", opts: ChangeOpts{}.Default(), want: [][4]string{
			{"a.txt", "a.txt", " M", "changes"},
			{"b.go", "b.go", "A ", "changes"},
			{"c.txt", "c.txt", " D", "changes"},
			{"d.md", "d.md", "??", "changes"},
			{"f.txt", "f.txt", "R ", "changes"},
			{"sub/g.md", "sub/g.md", "A ", "changes"},
		}},
		{name: "staged", opts: ChangeOpts{Staged: true}, want: [][4]string{
			{"b.go", "b.go", "A ", "changes"},
			{"f.txt", "f.txt", "R ", "changes"},
			{"sub/g.md", "sub/g.md", "A ", "changes"},
		}},
		{name: "unstaged", opts: ChangeOpts{Unstaged: true}, want: [][4]string{
			{"a.txt", "a.txt", " M", "changes"},
			{"c.txt", "c.txt", " D", "changes"},
			{"d.md", "d.md", "??", "changes"},
		}},
		{name: "dotted", opts: ChangeOpts{}.Default(), value: "./", want: [][4]string{
			{"./a.txt", "a.txt", " M", "changes"},
			{"./b.go", "b.go", "A ", "changes"},
			{"./c.txt", "c.txt", " D", "changes"},
			{"./d.md", "d.md", "??", "changes"},
			{"./f.txt", "f.txt", "R ", "changes"},
			{"./sub/g.md", "sub/g.md", "A ", "changes"},
		}},
		{name: "from-subdir", opts: ChangeOpts{}.Default(), wd: dir + "/sub", want: [][4]string{
			{"../a.txt", "../a.txt", " M", "changes"},
			{"../b.go", "../b.go", "A ", "changes"},
			{"../c.txt", "../c.txt", " D", "changes"},
			{"../d.md", "../d.md", "??", "changes"},
			{"../f.txt", "../f.txt", "R ", "changes"},
			{"g.md", "g.md", "A ", "changes"},
		}},
		{name: "from-subdir-dotted", opts: ChangeOpts{}.Default(), value: "./", wd: dir + "/sub", want: [][4]string{
			{"./../a.txt", "../a.txt", " M", "changes"},
			{"./../b.go", "../b.go", "A ", "changes"},
			{"./../c.txt", "../c.txt", " D", "changes"},
			{"./../d.md", "../d.md", "??", "changes"},
			{"./../f.txt", "../f.txt", "R ", "changes"},
			{"./g.md", "g.md", "A ", "changes"},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wd := test.wd
			if wd == "" {
				wd = dir
			}
			got := dumpValues(t, ActionChanges(test.opts), carapace.Context{Value: test.value, Dir: wd})
			if !slices.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestActionRefDiffs(t *testing.T) {
	dir := gitTestRepo(t,
		"write=a.txt", "write=c.txt", "write=e.txt",
		"add -A", "commit -m initial",
		"write=b.go", "write=sub/g.md", "modify=a.txt", "remove=c.txt",
		"add -A", "commit -m second",
		"modify=a.txt",
	)

	tests := []struct {
		name  string
		refs  []string
		value string
		wd    string
		want  [][4]string
	}{
		{name: "no-ref", want: [][4]string{
			{"a.txt", "a.txt", "M", "changed files"},
		}},
		{name: "one-ref", refs: []string{"HEAD~1"}, want: [][4]string{
			{"a.txt", "a.txt", "M", "changed files"},
			{"b.go", "b.go", "A", "changed files"},
			{"c.txt", "c.txt", "D", "changed files"},
			{"sub/g.md", "sub/g.md", "A", "changed files"},
		}},
		{name: "two-refs", refs: []string{"HEAD", "HEAD~1"}, want: [][4]string{
			{"a.txt", "a.txt", "M", "changed files"},
			{"b.go", "b.go", "D", "changed files"},
			{"c.txt", "c.txt", "A", "changed files"},
			{"sub/g.md", "sub/g.md", "D", "changed files"},
		}},
		{name: "dotted", value: "./", want: [][4]string{
			{"./a.txt", "a.txt", "M", "changed files"},
		}},
		{name: "from-subdir", wd: dir + "/sub", want: [][4]string{
			{"../a.txt", "../a.txt", "M", "changed files"},
		}},
		{name: "from-subdir-dotted", value: "./", wd: dir + "/sub", want: [][4]string{
			{"./../a.txt", "../a.txt", "M", "changed files"},
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wd := test.wd
			if wd == "" {
				wd = dir
			}
			got := dumpValues(t, ActionRefDiffs(test.refs...), carapace.Context{Value: test.value, Dir: wd})
			if !slices.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
