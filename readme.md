# MathBot

## Description

MathBot is a Telegram bot capable of solving arithmetic problems, 
equations of varying complexity, and plotting graphs. It utilizes 
the Wolfram API for processing mathematical queries. This is the 
first version of the bot, including a custom implementation of the 
Wolfram API client in Go, as there is no official client available.

## Prerequisites
Before running the code, ensure that you have the following prerequisites:

- [Golang](https://go.dev/dl/)
- [Git](https://git-scm.com/downloads)
- A valid [API key](https://t.me/botfather) for the Telegram API
- A valid [App ID](https://developer.wolframalpha.com/) for the Wolfram API
## Installation
1. To get started with MathBot, follow these steps:
```shell
git clone https://github.com/PavelUrusov/telegram-math-bot.git
```
2. Build the project:
```shell
cd TelegramMathBot
go build -o MathBot cmd/main/main.go
```

## Running
To run MathBot on different operating systems, 
use the following commands, specifying `-appid` for
the Wolfram API and `-apikey` for the Telegram API:

- On Linux
```shell
./TelegramMathBot -appid <YourAppID> -apikey <YourApiKey>
```
- On Windows
 ```shell
TelegramMathBot.exe -appid <YourAppID> -apikey <YourApiKey>
```

## Dependencies
Ensure that you have all the necessary dependencies 
installed, as listed in the `go.mod` file.


## About the Educational Project
MathBot is an educational project designed 
to demonstrate working with the Telegram API and integration 
with external APIs such as Wolfram. This project offers a great opportunity 
for learning Go language and developing bots for Telegram.

# Usage
MathBot offers a variety of commands for interactive mathematical computations:

- `/solve`: Solves equations.
- `/plot`: Plots the graph of a function.
- `/calculate`: Solves arithmetic problems
- `/start`: Displays start-up information about the bot.
- `/help`: Lists all available commands.

## About the  Project
Math Bot is an educational project designed
to demonstrate working with the Telegram API and integration
with external APIs such as Wolfram. This project offers a great opportunity
for learning Go language and developing bots for Telegram.
