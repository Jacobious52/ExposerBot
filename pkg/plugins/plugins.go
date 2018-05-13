package plugins

// RegisterAll registers all plugins
// Add your plugin here
func (p *PluginManager) RegisterAll() {
	p.RegisterPlugin("help", new(help))
	p.RegisterPlugin("fav", new(favorite))
}
