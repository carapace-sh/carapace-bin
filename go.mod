module github.com/rsteube/carapace-bin

go 1.15

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-ps v1.0.0
	github.com/pelletier/go-toml v1.8.1
	github.com/rsteube/carapace v0.6.1
	github.com/spf13/cobra v1.1.1
	golang.org/x/sys v0.0.0-20210426230700-d19ff857e887
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)

replace github.com/spf13/pflag => github.com/cornfeedhobo/pflag v1.1.0
