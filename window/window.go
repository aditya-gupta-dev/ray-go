package window

import rl "github.com/gen2brain/raylib-go/raylib"

type WindowLifeCycle interface {
	OnStart()
	OnUpdate()
	OnCleanup()
}

type Window struct {
	vsync            bool
	width            int32
	height           int32
	title            string
	lifeCycleHandler WindowLifeCycle
}

func (win *Window) Init(title string, lifeCycle WindowLifeCycle) {
	win.lifeCycleHandler = lifeCycle
	defer win.lifeCycleHandler.OnStart()

	if win.title == "" {
		win.title = title
	}
	rl.InitWindow(win.width, win.height, win.title)

	if win.vsync {
		rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))
	}
}

func (win *Window) Start() {
	for !rl.WindowShouldClose() {
		win.lifeCycleHandler.OnUpdate()
	}

	win.lifeCycleHandler.OnCleanup()
}
