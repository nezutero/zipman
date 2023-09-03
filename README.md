## Supported formats and functionality:

- `zip` (compression/extraction/adding)
- `tar` (compression/extraction/adding)
- `tar.gz` (compression/extraction/adding)
- `7z` (compression/extraction)
- `bz2` (compression/extraction)
- `rar` (extraction)

## Installation

```sh
go get github.com/nezutero/zipman
```

## Usage

```go
zipman.CompressToZip("./file.zip", []string{"./man.txt", "./hello.txt"})
```

```rust
➜  zipman git:(master) ✗ go run main.go
[SUCCESS] Zip compression completed successfully for file: ./file.zip
```

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

## License

- [MIT](./LICENSE)
