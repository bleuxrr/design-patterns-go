package main

import "fmt"

func main() {
	veggiemania := &veggiemania{}
	//add cheese
	veggiePizzaWithCheese := &cheeseTopping{
		pizza: veggiemania,
	}
	//add tomato
	veggiePizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: veggiePizzaWithCheese,
	}
	fmt.Printf("the price of veggieMania pizza with cheese and tomato toppings is %d\n", veggiePizzaWithCheeseAndTomato.getPrice())

	peppyPaneerPizza := &peppyPaneer{}

	peppyPaneerPizzaWithCheese := &cheeseTopping{
		pizza: peppyPaneerPizza,
	}
	fmt.Printf("the price of peppyPaneer pizza with cheese topping is %d\n", peppyPaneerPizzaWithCheese.getPrice())
}
