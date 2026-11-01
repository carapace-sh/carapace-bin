# DNF5 Daemon (D-Bus API)

The `dnf5daemon-server` is a D-Bus system service providing a programmatic interface to the libdnf5 package manager. The `dnf5daemon-client` is a CLI client that communicates with the daemon.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/dnf_daemon/index.html>. For the libdnf5 library, see [libdnf5.md](libdnf5.md).

## Architecture

### Service Activation

- **systemd unit**: `dnf5daemon-server.service` (Type=dbus, runs as root)
- **D-Bus service file**: `/usr/share/dbus-1/system-services/org.rpm.dnf.v0.service`
- **D-Bus bus name**: `org.rpm.dnf.v0`
- **Object path**: `/org/rpm/dnf/v0`

The daemon auto-starts via D-Bus activation when a client calls `open_session()`.

### D-Bus Bus Configuration

Only root can own the bus name. Any client can send messages (polkit handles authorization):

```xml
<busconfig>
  <policy user="root">
    <allow own="org.rpm.dnf.v0"/>
  </policy>
  <policy context="default">
    <allow send_destination="org.rpm.dnf.v0"/>
  </policy>
</busconfig>
```

## Session Management

### SessionManager Interface

**Object path**: `/org/rpm/dnf/v0` (singleton)

| Method | Signature | Description |
|--------|-----------|-------------|
| `open_session` | `(a{sv} options) → (o session_object_path)` | Open a new session. Returns a new object path (e.g., `/org/rpm/dnf/v0/session/<id>`). Each session has its own thread and libdnf5 `Base` instance. |
| `close_session` | `(o session_object_path) → (b success)` | Close a session and release resources. |

### `open_session` Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `load_system_repo` | bool | true | Load system repository (RPMDB) |
| `load_available_repos` | bool | true | Load available repos from metadata |
| `config` | map{string:string} | — | Configuration overrides (whitelisted since CVE-2024-1929) |
| `releasever` | string | — | Override release version |
| `locale` | string | — | Set locale |

Unknown options are ignored. Maximum 3 sessions per user (since CVE-2024-1930 fix).

### Session Interfaces

Each session object provides:

| Interface | Purpose |
|-----------|---------|
| `org.rpm.dnf.v0.Base` | Repo loading, cache cleaning, session reset |
| `org.rpm.dnf.v0.rpm.Rpm` | Package operations (list, install, upgrade, remove, etc.) |
| `org.rpm.dnf.v0.rpm.Repo` | Repository operations (list, enable, disable, key confirmation) |
| `org.rpm.dnf.v0.Goal` | Transaction resolution and execution |
| `org.rpm.dnf.v0.Offline` | Offline transaction management |
| `org.rpm.dnf.v0.comps.Group` | Comps group operations |
| `org.rpm.dnf.v0.Advisory` | Security advisory operations |
| `org.rpm.dnf.v0.History` | Transaction history queries |

## D-Bus API Methods

### Base Interface

| Method | Description |
|--------|-------------|
| `read_all_repos()` | Explicitly load repository metadata. |
| `clean_with_options(cache_type, options)` | Clean cache. `cache_type`: "all", "packages", "metadata", "dbcache", "expire-cache". |
| `clean(cache_type)` | Equivalent to `clean_with_options` with empty options. |
| `reset()` | Completely reset the session. |

### Repo Interface

| Method | Description |
|--------|-------------|
| `list(options)` | List repos matching filters. Options: `patterns`, `repo_attrs`, `enable_disable` ("enabled"/"disabled"/"all"), `interactive`. |
| `confirm_key_with_options(key_id, confirmed, options)` | Confirm repo OpenPGP key import. |
| `confirm_key(key_id, confirmed)` | Equivalent with empty options. |
| `enable_with_options(repo_ids, options)` | Enable repositories. |
| `enable(repo_ids)` | Equivalent with empty options. |
| `disable_with_options(repo_ids, options)` | Disable repositories. |
| `disable(repo_ids)` | Equivalent with empty options. |

### Rpm Interface

| Method | Description |
|--------|-------------|
| `list(options)` | List packages matching filters. Returns list of attribute maps. |
| `list_fd(options, fd)` | Stream large result sets via file descriptor (JSON objects). Returns `transfer_id`. 30-second read timeout. |
| `install(pkg_specs, options)` | Mark packages for installation. Options: `repo_ids`, `skip_broken`, `skip_unavailable`. |
| `upgrade(pkg_specs, options)` | Mark packages for upgrade. Options: `repo_ids`. |
| `remove(pkg_specs, options)` | Mark packages for removal. |
| `distro_sync(pkg_specs, options)` | Sync packages to latest available. |
| `downgrade(pkg_specs, options)` | Mark packages for downgrade. |
| `reinstall(pkg_specs, options)` | Mark packages for reinstall. |
| `system_upgrade(options)` | Prepare distribution release upgrade. Options: `mode` ("distrosync"/"upgrade"), `interactive`. |

