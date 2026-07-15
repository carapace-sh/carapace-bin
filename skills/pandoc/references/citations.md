# Citations and Bibliographies

Pandoc's citation system: citation syntax, citeproc, CSL styles, bibliography formats, and metadata fields.

> **Source of truth**: <https://pandoc.org/MANUAL.html#citations>. For the citeproc filter pipeline, see [filters.md](filters.md). For Markdown citation syntax, see [markdown.md](markdown.md).

## Overview

Pandoc processes citations with the built-in `--citeproc` filter, which uses the [citeproc](https://github.com/jgm/citeproc) library. The workflow:

1. Reader parses `@key` citation syntax into `Cite` AST elements
2. `--citeproc` filter resolves citations against bibliography data
3. Filter renders citations in the specified CSL style and inserts a bibliography

```bash
pandoc --citeproc --bibliography=refs.bib --csl=chicago.csl input.md -o output.html
```

## Citation Syntax

Citations use the `@key` syntax (requires the `citations` extension, enabled by default in pandoc Markdown):

| Syntax | Description |
|--------|-------------|
| `[@key]` | Normal citation |
| `@key` | Author-in-text citation |
| `[@key; @key2]` | Multiple citations |
| `[see @key, p. 33]` | With prefix and locator |
| `[-@key]` | Suppress author, show only date |
| `[@key1 -@key2]` | Mixed: show key1 author, suppress key2 author |

### Locator Terms

| Term | Abbreviation | Meaning |
|------|--------------|---------|
| `book` | `bk.` | Book |
| `chapter` | `chap.` | Chapter |
| `column` | `col.` | Column |
| `figure` | `fig.` | Figure |
| `folio` | `fol.` | Folio |
| `number` | `no.` | Number |
| `line` | `l.` | Line |
| `note` | `n.` | Note |
| `opus` | `op.` | Opus |
| `page` | `p.` | Page |
| `paragraph` | `para.` | Paragraph |
| `part` | `pt.` | Part |
| `section` | `sec.` | Section |
| `sub verbo` | `sv.` | Sub verbo |
| `verse` | `v.` | Verse |
| `volume` | `vol.` | Volume |

## Bibliography Formats

Pandoc supports several bibliography formats:

| Format | Extensions | Description |
|--------|------------|-------------|
| BibLaTeX | `.bib` | BibLaTeX format |
| BibTeX | `.bibtex` | BibTeX format |
| CSL JSON | `.json` | CSL JSON |
| CSL YAML | `.yaml` | CSL YAML |
| RIS | `.ris` | RIS format |
| EndNote XML | `.xml` | EndNote |

Specify with `--bibliography`:

```bash
pandoc --citeproc --bibliography=refs.bib input.md
```

Multiple bibliography files can be combined:

```bash
pandoc --citeproc --bibliography=refs1.bib --bibliography=refs2.json input.md
```

### Inline Bibliography

Bibliographic data can be included inline in YAML metadata:

```yaml
---
references:
  - type: article-journal
    id: smith2020
    author:
      - family: Smith
        given: Jane
    title: "A Revolutionary Paper"
    issued:
      date-parts:
        - [2020, 6, 15]
    container-title: "Journal of Important Things"
    volume: 42
    page: 100-115
---
```

## Citation Styles (CSL)

Citation styles are defined by [CSL (Citation Style Language)](https://citationstyles.org/) files. Pandoc bundles a default style (Chicago author-date).

```bash
pandoc --citeproc --csl=ieee.csl input.md
```

CSL styles can be fetched from the [Zotero style repository](https://www.zotero.org/styles). Place `.csl` files in the `csl` subdirectory of the data directory or specify a full path.

### Journal Abbreviations

For styles that use abbreviated journal names, provide an abbreviations file:

```bash
pandoc --citeproc --csl=ieee.csl --citation-abbreviations=abbrs.json input.md
```

## Bibliography Placement

By default, the bibliography is placed at the end of the document. To control placement, insert a Div with ID `refs` where you want the bibliography:

```markdown
# References

::: {#refs}
:::
```

## Uncited References

To include items in the bibliography that are not cited in the text, use the `nocite` metadata field:

```yaml
---
nocite: |
  @key1, @key2
---
```

To include all references from the bibliography file (cited or not):

```yaml
---
nocite: |
  @*
---
```

## Note-Style Citations

In note styles (e.g., Chicago notes), citations are automatically converted to footnotes or endnotes. The `--citeproc` filter handles this transparently.

## Metadata Fields for Citations

| Field | Description |
|-------|-------------|
| `link-citations` | Hyperlink citation text to bibliography entry (default: true) |
| `link-bibliography` | Hyperlink bibliography entries to citations (default: true) |
| `lang` | Language for citation localization (BCP 47) |
| `notes-after-punctuation` | Place citation notes after punctuation (default: true) |
| `references` | Inline bibliography data (YAML) |
| `nocite` | Items to include without citing |

## LaTeX-Specific Citation Options

For LaTeX output, pandoc can use native citation packages instead of citeproc:

| Flag | Description |
|------|-------------|
| `--natbib` | Use natbib for citations |
| `--biblatex` | Use biblatex for citations |

When using natbib or biblatex, citation keys are passed through as-is, and the LaTeX engine resolves them.

## Related References

- [filters.md](filters.md) — The citeproc filter and pipeline
- [markdown.md](markdown.md) — Citation syntax in pandoc Markdown
- [cli-options.md](cli-options.md) — `--citeproc`, `--bibliography`, `--csl` flags
- [defaults.md](defaults.md) — Specifying citation options in defaults files
