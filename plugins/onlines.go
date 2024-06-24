package plugins

func (p *Plugins) Online() {
}

func (p *Plugins) OnlineInfo() CollectionInfo {
	return CollectionInfo{
		Name:        "online",
		Description: "Check who is online",
	}
}
