package main

import "codeberg.org/anaseto/gruid"

func (md *model) KeyDown(msg gruid.Msg) gruid.Effect {
	switch msg := msg.(type) {
	case gruid.MsgKeyDown:
		switch msg.Key {
		case "Q":
			return gruid.End()
		}
	}
	return nil
}
