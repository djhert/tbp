package start

import (
	"github.com/hlfstr/ego2d/component/text"
	"github.com/hlfstr/ego2d/ego"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	Scene scene
)

const (
	pad    int32 = 16
	istart int32 = 64
	tab    int32 = 96
)

type scene struct {
	ego.Scene

	ticks  int
	loaded bool
}

func (m *scene) Start() bool {
	m.Init(1)
	ego.Log.Log(0, "Starting game")
	ego.LoadAssets()
	m.Background = ego.NewColor(0, 0, 0, 255)
	// load the fonts
	ego.LoadFont("score", "assets/fonts/pcsenior.ttf", 64)
	ego.LoadFont("large", "assets/fonts/pcsenior.ttf", 32)
	ego.LoadFont("med", "assets/fonts/pcsenior.ttf", 24)
	ego.LoadFont("small", "assets/fonts/pcsenior.ttf", 16)

	if f, ok := ego.Fonts["large"]; !ok {
		ego.Log.Log(1, "Unable to load font 'large'")
		return false
	} else {
		col := sdl.Color{255, 255, 255, 255}
		prev := int32(istart)
		// instructions
		txt := text.New("Keyboard Controls:", f, col)
		txt.SetXY(istart, prev)
		prev += txt.GetH() + pad
		m.Add(0, txt)
		if f2, ok := ego.Fonts["med"]; !ok {
			ego.Log.Log(1, "Unable to load font 'med'")
			return false
		} else {
			txt = text.New("W - Move Up", f2, col)
			txt.SetXY(tab, prev)
			prev += txt.GetH() + pad
			m.Add(0, txt)
		}
		if f2, ok := ego.Fonts["med"]; !ok {
			ego.Log.Log(1, "Unable to load font 'med'")
			return false
		} else {
			txt = text.New("S - Move Down", f2, col)
			txt.SetXY(tab, prev)
			prev += txt.GetH() + istart
			m.Add(0, txt)
		}
		txt = text.New("Mouse Controls:", f, col)
		txt.SetXY(istart, prev)
		prev += txt.GetH() + pad
		m.Add(0, txt)
		if f2, ok := ego.Fonts["med"]; !ok {
			ego.Log.Log(1, "Unable to load font 'med'")
			return false
		} else {
			txt = text.New("Left Click - Swing", f2, col)
			txt.SetXY(tab, prev)
			prev += txt.GetH() + pad
			m.Add(0, txt)
		}
		if f2, ok := ego.Fonts["small"]; !ok {
			ego.Log.Log(1, "Unable to load font 'small'")
			return false
		} else {
			txt = text.New("Swinging sends the ball towards the mouse", f2, col)
			txt.SetXY(tab+pad, prev)
			prev += txt.GetH() + pad
			m.Add(0, txt)
		}
		if f2, ok := ego.Fonts["med"]; !ok {
			ego.Log.Log(1, "Unable to load font 'med'")
			return false
		} else {
			txt = text.New("Right Click - Use Weapon", f2, col)
			txt.SetXY(tab, prev)
			prev += txt.GetH() + pad
			m.Add(0, txt)
		}
		if f2, ok := ego.Fonts["small"]; !ok {
			ego.Log.Log(1, "Unable to load font 'small'")
			return false
		} else {
			txt = text.New("Fires the weapon in the direction of the mouse", f2, col)
			txt.SetXY(tab+pad, prev)
			prev += txt.GetH() + pad
			m.Add(0, txt)
		}
		txt = text.New("Loading...", f, col)
		txt.SetXY((ego.Width/2)-(txt.GetW()/2), ego.Height-(txt.GetH()*4))
		m.Add(0, txt)
		return true
	}
}

func (m *scene) Update() {
	if m.ticks > 100 && !m.loaded {
		if ego.AssetsLoaded {
			ego.SetNextScene("battle")
			m.loaded = true
		}
	}
	if m.ticks > 150 && m.loaded {
		ego.NextScene()
	}
	m.ticks++
}

func (m *scene) Destroy() {
}
