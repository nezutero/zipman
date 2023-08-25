<h2 align="center">zipman - archive compressor/extractor golang library</h2>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## supported formats:

 - `zip`

## project structure:

```
├── core
│   ├── 7z.go
│   ├── bzip2.go
│   ├── gz.go
│   ├── rar.go
│   ├── tar.go
│   └── zip.go
├── go.mod
├── main.go
└── README.md
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
	zipman.Zip("./file.zip", []string{"./man.txt", "./hello.txt"})
}

➜  zipman git:(master) ✗ go run main.go
[SUCCESS] Zip compression completed successfully for file: ./file.zip
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
