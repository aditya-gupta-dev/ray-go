package window

import rl "github.com/gen2brain/raylib-go/raylib"

type WindowLifeCycle interface {
	OnStart()
	OnUpdate()
	OnCleanup()
}

type Window struct {
	Vsync            bool
	Width            int32
	Height           int32
	Title            string
	lifeCycleHandler WindowLifeCycle
}

func (win *Window) Init(title string, lifeCycle WindowLifeCycle) {
	win.lifeCycleHandler = lifeCycle
	defer win.lifeCycleHandler.OnStart()

	if win.Title == "" {
		win.Title = title
	}

	var monitor = rl.GetCurrentMonitor()

	if win.Height == 0 || win.Width == 0 {
		width := int32(rl.GetMonitorWidth(monitor))
		height := int32(rl.GetMonitorHeight(monitor))
		win.Width = width / 2
		win.Height = height / 2
	}
	rl.InitWindow(win.Width, win.Height, win.Title)

	if win.Vsync {
		rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(monitor)))
	}
}

func (win *Window) Start() {
	for !rl.WindowShouldClose() {
		win.lifeCycleHandler.OnUpdate()
	}

	win.lifeCycleHandler.OnCleanup()
}
