---
name: carapace-scrape
description: >
  Use when the user wants to create a carapace user spec by scraping/generating it from the source code
  of a CLI tool. Covers supported frameworks, the patch-and-container approach, and how to run the scraper.
  Triggers on: "scrape spec", "generate spec from source", "carapace scrape", "spec scraper", or any request to generate completions by parsing source code.
user-invocable: true
---

# Carapace Spec Scraping Guide

Generate carapace user specs by injecting a spec generator into CLI tools with available source code.

## How Scraping Works

The scraping approach:

1. **Checkout** the project with the corresponding base container
2. **Install** necessary dependencies
3. **Apply a patch** that injects the carapace spec generator into the command
4. **Set the entrypoint** to invoke the command so the spec is output on execution
5. **Capture** the generated YAML spec

This requires the CLI tool to use a [supported framework](#supported-frameworks) and have accessible source code.

## Supported Frameworks

| Framework | Scraper | Language | Notes |
|-----------|---------|----------|-------|
| **cobra** | carapace | Go | `carapace.Gen(cmd)` call injected via patch |
| **urfave/cli v2** | carapace-spec-urfavecli | Go | Register hidden `_carapace spec` subcommand |
| **kingpin v2** | carapace-spec-kingpin | Go | Register hidden `_carapace spec` subcommand |
| **kong** | carapace-spec-kong | Go | Register hidden `_carapace spec` subcommand |
| **clap (Rust)** | carapace-spec-clap | Rust | Implement `Generator` trait |
| **oclif** | oclif-carapace-spec-plugin | TypeScript | Install plugin, run `carapace-gen` |

## Example: Docker (Cobra-based)

The [scrape project](https://github.com/carapace-sh/scrape) demonstrates this pattern:

### Dockerfile

```dockerfile
FROM golang
ARG VERSION

WORKDIR /
RUN git clone https://github.com/docker/cli --branch "${VERSION}" --depth 1

ADD docker.patch /

WORKDIR /cli
RUN git apply ../docker.patch
RUN cp vendor.mod go.mod
RUN cp vendor.sum go.sum
RUN go get github.com/carapace-sh/carapace
RUN go mod tidy
RUN go mod vendor
RUN GOBIN=/usr/bin/ go install ./cmd/docker

ENTRYPOINT ["docker", "_carapace", "spec"]
```

### Patch (docker.patch)

```diff
diff --git a/cmd/docker/docker.go b/cmd/docker/docker.go
--- a/cmd/docker/docker.go
+++ b/cmd/docker/docker.go
@@ -10,6 +10,7 @@
 	"strings"
 	"syscall"

+	"github.com/carapace-sh/carapace"
 	"github.com/containerd/errdefs"
 	"github.com/docker/cli/cli"
 	pluginmanager "github.com/docker/cli/cli-plugins/manager"
@@ -147,6 +148,7 @@ func newDockerCommand(dockerCli *command.DockerCli) *cli.TopLevelCommand {
 	cmd.SetIn(dockerCli.In())
 	cmd.SetOut(dockerCli.Out())
 	cmd.SetErr(dockerCli.Err())
+	carapace.Gen(cmd)

 	opts, helpCmd = cli.SetupRootCommand(cmd)
```

**Key steps in the patch:**
- Add the `carapace-sh/carapace` import
- Call `carapace.Gen(cmd)` on the root command to register the `_carapace spec` subcommand

### Run the Container

```bash
docker build -t scrape-docker -f scrapers/docker/Dockerfile --build-arg VERSION=v29.5.2 scrapers/docker/
docker run --rm scrape-docker > docker.yaml
```

The resulting `docker.yaml` contains the full completion spec.

## Go-based Frameworks

### urfave/cli

```go
import (
    "github.com/urfave/cli/v2"
    "github.com/carapace-sh/carapace-spec-urfavecli/spec"
)

func main() {
    app := &cli.App{}
    spec.Register(app)  // registers hidden "_carapace spec" subcommand
    app.Run(os.Args)
}
```

Run: `myapp _carapace spec`

### kingpin

```go
import (
    "github.com/alecthomas/kingpin/v2"
    "github.com/carapace-sh/carapace-spec-kingpin/spec"
)

func main() {
    app := kingpin.New("myapp", "My application")
    spec.Register(app)
    kingpin.Parse()
}
```

Run: `myapp _carapace spec`

### kong

```go
import (
    "github.com/alecthomas/kong"
    "github.com/carapace-sh/carapace-spec-kong/spec"
)

type Plugin struct {
    Carapace struct {
        Spec spec `cmd:"" name:"spec"`
    } `cmd:"" name:"_carapace" hidden:""`
}

func main() {
    ctx := kong.Parse(&CLI{}, &kong.Help{
        Short: true,
    }, kong.Bind(&Plugin{}))
    ctx.Handle()
}
```

Run: `myapp _carapace spec`

## Rust (clap)

Implement the `Generator` trait:

```rust
use carapace_spec_clap::Spec;
use clap::{Arg, Command};
use clap_complete::generate;
use std::io;

fn main() {
    let mut cmd = Command::new("example")
        .about("example command")
        .arg(Arg::new("flag").long("flag"));

    generate(Spec, &mut cmd, "example", &mut io::stdout());
}
```

Build and run to output the spec.

## oclif (TypeScript)

```bash
<CLI> plugins install @cristiand391/oclif-carapace-spec-plugin
<CLI> carapace-gen
```

This generates a spec file and provides instructions to source it.

## Scraping Workflow

For new tools, follow this pattern:

1. **Identify the framework** (cobra, urfave/cli, kingpin, kong, clap, oclif)
2. **Create the Dockerfile:**
   ```dockerfile
   FROM golang  # or appropriate base image
   ARG VERSION
   WORKDIR /
   RUN git clone <repo> --branch "${VERSION}" --depth 1
   ADD <tool>.patch /
   WORKDIR /<project>
   RUN git apply ../<tool>.patch
   RUN go get github.com/carapace-sh/carapace
   RUN GOBIN=/usr/bin/ go install ./cmd/<tool>
   ENTRYPOINT ["<tool>", "_carapace", "spec"]
   ```
3. **Create the patch** — inject `carapace.Gen(cmd)` (Go) or implement the generator trait (Rust)
4. **Build and run** the container, capturing stdout

## Output Location

Generated specs are typically saved to:

```
${UserConfigDir}/carapace/specs/<command>.yaml
```

| OS | Path |
|----|------|
| Linux | `~/.config/carapace/specs/` |
| macOS | `~/Library/Application Support/carapace/specs/` |
| Windows | `%APPDATA%\carapace\specs\` |

## Limitations

- **Source code required** — closed-source tools cannot be scraped this way
- **Supported frameworks only** — only frameworks with available scrapers are supported
- **Build dependencies** — the tool must be buildable (dependencies must resolve)
- **Plugin approach** — for some tools (like oclif), a plugin must be installed first

## Next Steps

To convert a scraped YAML spec to a native Go completer, use the `codegen` MCP tool (see **carapace-mcp** skill).

## Reference Repositories

- [scrape](https://github.com/carapace-sh/scrape) — Docker-based scraping for Docker CLI, kubectl, gh, etc.
- [carapace-spec-clap](https://github.com/carapace-sh/carapace-spec-clap) — Rust clap scraper
- [carapace-spec-kingpin](https://github.com/carapace-sh/carapace-spec-kingpin) — Go kingpin scraper
- [carapace-spec-kong](https://github.com/carapace-sh/carapace-spec-kong) — Go kong scraper
- [carapace-spec-urfavecli](https://github.com/carapace-sh/carapace-spec-urfavecli) — Go urfave/cli scraper
- [oclif-carapace-spec-plugin](https://github.com/cristiand391/oclif-carapace-spec-plugin) — oclif plugin