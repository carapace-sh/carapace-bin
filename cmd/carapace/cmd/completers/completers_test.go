package completers

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/carapace-sh/carapace-bridge/pkg/choices"
)

func BenchmarkCompleters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Completers(choices.Choice{}, false)
	}
}

func TestSpecPath(t *testing.T) {
	if runtime.GOOS != "windows" && runtime.GOOS != "plan9" {
		t.Setenv("CARAPACE_SPECS_PATH", strings.Join([]string{"/user@user", "/unclassified", "/system@system"}, string(os.PathListSeparator)))
	} else {
		t.Setenv("CARAPACE_SPECS_PATH", strings.Join([]string{`C:\base@user`, `C:\unclassified`}, string(os.PathListSeparator)))
	}

	dirs := SpecPath()
	if runtime.GOOS == "windows" || runtime.GOOS == "plan9" {
		if len(dirs) != 2 {
			t.Fatalf("unexpected entries: %v", dirs)
		}
		if dirs[0].Path != `C:\base` || dirs[0].Group != "user" {
			t.Fatalf("unexpected first entry: %#v", dirs[0])
		}
		if dirs[1].Group != "system" {
			t.Fatalf("unclassified entry should default to system: %#v", dirs[1])
		}
		return
	}

	if len(dirs) != 3 {
		t.Fatalf("unexpected entries: %v", dirs)
	}
	if dirs[0].Path != "/user" || dirs[0].Group != "user" {
		t.Fatalf("unexpected first entry: %#v", dirs[0])
	}
	if dirs[1].Path != "/unclassified" || dirs[1].Group != "system" {
		t.Fatalf("unclassified entry should default to system: %#v", dirs[1])
	}
	if dirs[2].Path != "/system" || dirs[2].Group != "system" {
		t.Fatalf("unexpected third entry: %#v", dirs[2])
	}
}

func TestSpecDirs(t *testing.T) {
	if runtime.GOOS == "windows" || runtime.GOOS == "plan9" {
		t.Skip("unix-specific path layout test")
	}

	t.Setenv("HOME", "/home/test")
	t.Setenv("XDG_CONFIG_HOME", "/user/cfg")
	t.Setenv("XDG_DATA_DIRS", strings.Join([]string{"/data1", "/data2"}, string(os.PathListSeparator)))
	t.Setenv("CARAPACE_SPECS_PATH", strings.Join([]string{"/env/user@user", "/env/sys"}, string(os.PathListSeparator)))

	dirs, err := SpecDirs()
	if err != nil {
		t.Fatal(err)
	}

	expected := [...][2]string{
		{"/user/cfg/carapace/specs", "user"},
		{"/env/user", "user"},
		{"/env/sys", "system"},
		{"/data1/carapace/specs", "system"},
		{"/data2/carapace/specs", "system"},
	}
	if len(dirs) != len(expected) {
		t.Fatalf("expected %v entries, got: %v", len(expected), dirs)
	}
	for i, e := range expected {
		if dirs[i].Path != e[0] {
			t.Errorf("entry %v: expected path %#v, got %#v", i, e[0], dirs[i].Path)
		}
		if dirs[i].Group != e[1] {
			t.Errorf("entry %v: expected group %#v, got %#v", i, e[1], dirs[i].Group)
		}
	}
}

func TestSpecDirsDefaults(t *testing.T) {
	t.Setenv("HOME", "/home/test")
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("XDG_DATA_DIRS", "")
	t.Setenv("CARAPACE_SPECS_PATH", "")

	dirs, err := SpecDirs()
	if err != nil {
		t.Fatal(err)
	}
	if runtime.GOOS == "windows" || runtime.GOOS == "plan9" {
		// PROGRAMDATA dependent, just verify user dir is respected via os.UserConfigDir
		if dirs[0].Group != "user" {
			t.Fatalf("first entry should be the user spec dir: %v", dirs[0])
		}
		return
	}

	expected := [...][2]string{
		{"/home/test/.config/carapace/specs", "user"},
		{"/usr/local/share/carapace/specs", "system"},
		{"/usr/share/carapace/specs", "system"},
	}
	if len(dirs) != len(expected) {
		t.Fatalf("expected %v entries, got: %v", len(expected), dirs)
	}
	for i, e := range expected {
		if dirs[i].Path != e[0] || dirs[i].Group != e[1] {
			t.Errorf("entry %v: expected %#v, got %#v", i, e, dirs[i])
		}
	}
}

func TestDataDirsDefault(t *testing.T) {
	t.Setenv("XDG_DATA_DIRS", "")
	if runtime.GOOS == "windows" || runtime.GOOS == "plan9" {
		t.Skip("unix-specific default test")
	}

	dirs, err := specDataDirs()
	if err != nil {
		t.Fatal(err)
	}
	if len(dirs) != 2 || dirs[0] != "/usr/local/share" || dirs[1] != "/usr/share" {
		t.Fatalf("unexpected default data dirs: %v", dirs)
	}
}
