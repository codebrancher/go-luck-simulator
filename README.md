![Title](./screenshots/wf_game_title.png)

## Description

**The Go Slot Machine** is an open-source, interactive simulation crafted in Go, aimed at demystifying the core mechanics behind slot machines. You can play interactive or simulate high loads of spins to gain insights in game dynamics and statistics.

Designed with modularity in mind, the software allows you to experiment with different components such as random number generators (RNGs) and payout algorithms. Select the randomizer algorithm at the beginning, you can also write own algorithms and easily integrate them.

The goal is to enhance the program by different games, randomness implementations and more statistical methods in a modular way.

## Disclaimer
**Educational Purposes Only**: This project is intended strictly for educational purposes and is not to be used for gambling, betting, or any form of real money play.  It is not a guide for gambling or an endorsement of gambling and should not be used as such.


## Table of Contents
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Games](#games)
- [Features](#features)
- [Customization](#customization)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

### General Features

- **Advanced Randomization**:  Choose from several randomizers, including the standard RNG, LCG, and XorShift, to ensure game fairness and complexity.
- **Interactive Commands**: Utilize commands such as bet, auto, and stats for real-time game management and engagement.
- **Simulation Mode**: Toggle between interactive gameplay and simulation mode to conduct extensive tests or enjoy casual play.
- **Detailed Stats**: Gain detailed insights into randomization processes and game strategies.

## Getting Started

### Prerequisites

- Go 1.22 or later
- make sure your terminal supports Unicode (UTF-8) (required for playing)

### Installation

1. Clone the repository to your local machine:

```sh
git clone https://github.com/codebrancher/go-slot-machine.git
cd go-slot-machine/cmd
```

2. Build the project:

```sh
go build -o slotmachine .
```

## Usage

Run the Go Slot Machine with the following command:

```sh
./slotmachine
```

This command will start the game and you will see the initial game setup screen.

### Game Commands

First you have to select a randomizer, followed by the game. You can select them by entering the number of the randomizer.

In the next step you choose the game which you want to enter, followed by your starting amount (I recommend using higher amounts for simulations to increase granularity). 

### Game Mode Commands

Once the game is running, you will interact with it using the following commands:

- `play`: Starting the playing mode.
- `sim [spins] [bet]`: Perform a simulation.

### Game Commands (play-mode)

Once the game is running, you will interact with it using the following commands:

- `RETURN`: Hitting the return button re-bets with your last bet.
- `bet [bet amount]`: Places a bet of the specified amount.
- `auto [spins] [bet amount]`: Automatically play a specified number of spins with a fixed bet amount.
- `cashout`: Cash out if game has a winning pot.
- `info`: Displays current game information, including available symbols, your current cash balance, and the wild symbol.
- `stats`: Displays current game statistics, such as total spins, total wins, and current cash balance.
- `exit`: Exits the game.

If you are in the `info` or `stats` screen you should make a `bet` or `auto`-bet to get back to the slot machine.

**Remember**: Hitting the `RETURN`-button will re-bet with your last bet.

## Games

Currently two games are supported:

- Wild Fruits: moderate risk, -payout, regular payout, passive playing
- Peak Pursuit: high risk, -rewards, active playing

Generally I aimed to make the RTP (return to player) rates as balanced as I could.

A more detailed description is coming soon. Since then I recommend exploring the games on your own.

## Customization

The program is intended to be as modular as possible to encourage experimentation. Explore the following parts to learn more about their specific functionalities, configurations and potentials:

### Game Engine

The core logic and game mechanics of a specific game are implemented here. [Learn more.](internal/engine/README.md)

### Randomness Module

This module handles all random number generation used in the game. The program supports multiple algorithms to suit different needs. [Read about the randomness module.](internal/randomizer/README.md)

## Contributing

Thank you for your interest in contributing to the Go Slot Machine! This is my first open-source project, and I am excited to learn and grow with contributions from the community.

### How to Contribute

Since this is a learning experience for me as well, here are some basic steps to get involved:

- **Explore the Project**: Take a look around the codebase to see what's already there and where you might be able to add value.
- **Fork and Clone**: Make a fork of the repository, clone it to your machine, and set up a local development environment.
- **Make Changes**: Work on the changes you think would enhance the project. This could be adding comments, refactoring code, or implementing new features.
- **Submit a Pull Request**: Once you're happy with your changes, push them back to your fork and submit a pull request. I'll review it as soon as I can!

### Questions or Suggestions?

If you have questions about the project or need guidance on how to proceed, please feel free to open an issue in the repository. This is a learning experience for all of us, and every question or piece of feedback is valuable!

## License

The Go Slot Machine is open source and available under the GNU General Public License v3.0 (GPL-3.0). This license allows you to use, modify, and distribute the software, but all derivatives of this project must also be released under the same GPL-3.0 license.

For more details, see the [LICENSE](LICENSE) file included in the repository.

## Contact

If you have any questions, feedback, or would like to contribute to the Go Slot Machine project, please feel free to reach out through my [Twitter/X Account](https://twitter.com/codebrancher) or on GitHub.

For more detailed discussions, issue reporting, or to propose new features, consider opening an issue directly on the GitHub project page. This helps us keep track of all project-related discussions in one place.

Your input is incredibly valuable, and I look forward to hearing from you!
