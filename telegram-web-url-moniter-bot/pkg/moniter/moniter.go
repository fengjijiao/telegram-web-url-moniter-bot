package moniter

import (
	"time"
	dbop "telegram-web-url-moniter-bot/pkg/database"
	httpop "telegram-web-url-moniter-bot/pkg/httper"
	"github.com/jmoiron/sqlx"
	"log"
	"fmt"
	cm "telegram-web-url-moniter-bot/pkg/ctrip-message"
)

var sqliteDB *sqlx.DB

func GetSqliteDB() *sqlx.DB {
	return sqliteDB
}

func Run(forever bool, ch chan cm.CtripMessage) {
	sqliteDB, _ = dbop.OpenOrCreateDatabase(dBFilePath)
	defer sqliteDB.Close()
	for ok := true; ok; ok = forever {
		allUrls, err := dbop.GetUrls(sqliteDB)
		if err != nil {
			log.Fatal(err)
		}
		for _, item := range *allUrls {
			telegramId, err := dbop.GetTelegramIdViaUid(sqliteDB, item.Uid)
			if err != nil {
				log.Println(err)
				continue
			}
			dbop.UpdateLastTimestamp(sqliteDB, item.Rid)
			httpInfo, err := httpop.GetHTTPInfo(item.Url)
			if err != nil {
				//访问出错
				log.Println(err)
				newMessage := cm.CtripMessage {
					ChatId: int64(telegramId),
					Context: fmt.Sprintf("%s\n", err),
				}
				ch<-newMessage
				continue
			}
			if httpInfo.StatusCode != item.LastStatusCode {
				dbop.UpdateLastStatusCode(sqliteDB, httpInfo.StatusCode, item.Rid)
			}
			if httpInfo.Status != item.LastStatus {
				dbop.UpdateLastStatus(sqliteDB, httpInfo.Status, item.Rid)
			}
			if httpInfo.StatusCode != item.LastStatusCode {
				//状态码变化
				newMessage := cm.CtripMessage {
					ChatId: int64(telegramId),
					Context: fmt.Sprintf("%v  %s -  %s.\r\n\r\n----------------\r\n%s\r\nfrom %d to %d\r\n", getStatusEmoji(httpInfo.StatusCode), item.Url, getStatusString(httpInfo.StatusCode), httpInfo.StatusCode, getStatus(httpInfo.StatusCode), item.LastStatusCode, httpInfo.StatusCode),
				}
				ch<-newMessage
				log.Printf("change: %d ==> %d\n", item.LastStatusCode, httpInfo.StatusCode)
				continue
			}
		}
		if forever {
			time.Sleep(64 * time.Second)
		}
	}
}