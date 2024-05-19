# Customize Game Mechanics

**Accessing and Modifying Game Settings:**

The game's settings can be modified in the `symbols.go` file located at `/internal/engine/[game]/`. Here, you can adjust the frequency and payout values of each symbol used in the game.

**Steps to Modify Symbol Settings:**

1. Navigate to `/internal/engine/[game]/symbols.go`.
2. Look for the `Symbols` struct which contains the symbol definitions and their properties.
3. Modify the `Frequency` or `Payout` values to adjust how often a symbol appears and what payout it gives when part of a winning line.
4. Save your changes and re-run the game to see the effects.

**Example Modifications:**

- **Increase the Payout for Wild Symbols**: By increasing the payout for wild symbols, you could make the game more rewarding but also adjust the game's balance.
- **Change the Frequency of Bonus Symbols**: Increasing the frequency of bonus symbols might make the game more exciting and dynamic, but also impacts the RTP (Return to Player) greatly.


Feel free to share your experiments and their outcomes with the community or even suggest improvements to the game!

