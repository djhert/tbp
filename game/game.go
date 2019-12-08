package game

import (
	"errors"

	"github.com/hlfstr/ego2d/ego"
	"github.com/hlfstr/tbp"
	"github.com/hlfstr/tbp/scenes/battle"
	"github.com/hlfstr/tbp/scenes/start"

	"math/rand"
	"time"
)

func StartUp() error {
	ego.Create(1280, 720, ego.WINDOWED, tbp.NAME)
	rand.Seed(time.Now().UnixNano())

	ego.AddScene("init", &start.Scene)
	ego.AddScene("battle", &battle.Scene)

	if ok := ego.SetInitScene("init"); !ok {
		return errors.New("Unable to start")
	}

	return ego.Run()
}
