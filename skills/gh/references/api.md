# The `gh api` Command

How to make authenticated HTTP requests to the GitHub REST and GraphQL APIs using `gh api`.

## Overview

`gh api` makes authenticated HTTP requests to GitHub API v3 (REST) or v4 (GraphQL) and prints the response. It handles authentication, pagination, type conversion, and output formatting automatically.

## Endpoint Argument

The endpoint argument is a path of a GitHub API v3 endpoint, or `graphql` for API v4.

```sh
$ gh api repos/{owner}/{repo}/releases
$ gh api repos/{owner}/{repo}/issues/123/comments -f body='Hi from CLI'
$ gh api graphql -f query='...'
```

### Placeholder Values

Placeholders in the endpoint are replaced with values from the current directory's repository or `GH_REPO`:

| Placeholder | Replaced with |
|-------------|---------------|
| `{owner}` | Repository owner |
| `{repo}` | Repository name |
| `{branch}` | Current branch |

> In some shells (e.g. PowerShell), enclose values containing `{...}` in quotes to prevent special meaning of curly braces.

## HTTP Method

The default method is `GET`, or `POST` if any parameters were added. Override with `--method`.

```sh
$ gh api -X GET search/issues -f q='repo:cli/cli is:open remote'
$ gh api -X DELETE repos/{owner}/{repo}/branches/protection/main
```

## Parameters

### `-f, --raw-field key=value`

Add a **string** parameter. Adding parameters automatically switches the method to `POST` (unless overridden with `--method GET`).

```sh
$ gh api repos/{owner}/{repo}/issues/123/comments -f body='Hi from CLI'
```

### `-F, --field key=value`

Add a **typed** parameter with magic type conversion:

| Value format | Conversion |
|-------------|------------|
| `true`, `false`, `null` | JSON boolean/null |
| Integer numbers | JSON number |
| `{owner}`, `{repo}`, `{branch}` | Populated from current repo context |
| `@<path>` | Read value from file |
| `@-` | Read value from stdin |

```sh
$ gh api gists -F 'files[myfile.txt][content]=@myfile.txt'
$ gh api repos/{owner}/{repo} -F private=true
```

### Nested Parameters

| Syntax | Meaning |
|--------|---------|
| `key[subkey]=value` | Nested object parameter |
| `key[]=value1` `key[]=value2` | Array parameter |
| `key[]` | Empty array |

```sh
$ gh api -X PATCH /orgs/{org}/properties/schema \
   -F 'properties[][property_name]=environment' \
   -F 'properties[][default_value]=production' \
   -F 'properties[][allowed_values][]=staging' \
   -F 'properties[][allowed_values][]=production'
```

### `--input file`

Read a pre-constructed request body from a file. Use `-` for stdin. When using `--input`, parameters from field flags are added to the query string of the endpoint URL.

```sh
$ gh api repos/{owner}/{repo}/rulesets --input file.json
$ cat payload.json | gh api repos/{owner}/{repo}/rulesets --input -
```

## Headers

### `-H, --header key:value`

Add an HTTP request header.

```sh
$ gh api -H 'Accept: application/vnd.github.v3.raw+json' repos/{owner}/{repo}/readme
```

### `--hostname string`

Specify the GitHub hostname (default `github.com`). Useful for GHES.

## API Previews

### `-p, --preview strings`

Opt into GitHub API previews. Names should omit the `-preview` suffix. The API expects opt-in via the `Accept` header with format `application/vnd.github.<preview-name>-preview+json`.

```sh
$ gh api --preview baptiste,nebula ...
$ gh api --preview corsair --preview scarlet-witch ...
```

## Pagination

### `--paginate`

Fetch all pages of results sequentially until there are no more. Each page is a separate JSON array or object.

For **GraphQL** requests, `--paginate` requires:
- The query accepts an `$endCursor: String` variable
- The query fetches `pageInfo { hasNextPage, endCursor }` from a collection

### `--slurp`

Use with `--paginate` to wrap all pages of JSON arrays or objects into an outer JSON array.

```sh
# REST: list all releases with pagination
$ gh api --paginate repos/{owner}/{repo}/releases

# GraphQL: list all repos with pagination
$ gh api graphql --paginate -f query='
  query($endCursor: String) {
    viewer {
      repositories(first: 100, after: $endCursor) {
        nodes { nameWithOwner }
        pageInfo { hasNextPage, endCursor }
      }
    }
  }
'

# GraphQL: slurp all pages into one array
$ gh api graphql --paginate --slurp -f query='...' | jq '...'
```

## Output Control

| Flag | Description |
|------|-------------|
| `-i, --include` | Include HTTP response status line and headers |
| `--silent` | Do not print the response body |
| `--verbose` | Include full HTTP request and response |
| `-q, --jq string` | Query to select values using jq syntax |
| `-t, --template string` | Format using a Go template |
| `--cache duration` | Cache the response (e.g. `3600s`, `60m`, `1h`) |

For `--jq` and `--template` usage, see [output-formatting.md](output-formatting.md).

## GraphQL

Use `graphql` as the endpoint for GraphQL queries. The `query` field is required. For GraphQL, all fields other than `query` and `operationName` are interpreted as GraphQL variables.

```sh
# Simple GraphQL query
$ gh api graphql -F owner='{owner}' -F name='{repo}' -f query='
  query($name: String!, $owner: String!) {
    repository(owner: $owner, name: $name) {
      releases(last: 3) { nodes { tagName } }
    }
  }
'

# Paginated GraphQL query
$ gh api graphql --paginate -f query='
  query($endCursor: String) {
    viewer {
      repositories(first: 100, after: $endCursor) {
        nodes { nameWithOwner }
        pageInfo { hasNextPage, endCursor }
      }
    }
  }
'
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `GH_TOKEN`, `GITHUB_TOKEN` | Auth token for `github.com` (in precedence order) |
| `GH_ENTERPRISE_TOKEN`, `GITHUB_ENTERPRISE_TOKEN` | Auth token for GHES |
| `GH_HOST` | GitHub host other than `github.com` |

## Common Patterns

### List and Filter

```sh
# List releases
$ gh api repos/{owner}/{repo}/releases

# List issues with jq filtering
$ gh api repos/{owner}/{repo}/issues --jq '.[].title'
```

### Create Resources

```sh
# Post an issue comment
$ gh api repos/{owner}/{repo}/issues/123/comments -f body='Hi from CLI'

# Create a label
$ gh api repos/{owner}/{repo}/labels -f name='bug' -f color='ff0000'
```

### Template Output

```sh
$ gh api repos/{owner}/{repo}/issues --template \
  '{{range .}}{{.title}} ({{.labels | pluck "name" | join ", " | color "yellow"}}){{"\n"}}{{end}}'
```

## Related

- For JSON output formatting (`--jq`, `--template`), see [output-formatting.md](output-formatting.md).
- For authentication and tokens, see [auth.md](auth.md).
- For environment variables, see [config.md](config.md).
