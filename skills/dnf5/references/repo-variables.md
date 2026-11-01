# DNF5 Repo Variables

Repository variables available for substitution in DNF5 repository configuration files.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/dnf5.conf.5.html>.

## Built-in Variables

| Variable | Description | Overridable? |
|----------|-------------|--------------|
| `$arch` | System's CPU architecture (e.g., `aarch64`, `i586`, `i686`, `x86_64`) | No — cannot be overridden by variable files or env vars |
| `$basearch` | Base architecture. e.g., `i686`/`i586` -> `i386`; AMD64/Intel64 -> `x86_64` | No — cannot be overridden by variable files or env vars |
| `$releasever` | Release version of the OS, derived from RPMDB | Yes — via `--releasever` CLI option or `--setvar=releasever=VALUE` |
| `$releasever_major` | Major part of releasever (before first `.`) | Yes — via `--releasever-major` or auto-set from `--releasever` |
| `$releasever_minor` | Minor part of releasever (after first `.`) | Yes — via `--releasever-minor` or auto-set from `--releasever` |

## User-Defined Variables

### Via Variable Files

Place files in `/etc/dnf/vars/` (or `/usr/share/dnf5/vars.d/` for distribution). The filename is the variable name (lowercase, alphanumeric + underscores). The file content is the variable value.

```
# /etc/dnf//my_variable
my_value
```

Used as `$my_variable` in config files.

### Via Environment Variables

Set environment variables prefixed with `DNF_VAR_`:

```bash
export DNF_VAR_MY_VARIABLE=value
```

Used as `$MY_VARIABLE` in config files (without the prefix).

Variable name can only contain alphanumeric characters and underscores.

### Legacy Numbered Variables

`DNF0` through `DNF9` environment variables:

```bash
export DNF1=value
```

Used as `$DNF1` in config files.

### YUM Compatibility

Variables are also read from `/etc/yum/vars/` for YUM compatibility.

## Substitution

Variable substitution is performed on:
- Repository ID
- All repository configuration values (`baseurl`, `metalink`, `mirrorlist`, `gpgkey`, etc.)
- Repo paths from `--repofrompath`

## Configuration

| Option | Default | Description |
|--------|---------|-------------|
| `varsdir` | `/etc/dnf/vars` | Directories where variable definition files are looked for |

## Related

- [forcearch.md](forcearch.md) — `--forcearch` sets `arch` and `basearch` variables
- [installroot.md](installroot.md) — Variable resolution in installroot context
- [configuration.md](configuration.md) — `varsdir` option
