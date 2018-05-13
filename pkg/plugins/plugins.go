package plugins

// RegisterAll registers all plugins
// Add your plugin here
func (p *PluginManager) RegisterAll() {
	p.RegisterPlugin("help", new(help))
	p.RegisterPlugin("fav", new(favorite))
	// planned
	// --------
	// freq emojis
	// sentiment
	// top 10 words
	// most messages (model change)
	// top swears
	// top puncuation
	// longest messages (model change)
	// shortest messages (model change)
}
