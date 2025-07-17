package main

type StateMachine interface {
	SetState(string)
	GetState()
}

type Options struct {
	UIWidth int
	UIHeight int
	LogFile string
	States map[string]bool
}

func InitOptions() *Options {
	// Set all available menus
	opt := &Options{
		States: map[string]bool{
			"menu": true,
			"game": true,
		},
		UIWidth: 80,
		UIHeight: 24,
		LogFile: "./log.txt",
	}
	return opt
}
