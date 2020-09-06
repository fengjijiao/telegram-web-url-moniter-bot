package telegram_bot

import (
	"time"
	tb "gopkg.in/tucnak/telebot.v2"
	wm "telegram-web-url-moniter-bot/pkg/moniter"
	dbop "telegram-web-url-moniter-bot/pkg/database"
	"log"
	"fmt"
)

var b *tb.Bot

func Run() {
	var err error
	b, err = tb.NewBot(tb.Settings{
		Token: BOT_APIKEY,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}
	b.Handle("/start", func(m *tb.Message) {
		if !dbop.UserExists(wm.GetSqliteDB(), m.Sender.ID) {
			newUser := &dbop.UserVar {
				TelegramId: m.Sender.ID,
			}
			err := newUser.InsertUser(wm.GetSqliteDB())
			if err != nil {
				b.Send(m.Sender, fmt.Sprintf("%s\r\n", err))
			}
			b.Send(m.Sender, "Welcome to use this bot.")
		}
	})
	
	b.Handle("/list", func(m *tb.Message) {
		Uid, err := dbop.GetUidViaTelegramId(wm.GetSqliteDB(), m.Sender.ID)
		if err != nil {
			b.Send(m.Sender, fmt.Sprintf("%s\r\n", err))
			return
		}
		urls, err := dbop.GetUrlsViaUid(wm.GetSqliteDB(), Uid)
		if err != nil {
			b.Send(m.Sender, fmt.Sprintf("%s\r\n", err))
			return
		}
		if len(*urls) <= 0 {
			b.Send(m.Sender, "not found record.")
		}else {
			receive := "The URLs in all your monitoring are as follows: \r\n"
			for i, item := range *urls {
				receive += fmt.Sprintf("%d. [%d] %s\r\n", i, item.LastStatusCode, item.Url)
			}
			b.Send(m.Sender, receive)
		}
	})
	
	b.Handle("/add", func(m *tb.Message) {
		if m.Payload == "" {
			b.Send(m.Sender, "/add url is the format.")
		}else {
			Uid, err := dbop.GetUidViaTelegramId(wm.GetSqliteDB(), m.Sender.ID)
			if err != nil {
				b.Send(m.Sender, fmt.Sprintf("%s\r\n", err))
				return
			}
			if dbop.UrlExists(wm.GetSqliteDB(), m.Payload, Uid) {
				b.Send(m.Sender, "url already exists.")
			}else {
				newUrl := &dbop.UrlVar {
					Url: m.Payload,
					Uid: Uid,
				}
				err = newUrl.InsertUrl(wm.GetSqliteDB())
				if err != nil {
					b.Send(m.Sender, fmt.Sprintf("Add failed, cause by: %s\r\n", err))
				}else {
					b.Send(m.Sender, "Added successfully.")
				}
			}
		}
	})
	
	b.Handle("/del", func(m *tb.Message) {
		if m.Payload == "" {
			b.Send(m.Sender, "/del url is the format.")
		}else {
			Uid, err := dbop.GetUidViaTelegramId(wm.GetSqliteDB(), m.Sender.ID)
			if err != nil {
				b.Send(m.Sender, fmt.Sprintf("%s\r\n", err))
				return
			}
			if dbop.UrlExists(wm.GetSqliteDB(), m.Payload, Uid) {
				err = dbop.DeleteUrl(wm.GetSqliteDB(), m.Payload, Uid)
				if err != nil {
					b.Send(m.Sender, fmt.Sprintf("Delete failed, cause by: %s\r\n", err))
				}else {
					b.Send(m.Sender, "Deleted successfully.")
				}
			}else {
				b.Send(m.Sender, "url does not exist.")
			}
		}
	})
	
	b.Handle("/drop", func(m *tb.Message) {
		Uid, err := dbop.GetUidViaTelegramId(wm.GetSqliteDB(), m.Sender.ID)
		if err != nil {
			b.Send(m.Sender, fmt.Sprintf("%s\r\n", err))
			return
		}
		err = dbop.DeleteUrls(wm.GetSqliteDB(), Uid)
		if err != nil {
			b.Send(m.Sender, fmt.Sprintf("Drop failed, cause by: %s\r\n", err))
		}else {
			b.Send(m.Sender, "Droped successfully.")
		}
	})
	
	b.Start()
}

func Send(chatId int64, context string) {
	b.Send(tb.ChatID(chatId), context)
}