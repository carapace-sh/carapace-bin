package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
	"github.com/rsteube/carapace-bin/pkg/actions/net/ssh"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/os/usb"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/adb"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/apt"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/aws"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jaeger"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/journalctl"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kak"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pub"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tshark"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/virtualbox"
)

// TODO generate with go:generate
// TODO support Actions with arguments
var macros = map[string]func(string) carapace.Action{
	"fs.BlockDevices":                        func(s string) carapace.Action { return fs.ActionBlockDevices() },
	"fs.FileModes":                           func(s string) carapace.Action { return fs.ActionFileModes() },
	"fs.FileModesNumeric":                    func(s string) carapace.Action { return fs.ActionFileModesNumeric() },
	"fs.FileModesSymbolic":                   func(s string) carapace.Action { return fs.ActionFileModesSymbolic() },
	"fs.FilenameExtensions":                  func(s string) carapace.Action { return fs.ActionFilenameExtensions() },
	"fs.FilesystemTypes":                     func(s string) carapace.Action { return fs.ActionFilesystemTypes() },
	"fs.Labels":                              func(s string) carapace.Action { return fs.ActionLabels() },
	"fs.Mounts":                              func(s string) carapace.Action { return fs.ActionMounts() },
	"fs.PartLabels":                          func(s string) carapace.Action { return fs.ActionPartLabels() },
	"fs.PartUuids":                           func(s string) carapace.Action { return fs.ActionPartUuids() },
	"fs.Uuids":                               func(s string) carapace.Action { return fs.ActionUuids() },
	"net.Bssids":                             func(s string) carapace.Action { return net.ActionBssids() },
	"net.Connections":                        func(s string) carapace.Action { return net.ActionConnections() },
	"net.Hosts":                              func(s string) carapace.Action { return net.ActionHosts() },
	"net.Ipv4Addresses":                      func(s string) carapace.Action { return net.ActionIpv4Addresses() },
	"net.Ports":                              func(s string) carapace.Action { return net.ActionPorts() },
	"net.Ssids":                              func(s string) carapace.Action { return net.ActionSsids() },
	"net.Subnets":                            func(s string) carapace.Action { return net.ActionSubnets() },
	"net.http.CacheControlRequestDirectives": func(s string) carapace.Action { return http.ActionCacheControlRequestDirectives() },
	"net.http.ContentEncodingTokens":         func(s string) carapace.Action { return http.ActionContentEncodingTokens() },
	"net.http.RequestHeaderNames":            func(s string) carapace.Action { return http.ActionRequestHeaderNames() },
	"net.http.RequestHeaders":                func(s string) carapace.Action { return http.ActionRequestHeaders() },
	"net.http.Tags":                          func(s string) carapace.Action { return http.ActionTags() },
	"net.http.MediaTypes":                    func(s string) carapace.Action { return http.ActionMediaTypes() },
	"net.http.RequestMethods":                func(s string) carapace.Action { return http.ActionRequestMethods() },
	"net.http.StatusCodes":                   func(s string) carapace.Action { return http.ActionStatusCodes() },
	"net.http.TransferEncodingTokens":        func(s string) carapace.Action { return http.ActionTransferEncodingTokens() },
	"net.http.UserAgents":                    func(s string) carapace.Action { return http.ActionUserAgents() },
	"os.Cgroups":                             func(s string) carapace.Action { return os.ActionCgroups() },
	"os.Displays":                            func(s string) carapace.Action { return os.ActionDisplays() },
	"os.EnvironmentVariables":                func(s string) carapace.Action { return os.ActionEnvironmentVariables() },
	"os.GpgKeyIds":                           func(s string) carapace.Action { return os.ActionGpgKeyIds() },
	"os.Groups":                              func(s string) carapace.Action { return os.ActionGroups() },
	"os.HexColors":                           func(s string) carapace.Action { return os.ActionHexColors() },
	"os.KernelModulesLoaded":                 func(s string) carapace.Action { return os.ActionKernelModulesLoaded() },
	"os.KillSignals":                         func(s string) carapace.Action { return os.ActionKillSignals() },
	"os.Languages":                           func(s string) carapace.Action { return os.ActionLanguages() },
	"os.Locales":                             func(s string) carapace.Action { return os.ActionLocales() },
	"os.MouseButtons":                        func(s string) carapace.Action { return os.ActionMouseButtons() },
	"os.PathExecutables":                     func(s string) carapace.Action { return os.ActionPathExecutables() },
	"os.ProcessExecutables":                  func(s string) carapace.Action { return os.ActionProcessExecutables() },
	"os.ProcessIds":                          func(s string) carapace.Action { return os.ActionProcessIds() },
	"os.ProcessStates":                       func(s string) carapace.Action { return os.ActionProcessStates() },
	"os.SessionIds":                          func(s string) carapace.Action { return os.ActionSessionIds() },
	"os.Shells":                              func(s string) carapace.Action { return os.ActionShells() },
	"os.SoundCards":                          func(s string) carapace.Action { return os.ActionSoundCards() },
	"os.Terminals":                           func(s string) carapace.Action { return os.ActionTerminals() },
	"os.UserGroup":                           func(s string) carapace.Action { return os.ActionUserGroup() },
	"os.Users":                               func(s string) carapace.Action { return os.ActionUsers() },
	"os.XtermColorNames":                     func(s string) carapace.Action { return os.ActionXtermColorNames() },
	"os.usb.DeviceNumbers":                   func(s string) carapace.Action { return usb.ActionDeviceNumbers() },
	"os.usb.ProductNumbers":                  func(s string) carapace.Action { return usb.ActionProductNumbers() },
	"ssh.Ciphers":                            func(s string) carapace.Action { return ssh.ActionCiphers() },
	"ssh.HostKeyAlgorithms":                  func(s string) carapace.Action { return ssh.ActionHostKeyAlgorithms() },
	"ssh.Options":                            func(s string) carapace.Action { return ssh.ActionOptions() },
	"time.Date":                              func(s string) carapace.Action { return time.ActionDate() },
	"time.DateTime":                          func(s string) carapace.Action { return time.ActionDateTime() },
	"time.Months":                            func(s string) carapace.Action { return time.ActionMonths() },
	"time.Time":                              func(s string) carapace.Action { return time.ActionTime() },
	"time.TimeS":                             func(s string) carapace.Action { return time.ActionTimeS() },
	"tools.adb.Packages":                     func(s string) carapace.Action { return adb.ActionPackages() },
	"tools.adb.Users":                        func(s string) carapace.Action { return adb.ActionUsers() },
	"tools.apt.PackageSearch":                func(s string) carapace.Action { return apt.ActionPackageSearch() },
	"tools.apt.Packages":                     func(s string) carapace.Action { return apt.ActionPackages() },
	"tools.aws.Profiles":                     func(s string) carapace.Action { return aws.ActionProfiles() },
	"tools.aws.Regions":                      func(s string) carapace.Action { return aws.ActionRegions() },
	"tools.docker.Configs":                   func(s string) carapace.Action { return docker.ActionConfigs() },
	"tools.docker.ContainerIds":              func(s string) carapace.Action { return docker.ActionContainerIds() },
	"tools.docker.ContainerPath":             func(s string) carapace.Action { return docker.ActionContainerPath() },
	"tools.docker.Containers":                func(s string) carapace.Action { return docker.ActionContainers() },
	"tools.docker.Contexts":                  func(s string) carapace.Action { return docker.ActionContexts() },
	"tools.docker.DetachKeys":                func(s string) carapace.Action { return docker.ActionDetachKeys() },
	"tools.docker.LogDrivers":                func(s string) carapace.Action { return docker.ActionLogDrivers() },
	"tools.docker.Networks":                  func(s string) carapace.Action { return docker.ActionNetworks() },
	"tools.docker.Nodes":                     func(s string) carapace.Action { return docker.ActionNodes() },
	"tools.docker.Plugins":                   func(s string) carapace.Action { return docker.ActionPlugins() },
	"tools.docker.Ports":                     func(s string) carapace.Action { return docker.ActionPorts() },
	"tools.docker.Repositories":              func(s string) carapace.Action { return docker.ActionRepositories() },
	"tools.docker.RepositoryTags":            func(s string) carapace.Action { return docker.ActionRepositoryTags() },
	"tools.docker.Secrets":                   func(s string) carapace.Action { return docker.ActionSecrets() },
	"tools.docker.Services":                  func(s string) carapace.Action { return docker.ActionServices() },
	"tools.docker.Volumes":                   func(s string) carapace.Action { return docker.ActionVolumes() },
	"tools.gh.ConfigHosts":                   func(s string) carapace.Action { return gh.ActionConfigHosts() },
	"tools.gh.OwnerRepositories":             func(s string) carapace.Action { return gh.ActionOwnerRepositories() },
	"tools.git.Authors":                      func(s string) carapace.Action { return git.ActionAuthors() },
	"tools.git.CleanupMode":                  func(s string) carapace.Action { return git.ActionCleanupMode() },
	"tools.git.ColorConfigs":                 func(s string) carapace.Action { return git.ActionColorConfigs() },
	"tools.git.Colors":                       func(s string) carapace.Action { return git.ActionColors() },
	"tools.git.ConfigTypes":                  func(s string) carapace.Action { return git.ActionConfigTypes() },
	"tools.git.Configs":                      func(s string) carapace.Action { return git.ActionConfigs() },
	"tools.git.CurrentBranch":                func(s string) carapace.Action { return git.ActionCurrentBranch() },
	"tools.git.FieldNames":                   func(s string) carapace.Action { return git.ActionFieldNames() },
	"tools.git.MergeStrategy":                func(s string) carapace.Action { return git.ActionMergeStrategy() },
	"tools.git.Remotes":                      func(s string) carapace.Action { return git.ActionRemotes() },
	"tools.git.RepositorySearch":             func(s string) carapace.Action { return git.ActionRepositorySearch() },
	"tools.git.Stashes":                      func(s string) carapace.Action { return git.ActionStashes() },
	"tools.git.SubmoduleNames":               func(s string) carapace.Action { return git.ActionSubmoduleNames() },
	"tools.git.SubmodulePaths":               func(s string) carapace.Action { return git.ActionSubmodulePaths() },
	"tools.git.TextAttributes":               func(s string) carapace.Action { return git.ActionTextAttributes() },
	"tools.git.WhitespaceModes":              func(s string) carapace.Action { return git.ActionWhitespaceModes() },
	"tools.golang.BuildTags":                 func(s string) carapace.Action { return golang.ActionBuildTags() },
	"tools.jaeger.SamplingTypes":             func(s string) carapace.Action { return jaeger.ActionSamplingTypes() },
	"tools.journalctl.Outputs":               func(s string) carapace.Action { return journalctl.ActionOutputs() },
	"tools.kak.Sessions":                     func(s string) carapace.Action { return kak.ActionSessions() },
	"tools.pacman.PackageGroups":             func(s string) carapace.Action { return pacman.ActionPackageGroups() },
	"tools.pacman.Repositories":              func(s string) carapace.Action { return pacman.ActionRepositories() },
	"tools.pub.Dependencies":                 func(s string) carapace.Action { return pub.ActionDependencies() },
	"tools.pub.PackageSearch":                func(s string) carapace.Action { return pub.ActionPackageSearch() },
	"tools.tshark.FileTypes":                 func(s string) carapace.Action { return tshark.ActionFileTypes() },
	"tools.tshark.Interfaces":                func(s string) carapace.Action { return tshark.ActionInterfaces() },
	"tools.tshark.Protocols":                 func(s string) carapace.Action { return tshark.ActionProtocols() },
	"tools.tshark.ReadFormats":               func(s string) carapace.Action { return tshark.ActionReadFormats() },
	"tools.tshark.Selectors":                 func(s string) carapace.Action { return tshark.ActionSelectors() },
	"tools.tshark.Statistics":                func(s string) carapace.Action { return tshark.ActionStatistics() },
	"tools.virtualbox.Machines":              func(s string) carapace.Action { return virtualbox.ActionMachines() },
}
