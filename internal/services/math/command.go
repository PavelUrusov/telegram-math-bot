package math

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-math-bot/pkg/wfclientapi"
)

type Executor interface {
	Execute(input string, chatId int64, solver *Solver) []tgbotapi.Chattable
}

type calculateCommand struct {
	BaseCommand
}

type equationCommand struct {
	BaseCommand
}

type plotCommand struct {
	BaseCommand
}

type startCommand struct {
	BaseCommand
}

type helpCommand struct {
	BaseCommand
}
type BaseCommand struct{}

func (c *helpCommand) Execute(_ string, chatId int64, _ *Solver) []tgbotapi.Chattable {
	return c.CreateStrResponse(chatId, nil, invalidCommandResponse)
}
func (c *startCommand) Execute(_ string, chatId int64, _ *Solver) []tgbotapi.Chattable {
	return c.CreateStrResponse(chatId, nil, startCommandResponse)
}

func (c *calculateCommand) Execute(input string, chatId int64, solver *Solver) []tgbotapi.Chattable {
	sln, err := solver.wfClient.MakeElementaryMathRequest(input)
	return c.CreateSlnResponse(chatId, err, sln)
}

func (c *equationCommand) Execute(input string, chatId int64, solver *Solver) []tgbotapi.Chattable {
	sln, err := solver.wfClient.MakeEquationRequest(input)
	return c.CreateSlnResponse(chatId, err, sln)
}

func (c *plotCommand) Execute(input string, chatId int64, solver *Solver) []tgbotapi.Chattable {
	sln, err := solver.wfClient.MakePlotRequest(input)
	return c.CreateSlnResponse(chatId, err, sln)
}

func (c *BaseCommand) CreateSlnResponse(chatId int64, err error, sln *wfclientapi.Solution) []tgbotapi.Chattable {
	var msgs []tgbotapi.Chattable
	if err != nil {
		msgs = append(msgs, tgbotapi.NewMessage(chatId, errorResponse))
		return msgs
	}
	for _, a := range sln.Answers {
		msgs = append(msgs, tgbotapi.NewMessage(chatId, a))
	}
	for _, i := range sln.ImageStepsURL {
		msgs = append(msgs, tgbotapi.NewPhoto(chatId, tgbotapi.FileURL(i)))
	}
	return msgs
}
func (c *BaseCommand) CreateStrResponse(chatId int64, err error, str string) []tgbotapi.Chattable {
	var msgs []tgbotapi.Chattable
	if err != nil {
		//handle error )))))
		msgs = append(msgs, tgbotapi.NewMessage(chatId, errorResponse))
		return msgs
	}
	msgs = append(msgs, tgbotapi.NewMessage(chatId, str))
	return msgs
}
