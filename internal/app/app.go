package app

import "telegram-math-bot/internal/services/math"

type WolframAppID string
type TelegramAPIKey string

type App interface {
	Run()
}
type app struct {
	solver *math.Solver
}

func NewApp(wolframAppID WolframAppID, telegramAPIKey TelegramAPIKey) App {
	solver, err := math.NewSolver(string(wolframAppID), string(telegramAPIKey))
	if err != nil {
		panic(err.Error())
	}
	newApp := &app{solver: solver}
	return newApp
}
func (a *app) Run() {
	a.solver.Listen()
}
