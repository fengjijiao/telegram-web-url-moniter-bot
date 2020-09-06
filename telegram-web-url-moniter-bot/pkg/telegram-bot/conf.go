package telegram_bot

import (
)

var BOT_APIKEY string

func init() {
	BOT_APIKEY = ""	
}

func SetBOTAPIKEY(apikey string) {
	BOT_APIKEY = apikey
}