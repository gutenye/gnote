Install
--------

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
go get -u github.com/gutenye/gnote

# Fix bad file descriptor
launchctl limit maxfiles 90000
Restart system
```

Development
--------

```
go get -u github.com/gutenye/gnote
cd ~/go/src/github.com/gutenye/gnote
dep ensure
go run .
./ake build
```
