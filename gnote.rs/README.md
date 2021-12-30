# gnote.rs

## Getting Started

```
brew install gutenye/alt/gnote.rs
brew services info gnote.rs
tail -f $(brew --prefix)/log/gnote.rs.log

gnote.rs -h
gnote.rs tags
gnote.rs tags --watch
gnote.rs tags --note-dir ~/note --note-marker '*' --note-extension .note --output ~/tags --cache-dir ~/.cache/note
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