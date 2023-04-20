package main

func main() {
	tv := &tv{}
	onCommand := &onCommand{
		device: tv,
	}
	offCommand := &offCommand{
		device: tv,
	}
	onButton := &button{
		command: onCommand,
	}
	offButoon := &button{
		command: offCommand,
	}
	onButton.press()
	offButoon.press()
}
