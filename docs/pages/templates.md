---
title: "Templates - Notebutler"
---

# Templates

Templates are quite limited at the moment, but I'm planning to add more features
soon. At the moment, there are two variables that you can parse into your template.
Frontmatter isn't function yet, but included in the default template. Templates
are written in [Go's text/template](https://golang.org/pkg/text/template/) format,
and stored in the `.notebutler/templates` directory. Any markdown file in that
folder will be automatically detected and can be used when creating new notes.

## Variables

`Name`: The name of the note.

`CreatedAt`: The date and time the note was created.

## Example

```md
# {{ .Name }}

This is a new note created at {{ .CreatedAt }}.
```
