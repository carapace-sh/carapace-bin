# SSL/TLS Configuration

HTTPie's SSL/TLS options for certificate verification, protocol versions, ciphers, and client certificates.

## Certificate Verification (`--verify`)

| Value | Description |
|-------|-------------|
| `yes` (default) | Verify the host's SSL certificate against the system CA store |
| `no` (or `false`) | Skip certificate verification |
| `/path/to/ca_bundle` | Use a custom CA bundle file for verification |

```bash
# Skip verification
$ http --verify=no https://self-signed.example.org

# Use custom CA bundle
$ http --verify=/path/to/ca_bundle https://internal.example.org

# Default (verify)
$ https example.org
```

The carapace completer completes `--verify` with `yes` and `no`, styled with `style.ForKeyword` (green/red).

## Protocol Version (`--ssl`)

| Value | Protocol |
|-------|---------|
| `ssl2.3` | SSL 2.3 (default — negotiates highest supported) |
| `ssl3` | SSL v3 (insecure, for legacy servers) |
| `tls1` | TLS 1.0 |
| `tls1.1` | TLS 1.1 |
| `tls1.2` | TLS 1.2 |
| `tls1.3` | TLS 1.3 |

```bash
$ http --ssl=tls1.2 https://example.org
```

The carapace completer completes `--ssl` with `ssl2.3`, `tls1`, `tls1.1`, `tls1.2`. Note: `tls1.3` and `ssl3` are missing from the completer's values.

## Cipher Suites (`--ciphers`)

A string in the OpenSSL cipher list format:

```bash
$ http --ciphers=ECDHE-RSA-AES128-GCM-SHA256 https://example.org
```

Multiple ciphers are colon-separated. The carapace completer has a TODO for cipher completion.

## Client Certificates

| Flag | Description |
|------|-------------|
| `--cert` | Client-side SSL certificate file |
| `--cert-key` | Private key file (if not included in the cert file) |
| `--cert-key-pass` | Passphrase for the private key (auto-prompts if omitted) |

```bash
$ http --cert=client.pem --cert-key=client.key https://mtls.example.org
$ http --cert=client.pem --cert-key=client.key --cert-key-pass=secret https://mtls.example.org
```

If `--cert-key` is omitted, the private key is read from the certificate file (if combined).

The carapace completer completes `--cert` and `--cert-key` with `carapace.ActionFiles()`.

## Environment Variables

| Variable | Description |
|----------|-------------|
| `REQUESTS_CA_BUNDLE` | Path to CA bundle for certificate verification (overrides system default) |

## Edge Cases and Known Issues

- The carapace completer's `--ciphers` completion is a TODO (returns empty `ActionMultiParts`)
- The carapace completer is missing `ssl3` and `tls1.3` from the `--ssl` values
- `--cert-key-pass` is not present in the carapace completer's flag definitions

## References

- Source of truth: <https://httpie.io/docs/cli> (SSL options section)
- Carapace completer: `completers/common/http_completer/cmd/root.go`

## Related Skills

- [cli.md](cli.md) — full flag reference
- [config.md](config.md) — environment variable configuration
