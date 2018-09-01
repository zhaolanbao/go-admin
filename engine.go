package goAdmin

import (
	"goAdmin/plugins"
	"goAdmin/framework"
	"goAdmin/modules/config"
)

type Engine struct {
	PluginsList []plugins.Plugin
	Config      config.Config
}

func Default() *Engine {
	return new(Engine)
}

func (eng *Engine) Use(fw framework.WebFrameWork, router interface{}) error {
	return fw.Use(router, eng.PluginsList)
}

func (eng *Engine) AddPlugins(plugs ... plugins.Plugin) *Engine {

	for _, plug := range plugs {
		plug.InitPlugin(eng.Config)
	}

	eng.PluginsList = append(eng.PluginsList, plugs...)
	return eng
}

func (eng *Engine) AddConfig(cfg config.Config) *Engine {
	eng.Config = cfg
	return eng
}