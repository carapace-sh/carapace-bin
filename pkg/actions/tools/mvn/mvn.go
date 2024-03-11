package mvn

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/carapace-sh/carapace/pkg/util"
)

type plugin struct {
	XMLName    xml.Name `xml:"plugin"`
	GoalPrefix string   `xml:"goalPrefix"`
	Mojos      []struct {
		Goal        string `xml:"goal"`
		Description string `xml:"description"`
	} `xml:"mojos>mojo"`
}

func (p plugin) FormattedGoals() map[string]string {
	goals := make(map[string]string)
	for _, mojo := range p.Mojos {
		goal := fmt.Sprintf("%v:%v", p.GoalPrefix, mojo.Goal)
		description := strings.SplitAfter(mojo.Description, ".")[0]
		if len(description) > 60 {
			description = description[:57] + "..."
		}
		goals[goal] = description
	}
	return goals
}

type artifact struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
}

func (a artifact) Location(repository string) string {
	return fmt.Sprintf("%v/%v/%v/%v/%v-%v.jar", repository, strings.Replace(a.GroupId, ".", "/", -1), a.ArtifactId, a.Version, a.ArtifactId, a.Version)
}

type project struct {
	// TODO parent pom plugins
	// TODO plugins locatad in pluginmanagement and profiles
	XMLName  xml.Name   `xml:"project"`
	Plugins  []artifact `xml:"build>plugins>plugin"`
	Profiles []string   `xml:"profiles>profile>id"`
}

func repositoryLocation() string {
	// TODO environment variable / settings override
	if home, err := os.UserHomeDir(); err == nil {
		return home + "/.m2/repository"
	}
	return ""
}

func locatePom(file string) (pom string) {
	if file != "" {
		return file
	}
	pom, _ = util.FindReverse("", "pom.xml")
	return
}

func loadProject(file string) (loadedProject *project, err error) {
	var content []byte
	if content, err = os.ReadFile(locatePom(file)); err == nil {
		err = xml.Unmarshal(content, &loadedProject)
	}
	return
}

func loadPlugin(file string) (loadedPlugin *plugin) {
	if reader, err := zip.OpenReader(file); err == nil {
		defer reader.Close()
		for _, f := range reader.File {
			if f.Name == "META-INF/maven/plugin.xml" {
				if pluginFile, err := f.Open(); err == nil {
					defer pluginFile.Close()
					if content, err := io.ReadAll(pluginFile); err == nil {
						_ = xml.Unmarshal(content, &loadedPlugin)
					}
				}
			}
		}
	}
	return
}
