module github.com/rsteube/carapace-bin

go 1.12

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-ps v1.0.0
	github.com/pelletier/go-toml v1.2.0
	github.com/rsteube/carapace v0.1.6
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/traefik/yaegi v0.9.5
	gopkg.in/ini.v1 v1.60.2
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)

replace github.com/spf13/pflag => github.com/cornfeedhobo/pflag v1.0.0
