package main

type client struct {
}

func (c *client) insertSquareUSBInComputer(com computer) {
	com.insertInSquarePort()
}
