package infrastructure

import (
	"github.com/gozelle/_exec"
)

// SSHRemoteExec implements domain.RemoteExec using SSH.
type SSHRemoteExec struct {
	server *_exec.Server
}

// NewSSHRemoteExec creates a new SSHRemoteExec.
func NewSSHRemoteExec(server *_exec.Server) *SSHRemoteExec {
	return &SSHRemoteExec{server: server}
}

func (r *SSHRemoteExec) Exec(command string) error {
	return r.server.CombinedExec(command)
}

func (r *SSHRemoteExec) ExecPipe(command string) error {
	return r.server.PipeExec(command)
}
