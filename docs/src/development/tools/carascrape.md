# carascrape

[carascrape](https://github.com/rsteube/carapace-bin/tree/master/pkg/carascrape) is a scraper for projects that use [spf13/cobra](https://github.com/spf13/cobra).
For most cases it is best to add a `Run` function to the root command. Then build and run the project. Generated source files are written to `/tmp/carascrape##########/`.


```go
import (
    "github.com/rsteube/carapace-bin/pkg/carascrape"
)

rootCmd := &cobra.Command{
    Run: func(cmd *cobra.Command, args []string) {
        carascrape.Scrape(cmd)
    },
}
```
