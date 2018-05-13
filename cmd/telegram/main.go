package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Jacobious52/expose/pkg/plugins"
	"github.com/Jacobious52/expose/pkg/storage"
	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	tb "gopkg.in/tucnak/telebot.v2"
)

var telegramToken = kingpin.Flag("token", "telegram bot token").Envar("TBTOKEN").Required().String()
var dataStorePath = kingpin.Flag("store", "path to save and read store file").Default("/usr/share/store/store.json").String()

func main() {
	kingpin.Parse()

	// create the bot
	bot, err := tb.NewBot(tb.Settings{
		Token: *telegramToken,
		Poller: &tb.LongPoller{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		log.Fatalf("failed to created bot: %v", err)
	}
	log.Infoln("created bot!")

	// load data and start syncing to disk
	dataStore := storage.NewDataStore(*dataStorePath)
	err = dataStore.Load()
	if err != nil {
		log.Warningf("data not loaded. %v", err)
	}
	dataStore.StartSync(nil)

	// load plugins
	log.Infoln("loading plugins")
	pluginManager := plugins.NewPluginManager(bot, dataStore)
	pluginManager.RegisterAll()

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, fmt.Sprintf("Hello, %v. Welcome to the bot!", m.Sender.FirstName))
	})

	// handle reading of messages
	bot.Handle(tb.OnText, func(m *tb.Message) {
		log.WithField("chat", m.Chat.Title).Infoln("message from", m.Sender.Username)
		basicWords := strings.Split(m.Text, " ")
		for _, word := range basicWords {
			dataStore.AddWord(m.Sender.Username, word)
		}
	})

	log.Infoln("starting bot")
	bot.Start()
}
