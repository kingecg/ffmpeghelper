package ffmpeghelper

import (
	. "m7s.live/engine/v4"
	"m7s.live/engine/v4/config"
)

type FFMpegHelper struct {
	config.Plugin
	FFMpegPath      string
	HardwareDecoder string
	PreferedFormat  string
}

func (c *FFMpegHelper) OnEvent(event any) {
	plugin.Info("FFMpeg helper loaded")
}

var conf FFMpegHelper
var plugin = InstallPlugin(&conf)
