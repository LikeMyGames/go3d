package config

import (
	eng "go3d/engine"
)

var (
	InputMapping_Player eng.InputMapping = eng.InputMapping{
		"spacebar_press": "jump_forward",
		"w_press":        "jump_forward",
		"a_press":        "jump_left",
		"s_press":        "jump_back",
		"d_press":        "jump_right",
	}
	InputMapping_Menu eng.InputMapping = eng.InputMapping{}
)

var InputMappings map[string]eng.InputMapping = map[string]eng.InputMapping{
	"player": InputMapping_Player,
	// "menu":   InputMapping_Menu,
}
