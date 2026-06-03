# Seska

**Seska** (*Stephen's Eternal Scrap Keeper Application*) is a note-taking database, written in Go 1.26 by Stephen Malone.

## Database

Seska stores all information in a single `.seska` database in your home directory. Each note has a lowercase name and the following data:

Field  | Type    | Description
------ | ------- | -----------
`body` | string  | A plaintext Markdown body.
`hash` | string  | A SHA256 hash of the note's body.
`flag` | string  | A space-separated list of metadata flags.
`init` | integer | A creation time as a Unix integer.
`last` | integer | A latest update time as a Unix integer.

### Secret Notes

Creating a note that starts with a dot (e.g.: `.name`) will result in that note being hidden from view, unless `show_hidden` is true.

## Configuration

All configuration is stored in a hidden `.conf` note, which starts empty but can be written with standard TOML. The defaults are:

```toml
# Check all data hashes on startup, every time.
boot_checks = false

# Show secret and deleted notes in all list views.
show_hidden = false

# Flag notes as "deleted" instead of actually deleting them.
soft_delete = true
```

## Usage

Seska opens in an interactive session, showing you a list of all your notes, a status bar and a search field. Typing anything will filter the list to notes matching your search, the arrow keys will select a note and pressing enter will open your preferred editor, as indicated by the environment variables `EDITOR` or `VISUAL`. You can also press Tab to toggle `show_hidden` for this session only and Escape to clear the search. Pressing Escape on an empty search will exit the application.
