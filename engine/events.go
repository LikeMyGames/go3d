package engine

var actions map[string][]func()

func Listen(event string, action func()) {
	actions[event] = append(actions[event], action)
}

func Trigger(event string) {
	for _, v := range actions[event] {
		v()
	}
}
