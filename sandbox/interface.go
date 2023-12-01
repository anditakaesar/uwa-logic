package sandbox

import "github.com/anditakaesar/uwa-logic/sandbox/differences"

type SandboxInterface interface {
	RunDifferences()
}

type Sandbox struct{}

func (s *Sandbox) RunDifferences() {
	differences.Run()
}
