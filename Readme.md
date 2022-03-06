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
