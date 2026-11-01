# Setup

Installation, server deployment, certificates, and configuration.

> **Source of truth**: <https://epicgames.github.io/lore/how-to/install-lore-cli/> and <https://epicgames.github.io/lore/how-to/deploy-local-lore-server/>

## Install the Lore CLI

### Prebuilt Binary (Recommended)

```bash
# macOS / Linux
curl -fsSL https://raw.githubusercontent.com/EpicGames/lore/main/scripts/install.sh | bash

# Windows (PowerShell)
irm https://raw.githubusercontent.com/EpicGames/lore/main/scripts/install.ps1 | iex
```

Open a new terminal session for PATH changes. Verify with `lore --version`.

### Build from Source

```bash
cargo build --release -p lore-client --bin lore
sudo cp target/release/lore /usr/local/bin/lore   # or ~/bin/
```

### Shell Completions

```bash
lore completions zsh ~/.zsh/completions
lore completions bash > ~/.local/share/bash-completion/completions/lore
lore completions powershell | Out-String | Invoke-Expression
```

Supported shells: `bash`, `zsh`, `fish`, `powershell`, `elvish`. If you omit the output path, the completion script is printed to stdout.

## Start a Local Server (Demo Mode)

Runs a throwaway server with ephemeral self-signed certificate, temporary store, and no auth.

```bash
# macOS / Linux — install CLI + server + start demo
curl -fsSL https://raw.githubusercontent.com/EpicGames/lore/main/scripts/install.sh | bash -s -- --demo

# Windows
$env:LORE_DEMO=1; irm https://raw.githubusercontent.com/EpicGames/lore/main/scripts/install.ps1 | iex
```

The server listens on ports **41337** (QUIC/UDP + gRPC/TCP) and **41339** (HTTP).

## Deploy a Persistent Server

### From Binary

```bash
# Install the server binary
curl -fsSL https://raw.githubusercontent.com/EpicGames/lore/main/scripts/install.sh | bash -s -- --server
```

### From Docker

```bash
docker build --platform linux/amd64 -f lore-server/Dockerfile -t lore-server .
docker run -d --name lore-server \
    -p 41337:41337/tcp -p 41337:41337/udp -p 41339:41339 \
    lore-server
```

> Apple Silicon requires `--platform linux/amd64` — the `linux/arm64` image targets AWS Graviton3 (SVE).

### Ports

| Port | Protocol | Purpose |
|------|----------|---------|
| 41337 | TCP (gRPC) + UDP (QUIC) | Primary server communication |
| 41339 | TCP | HTTP (health checks) |

### Health Check

```bash
curl -i http://127.0.0.1:41339/health_check
# → HTTP/1.1 200 OK
```

### Certificates

Generate a self-signed certificate for persistent deployments:

```bash
openssl req -x509 -newkey rsa:2048 -nodes \
    -keyout key.pem -out cert.pem -days 365 \
    -subj "/CN=localhost" -addext "subjectAltName=IP:127.0.0.1,DNS:localhost"
```

### Durable Storage

Both stores need a persistent path:

```toml
[immutable_store.local]
path = "/opt/loreserver/store"
flush_delay_seconds = 10

[mutable_store.local]
path = "/opt/loreserver/store"
flush_delay_seconds = 10
```

### Minimal Persistent Config

Place in `local.toml` in the config directory:

```toml
[server.quic.certificate]
cert_file = "/opt/loreserver/certs/cert.pem"
pkey_file = "/opt/loreserver/certs/key.pem"

[immutable_store.local]
path = "/opt/loreserver/store"
flush_delay_seconds = 10

[mutable_store.local]
path = "/opt/loreserver/store"
flush_delay_seconds = 10
```

Start with: `loreserver --config /opt/loreserver/config`

### Docker Persistent Setup

```bash
docker run -d --name lore-server \
    -p 41337:41337/tcp -p 41337:41337/udp -p 41339:41339 \
    -v "$PWD/cert.pem:/etc/lore/cert.pem:ro" \
    -v "$PWD/key.pem:/etc/lore/key.pem:ro" \
    -v "$PWD/local.toml:/etc/lore/config/local.toml:ro" \
    -v ~/lore-data:/data \
    lore-server
```

## Recommended Directory Layout

```
/opt/loreserver/
├── config/
│   └── local.toml
├── certs/
│   ├── cert.pem
│   └── key.pem
└── store/
```