### Rpm.list Package Attributes

`name`, `epoch`, `version`, `release`, `arch`, `repo_id`, `from_repo_id`, `is_installed`, `install_size`, `download_size`, `buildtime`, `sourcerpm`, `summary`, `url`, `license`, `description`, `files`, `changelogs`, `provides`, `requires`, `requires_pre`, `conflicts`, `obsoletes`, `recommends`, `suggests`, `enhances`, `supplements`, `evr`, `nevra`, `full_nevra`, `reason`, `vendor`, `group`.

### Rpm.list Filters

`patterns`, `scope` ("all"/"installed"/"available"/"upgrades"/"upgradable"), `arch`, `repo`, `latest-limit`, `with_nevra`, `with_provides`, `with_filenames`, `with_binaries`, `with_src`, `icase`, `whatprovides`, `whatdepends`, `whatrequires`, `whatrecommends`, `whatenhances`, `whatsuggests`, `whatsupplements`, `whatobsoletes`, `whatconflicts`, `interactive`.

### Goal Interface

| Method | Description |
|--------|-------------|
| `resolve(options)` | Resolve transaction. Returns `(transaction_items, result)`. Options: `allow_erasing` (bool), `interactive` (bool). `result`: 0=no problem, 1=info/warnings, 2=failed. |
| `get_transaction_problems_string()` | Human-readable problems from resolution. |
| `get_transaction_problems()` | Structured problems with keys: `action`, `problem`, `goal_job_settings`, `spec`, `additional_data`, `solver_problems`. |
| `do_transaction(options)` | Execute resolved transaction. Options: `comment` (string), `offline` (bool), `interactive` (bool), `downloadonly` (bool, since 5.2.16). |
| `cancel()` | Cancel running transaction (only during download phase). |
| `reset()` | Reset prepared transaction for another operation. |

### Offline Interface

| Method | Description |
|--------|-------------|
| `get_status()` | Check for pending offline transaction. Returns `(pending, transaction_status)`. |
| `cancel_with_options(options)` | Cancel offline transaction. |
| `cancel()` | Equivalent with empty options. |
| `clean_with_options(options)` | Cancel and remove all offline transaction data. |
| `clean(options)` | Equivalent with empty options. |
| `set_finish_action_with_options(action, options)` | Set post-transaction action: "poweroff" or "reboot". |
| `set_finish_action(action)` | Equivalent with empty options. |
| `schedule_for_next_boot(options)` | Schedule offline transaction for next boot (creates `/system-update` symlink). |

### comps.Group Interface

| Method | Description |
|--------|-------------|
| `list(options)` | List groups. Options: `attributes`, `match_group_id`, `match_group_name`, `scope`, `with_hidden`, `patterns`, `contains_pkgs`, `lang`, `interactive`. |

### Advisory Interface

| Method | Description |
|--------|-------------|
| `list(options)` | List advisories. Options: `advisory_attrs`, `availability`, `names`, `types`, `contains_pkgs`, `severities`, `reference_bzs`, `reference_cves`, `with_bz`, `with_cve`, `interactive`. |

### History Interface

| Method | Description |
|--------|-------------|
| `recent_changes(options)` | Get recently changed packages. Returns dict with "installed", "removed", "upgraded", "downgraded". Options: `upgraded_packages`, `downgraded_packages`, `installed_packages`, `removed_packages`, `include_advisory`, `all_advisories`, `package_attrs`, `since` (unix timestamp), `interactive`. |

## Signals

### Download Signals (Base interface)

| Signal | Description |
|--------|-------------|
| `download_add_new` | A new download has started. |
| `download_progress` | Download progress update. |
| `download_mirror_failure` | Mirror failure during download. |
| `download_end` | Download ended. `transfer_status`: 0=success, 1=already exists, 2=error. |
| `repo_key_import_request` | Request for repository key import confirmation. |
| `repo_key_imported` | Informational signal after key import. |

### Transaction Signals (Rpm interface)

| Signal | Description |
|--------|-------------|
| `transaction_elem_progress` | Overall progress in transaction item processing. |
| `transaction_before_begin` | Right before the RPM transaction runs. |
| `transaction_after_complete` | Right after the RPM transaction finishes. |
| `transaction_action_start` | Package installation/removal started. |
| `transaction_action_progress` | Progress in package processing (max 1 signal per 400ms). |
| `transaction_action_stop` | Package processing finished. |
| `transaction_script_start` | RPM scriptlet started. |
| `transaction_script_stop` | RPM scriptlet finished successfully. |
| `transaction_script_error` | RPM scriptlet finished with error. |
| `transaction_verify_start` | Package files verification started. |
| `transaction_verify_progress` | Verification progress. |
| `transaction_verify_stop` | Verification finished. |
| `transaction_transaction_start` | Transaction package preparation started. |
| `transaction_transaction_progress` | Preparation progress. |
| `transaction_transaction_stop` | Preparation finished. |
| `transaction_unpack_error` | Error unpacking a package. |
| `write_to_fd_finished` | File descriptor transfer finished. |

