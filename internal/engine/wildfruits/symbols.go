package wildfruits

type Symbol struct {
	Representation string
	Payout         int
	Frequency      int
}

var Symbols = []Symbol{
	{"🍉", 10, 5},
	{"🥥", 5, 7},
	{"🌟", 5, 7}, // wild triggers free games
	{"🫐", 3, 8},
	{"🍇", 3, 8},
	{"🍏", 2, 10},
	{"🍓", 2, 10},
	{"🍊", 2, 10},
	{"🍒", 1, 35},
}
