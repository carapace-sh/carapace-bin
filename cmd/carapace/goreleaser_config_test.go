package main

import (
	"os"
	"testing"

	"go.yaml.in/yaml/v3"
)

type goreleaserConfig struct {
	Archives []struct {
		ID    string `yaml:"id"`
		Files []struct {
			Src string `yaml:"src"`
			Dst string `yaml:"dst"`
		} `yaml:"files"`
	} `yaml:"archives"`
}

func TestArchiveFileDestinationsAreNotDot(t *testing.T) {
	raw, err := os.ReadFile("../../.goreleaser.yml")
	if err != nil {
		t.Fatal(err)
	}
	var cfg goreleaserConfig
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		t.Fatal(err)
	}
	for _, archive := range cfg.Archives {
		for _, file := range archive.Files {
			if file.Dst == "." {
				t.Errorf("archive %q: file %q has dst %q, which creates a regular file named '.' in the archive", archive.ID, file.Src, file.Dst)
			}
		}
	}
}
