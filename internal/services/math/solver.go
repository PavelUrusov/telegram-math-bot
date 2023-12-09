package math

import (
	"crypto/sha256"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"telegram-math-bot/pkg/wfclientapi"
)

type Solver struct {
	wfClient      wfclientapi.WfClient
	tgClient      *tgbotapi.BotAPI
	OffsetUpdate  int
	TimeoutUpdate int
	Debug         bool
	cache         CacheManager
	commands      map[string]Executor
}

func NewSolver(wolframToken, telegramToken string) (*Solver, error) {
	solver := &Solver{}
	wfClient, err := wfclientapi.NewWfAPIClient(wolframToken)
	if err != nil {
		return nil, err
	}
	tgClient, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		return nil, err
	}
	solver.tgClient = tgClient
	solver.wfClient = wfClient
	solver.TimeoutUpdate = DefaultTimeoutUpdate
	solver.OffsetUpdate = DefaultOffsetUpdate
	solver.Debug = false
	solver.cache = NewCacheManager()

	solver.initCommands()

	return solver, nil
}

func (s *Solver) Listen() {
	updates := s.tgClient.GetUpdatesChan(s.UpdateConfig())
	for update := range updates {
		go s.handleUpdate(&update)
	}
}

func (s *Solver) handleUpdate(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}
	var msgs []tgbotapi.Chattable
	if s.Debug {
		logrus.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
	cmd, args := s.getCommandAndArgs(update)
	msgs = s.ExecuteCommand(cmd, args, update.Message.Chat.ID)
	s.sendMessage(msgs)
}

func (s *Solver) UpdateConfig() tgbotapi.UpdateConfig {
	cfg := tgbotapi.NewUpdate(s.OffsetUpdate)
	cfg.Timeout = s.TimeoutUpdate

	return cfg
}

func (s *Solver) ExecuteCommand(cmd Executor, input string, chatId int64) []tgbotapi.Chattable {
	hashKey := hashInput(input)
	if cachedResponse, found := s.cache.GetFromCache(hashKey); found {
		return cachedResponse
	}
	response := cmd.Execute(input, chatId, s)
	s.cache.AddToCache(hashKey, response)
	return response
}

func hashInput(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	return fmt.Sprintf("%x", h.Sum(nil))
}
func (s *Solver) initCommands() {
	s.commands = make(map[string]Executor)
	s.commands[CalculateCommandID] = &calculateCommand{}
	s.commands[SolveEquationCommandID] = &equationCommand{}
	s.commands[PlotCommandID] = &plotCommand{}
	s.commands[StartCommandID] = &startCommand{}
	s.commands[HelpCommandID] = &helpCommand{}
}

func (s *Solver) сommand(upd *tgbotapi.Update) (Executor, bool) {
	cmd, ok := s.commands[upd.Message.Command()]
	if ok {
		return cmd, true
	}
	cmd, _ = s.commands[HelpCommandID]
	return cmd, false
}

func (s *Solver) getCommandAndArgs(update *tgbotapi.Update) (Executor, string) {
	cmd, ok := s.сommand(update)
	if !ok {
		cmd, _ = s.commands[HelpCommandID]
	}
	args := update.Message.CommandArguments()
	return cmd, args
}

func (s *Solver) sendMessage(msgs []tgbotapi.Chattable) {
	for _, m := range msgs {
		if _, err := s.tgClient.Send(m); err != nil {
			logrus.Errorf("Error sending message: %v", err)
		}
	}
}
