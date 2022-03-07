# Codeblock Extractor

This simply extracts the following HTML into a snippet file given in `data-file` attribute.

```
<code class="snippet" data-file="myfile">
my snippet code
</code>
```

The workflow involves

- adding `<pre><code>` blocks to markdown files as I follow tutorials
- extract them to separate snippet files
- fzf and xclip to copy snippet to clipboard when programming

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
