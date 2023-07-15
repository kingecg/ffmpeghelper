package ffmpeghelper

import (
	"io"
	"os/exec"
)

type Process struct {
	cmd *exec.Cmd
	In  io.Writer
	Out io.Reader
	Err io.Reader
}

func NewProcess(name string, args ...string) *Process {

	p := &Process{

		cmd: exec.Command(name, args...),
	}
	p.In, _ = p.cmd.StdinPipe()
	p.Out, _ = p.cmd.StdoutPipe()
	p.Err, _ = p.cmd.StderrPipe()
	return p
}
