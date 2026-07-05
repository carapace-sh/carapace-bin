# Services

`brew services` manages background services via `launchctl` (macOS) or `systemctl` (Linux). This document covers the `services` command, the `service do` DSL block in formulae, and the launchd/systemd integration.

## `brew services` Subcommands

| Subcommand | Purpose |
|------------|---------|
| `brew services list` | List all managed services |
| `brew services start <formula\|--all>` | Start and register service (auto-launch at login/boot) |
| `brew services stop <formula\|--all>` | Stop and unregister service |
| `brew services restart <formula\|--all>` | Stop (if needed) and start, register |
| `brew services run <formula\|--all>` | Run without registering (no auto-launch) |
| `brew services kill <formula\|--all>` | Stop but keep registered |
| `brew services info <formula\|--all>` | Show service info |
| `brew services cleanup` | Remove all unused services |

### Key Distinction: `start` vs `run`

| Command | Runs now? | Registers for auto-launch? |
|---------|----------|---------------------------|
| `start` | Yes | Yes |
| `run` | Yes | No |
| `stop` | No (stops) | No (unregisters) |
| `kill` | No (stops) | Yes (keeps registration) |
| `restart` | Yes | Yes |

### Flags

| Flag | Applies to | Purpose |
|------|-----------|---------|
| `--all` | start, stop, restart, run, kill, info | Operate on all services |
| `--file <path>` | start, restart, run | Use a custom service file |
| `--json` | list, info | JSON output |
| `--keep` | stop | Don't unregister from launchd/systemd |
| `--no-wait` | stop | Don't wait for service to stop |
| `--sudo-service-user <user>` | start (root/macOS) | Run service as specific user |

### User vs System Services

| Scope | macOS Location | Linux Location | Starts at |
|-------|---------------|----------------|-----------|
| **User** (no sudo) | `~/Library/LaunchAgents` | `~/.config/systemd/user` | Login |
| **System** (with sudo) | `/Library/LaunchDaemons` | `/usr/lib/systemd/system` | Boot |

Prepend `sudo` to operate system-wide:

```bash
sudo brew services start nginx      # system-wide (starts at boot)
brew services start nginx            # user-level (starts at login)
```

## launchctl vs systemctl Integration

| Aspect | macOS | Linux |
|--------|-------|-------|
| **Manager** | `launchctl` (via `launchd`) | `systemctl` (via `systemd`) |
| **Service files** | `.plist` files | `.service` files |
| **User scope** | `~/Library/LaunchAgents` | `~/.config/systemd/user` |
| **System scope** | `/Library/LaunchDaemons` | `/usr/lib/systemd/system` |

Homebrew generates `.plist` (macOS) or `.service` (Linux) files from the formula's `service do` block and places them in the appropriate location.

## The `service do` Block in Formulae

The `service` block in a formula defines background services. There are two approaches:

### 1. Package-Provided Service File

If the software already ships a `launchd` plist or `systemd` service file, install it into the prefix and reference it by name:

```ruby
service do
  name macos: "custom.launchd.name",
       linux: "custom.systemd.name"
end
```

Homebrew appends `.plist` (macOS) or `.service` (Linux) to the name to locate the file.

### 2. Formula-Defined Service File

Generate a service file inline:

```ruby
# Single command
service do
  run opt_bin/"script"
end

# Command with arguments
service do
  run [opt_bin/"script", "--config", etc/"config.yml"]
end

# OS-specific commands
service do
  run macos: [opt_bin/"macos_script", "standalone"],
      linux: var/"special_linux_script"
end
```

## Service Block Methods

The `run` or `name` field **must** be defined. If `name` is defined without `run`, Homebrew uses the package-provided file unchanged.

| Method | Default | macOS | Linux | Description |
|--------|---------|-------|-------|-------------|
| `run` | — | yes | yes | Command to execute (array with args or path) |
| `run_type` | `:immediate` | yes | yes | `:immediate`, `:interval`, or `:cron` |
| `interval` | — | yes | yes | Start interval in seconds (for `:interval`) |
| `cron` | — | yes | yes | Crontab syntax (for `:cron`) |
| `keep_alive` | `false` | yes | yes | Conditions for staying alive |
| `launch_only_once` | `false` | yes | yes | Run only once, no restart |
| `require_root` | `false` | yes | yes | Hint that service needs root |
| `environment_variables` | — | yes | yes | Hash of env vars |
| `working_dir` | — | yes | yes | Working directory |
| `root_dir` | — | yes | yes | Chroot directory |
| `input_path` | — | yes | yes | Input path |
| `log_path` | — | yes | yes | stdout log path |
| `error_log_path` | — | yes | yes | stderr log path |
| `restart_delay` | — | yes | yes | Seconds before restart |
| `throttle_interval` | — | yes | **no-op** | Min seconds between invocations (macOS default 10) |
| `process_type` | — | yes | **no-op** | `:background`, `:standard`, `:interactive`, `:adaptive` |
| `macos_legacy_timers` | — | yes | **no-op** | Don't coalesce launchd timers |
| `sockets` | — | yes | **no-op** | Socket definition: `<type>://<host>:<port>` |
| `nice` | — | yes | yes | Scheduling priority (-20 to 19; negative requires `require_root`) |
| `name` | — | yes | yes | Hash with `macos:` and/or `linux:` service name |

### `run_type`

| Type | Behavior |
|------|---------|
| `:immediate` (default) | Start immediately, keep alive |
| `:interval` | Run on repeating interval (requires `interval`) |
| `:cron` | Run at specific times (requires `cron` in crontab syntax) |

```ruby
service do
  run [opt_bin/"cleanup"]
  run_type :interval
  interval 500
end

service do
  run [opt_bin/"backup"]
  run_type :cron
  cron "5 * * * *"
end
```

### `keep_alive`

| Form | Behavior |
|------|---------|
| `keep_alive true` | Always keep alive |
| `keep_alive false` | Don't keep alive (run once) |
| `keep_alive always: true` | Same as `true` (hash form) |
| `keep_alive successful_exit: true` | Keep alive until non-zero exit |
| `keep_alive crashed: true` | Keep alive only if crashed |
| `keep_alive path: "/some/path"` | Keep alive while file exists |

### `environment_variables`

```ruby
service do
  run opt_bin/"beanstalkd"
  environment_variables PATH: std_service_path_env
end
```

The helper `std_service_path_env` returns `#{HOMEBREW_PREFIX}/bin:#{HOMEBREW_PREFIX}/sbin:/usr/bin:/bin:/usr/sbin:/sbin`.

### `sockets`

Format: `<type>://<host>:<port>` where type is `udp` or `tcp`.

```ruby
# Single socket (default name "listeners")
service do
  run [opt_bin/"server"]
  sockets "tcp://127.0.0.1:80"
end

# Multiple named sockets
service do
  run [opt_bin/"server"]
  sockets http: "tcp://0.0.0.0:80", https: "tcp://0.0.0.0:443"
end
```

> `sockets`, `throttle_interval`, `process_type`, and `macos_legacy_timers` are **no-op on Linux**.

## References

- [Formula Cookbook — Service files](https://docs.brew.sh/Formula-Cookbook#service-files)
- [Manpage — services subcommand](https://docs.brew.sh/Manpage#services-subcommand)
- For formula DSL: [formula.md](formula.md)
- For bundle service options: [bundle.md](bundle.md)
