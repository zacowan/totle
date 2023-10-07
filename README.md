# totle

> Aristotle was an Ancient Greek philosopher and polymath who made many important contributions to various subjects of thinking. Importantly, Aristotle wrote down his thoughts! **totle** is a simple tool to allow developers to jot down their thoughts for safe-keeping in a transferrable format.

## Features

- 🗓️ Simple, date-based note-taking solution
- ⬇️ Uses Markdown syntax
- 💻 Add notes via the CLI or your favorite editor
- ⚙️ Configurable

## Installation

### Homebrew

```sh
brew tap zacowan/tap
brew install zacowan/tap/totle
```

### Scoop

```sh
scoop bucket add zacowan https://github.com/zacowan/scoop-bucket.git
scoop install zacowan/totle
```

## Usage

### Create

```sh
totle create
```

Creates a note file for today and opens it using the configured `open_cmd` command. By default, the note file is opened using the `code` command provided by VSCode.

If a note file already exists for today, that file is opened. Notes are named by today's date (`yyyy-mm-dd`) based on your local timezone. They are placed in the configured notes directory, and organized into folders based on the year and month of the note.

For example, for a note created on January 1st, 1970, the notes directory would look like the following:

```
1970/
  01/
    1970-01-01.md
```

The initial contents of a note include only the title of the note, which is the date it was created:

```md
# 1970-01-01
```

### Add

```sh
totle add [note]
```

Adds a new note to today's note file. If no note file exists for today, a new note file is created with the contents of `[note]`. If a note file already exists for today, the contents of the `[note]` are appended to that file. For new notes, the same behavior of `totle create` is followed. Then, the `[note]` is added to the file as a markdown bullet point.

For example, `totle add "Hello, world"` would produce the following note file contents:

```md
# 1970-01-01

- Hello, world
```

### Open

```sh
totle open
```

Opens the note file for today using the configured `open_cmd` command. By default, the note file is opened using the `code` command provided by VSCode.

### Configuration

`totle` supports being configured using a `.totle.yaml` file in your home directory. Alternatively, you can specify a different file the load configuration options from using the `--config` flag. The following configuration options are supported:

```yaml
# The command to use when opening a note file. The command is passed
# the path to the note file as the first and only argument.
#
# Default: code
open_cmd: open

# The directory that your notes are stored in.
#
# Default: $HOME/Documents/totle
notes_dir: path/to/notes
```

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md).
