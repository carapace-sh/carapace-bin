# Specialized Output Formats

Pandoc's specialized output formats: slide shows, EPUBs, chunked HTML, Jupyter notebooks, Vim documentation, and custom Lua readers/writers.

> **Source of truth**: <https://pandoc.org/MANUAL.html>. For format lists, see [formats.md](formats.md). For CLI flags, see [cli-options.md](cli-options.md).

## Slide Shows

Pandoc can generate several slide show formats:

| Format | Description |
|--------|-------------|
| `revealjs` | reveal.js HTML slides (most feature-rich) |
| `beamer` | LaTeX Beamer slides |
| `pptx` | Microsoft PowerPoint |
| `slidy` | Slidy HTML slides |
| `slideous` | Slideous HTML slides |
| `dzslides` | DZSlides HTML slides |
| `s5` | S5 HTML slides |

### Slide Level

The **slide level** is the highest heading level followed by content. Pandoc splits the document into slides at this heading level. Control it with `--slide-level`.

```markdown
# Section Title          <-- If slide level is 1, this starts a new slide section

## Slide Title            <-- If slide level is 2, this starts a new slide

Content of the slide.

## Next Slide

More content.
```

If not specified, pandoc auto-detects the slide level.

### Incremental Lists

Use `-i`/`--incremental` to display list items one at a time:

```bash
pandoc -t revealjs -i input.md -o slides.html
```

Alternatively, use Divs for per-list control:

```markdown
::: incremental
- First
- Second
- Third
:::

::: nonincremental
- All at once
:::
```

### Pauses

A paragraph containing only `. . .` creates a pause (reveal subsequent content on next click):

```markdown
First content

. . .

More content after pause
```

### Speaker Notes

```markdown
::: notes
This is speaker notes content.
:::
```

### Columns

```markdown
::: {.columns}

::: {.column width="40%"}
Left column content
:::

::: {.column width="60%"}
Right column content
:::

:::
```

### Beamer Frame Attributes

Apply attributes to Beamer frames:

```markdown
## Fragile Frame {.fragile}

## Frame with Page Break {.allowframebreaks}

## Frame with Label {#mylabel}
```

### Reveal.js Backgrounds

Per-slide backgrounds via attributes:

```markdown
## Slide with Image Background {background-image="url(image.jpg)"}
```

Global backgrounds via variables:

```yaml
---
variables:
  parallaxBackgroundImage: "url(bg.jpg)"
---
```

### PowerPoint Layouts

Pandoc maps content to PowerPoint layouts:

| Layout | Used for |
|--------|----------|
| Title Slide | First slide (title, author, date) |
| Section Header | Section divider slides |
| Two Content | Content with columns |
| Comparison | Side-by-side comparison |
| Content with Caption | Content with a caption |
| Blank | Empty slide |
| Title and Content | Standard content slide |

Use a reference document (`--reference-doc`) to customize layouts, fonts, and colors.

## EPUB

Pandoc generates EPUB v3 (`epub`/`epub3`) or v2 (`epub2`) e-books.

```bash
pandoc input.md -o book.epub --toc --split-level=2
```

### EPUB Metadata

Dublin Core metadata can be specified via YAML metadata or `--epub-metadata` (XML):

```yaml
---
title: My Book
author: Jane Doe
identifier: "urn:isbn:9781234567890"
lang: en
date: "2024-01-15"
rights: "Copyright 2024 Jane Doe"
cover-image: cover.jpg
---
```

| Field | Description |
|-------|-------------|
| `identifier` | Unique identifier (URN/ISBN) |
| `title` | Book title |
| `creator` | Author |
| `contributor` | Contributor |
| `date` | Publication date |
| `lang` | Language (BCP 47) |
| `subject` | Subject/keyword |
| `description` | Description |
| `type` | Resource type |
| `format` | Format |
| `relation` | Related resource |
| `coverage` | Coverage |
| `rights` | Rights statement |
| `belongs-to-collection` | Collection membership |
| `group-position` | Position in collection |
| `cover-image` | Cover image file |
| `css` | Custom CSS file |
| `page-progression-direction` | `ltr` or `rtl` |

### epub:type Attribute

The `epub:type` attribute marks semantic roles for EPUB content:

```markdown
# Foreword {epub:type="foreword"}

# Chapter 1 {epub:type="chapter"}

# Appendix A {epub:type="appendix"}

# Bibliography {epub:type="bibliography"}

# Index {epub:type="index"}
```

