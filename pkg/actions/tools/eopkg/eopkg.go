// package eopkg contains Solus package manager related actions.
package eopkg

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace"
)

const (
	indexPattern     = "/var/lib/eopkg/index/*/eopkg-index.xml"
	metadataPattern  = "/var/lib/eopkg/package/*/metadata.xml"
	repositoryFolder = "/var/lib/eopkg/index"
)

type packageEntry struct {
	Name        string `xml:"Name"`
	Summary     string `xml:"Summary"`
	Description string `xml:"Description"`
	PartOf      string `xml:"PartOf"`
}

// ActionAvailablePackages completes packages available from configured repositories.
func ActionAvailablePackages() carapace.Action {
	return actionPackages("available packages", indexPattern)
}

// ActionInstalledPackages completes installed packages.
func ActionInstalledPackages() carapace.Action {
	return actionPackages("installed packages", metadataPattern)
}

func actionPackages(tag string, patterns ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		packages := make(map[string]string)

		for _, entry := range readPackageEntries(patterns...) {
			if entry.Name == "" {
				continue
			}

			if description := firstNonEmpty(entry.Summary, entry.Description); description != "" {
				packages[entry.Name] = description
			} else if _, ok := packages[entry.Name]; !ok {
				packages[entry.Name] = ""
			}
		}

		vals := make([]string, 0, len(packages)*2)
		for _, name := range sortedKeys(packages) {
			vals = append(vals, name, packages[name])
		}
		return carapace.ActionValuesDescribed(vals...).Tag(tag)
	})
}

// ActionComponents completes package components found in local eopkg metadata.
func ActionComponents() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		components := make(map[string]bool)

		for _, entry := range readPackageEntries(indexPattern, metadataPattern) {
			if component := strings.TrimSpace(entry.PartOf); component != "" {
				components[component] = true
			}
		}

		vals := make([]string, 0, len(components))
		for component := range components {
			vals = append(vals, component)
		}
		sort.Strings(vals)
		return carapace.ActionValues(vals...).Tag("components")
	})
}

// ActionRepositories completes configured repository names.
func ActionRepositories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		entries, err := os.ReadDir(repositoryFolder)
		if err != nil {
			return carapace.ActionValues()
		}

		vals := make([]string, 0, len(entries))
		for _, entry := range entries {
			if entry.IsDir() {
				vals = append(vals, entry.Name())
			}
		}
		sort.Strings(vals)
		return carapace.ActionValues(vals...).Tag("repositories")
	})
}

func readPackageEntries(patterns ...string) []packageEntry {
	entries := make([]packageEntry, 0)
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			continue
		}

		for _, match := range matches {
			entries = append(entries, readPackageFile(match)...)
		}
	}
	return entries
}

func readPackageFile(path string) []packageEntry {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	entries := make([]packageEntry, 0)
	decoder := xml.NewDecoder(file)
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return entries
		}

		start, ok := token.(xml.StartElement)
		if !ok || start.Name.Local != "Package" {
			continue
		}

		var entry packageEntry
		if err := decoder.DecodeElement(&entry, &start); err == nil {
			entries = append(entries, entry)
		}
	}
	return entries
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value = normalizeWhitespace(value); value != "" {
			return value
		}
	}
	return ""
}

func normalizeWhitespace(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

func sortedKeys(values map[string]string) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
