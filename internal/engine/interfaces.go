package engine

type GameEngine interface {
	Spin() error                                         // Conducts the spin logic without displaying results
	DisplayResults()                                     // Handles all output logic separate from game mechanics
	CalculateWinnings() (description string, amount int) // Calculates winnings based on the last spin
	RequestCommand(input string) (int, error)            // Handles the bet input and validation logic
}

type Observer[T any] interface {
	Update(state T)
}
