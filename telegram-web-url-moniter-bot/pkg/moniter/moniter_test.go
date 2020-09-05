package moniter

import(
	"testing"
	dbop "telegram-web-url-moniter-bot/pkg/database"
)

func TestRun(t *testing.T) {
	t.Parallel()
	sqliteDB, _ = dbop.OpenOrCreateDatabase(dBFilePath)
	user := &dbop.UserVar {
		TelegramId : 2222,
	}
	err := user.InsertUser(sqliteDB)
	if err != nil {
		t.Error(err)
	}
	k1 := &dbop.UrlVar {
		Url : "http://www.google.com",
		Uid : 1,
	}
	err = k1.InsertUrl(sqliteDB)
	if err != nil {
		t.Error(err)
	}
	k2 := &dbop.UrlVar {
		Url : "http://www.facebook.com",
		Uid : 1,
	}
	err = k2.InsertUrl(sqliteDB)
	if err != nil {
		t.Error(err)
	}
	k3 := &dbop.UrlVar {
		Url : "https://httpstat.us/505",
		Uid : 1,
	}
	err = k3.InsertUrl(sqliteDB)
	if err != nil {
		t.Error(err)
	}
	Run(false)
	dbop.RemoveDatabase("test.db")
}