Common types: `prologue`, `abstract`, `acknowledgments`, `copyright-page`, `dedication`, `credits`, `keywords`, `imprint`, `contributors`, `errata`, `revision-history`, `titlepage`, `foreword`, `preface`, `frontispiece`, `appendix`, `colophon`, `bibliography`, `index`.

### EPUB Splitting

Control how the EPUB is split into chapters with `--split-level` (formerly `--epub-chapter-level`):

```bash
pandoc input.md -o book.epub --split-level=2
```

### EPUB Styling

Custom CSS via `--css` or the `css` metadata field. The default stylesheet is `epub.css` in pandoc's data files. Override by placing a custom `epub.css` in the data directory.

### Embedding Fonts

```bash
pandoc input.md -o book.epub --epub-embed-font=fonts/serif.ttf --epub-embed-font=fonts/sans.ttf
```

### Linked External Media

Mark media as external (not embedded):

```markdown
![Caption](https://example.com/video.mp4){data-external="1"}
```

## Chunked HTML

```bash
pandoc input.md -t chunkedhtml -o output.zip
```

Produces a ZIP archive of linked HTML files with:

- `sitemap.json` — metadata about the chunks
- Navigation links between chunks (prev/next)
- Auto-adjusted internal links pointing to the correct chunk
- Split controlled by `--split-level`

```bash
pandoc input.md -t chunkedhtml --split-level=2 -o output.zip
```

Customize chunk filenames with `--chunk-template`:

```bash
pandoc input.md -t chunkedhtml --chunk-template="page-{n}" -o output.zip
```

## Jupyter Notebooks

```bash
pandoc input.md -t ipynb -o notebook.ipynb
```

### Notebook Structure

- Code blocks with class `code` become code cells
- Intervening content becomes Markdown cells
- Cell grouping via Divs:

```markdown
::: {.cell .code}
```python
print("Hello")
```
:::

::: {.cell .markdown}
This is a markdown cell.
:::
```

### Output Cells

```markdown
::: {.output .stream .stdout}
```
Hello
```
:::
```

### Notebook Metadata

```yaml
---
jupyter:
  kernelspec:
    name: python3
    display_name: Python 3
  language_info:
    name: python
    version: "3.11"
---
```

### Controlling Output

Use `--ipynb-output` to control how existing outputs are handled:

| Mode | Behavior |
|------|----------|
| `all` | Keep all outputs |
| `none` | Strip all outputs |
| `best` | Keep the best output for each cell (default) |

## Vimdoc

```bash
pandoc input.md -t vimdoc -o helpfile.txt
```

### Vimdoc Metadata

| Field | Description |
|-------|-------------|
| `abstract` | Document abstract |
| `author` | Author |
| `title` | Document title |
| `filename` | Output filename |
| `vimdoc-prefix` | Prefix for tags |

Ex-commands (starting with `:`) are not prefixed. Links to vimhelp.org become cross-references.

## Custom Readers and Writers

Custom readers and writers are Lua scripts placed in the `custom` subdirectory of the data directory. They are invoked by using the script path as the format name.

### Custom Lua Reader

A Lua script defining a `Reader(input)` function:

```lua
-- ~/.local/share/pandoc/custom/myformat.lua
function Reader(input)
    local blocks = {}
    for line in input:lines() do
        table.insert(blocks, pandoc.Para{pandoc.Str(line)})
    end
    return pandoc.Pandoc(blocks)
end
```

```bash
pandoc -f myformat input.txt -o output.html
```

The `lpeg` parsing library is available by default for writing parsers.

### Custom Lua Writer

A Lua script defining rendering functions for each AST element:

```lua
-- ~/.local/share/pandoc/custom/mywriter.lua
function Str(s)
    return s.text
end

function Para(s)
    return s .. "\n\n"
end

function Plain(s)
    return s
end
```

Custom writers have no default template; provide one with `--template`:

```bash
pandoc -t mywriter --template=my-template input.md -o output.txt
```

For more on Lua filters and the Lua interpreter mode, see [filters.md](filters.md).

## Related References

- [formats.md](formats.md) — Full format lists
- [cli-options.md](cli-options.md) — Slide, EPUB, and notebook flags
- [templates.md](templates.md) — Slide and EPUB template variables
- [filters.md](filters.md) — Custom Lua readers/writers, Lua interpreter mode
