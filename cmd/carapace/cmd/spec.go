package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	spec "github.com/carapace-sh/carapace-spec"
	"gopkg.in/yaml.v3"
)

func loadSpec(path string) (string, *spec.Command, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", nil, err
	}

	content, err := os.ReadFile(abs)
	if err != nil {
		return "", nil, err
	}

	var specCmd spec.Command
	if err := yaml.Unmarshal(content, &specCmd); err != nil {
		return "", nil, err
	}
	return abs, &specCmd, nil
}

func codegen(path string) error {
	// TODO yuck - all this needs some cleanup
	_, spec, err := loadSpec(path)
	if err != nil {
		return err
	}
	return spec.Codegen()
}

func specCompletion(path string, args ...string) (string, error) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	abs, spec, err := loadSpec(path)
	if err != nil {
		return "", err
	}

	cmd := spec.ToCobra()

	a := []string{"_carapace"}
	a = append(a, args...)
	cmd.SetArgs(a)
	cmd.Execute()

	w.Close()
	out := <-outC
	os.Stdout = old

	executable, err := os.Executable()
	if err != nil {
		return "", err
	}

	executableName := filepath.Base(executable)
	patched := strings.Replace(string(out), fmt.Sprintf("%v _carapace", executableName), fmt.Sprintf("%v --spec '%v'", executableName, abs), -1)       // general callback
	patched = strings.Replace(patched, fmt.Sprintf("'%v', '_carapace'", executableName), fmt.Sprintf("'%v', '--spec', '%v'", executableName, abs), -1) // xonsh callback
	return patched, nil
}
