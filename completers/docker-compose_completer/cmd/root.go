package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var rootCmd = &cobra.Command{
	Use:   "docker-compose",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("compatibility", false, "If set, Compose will attempt to convert keys")
	rootCmd.Flags().StringP("context", "c", "", "Specify a context name")
	rootCmd.Flags().String("env-file", "", "Specify an alternate environment file")
	rootCmd.Flags().StringP("file", "f", "", "Specify an alternate compose file")
	rootCmd.Flags().StringP("host", "H", "", "Daemon socket to connect to")
	rootCmd.Flags().String("log-level", "", "Set log level (DEBUG, INFO, WARNING, ERROR, CRITICAL)")
	rootCmd.Flags().Bool("no-ansi", false, "Do not print ANSI control characters")
	rootCmd.Flags().String("project-directory", "", "Specify an alternate working directory")
	rootCmd.Flags().StringP("project-name", "p", "", "Specify an alternate project name")
	rootCmd.Flags().Bool("skip-hostname-check", false, "Don't check the daemon's hostname against the")
	rootCmd.Flags().Bool("tls", false, "Use TLS; implied by --tlsverify")
	rootCmd.Flags().String("tlscacert", "", "Trust certs signed only by this CA")
	rootCmd.Flags().String("tlscert", "", "Path to TLS certificate file")
	rootCmd.Flags().String("tlskey", "", "Path to TLS key file")
	rootCmd.Flags().Bool("tlsverify", false, "Use TLS and verify the remote")
	rootCmd.Flags().Bool("verbose", false, "Show more output")
	rootCmd.Flags().BoolP("version", "v", false, "Print version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file":              carapace.ActionFiles(),
		"log-level":         carapace.ActionValues("DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"),
		"tlscert":           carapace.ActionFiles(".crt"),
		"tlskey":            carapace.ActionFiles(".key"),
		"project-directory": carapace.ActionDirectories(),
		"env-file":          carapace.ActionFiles(),
	})
}

func ActionServices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if services, _, err := loadFile(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(services...)
		}
	})
}

func ActionVolumes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if _, volumes, err := loadFile(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(volumes...)
		}
	})
}

type compose struct {
	Version  string
	Services map[string]interface{}
	Volumes  map[string]interface{}
}

func fileLocation() (string, error) {
	if rootCmd.Flag("file").Changed {
		return rootCmd.Flag("file").Value.String(), nil
	} else {
		if path, err := os.Getwd(); err == nil {
			return traverse(path)
		} else {
			return "", err
		}
	}
}

func traverse(path string) (string, error) {
	if _, err := os.Stat(path + "/docker-compose.yml"); err == nil {
		return path + "/docker-compose.yml", nil
	} else {
		if path == "/" {
			return "", errors.New("no docker-compose.yml present")
		}
		return traverse(filepath.Dir(path))
	}
}

func loadFile() ([]string, []string, error) {
	file, err := fileLocation()
	if err != nil {
		return nil, nil, err
	}

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, nil, err
	}

	var content compose
	if err = yaml.Unmarshal(yamlFile, &content); err != nil {
		return nil, nil, err
	}
	return keys(content.Services), keys(content.Volumes), nil
}

func keys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	index := 0
	for key := range m {
		keys[index] = key
		index += 1
	}
	return keys
}
