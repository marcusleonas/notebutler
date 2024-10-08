<p align="center">
  <img src="https://github.com/marcusleonas/notebutler/blob/main/logo.png?raw=true" alt="Notebutler logo" width="200">
  <br/><br/>
  <h2><strong><a href="https://nb.flvffy.xyz/">Read the docs</a></strong></h2>
</p>

# Notebutler

A little cli tool to manage your notes. Written in Golang.

## Installation

Download the binary from the [releases page](https://github.com/marcusleonas/notebutler/releases) and put it somewhere in your `$PATH`.

## Usage

### Initialise a notebook:

```sh
notebutler init
```

### Create a new note

```sh
notebutler new
```

**Flags:**

(all flags are optional)

- `-n, --name`: Name of the note.
- `-t, --template`: Template to use for the note (found inside `.notebutler/templates`). Do not include the `.md` extension. Defaults to `default`.
- `-o, --open`: Open the file in the default editor after creation. Defaults to `false`.

### Read a note

```sh
notebutler read <name>
```

## Build

Install [Go](https://golang.org/doc/install).

### For current architecture

```sh
go build .
```

### For all architectures

```sh
./build.sh github.com/marcusleonas/notebutler
```

## License

Licensed under the MIT License. See [LICENSE](LICENSE) for details.
