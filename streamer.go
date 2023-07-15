package ffmpeghelper

import "io"

type Streamer struct {
	ffmpegProcess *Process
	InputOptions  []string
	Inputs        string
	VideoCodec    string
	AudioCodec    string
	OutputFmt     string
	OutputTarget  string
}

func toperator[Op any](cond bool, a Op, b Op) Op {
	if cond {
		return a
	} else {
		return b
	}
}
func (s *Streamer) Build() (io.Writer, io.Reader, io.Reader) {
	args := []string{}
	args = append(args, s.InputOptions...)

	inputSrc := toperator(s.Inputs != "", s.Inputs, "pipe:0")
	args = append(args, "-i", inputSrc)
	vcodec := toperator(s.VideoCodec != "", s.VideoCodec, "copy")
	args = append(args, vcodec)
	acodec := toperator(s.AudioCodec != "", s.AudioCodec, "copy")
	if acodec != "an" {
		args = append(args, "-c:a", acodec)
	} else {
		args = append(args, "-an")
	}

	if s.OutputFmt != "" {
		args = append(args, "-f", s.OutputFmt)
	}
	outArg := toperator(s.OutputTarget != "", s.OutputTarget, "pipe:1")
	args = append(args, outArg)
	s.ffmpegProcess = NewProcess("ffmpeg", args...)
	return s.ffmpegProcess.In, s.ffmpegProcess.Out, s.ffmpegProcess.Err
}

func (s *Streamer) Run() error {
	return s.ffmpegProcess.cmd.Run()
}
