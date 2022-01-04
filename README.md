# Gnote

Same cli app written in Typescript, Rust and Go

## Comparision

| Language   | Startup Time | Single Binary size |
|------------|--------------|--------------------|
| Rust       |  10,  40ms   | 0.5 MB             |
| Go         |  30,  50ms   | 3 MB               |
| Typescript | 700, 600ms   | 40 MB              |

- Startup time: restart system, run with `time <cmd> -h` and `hyperfine --warmup 3  '<cmd> tags'`, >500ms is lag
- Single Binary size: macos-amd64
- Node: use pkg to create a single binary
- See [Language Syntax Comparision](https://github.dev/gutenye/gnote/tree/main/lang)