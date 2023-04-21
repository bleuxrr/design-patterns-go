package main

func main() {
	shirtItem := newItem("nike shirt")
	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateAvailability()
	shirtItem.deregister(observerSecond)
	shirtItem.updateAvailability()
}
