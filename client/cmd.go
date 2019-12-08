package client

import (
	"github.com/hlfstr/flagger"
	"github.com/hlfstr/tbp/game"
)

type Cmd struct {
}

func (c *Cmd) Prepare(flags *flagger.Flags) {
}

func (c *Cmd) Action(s []string, f *flagger.Flags) error {
	return game.StartUp()
}

func (c *Cmd) Print() {
}
