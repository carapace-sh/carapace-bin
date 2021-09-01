module github.com/rsteube/carapace-bin

go 1.17

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-ps v1.0.0
	github.com/pelletier/go-toml v1.9.3
	github.com/rsteube/carapace v0.8.4
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require github.com/inconshreveable/mousetrap v1.0.0 // indirect

replace github.com/spf13/pflag => github.com/cornfeedhobo/pflag v1.1.0
