# Configuration

CLI configuration (`config.toml`, `cli.toml`) and server configuration (layered TOML files, environment variables).

> **Source of truth**: <https://epicgames.github.io/lore/reference/lore-cli-config/> and <https://epicgames.github.io/lore/reference/lore-server-config/>

## CLI Configuration

### On-Disk Locations

| File | Location |
|------|----------|
| Per-repo | `<repo>/.lore/config.toml` |
| User-level (Linux) | `~/.config/lore/cli.toml` or `$XDG_CONFIG_HOME/lore/cli.toml` |
| User-level (macOS) | `~/Library/Application Support/com.epicgames.lore/cli.toml` |
| User-level (Windows) | `%LOCALAPPDATA%\Epic Games\lore\config\cli.toml` |

### Per-Repository `config.toml`

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `remote_url` | string | none | URL of the remote (e.g. `lore://127.0.0.1:41337/my-project`) |
| `identity` | string | none | Commit identity. Resolved at create/clone: `--identity` flag → server connection → unset |

#### `[store]` Table

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `max_capacity` | integer (bytes) | 10485760 (10 MiB) | Max local store capacity before eviction |
| `eviction_delay` | integer (seconds) | 10 | Delay before evicting over-capacity fragments |
| `max_size` | integer (bytes) | 10737418240 (10 GiB) | Max total store size before compaction |
| `compaction_delay` | integer (seconds) | 30 | Delay between background compaction passes |
| `verify_write` | Boolean | unset (false) | Verify writes by re-reading and re-hashing |

#### `[file]` Table

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `direct_write` | Boolean | false | Write directly to target files (non-atomic) |
| `flush_write` | Boolean | false | Flush file data to disk after each write (not yet wired) |

#### `[shared_store_to_use]` Table

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `use_shared_store` | Boolean | unset | Whether repo uses a shared store |
| `shared_store_path` | string | unset | Filesystem path of the shared store |

Legacy aliases: `global_store_to_use`, `use_global_store`, `global_store_path`.

### User-Level `cli.toml`

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `pager` | string | `less -R` (Unix), `more.com` (Windows) | Pager program for long output |

This is the only field read from `cli.toml`. All other behavioral settings are per-invocation CLI flags.

### Example Configs

```toml
# Minimal config.toml
remote_url = "lore://127.0.0.1:41337/my-project"
identity = "alex@example.com"

# Custom store limits
[store]
max_capacity = 1073741824   # 1 GiB
eviction_delay = 30

# Shared store
[shared_store_to_use]
use_shared_store = true
shared_store_path = "/srv/lore/shared-store"
```

```toml
# User-level cli.toml
pager = "bat --paging=always"
```

## Server Configuration

### CLI Flags

| Flag | Env Var | Default | Description |
|------|---------|---------|-------------|
| `--config <DIR>` | `LORE_CONFIG_PATH` | `lore-server/config` | Directory of TOML config files |
| `--env <ENV>` | `LORE_ENV` | `local` | Environment name selecting `<environment>.toml` |

### Config-File Layering (4 layers, merged field-by-field)

| Order | Source | Description |
|-------|--------|-------------|
| 1 | `default.toml` | Baked into the binary at compile time |
| 2 | `<environment>.toml` | Named by `--env` / `LORE_ENV` (default: `local`) |
| 3 | `local.toml` | Local overrides, applied last among file layers |
| 4 | `LORE__`-prefixed env vars | Override every file layer. `__` is nested key separator |

Missing files/directories are silently skipped. Arrays are replaced wholesale by the last source. Only scalar fields can be set via environment variables.

### Server Settings (`[server]`)

| Field | Default | Description |
|-------|---------|-------------|
| `connection_close_timeout_seconds` | 5 | Wait for open connections after shutdown |
| `runtime_shutdown_timeout_seconds` | 25 | Wait for async runtime shutdown |

#### QUIC Endpoints (`[server.quic]`, `[server.quic_internal]`)

