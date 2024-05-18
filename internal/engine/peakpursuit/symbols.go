package peakpursuit

type Symbol struct {
	Representation string
	Payouts        []int
	Frequency      int
}

var Symbols = []Symbol{
	{"💫", []int{1, 2, 4, 7, 10}, 20},
	{"🪐", []int{3, 8, 16, 29, 50}, 12},
	{"🚀", []int{10, 40, 100, 200, 350}, 4},
	{"👽", []int{0}, 12},
	{"🌠", []int{0}, 2},
	{"  ", []int{0}, 50},
}
