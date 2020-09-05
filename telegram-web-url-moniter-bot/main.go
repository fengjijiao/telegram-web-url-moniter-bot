package main

import (
	wm "telegram-web-url-moniter-bot/pkg/moniter"
	tb "telegram-web-url-moniter-bot/pkg/telegram-bot"
)

func init() {
	wm.SetDBFilePath("test.db")
}

func main() {
	go wm.Run(true)
	tb.Run()
}