package pkg

import (
	"github.com/anditakaesar/uwa-logic/pkg/getnames"
	"github.com/anditakaesar/uwa-logic/pkg/midlinkedlist"
	"github.com/anditakaesar/uwa-logic/pkg/richest"
	"github.com/anditakaesar/uwa-logic/pkg/stepstozero"
)

type PkgInterface interface {
	RunGetNames()
	RunRichest()
	RunStepsToZero()
	RunMidLinkedList()
}

type Pkg struct{}

func (p *Pkg) RunGetNames() {
	getnames.Run()
}

func (p *Pkg) RunRichest() {
	richest.Run()
}

func (p *Pkg) RunStepsToZero() {
	stepstozero.Run()
}

func (p *Pkg) RunMidLinkedList() {
	midlinkedlist.Run()
}