| Field | Default (quic) | Default (internal) | Description |
|-------|----------------|--------------------|-------------|
| `enabled` | true | false | Enable endpoint |
| `host` | 0.0.0.0 | 0.0.0.0 | Bind address |
| `port` | 41337 | 41340 | Listen port |
| `verify_client_certs` | false | true | mTLS client certs |
| `idle_timeout` | 30000 | 30000 | Connection idle timeout (ms) |
| `keep_alive` | 500 | 500 | Keep-alive interval (ms) |
| `max_bidi_streams` | 8 | 8 | Max concurrent bidirectional streams |
| `num_listeners` | 10 | 10 | Number of listener tasks |
| `transport_bits_per_second` | 1073741824 (1 Gbps) | 10737418240 (10 Gbps) | Bandwidth estimate |
| `transport_rtt` | 100 | 100 | Expected RTT (ms) |
| `handler_timeout_seconds` | 50 | 50 | Per-request handler timeout |

**Certificate block** (`[server.quic.certificate]`): `cert_file` (required), `pkey_file` (required), `cert_chain` (optional).

#### HTTP Endpoint (`[server.http]`)

| Field | Default | Description |
|-------|---------|-------------|
| `enabled` | true | Enable endpoint |
| `host` | 0.0.0.0 | Bind address |
| `port` | 41339 | Listen port |
| `max_file_size` | 10485760 (10 MiB) | Max upload size |
| `request_timeout_seconds` | 300 | Overall request timeout |
| `request_body_timeout_seconds` | 3600 | Body read timeout |
| `presigned_url_hmac_key` | none | Hex HMAC key (32+ bytes) enabling presigned URLs |
| `presigned_url_min_ttl_seconds` | 1 | Min presigned URL lifetime |
| `presigned_url_default_ttl_seconds` | 3600 | Default presigned URL lifetime |
| `presigned_url_max_ttl_seconds` | 86400 (24h) | Max presigned URL lifetime |

#### gRPC Endpoints (`[server.grpc]`, `[server.grpc_internal]`)

| Field | Default (grpc) | Default (internal) | Description |
|-------|----------------|--------------------|-------------|
| `host` | 0.0.0.0 | 0.0.0.0 | Bind address |
| `port` | 41337 | 41340 | Listen port |
| `request_handler_timeout_seconds` | 50 | 50 | Per-request handler timeout |
| `certificate` | none | none | Optional cert block |

Internal only: `enabled` (default false), `verify_client_certs` (default true).

#### Authentication (`[server.auth]`)

| Field | Description |
|-------|-------------|
| `jwt_issuer` | Expected JWT `iss` claim |
| `jwt_audience` | Array of accepted JWT `aud` values |
| `[server.auth.jwk]` | Sub-table with `endpoint` (JWKS URL). Presence enables JWT verification |

When absent, JWT verification is disabled.

### Store Backends

| Table | Mode | Options |
|-------|------|---------|
| `[immutable_store]` | `local` (default) | `local`, `composite`, `replicated`, `remote`, or plugin |
| `[mutable_store]` | `local` (default) | `local`, `remote`, or plugin |
| `[lock_store]` | `local` (default) | `local` or plugin |

**Plugins are compiled into the server binary, not loaded at runtime.** The stock `loreserver` registers no plugins. Selecting a plugin backend fails with `PluginNotFound` unless using a custom binary.

### Topology (`[topology]`)

| Provider | Description |
|----------|-------------|
| `none` (default) | Single-node, no peers |
| `fixed` | Static peer list (`[topology.fixed]` with `peers` array) |
| `rotating_id_fixed` | Static peer list with rotating IDs |
| `composite` | Merges peers from multiple sources |
| `consul` | Dynamic discovery via Consul (plugin) |

### Hooks (`[hooks.<name>]`)

**Fire points**: `BranchPush`, `BranchCreate`, `BranchDelete`, `RepositoryCreate`, `Obliterate`.

Each hook has `enabled` (default false) plus hook-specific fields. The base `loreserver` registers no hooks — custom hooks are wired in at compile time via Rust traits.

### Telemetry (`[telemetry]`)

| Subsection | Key fields |
|------------|------------|
| `logger` | `format` (ansi/json/text), `output` (stdout/stderr/file), `enable_otlp` |
| `exporter` | `endpoint`, `queue_size`, `timeout` (all required if table present) |
| `metrics` | `export_interval_millis` (30000), `sample_interval_millis` (10000) |
| `traces` | `sample_rate` (0.05), `sample_rate_low_tier` (0.001), `service_name` |
| `additional_labels` | Extra key-value labels |