package shim

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/xdg"
)

func Update() error {
	if err := removeObsolete(); err != nil {
		return err
	}

	carapace.LOG.Println("updating shims")

	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(configDir + "/carapace/specs")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".yaml") {
			s := shim{Target: fmt.Sprintf("%v/carapace/specs/%v", configDir, entry.Name())}
			if err := s.Update(); err != nil {
				return err
			}
		}
	}
	return nil
}

type shim struct {
	Target string
}

func (s shim) Name() string {
	return strings.TrimSuffix(filepath.Base(s.Target), ".yaml")
}

func (s shim) Path() (string, error) {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return "", err
	}

	name := s.Name()
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	return fmt.Sprintf("%v/carapace/bin/%v", configDir, name), nil
}

func (s shim) DummyPath() (string, error) {
	path, err := xdg.UserCacheDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v/carapace/.shims/%v", path, s.Name()), nil
}

func (s shim) NeedsUpdate() bool {
	path, err := s.Path()
	if err != nil {
		return true
	}

	targetInfo, err := os.Stat(s.Target)
	if err != nil {
		return true
	}

	shimInfo, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return true
		}

		dummyPath, err := s.DummyPath()
		if err != nil {
			return true
		}

		shimInfo, err = os.Stat(dummyPath)
		if err != nil {
			return true
		}
	}

	return shimInfo.ModTime().Before(targetInfo.ModTime())
}

func (s shim) IsRunnable() bool {
	file, err := os.Open(s.Target)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	r := regexp.MustCompile(`^ *run:`)
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			return true
		}
	}
	return false
}

func (s shim) Update() error {
	if !s.NeedsUpdate() {
		return nil
	}

	if s.IsRunnable() {
		return s.WriteShim()
	}

	if err := s.RemoveShim(); err != nil {
		return err
	}
	return s.WriteDummy()
}

func (s shim) RemoveShim() error {
	path, err := s.Path()
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	carapace.LOG.Printf("removing shim %#v", path)
	return os.Remove(path)
}

func (s shim) WriteDummy() error {
	path, err := s.DummyPath()
	if err != nil {
		return err
	}

	dummyDir := filepath.Dir(path)
	if _, err := os.Stat(dummyDir); os.IsNotExist(err) {
		carapace.LOG.Printf("creating shim dummy directory %#v\n", dummyDir)
		if err := os.MkdirAll(dummyDir, 0700); err != nil {
			return err
		}
	}

	carapace.LOG.Printf("writing dummy to %#v\n", path)
	if err := os.WriteFile(path, []byte{}, 0644); err != nil {
		return err
	}
	return os.Chmod(path, 0644)

}

func (s shim) WriteShim() error {
	path, err := s.Path()
	if err != nil {
		return err
	}

	binDir := filepath.Dir(path)
	if _, err := os.Stat(binDir); os.IsNotExist(err) {
		carapace.LOG.Printf("creating shim directory %#v\n", binDir)
		if err := os.MkdirAll(binDir, os.ModePerm); err != nil {
			return err
		}
	}

	carapace.LOG.Printf("writing shim to %#v\n", path)
	if err := os.WriteFile(path, body(s.Target), 0755); err != nil {
		return err
	}
	return os.Chmod(path, 0755)
}
