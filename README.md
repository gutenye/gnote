# Gnote

Same cli app written in Typescript, Rust and Go

## Comparision

| Language   | Startup Time | Single Binary size | Note    |
|------------|--------------|--------------------|---------|
| Typescript | 3s           | 40 MB              | use pkg | 
| Rust       | 0.3s         | 2.4MB              |         |
| Go         |              |                    |         |

- TODO: use hyperfine for benchmark, do cold start
- Startup time: restart system, run with `time <cmd> -h` for the first time.
- Single Binary size: on Apple Silicon
- See [Language Syntax Comparision](https://github.dev/gutenye/gnote/tree/main/lang)
