package cmd

import (
	"errors"
	_ "fmt"
	_ "os"

	"github.com/spf13/cobra"

	"github.com/caldog20/stegan/lib"
)

var (
	encode     bool
	decode     bool
	inputFile  string
	outputFile string
	message    string
	rootCmd    *cobra.Command
)

func validateArgs() error {
	if encode && decode {
		return errors.New("encode and decode options cannot be used at the same time")
	}
	if !encode && !decode {
		return errors.New("encode or decode option must be specified")
	}
	if inputFile == "" {
		return errors.New("input file must be specified")
	}
	if encode && message == "" {
		return errors.New("message cannot be empty while encode option is set")
	}
	if encode && outputFile == "" {
		return errors.New("must specify output file when using encode option")
	}
	return nil
}

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stegan",
		Short: "A steganography tool to encode/decode messages into images",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateArgs(); err != nil {
				return err
			}

			if encode {
				return lib.Encode(inputFile, outputFile, message)
			}

			if decode {
				return lib.Decode(inputFile)
			}

			return nil
		},
	}
}

func init() {
	rootCmd = NewRootCmd()
	rootCmd.PersistentFlags().
		BoolVarP(&encode, "encode", "e", false, "encode a message into the specified image file, resulting in a new encoded image file")
	rootCmd.PersistentFlags().
		BoolVarP(&decode, "decode", "d", false, "decode a message from the specified image file, resulting in a printed message to stdout")
	rootCmd.PersistentFlags().StringVarP(&inputFile, "in", "i", "", "input image file")
	rootCmd.PersistentFlags().
		StringVarP(&outputFile, "out", "o", "", "output image file (if encoding)")
	rootCmd.PersistentFlags().
		StringVarP(&message, "message", "m", "", "message to encode into image (if encoding)")
}

func Execute() error {
	return rootCmd.Execute()
}
