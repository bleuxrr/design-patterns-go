package main

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 15,
		},
	}
}

func (n *nike) makeShort() iShort {
	return &nikeShort{
		short: short{
			logo: "nike",
			size: 15,
		},
	}
}
