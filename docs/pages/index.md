---
title: "Getting Started - Notebutler"
---

# Getting Started

## Installation

The latest binaries are available on the [releases page](https://github.com/marcusleonas/notebutler/releases).

On Mac/Linux you can install the binary with:

```sh
curl -L https://github.com/marcusleonas/notebutler/releases/latest/download/<version>.tgz -o notebutler

install notebtler /usr/local/bin
```

On windows, you can download the binary from the [releases page](https://github.com/marcusleonas/notebutler/releases)
and add it to your PATH.

## Initialise a new notebook

Initialising a new notebook is as simple as running:

```sh
notebutler init
```

Make sure to fill in all the fields otherwise the notebook will not be created.

## Create a new note

To create a new note, run:

```sh
notebutler new
```

Once again, fill in all the fields. This time, the fields can be filled using flags.

**Available flags:**

`-n, --name`: Name of the note.

`-t, --template`: Template to use for the note (found inside .notebutler/templates).
Do not include the .md extension. Defaults to default.

`-o, --open`: Open the file in the default editor after creation. Defaults to false.

## Read a note

Reading a note renders it using [glamour](https://github.com/charmbracelet/glamour)
and outputs it to stdout.

```sh
notebuter read <name>
```

## Build your notes to html

To build your notes to a static html site, run:

```sh
notebutler build
```

This will create a `html` directory in the current directory, with all your notes
converted to html. This includes [picocss](https://picocss.com/) for styling.

## Serve your notes for local development

To serve your notes for local development, run:

```sh
notebutler serve
```

This will start a local server on port 8080.

## Issues / Feature Requests

If you have any issues or feature requests, please open an issue on the
[issues page](https://github.com/marcusleonas/notebutler/issues). I'm open
to any suggestions and contributions.

## License

Licensed under the MIT License.
