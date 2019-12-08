package battle

import (
	"fmt"

	"github.com/hlfstr/ego2d/component/pointer"
	"github.com/hlfstr/ego2d/component/sprite"
	"github.com/hlfstr/ego2d/component/text"
	"github.com/hlfstr/ego2d/ego"
	"github.com/hlfstr/tbp/game/player"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	Scene scene
	Ptr   *pointer.Pointer
)

type scene struct {
	ego.Scene

	// Player info
	pName  *text.Text
	pScore *text.Text

	// Oponent info
	oName  *text.Text
	oScore *text.Text
}

const (
	pad int32 = 32
)

func (m *scene) Start() bool {
	// layer 0 is first, background elements
	// layer 1, player + weapons fire
	// layer 2, ball
	// layer 3, effects (maybe)
	// layer 4 is the top layer, reserved for UI
	m.Init(5)
	ego.Log.Log(0, "Battle Scene")
	m.Background = ego.NewColor(0, 0, 0, 255)

	Ptr = pointer.New()
	Ptr.SetTexture("middlefinger")
	Ptr.Start()
	m.Add(4, Ptr)
	// Top edge
	obj := sprite.New()
	obj.SetTexture("battle/edge")
	obj.SetXY(0, 0)
	obj.SetW(ego.Width)
	m.Add(0, obj)
	// Bottom Edge
	obj = sprite.New()
	obj.SetTexture("battle/edge")
	obj.Flip(sprite.VERTICAL)
	obj.SetXY(0, ego.Height-obj.GetH())
	obj.SetW(ego.Width)
	m.Add(0, obj)
	// Middle "Net"
	obj = sprite.New()
	obj.SetTexture("battle/middle")
	obj.SetXY((ego.Width/2)-(obj.GetW()/2), 0)
	obj.SetH(ego.Height)
	m.Add(0, obj)

	// Player stuff
	player.New()
	player.Player.Start()
	m.Add(1, player.Player)

	m.pName = text.New(player.Player.Name, ego.Fonts["small"], sdl.Color{255, 255, 255, 255})
	m.pName.SetXY((ego.Width/2)-(m.pName.GetW()+(pad*2)), m.pName.GetH())
	m.Add(4, m.pName)

	m.pScore = text.New(fmt.Sprintf("%02d", player.Player.Score), ego.Fonts["score"], sdl.Color{255, 255, 255, 255})
	m.pScore.SetXY((ego.Width/2)-(m.pScore.GetW()+pad), m.pScore.GetH())
	m.Add(4, m.pScore)

	m.oName = text.New("Player two over here they are", ego.Fonts["small"], sdl.Color{255, 255, 255, 255})
	m.oName.SetXY((ego.Width/2)+(pad*2), m.oName.GetH())
	m.Add(4, m.oName)

	m.oScore = text.New(fmt.Sprintf("%02d", 2), ego.Fonts["score"], sdl.Color{255, 255, 255, 255})
	m.oScore.SetXY((ego.Width/2)+(pad), m.oScore.GetH())
	m.Add(4, m.oScore)

	m.AddInput(player.Player)
	return true
}

func (m *scene) Update() {
	Ptr.Update()
	player.Player.Update()
}

func (m *scene) WindowGainFocus(e *sdl.WindowEvent) {
	Ptr.Focus(true)
}

func (m *scene) WindowLoseFocus(e *sdl.WindowEvent) {
	Ptr.Focus(false)
}

func (m *scene) Destroy() {
}
