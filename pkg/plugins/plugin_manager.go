package plugins

import (
	"github.com/Jacobious52/expose/pkg/exposer"
	"github.com/Jacobious52/expose/pkg/storage"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

// PluginManager manages all the plugins with the datastore and bot
type PluginManager struct {
	// Plugins are all the exposers plugged in
	plugins map[string]exposer.Exposer
	bot     *tb.Bot
	store   storage.ReadStore
}

// NewPluginManager creates a new plugin manager
func NewPluginManager(bot *tb.Bot, store storage.ReadStore) *PluginManager {
	return &PluginManager{
		plugins: make(map[string]exposer.Exposer),
		bot:     bot,
		store:   store,
	}
}

// RegisterPlugin adds the plugin to the plugins list
func (p *PluginManager) RegisterPlugin(name string, e exposer.Exposer) {
	err := e.Setup()
	if err != nil {
		log.Errorf("failed to load plugin %v: %v", name, err)
		return
	}
	p.plugins[name] = e

	// create the endpoint
	p.bot.Handle("/expose_"+name, func(m *tb.Message) {
		log.Infoln("plugin starting:", name)
		str, err := e.Expose(p.store)
		if err != nil {
			log.Errorf("failed to expose metric %v: %v", name, err)
			return
		}
		p.bot.Send(m.Chat, str)
	})
}
