package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"os"

	"github.com/auyer/steganography"
)

func Decode(filepath string) error {
	in, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("Error opening file %s: %v", filepath, err)
	}
	defer in.Close()

	reader := bufio.NewReader(in)
	img, _, err := image.Decode(reader)
	if err != nil {
		return fmt.Errorf("error decoding image ", img)
	}

	size := steganography.GetMessageSizeFromImage(img)

	msg := steganography.Decode(size, img)

	for i := range msg {
		fmt.Printf("%c", msg[i])
	}
	fmt.Println()
	return nil
}

func Encode(input string, output string, message string) error {
	in, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("Error opening file %s: %v", input, err)
	}
	defer in.Close()

	reader := bufio.NewReader(in)
	img, _, err := image.Decode(reader)
	if err != nil {
		return fmt.Errorf("error decoding image ", img)
	}

	encoded := new(bytes.Buffer)
	err = steganography.Encode(encoded, img, []byte(message))
	if err != nil {
		return fmt.Errorf("error encoding message into image: %v", err)
	}

	out, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("error creating output file %s: %v", input, err)
	}

	bufio.NewWriter(out).Write(encoded.Bytes())
	return nil
}
