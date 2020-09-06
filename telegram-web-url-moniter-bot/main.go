package main

import (
	wm "telegram-web-url-moniter-bot/pkg/moniter"
	tb "telegram-web-url-moniter-bot/pkg/telegram-bot"
	cm "telegram-web-url-moniter-bot/pkg/ctrip-message"
	"github.com/jinzhu/configor"
	"flag"
)

var ConfigFilePath string
var Config = struct {
	DBFilePath string `default:"test.db"`
	BotApiKey string `default:""`
}{}

func init() {
	flag.StringVar(&ConfigFilePath, "c", "config.yaml", "config location.")
	flag.Parse()
	configor.Load(&Config, ConfigFilePath)
	wm.SetDBFilePath(Config.DBFilePath)
	tb.SetBOTAPIKEY(Config.BotApiKey)
}

var ch chan cm.CtripMessage
func main() {
	ch = make(chan cm.CtripMessage)
	go wm.Run(true, ch)
	go tb.Run()
	for {
		select {
			case res := <-ch:
				tb.Send(res.ChatId, res.Context)
			default:
			    continue
    	}
	}
}