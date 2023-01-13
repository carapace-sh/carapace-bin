package cmd

import (
	"bytes"
	"fmt"
	"github.com/rsteube/carapace-spec"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func scrape(path string) {
	// TODO yuck - all this needs some cleanup
	if _, spec, err := loadSpec(path); err == nil {
		spec.Scrape()
	}
}

func specCompletion(path string, args ...string) {
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
		return /// TODO handle error
	}
	cmd, err := spec.ToCobra()
	if err != nil {
		return /// TODO handle error
	}

	a := []string{"_carapace"}
	a = append(a, args...)
	cmd.SetArgs(a)
	cmd.Execute()

	w.Close()
	out := <-outC
	os.Stdout = old

	executable, err := os.Executable()
	if err != nil {
		panic(err.Error()) // TODO exit with error message
	}

	executableName := filepath.Base(executable)
	patched := strings.Replace(string(out), fmt.Sprintf("%v _carapace", executableName), fmt.Sprintf("%v --spec '%v'", executableName, abs), -1)       // general callback
	patched = strings.Replace(patched, fmt.Sprintf("'%v', '_carapace'", executableName), fmt.Sprintf("'%v', '--spec', '%v'", executableName, abs), -1) // xonsh callback
	fmt.Print(patched)

}
