package secret

const (
	wink          = 1
	doubleBlink   = 2
	closeYourEyes = 4
	jump          = 8
	reverse       = 16
)

func Handshake(code int) []string {
	commands := make([]string, 0, 0)
	if code < 0 || code > 32 {
		return commands
	}
	if wink&code > 0 {
		commands = append(commands, "wink")
	}
	if doubleBlink&code > 0 {
		commands = append(commands, "double blink")
	}
	if closeYourEyes&code > 0 {
		commands = append(commands, "close your eyes")
	}
	if jump&code > 0 {
		commands = append(commands, "jump")
	}
	if reverse&code > 0 {
		commands = rev(commands)
	}
	return commands
}

func rev(orig []string) []string {
	reversed := make([]string, len(orig), cap(orig))
	for i, _ := range orig {
		ix := len(orig) - i - 1
		reversed[ix] = orig[i]
	}
	return reversed
}
