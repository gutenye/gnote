# gnote.rs

## Getting Started

```
gnote.rs -h
gnote.rs tags --note-dir ~/note --note-marker '*' --note-extension .note --output ~/tags --cache-dir ~/.cache/note
gnote.rs tags --watch
```

## Development

```
./ake [...args]
./ake test
./ake test:cmd
```

## Release

```
vi Cargo.toml
	version = "1.0.0"
git commit -m 'rs1.0.0'
git tag rs1.0.0
git push --tags
```

