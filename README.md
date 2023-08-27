<h2 align="center">zipman - archive compressor/extractor golang library</h2>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## supported formats and functionality:

- `zip` (compression/extraction/adding)
- `tar` (compression/extraction/adding)
- `tar.gz` (compression/extraction/adding)
- `7z` (compression/extraction)
- `bz2` (compression/extraction)
- `rar` (extraction)

## project structure:

```
├── 7z.go
├── bzip2.go
├── go.mod
├── go.sum
├── LICENSE
├── rar.go
├── README.md
├── tar.go
├── targz.go
└── zip.go
```

## installation

- use go get:

```
go get github.com/kenjitheman/zipman
```

## usage

```
package main

import "github.com/kenjitheman/zipman"

func main() {
	zipman.CompressToZip("./file.zip", []string{"./man.txt", "./hello.txt"})

    // zipman.AddFileToZip(zipWriter *zip.Writer, filename string)
	// zipman.ExtractZip("./file.zip", "./man")
}
```

```
➜  zipman git:(master) ✗ go run main.go
[SUCCESS] Zip compression completed successfully for file: ./file.zip
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