## Polkit Authorization

### Policy Actions

Five polkit actions, all requiring `auth_admin`:

| Action ID | Description | Active Session |
|-----------|-------------|----------------|
| `org.rpm.dnf.v0.rpm.Repo.conf_write` | Write repo config file | `auth_admin_keep` |
| `org.rpm.dnf.v0.rpm.execute_trusted_transaction` | Execute transaction (signed, verified packages) | `auth_admin_keep` |
| `org.rpm.dnf.v0.rpm.execute_transaction` | Execute untrusted transaction (local/unverified) | `auth_admin_keep` |
| `org.rpm.dnf.v0.rpm.Repo.confirm_key` | Repository key import | `auth_admin` |
| `org.rpm.dnf.v0.base.Config.override` | Override config options | `auth_admin` |

`auth_admin_keep` caches the authorization for active sessions.

### Default Authorization Rule

Members of the `wheel` group get automatic authorization (no password) for trusted transactions and key confirmations:

```javascript
polkit.addRule(function(action, subject) {
    if ((action.id == "org.rpm.dnf.v0.rpm.execute_trusted_transaction" ||
         action.id == "org.rpm.dnf.v0.rpm.Repo.confirm_key") &&
        subject.active == true && subject.local == true &&
        subject.isInGroup("wheel")) {
            return polkit.Result.YES;
    }
});
```

### How Polkit is Invoked

DNF5 calls `CheckAuthorization()` on `org.freedesktop.PolicyKit1.Authority`. The `AllowUserInteraction` flag blocks the call while the user is prompted to authenticate. Hardcoded timeout: 2 minutes.

## Configuration

### Daemon Config File: `/etc/dnf/dnf5daemon-server.conf`

```ini
[main]
# Override any main DNF5 configuration option
cachedir = /var/cache/dnf5daemon-server
```

The `[main]` section can override any option from the main DNF5 configuration. See [configuration.md](configuration.md) for available options.

### Session-Level Overrides

Per-session config can be provided via the `config` map in `open_session()`. Since CVE-2024-1929, a whitelist enforces which options unprivileged clients can override.

## Client CLI

```
dnf5daemon-client [options] <command> [<args>...]
```

Commands: `advisory`, `clean`, `distro-sync`, `downgrade`, `group`, `install`, `reinstall`, `remove`, `repo`, `repolist`, `repoquery`, `system-upgrade`, `upgrade`.

**Note**: The client currently requires root for `install`, `remove`, `upgrade`, `downgrade`, `distro-sync`, `reinstall` (client-side `am_i_root()` check). This does not leverage the server-side polkit authorization.

## Example: Python D-Bus Usage

```python
import dbus

DNFDAEMON_BUS_NAME = 'org.rpm.dnf.v0'
DNFDAEMON_OBJECT_PATH = '/' + DNFDAEMON_BUS_NAME.replace('.', '/')

bus = dbus.SystemBus()
iface_session = dbus.Interface(
    bus.get_object(DNFDAEMON_BUS_NAME, DNFDAEMON_OBJECT_PATH),
    dbus_interface='org.rpm.dnf.v0.SessionManager')

# Open a session
session = iface_session.open_session(
    dbus.Dictionary({}, signature=dbus.Signature('sv')))

# Get Rpm interface
iface_rpm = dbus.Interface(
    bus.get_object(DNFDAEMON_BUS_NAME, session),
    dbus_interface='org.rpm.dnf.v0.rpm.Rpm')

# List upgrades
options = {"package_attrs": ["nevra", "repo_id"], "scope": "upgrades", "latest-limit": 1}
upgrades = iface_rpm.list(options)
for pkg in upgrades:
    print("{} (@{})".format(pkg["nevra"], pkg["repo_id"]))
```

### System Upgrade Example

```python
# Set releasever to the new distribution release
session = iface_session.open_session(
    dbus.Dictionary({"releasever": "40"}, signature=dbus.Signature('sv')))

iface_rpm = dbus.Interface(bus.get_object(DNFDAEMON_BUS_NAME, session),
    dbus_interface='org.rpm.dnf.v0.rpm.Rpm')
iface_goal = dbus.Interface(bus.get_object(DNFDAEMON_BUS_NAME, session),
    dbus_interface='org.rpm.dnf.v0.Goal')

# Add system upgrade
iface_rpm.system_upgrade({"mode": "distrosync"})

# Resolve
resolved, result = iface_goal.resolve({})
if result == 0:
    iface_goal.do_transaction({"offline": True}, timeout=2000)
else:
    errors = iface_goal.get_transaction_problems_string()
```
