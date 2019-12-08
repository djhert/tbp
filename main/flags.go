package main

import (
	"github.com/hlfstr/flagger"
	"github.com/hlfstr/flagger/commands"

	"github.com/hlfstr/tbp"
	"github.com/hlfstr/tbp/client"
	"github.com/hlfstr/tbp/server"
	"github.com/hlfstr/tbp/single"

	"fmt"
	"os"
)

var (
	cmd *commands.Commands
)

func setup() error {
	cmd = commands.New()
	cmd.Add("help", &helpCmd{})
	cmd.Add("version", &versionCmd{})
	cmd.Add("client", &client.Cmd{})
	cmd.Add("server", &server.Cmd{})
	cmd.Add("single", &single.Cmd{})
	err := cmd.Parse(os.Args[1:])
	if err != nil {
		return err
	}
	return nil
}

type helpCmd struct {
}

func (h *helpCmd) Print() {
	fmt.Printf("Usage: %s [OPTION] [FLAGS]...\n\nAvailable Options:\n", os.Args[0])
	fmt.Println("  client  - Start the game as a client")
	fmt.Println("  server  - Start the game as a server")
	fmt.Println("  single  - Start a single player game")
	fmt.Println("  help    - Show this help")
	fmt.Println("  version - Prints the version")
}

func (h *helpCmd) Prepare(flags *flagger.Flags) {
}

func (h *helpCmd) Action(s []string, f *flagger.Flags) error {
	h.Print()
	return nil
}

type versionCmd struct {
}

func (v *versionCmd) Print() {
	fmt.Printf("%s\n\tA Game of Pong with a twist\n", tbp.Version())
}

func (v *versionCmd) Prepare(flags *flagger.Flags) {
}

func (v *versionCmd) Action(s []string, f *flagger.Flags) error {
	v.Print()
	return nil
}
