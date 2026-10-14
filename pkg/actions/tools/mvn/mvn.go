package mvn

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	return fmt.Sprintf("%v/%v/%v/%v/%v-%v.jar", repository, strings.ReplaceAll(a.GroupId, ".", "/"), a.ArtifactId, a.Version, a.ArtifactId, a.Version)
}

type parent struct {
	GroupId      string `xml:"groupId"`
	ArtifactId   string `xml:"artifactId"`
	Version      string `xml:"version"`
	RelativePath string `xml:"relativePath"`
}

type profile struct {
	XMLName xml.Name   `xml:"profile"`
	Id      string     `xml:"id"`
	Plugins []artifact `xml:"build>plugins>plugin"`
}

type project struct {
	XMLName        xml.Name   `xml:"project"`
	Parent         parent     `xml:"parent"`
	Plugins        []artifact `xml:"build>plugins>plugin"`
	ManagedPlugins []artifact `xml:"build>pluginManagement>plugins>plugin"`
	Profiles       []profile  `xml:"profiles>profile"`
}

// profilesFromIds returns the list of profile ids for backward compatibility.
func (p project) profileIds() []string {
	ids := make([]string, len(p.Profiles))
	for i, profile := range p.Profiles {
		ids[i] = profile.Id
	}
	return ids
}

type settings struct {
	XMLName         xml.Name `xml:"settings"`
	LocalRepository string   `xml:"localRepository"`
}

func repositoryLocation() string {
	if repo := os.Getenv("MAVEN_REPO_LOCAL"); repo != "" {
		return repo
	}
	if home, err := os.UserHomeDir(); err == nil {
		repo := home + "/.m2/repository"
		if localRepo := localRepositoryFromSettings(home); localRepo != "" {
			if filepath.IsAbs(localRepo) {
				return localRepo
			}
			return filepath.Join(home, localRepo)
		}
		return repo
	}
	return ""
}

func localRepositoryFromSettings(home string) string {
	for _, path := range []string{
		home + "/.m2/settings.xml",
		globalSettingsPath(),
	} {
		if path == "" {
			continue
		}
		if content, err := os.ReadFile(path); err == nil {
			var s settings
			if xml.Unmarshal(content, &s) == nil && s.LocalRepository != "" {
				return s.LocalRepository
			}
		}
	}
	return ""
}

func globalSettingsPath() string {
	if mavenHome := os.Getenv("MAVEN_HOME"); mavenHome != "" {
		return mavenHome + "/conf/settings.xml"
	}
	if mavenHome := os.Getenv("M2_HOME"); mavenHome != "" {
		return mavenHome + "/conf/settings.xml"
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

// loadProjectWithParent loads a project and, if it has a parent with a
// relativePath, loads the parent POM and merges its plugins.
func loadProjectWithParent(file string) (*project, error) {
	loadedProject, err := loadProject(file)
	if err != nil {
		return nil, err
	}
	if loadedProject.Parent.RelativePath != "" {
		baseDir := filepath.Dir(locatePom(file))
		parentPom := filepath.Join(baseDir, loadedProject.Parent.RelativePath)
		if parentProject, err := loadProject(parentPom); err == nil && parentProject != nil {
			loadedProject.Plugins = append(parentProject.Plugins, loadedProject.Plugins...)
			loadedProject.ManagedPlugins = append(parentProject.ManagedPlugins, loadedProject.ManagedPlugins...)
		}
	}
	return loadedProject, nil
}

// allPlugins returns plugins from build>plugins, build>pluginManagement,
// and all profiles' build>plugins.
func (p project) allPlugins() []artifact {
	plugins := make([]artifact, 0, len(p.Plugins)+len(p.ManagedPlugins))
	plugins = append(plugins, p.Plugins...)
	plugins = append(plugins, p.ManagedPlugins...)
	for _, profile := range p.Profiles {
		plugins = append(plugins, profile.Plugins...)
	}
	return plugins
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
