# Gnote

Same cli app written in Typescript, Rust and Go

## Comparision

| Language   | Startup Time | Single Binary size | Note    |
|------------|--------------|--------------------|---------|
| Typescript | 700ms (lag)  | 40 MB              | use pkg | 
| Rust       | 10ms         | 0.5 MB             |         |
| Go         |              |                    |         |

- Startup time: restart system, run with `time <cmd> -h`
- Single Binary size: on Apple Silicon
- See [Language Syntax Comparision](https://github.dev/gutenye/gnote/tree/main/lang)