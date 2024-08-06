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

## Issues / Feature Requests

If you have any issues or feature requests, please open an issue on the
[issues page](https://github.com/marcusleonas/notebutler/issues). I'm open
to any suggestions and contributions.

## License

Licensed under the MIT License.
