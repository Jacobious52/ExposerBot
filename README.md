# ExposerBot
exposer is a modular telegram bot for exposing metrics of chats

## Add a plugin
- goto pkg/plugins/
- make a new go file for example `sentiment.go`
- implement the `Exposer` interface
- register your struct in `plugins.go` with  `p.RegisterPlugin("sentiment", new(sentiment))`
- you can now summon your plugin through the bot with `/expose_setiment`
