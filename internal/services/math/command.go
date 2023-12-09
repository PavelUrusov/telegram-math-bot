package math

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-math-bot/pkg/wfclientapi"
)

type Executor interface {
	Execute(message *tgbotapi.Message, solver *Solver) []tgbotapi.Chattable
	BaseExecute(chatId int64, err error, sln *wfclientapi.Solution) []tgbotapi.Chattable
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

func (c *helpCommand) Execute(message *tgbotapi.Message, s *Solver) []tgbotapi.Chattable {
	sln := &wfclientapi.Solution{Answers: []string{invalidCommandResponse}}
	s.cache.AddToCache(hashString(message.Text), sln)
	return c.BaseExecute(message.Chat.ID, nil, sln)
}
func (c *startCommand) Execute(message *tgbotapi.Message, s *Solver) []tgbotapi.Chattable {
	sln := &wfclientapi.Solution{Answers: []string{startCommandResponse}}
	s.cache.AddToCache(hashString(message.Text), sln)
	return c.BaseExecute(message.Chat.ID, nil, sln)
}

func (c *calculateCommand) Execute(message *tgbotapi.Message, s *Solver) []tgbotapi.Chattable {
	sln, err := s.wfClient.MakeElementaryMathRequest(message.CommandArguments())
	s.cache.AddToCache(hashString(message.Text), sln)
	return c.BaseExecute(message.Chat.ID, err, sln)
}

func (c *equationCommand) Execute(message *tgbotapi.Message, s *Solver) []tgbotapi.Chattable {
	sln, err := s.wfClient.MakeEquationRequest(message.CommandArguments())
	s.cache.AddToCache(hashString(message.Text), sln)
	return c.BaseExecute(message.Chat.ID, err, sln)
}

func (c *plotCommand) Execute(message *tgbotapi.Message, s *Solver) []tgbotapi.Chattable {
	sln, err := s.wfClient.MakePlotRequest(message.CommandArguments())
	s.cache.AddToCache(hashString(message.Text), sln)
	return c.BaseExecute(message.Chat.ID, err, sln)
}

func (c *BaseCommand) BaseExecute(chatId int64, err error, sln *wfclientapi.Solution) []tgbotapi.Chattable {
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
