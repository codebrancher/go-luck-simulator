# Customizable Randomization

#### Available Randomizers

- **Default RNG**: Utilizes Go's standard random number generation based on a pseudo-random algorithm, suitable for general use in games and simulations.

- **LCG (Linear Congruential Generator)**: This generator is fast and has a small memory footprint but may exhibit patterns over long-term use. It's defined by the formula `Xn+1 = (aXn + c) % m`.

- **XorShift**: Known for its excellent randomness properties and speed, this algorithm uses bit shifting to generate high-quality random numbers and is excellent for more complex simulations and is currently in use.

#### How to Switch Randomizers

To switch the randomizer used in the game, modify the instantiation of the RNG in the `main` function of your game. Here is a quick guide:

1. Open the `main.go` file.
2. Locate the RNG initialization line, `rng := randomizer.NewXorShiftRNG(seed)`.
3. Replace `NewXorShiftRNG` with either `NewDefaultRNG` or `NewLCG` as needed.
4. Optionally, adjust the `seed` variable to control the randomness seed.

Example:
```go
// To switch to the LCG randomizer
rng := randomizer.NewLCG(seed)
``` 

## Exploration

Experiment with these different randomizers to observe how they influence the game's behavior. Try using the LCG for a faster but less random experience, or switch to XorShift for high-quality randomness. Additionally, feel free to implement and test your own RNG algorithms to further explore game dynamics.

**Feel free to share your algorithms!**