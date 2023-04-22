package main

func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUSBInComputer(mac)
	windowMachine := &windows{}
	windowMachineAdapter := &windowsAdapter{
		windowMachine: windowMachine,
	}
	client.insertSquareUSBInComputer(windowMachineAdapter)
}
