# Customizable Randomization

#### Available Randomizers

- **Default RNG**: Utilizes Go's standard random number generation based on a pseudo-random algorithm, suitable for general use in games and simulations.

- **LCG (Linear Congruential Generator)**: This generator is fast and has a small memory footprint but may exhibit patterns over long-term use. It's defined by the formula `Xn+1 = (aXn + c) % m`.

- **XorShift**: Known for its excellent randomness properties and speed, this algorithm uses bit shifting to generate high-quality random numbers and is excellent for more complex simulations and is currently in use.

## Exploration

Experiment with these different randomizers to observe how they influence the game's behavior. Try using the LCG for a faster but less random experience, or switch to XorShift for high-quality randomness. Additionally, feel free to implement and test your own RNG algorithms to further explore game dynamics.

**Feel free to share your algorithms!**