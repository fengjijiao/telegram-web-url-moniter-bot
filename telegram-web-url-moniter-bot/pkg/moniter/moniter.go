package moniter

import (
	"time"
	dbop "telegram-web-url-moniter-bot/pkg/database"
	httpop "telegram-web-url-moniter-bot/pkg/httper"
	"github.com/jmoiron/sqlx"
	"log"
)

var sqliteDB *sqlx.DB

func GetSqliteDB() *sqlx.DB {
	return sqliteDB
}

type Message struct {
	ChatId int64
	Context string
}

func Run(forever bool) {
	sqliteDB, _ = dbop.OpenOrCreateDatabase(dBFilePath)
	defer sqliteDB.Close()
	for ok := true; ok; ok = forever {
		allUrls, err := dbop.GetUrls(sqliteDB)
		if err != nil {
			log.Fatal(err)
		}
		for _, item := range *allUrls {
			dbop.UpdateLastTimestamp(sqliteDB, item.Rid)
			httpInfo, err := httpop.GetHTTPInfo(item.Url)
			if err != nil {
				//访问出错
				log.Println(err)
				continue
			}
			if httpInfo.StatusCode != item.LastStatusCode {
				dbop.UpdateLastStatusCode(sqliteDB, httpInfo.StatusCode, item.Rid)
			}
			if httpInfo.Status != item.LastStatus {
				dbop.UpdateLastStatus(sqliteDB, httpInfo.Status, item.Rid)
			}
			if httpInfo.StatusCode != 200 && httpInfo.StatusCode != item.LastStatusCode {
				//状态码变化
				log.Printf("change: %d ==> %d\n", item.LastStatusCode, httpInfo.StatusCode)
				tb.Send(260685916, fmt.Sprintf("change: %d ==> %d\n", item.LastStatusCode, httpInfo.StatusCode))
				continue
			}
		}
		if forever {
			time.Sleep(64 * time.Second)
		}
	}
}