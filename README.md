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

## Full Usage
`Usage:
  stegan [flags]

Flags:
  -d, --decode           decode a message from the specified image file, resulting in a printed message to stdout
  -e, --encode           encode a message into the specified image file, resulting in a new encoded image file
  -h, --help             help for stegen
  -i, --in string        input image file
  -m, --message string   message to encode into image (if encoding)
  -o, --out string       output image file (if encoding)`

