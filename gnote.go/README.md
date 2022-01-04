## Install

```
edit ~/.gnoterc
  dir: ~/note
  output: ~/tags
  mark: ∗


# macOS
brew install gutenye/alt-gnote
brew start  gnote
tail -f /usr/local/var/log/gnote.log

# Linux
pacman -S gnote

# From source
go get -u github.com/gutenye/gnote/gnote.go

# Fix bad file descriptor
launchctl limit maxfiles 90000
Restart system
```

## Development

```
go get -u github.com/gutenye/gnote/gnote.go
cd ~/go/src/github.com/gutenye/gnote/gnote.go
dep ensure
go run .
./ake build
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
