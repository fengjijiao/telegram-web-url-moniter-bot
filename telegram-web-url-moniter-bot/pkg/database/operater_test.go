package database

import (
	"testing"
)

func TestInsertUrl(t *testing.T) {
	t.Parallel()
	sqliteDatabase, _ := OpenOrCreateDatabase("test1.db")
	k1 := &UrlVar {
		Url : "http://y.qq.com",
		Uid : 2,
	}
	err := k1.InsertUrl(sqliteDatabase)
	if err != nil {
		t.Error(err)
	}
	k2 := &UrlVar {
		Url : "http://kf.qq.com",
		Uid : 2,
	}
	err = k2.InsertUrl(sqliteDatabase)
	if err != nil {
		t.Error(err)
	}
	RemoveDatabase("test1.db")
}

func TestGetUrlsViaUid(t *testing.T) {
	t.Parallel()
	sqliteDatabase, _ := OpenOrCreateDatabase("test2.db")
	res, err := GetUrlsViaUid(sqliteDatabase, 2)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v\n", res)
	RemoveDatabase("test2.db")
}

// func TestGetUrlsViaUid2(t *testing.T) {
//  t.Parallel()
//  sqliteDatabase, _ := OpenOrCreateDatabase("test3.db")
// 	searchSQL := "SELECT * FROM url WHERE `uid` = ?"
// 	row, err := sqliteDatabase.Queryx(searchSQL, 2)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer row.Close()
// 	for row.Next() { // Iterate and fetch the records from result cursor
// 		var item UrlVar
// 		err := row.StructScan(&item)
//         if err != nil {
//             t.Fatal(err)
//         }
// 		t.Logf("%v\n", item)
// 	}
//  RemoveDatabase("test3.db")
// }

func TestUrlExists(t *testing.T) {
	t.Parallel()
	sqliteDatabase, _ := OpenOrCreateDatabase("test4.db")
	res := UrlExists(sqliteDatabase, "http://kf.qq.com", 2)
	t.Logf("result:%t\n", res)
	RemoveDatabase("test4.db")
}

func TestDeleteUrl(t *testing.T) {
	t.Parallel()
	sqliteDatabase, _ := OpenOrCreateDatabase("test5.db")
	err := DeleteUrl(sqliteDatabase, "http://kf.qq.com", 2)
	if err != nil {
		t.Error(err)
	}
	RemoveDatabase("test5.db")
}

func TestDeleteUrls(t *testing.T) {
	t.Parallel()
	sqliteDatabase, _ := OpenOrCreateDatabase("test6.db")
	err := DeleteUrls(sqliteDatabase, 2)
	if err != nil {
		t.Error(err)
	}
	RemoveDatabase("test6.db")
}

func TestInsertUser(t *testing.T) {
	t.Parallel()
	sqliteDatabase, _ := OpenOrCreateDatabase("test7.db")
	user := &UserVar {
		TelegramId : 2222,
	}
	err := user.InsertUser(sqliteDatabase)
	if err != nil {
		t.Error(err)
	}
	RemoveDatabase("test7.db")
}

func TestUserExists(t *testing.T) {
	t.Parallel()
	sqliteDatabase, _ := OpenOrCreateDatabase("test8.db")
	user := &UserVar {
		TelegramId : 2222,
	}
	err := user.InsertUser(sqliteDatabase)
	if err != nil {
		t.Error(err)
	}
	res := UserExists(sqliteDatabase, 2222)
	t.Logf("result: %t\n", res)
	RemoveDatabase("test8.db")
}

func TestGetUserInfoViaTelegramId(t *testing.T) {
	t.Parallel()
	sqliteDatabase, _ := OpenOrCreateDatabase("test9.db")
	user := &UserVar {
		TelegramId : 2222,
	}
	err := user.InsertUser(sqliteDatabase)
	if err != nil {
		t.Error(err)
	}
	res, err := GetUserInfoViaTelegramId(sqliteDatabase, 2222)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", res)
	RemoveDatabase("test9.db")
}

func TestGetUidViaTelegramId(t *testing.T) {
	t.Parallel()
	sqliteDatabase, _ := OpenOrCreateDatabase("test10.db")
	user := &UserVar {
		TelegramId : 2222,
	}
	err := user.InsertUser(sqliteDatabase)
	if err != nil {
		t.Error(err)
	}
	res, err := GetUidViaTelegramId(sqliteDatabase, 2222)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("uid: %d\n", res)
	RemoveDatabase("test10.db")
}