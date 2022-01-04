## Getting Started

```
brew install gutenye/alt/gnote.go
brew services info gnote.go
tail -f $(brew --prefix)/log/gnote.go.log

gnote.go -h
gnote.go tags
gnote.go tags --watch
gnote.go tags --note-dir ~/note --note-marker '*' --note-extension .note --output ~/tags --cache-dir ~/.cache/note

# Fix bad file descriptor
launchctl limit maxfiles 90000
Restart system
```

## Development

```
./ake [...args]
./ake test
./ake test:cmd
```

## Release

```
git tag go1.0.0
git push --tags
```

## Test

```
./ake t:cmd
cd testdata/note
mkdir -p a
echo '*a*' > a/f
echo '*a*' > a/f.gnote
mv a/f.gnote a/f2.gnote
rm a/f2.gnote
echo '*a*' > a/f.gnote
mv a b
rm -r b
```
