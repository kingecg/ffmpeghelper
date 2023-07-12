package ffmpeghelper

import (
	"io"
)

type Streamer struct {
	input io.Reader
	out   io.Writer
	Args  []string
}

func (s *Streamer) SetInput(input io.Reader) {
	s.input = input
}
func (s *Streamer) SetOutput(out io.Writer) {
	s.out = out
}

// init io if not set before
func (s *Streamer) Init() (input io.Writer, output io.Reader) {
	var inputReader io.Reader
	var inputWriter io.Writer
	if s.input == nil {
		inputReader, inputWriter = io.Pipe()
		s.input = inputReader
	}
	var outputReader io.Reader
	var outputWriter io.Writer
	if s.out == nil {
		outputReader, outputWriter = io.Pipe()
		s.out = outputWriter
	}

	return inputWriter, outputReader
}
