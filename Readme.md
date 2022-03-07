# Codeblock Extractor

This simply extracts the following HTML into a snippet file given in `data-file` attribute.

```
<code class="snippet" data-file="myfile">
my snippet code
</code>
```

The workflow involves

- adding `<pre><code>` blocks to markdown files as I follow tutorials
- Extract them to separate snippet files.
  - Run as a post commit hook on all changed files
- fzf and xclip to copy snippet to clipboard when programming

## Install

- Clone repo
- Run `go build` and put the executable in `PATH`

## Usage

```
snippet-extractor -f "my/file/to/extract/from" -d "directory/to/extract/to"
```

Accompanying [Templater] Template for [Obsidian]

```
<%*
const languages = {
    "Zsh": "zsh",
    "JavaScript": "javaScript",
    "Golang": "golang"
}
lang = await tp.system.suggester(Object.keys(languages), Object.values(languages))
if (lang) {
%>

## <% tp.file.selection() %>

<pre>
	<code class="snippet language-<% lang %>"  data-file="<% lang %>/<% tp.file.selection() %>"	>

	</code>
</pre>
<%* } %>

```

[templater]: https://github.com/SilentVoid13/Templater
[obsidian]: https://obsidian.md/
