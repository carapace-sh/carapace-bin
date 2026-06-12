# Viper Integration

How to integrate cobra with [viper](https://github.com/spf13/viper) for 12-factor application configuration.

## The Pattern

Cobra handles the CLI surface. Viper handles configuration from files, environment variables, and remote sources. The bridge is `BindPFlag` and `OnInitialize`.

## OnInitialize

Register a config initialization function that runs before every command:

```go
var cfgFile string

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.app.yaml)")
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        home, _ := os.UserHomeDir()
        viper.AddConfigPath(home)
        viper.SetConfigName(".app")
        viper.SetConfigType("yaml")
    }
    viper.AutomaticEnv()
    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
```

## Binding Flags to Viper

```go
rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name")
viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
```

After binding, `viper.GetString("author")` returns the flag value (if set) or the config file value (if set) or the default.

## Precedence Order

Viper resolves values in this order (highest wins):

1. **Flag value** (explicitly set on command line)
2. **Environment variable** (if `viper.AutomaticEnv()` is called)
3. **Config file value**
4. **Default value** (from `SetDefault` or flag default)

## Full Integration Example

```go
package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var cfgFile string
var author string

var rootCmd = &cobra.Command{
    Use:   "app",
    Short: "My application",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("author:", viper.GetString("author"))
    },
}

func init() {
    cobra.OnInitialize(initConfig)

    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
    rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name")
    viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
    viper.SetDefault("author", "unknown")
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        home, _ := os.UserHomeDir()
        viper.AddConfigPath(home)
        viper.SetConfigType("yaml")
        viper.SetConfigName(".app")
    }
    viper.AutomaticEnv()
    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}

func Execute() error {
    return rootCmd.Execute()
}
```

## Environment Variable Binding

`viper.AutomaticEnv()` makes viper check environment variables. By default, it maps config keys to env vars by uppercasing and replacing `.` with `_`:

```go
// Config key: "database.host"
// Env var:    DATABASE_HOST
```

Use `viper.BindEnv("database.host", "DB_HOST")` for custom mappings.

## Edge Cases

- **Bound variable vs viper.Get**: When you use `StringVar` with `BindPFlag`, the Go variable is set by pflag during flag parsing. But if the user doesn't provide the flag, the variable gets the pflag default — **not** the viper config value. Always use `viper.GetString()` to get the merged value.
- **Config file errors**: `ReadInConfig()` returns an error if the config file exists but can't be parsed. Check `err != nil` and handle appropriately.
- **OnInitialize timing**: `OnInitialize` functions run in `ExecuteC()` before command resolution. They run on every execution, even for completions and help.
- **Multiple config paths**: `AddConfigPath` can be called multiple times. Viper searches paths in reverse order (last added = first searched).

## References

- [viper documentation](https://github.com/spf13/viper)
- [cobra user guide: viper integration](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)

## Related Skills

- For `OnInitialize` internals, see the **cobra-dev** skill → `references/global-state.md`.
