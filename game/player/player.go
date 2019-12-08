package player

import (
	"github.com/hlfstr/ego2d/component/sprite"
	"github.com/hlfstr/ego2d/ego"
	"github.com/hlfstr/ego2d/math"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	Player *play
)

const (
	width  int32 = 16
	height int32 = 96
	pad    int32 = 4
)

type play struct {
	*sprite.Sprite

	// spped of moving up/down
	speed int32

	// 0 == not moving
	// 2 == moving
	// anything else == not moving, set to 0
	moving int8
	// can only be 0 or 1
	// if 2+ then set to 0
	frame int8

	// ticks counts to 4
	ticks int8

	// 0 = nothing
	// 1 = up
	// 2 = down
	direction int8

	// player name
	Name string

	Score int

	regions []math.Rect
}

func New() {
	Player = &play{
		Sprite: sprite.New(),
	}
}

func (p *play) Start() {
	p.moving = 0
	p.frame = 0
	p.speed = 10
	p.ticks = 0
	// Setup the animation regions
	p.regions = make([]math.Rect, 4)
	p.regions[0] = math.NewRect(pad, pad, width, height)
	p.regions[1] = math.NewRect(width+(pad*3), pad, width, height)
	p.regions[2] = math.NewRect(pad, height+(pad*3), width, height)
	p.regions[3] = math.NewRect(width+(pad*3), height+(pad*3), width, height)
	p.SetTexture("player/player_sheet")
	p.SetWH(width, height)
	p.SetXY(width*2, (ego.Height/2)-(p.GetH()/2))
	p.Name = "Player One over here I am"
}

func (p *play) Update() {
	if p.ticks%20 == 0 {
		if p.frame == 0 {
			p.frame = 1
		} else {
			p.frame = 0
		}
	}

	if p.moving != 0 && p.moving != 2 {
		p.moving = 0
	}

	if p.direction == 1 {
		p.AddY(-p.speed)
	} else if p.direction == 2 {
		p.AddY(p.speed)
	}

	if p.GetY() <= 20 {
		p.SetY(21)
	} else if p.GetY() >= ego.Height-height-20 {
		p.SetY(ego.Height - height - 21)
	}

	p.SetTextureRegionRect(&p.regions[p.frame+p.moving])
	p.ticks++
	if p.ticks >= 59 {
		p.ticks = 0
	}
}

func (p *play) MouseMotion(e *sdl.MouseMotionEvent) bool {
	return false
}

func (p *play) MouseButton(e *sdl.MouseButtonEvent) bool {
	return false
}

func (p *play) MouseWheel(e *sdl.MouseWheelEvent) bool {
	return false
}

func (p *play) Keyboard(e *sdl.KeyboardEvent) bool {
	if e.Type == sdl.KEYDOWN {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_W:
			p.moving = 2
			p.direction = 1
			return true
		case sdl.SCANCODE_S:
			p.moving = 2
			p.direction = 2
			return true
		default:
			return false
		}
	} else {
		switch e.Keysym.Scancode {
		case sdl.SCANCODE_W:
			fallthrough
		case sdl.SCANCODE_S:
			p.moving = 0
			p.direction = 0
			return true
		default:
			return false
		}
	}
}
