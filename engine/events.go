package engine

type action struct {
	function func(ref any)
	ref      any
}

var actions map[string][]action

func Listen(event string, function func(ref any), selfRef any) {
	actions[event] = append(actions[event], action{function: function, ref: selfRef})
}

func Trigger(event string) {
	for _, v := range actions[event] {
		v.function(v.ref)
	}
}
