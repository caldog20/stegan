# stegan
CLI tool to encode/decode messages to/from images

## Installing
`go install github.com/caldog20/stegan@latest`

## Building
The project doesn't require many dependencies. Just clone the repo, run `go mod download` then run `go build` to build the resulting binary `stegan` in the root of the repo.

## Encoding
`stegan --encode --in <input image file> --out <output image file> --message <message>`

## Decoding
`stegan --decode --in <input image file>`
