package database

import (
	//"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UrlVar struct {
	Rid int `db:"rid"`
	Url string `db:"url"`
	CreateTimestamp int `db:"create_timestamp"`
	LastTimestamp int `db:"last_timestamp"`
	Uid int `db:"uid"`
	LastStatusCode int `db:"last_statuscode"`
	LastStatus string `db:"last_status"`
}

type UserVar struct {
	Uid int `db:"uid"`
	TelegramId int `db:"telegram_id"`
	StartTimestamp int `db:"start_timestamp"`
	LastTimestamp int `db:"last_timestamp"`
}

func UrlExists(db * sqlx.DB, url string, uid int) bool {
    sqlStmt := "SELECT * FROM url WHERE `url` = ? AND `uid` = ?"
    count := 0
    rows, err := db.Query(sqlStmt, url, uid)
	if err != nil {
		return true
	}
    for rows.Next() {  // Can I just check if rows is non-zero somehow?
        count++
    }
    return count != 0
}

func UserExists(db * sqlx.DB, telegram_id int) bool {
    sqlStmt := "SELECT * FROM user WHERE `telegram_id` = ?"
    count := 0
    rows, err := db.Query(sqlStmt, telegram_id)
	if err != nil {
		return true
	}
    for rows.Next() {  // Can I just check if rows is non-zero somehow?
        count++
    }
    return count != 0
}

func GetUserInfoViaTelegramId(db *sqlx.DB, telegram_id int) (*UserVar, error) {
	searchSQL := "SELECT * FROM user WHERE `telegram_id` = ? LIMIT 1"
	rows, err := db.Queryx(searchSQL, telegram_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result UserVar
	rows.Next()
	err = rows.StructScan(&result)
    if err != nil {
        return nil, err
    }
	return &result, nil
}

func GetUserInfoViaUid(db *sqlx.DB, uid int) (*UserVar, error) {
	searchSQL := "SELECT * FROM user WHERE `uid` = ? LIMIT 1"
	rows, err := db.Queryx(searchSQL, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result UserVar
	rows.Next()
	err = rows.StructScan(&result)
    if err != nil {
        return nil, err
    }
	return &result, nil
}

func GetUidViaTelegramId(db *sqlx.DB, telegram_id int) (int, error) {
	userInfo, err := GetUserInfoViaTelegramId(db, telegram_id)
	if err != nil {
		return 0, err
	}
	return userInfo.Uid, nil
}

func GetTelegramIdViaUid(db *sqlx.DB, uid int) (int, error) {
	userInfo, err := GetUserInfoViaUid(db, uid)
	if err != nil {
		return 0, err
	}
	return userInfo.TelegramId, nil
}

func (uv *UserVar) InsertUser(db *sqlx.DB) error {
	if UserExists(db, uv.TelegramId) {
		return errors.New("user was exists.")
	}
	insertSQL := `INSERT INTO user (telegram_id,start_timestamp,last_timestamp) VALUES (?,?,?); `
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(uv.TelegramId, timestampNano(), timestampNano())
	if err != nil {
		return err
	}
	return nil
}

func (uv *UrlVar) InsertUrl(db *sqlx.DB) error {
	if len(uv.Url) > 256 {
		return errors.New("url length more then 256.")
	}
	if UrlExists(db, uv.Url, uv.Uid) {
		return errors.New("url was exists.")
	}
	insertSQL := `INSERT INTO url (url,create_timestamp,last_timestamp,uid,last_statuscode,last_status) VALUES (?, ?, ?, ?, 200, '200 OK'); `
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(uv.Url, timestampNano(), timestampNano(), uv.Uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLastTimestamp(db *sqlx.DB, rid int) error {
	updateSQL := "UPDATE user SET `last_timestamp`=? WHERE `rid` = ?"
	_, err := db.Exec(updateSQL, timestampNano(), rid)
	if err != nil {
		return err
	}
	return nil
}


func UpdateLastStatus(db *sqlx.DB, last_status string, rid int) error {
	updateSQL := "UPDATE url SET `last_status` = ? WHERE `rid` = ?"
	_, err := db.Exec(updateSQL, last_status, rid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLastStatusCode(db *sqlx.DB, last_statuscode, rid int) error {
	updateSQL := "UPDATE url SET `last_statuscode` = ? WHERE `rid` = ?"
	_, err := db.Exec(updateSQL, last_statuscode, rid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUrl(db *sqlx.DB, url string, uid int) error {
	deleteSQL := "DELETE FROM url WHERE `url` = ? AND `uid` = ?"
	_, err := db.Exec(deleteSQL, url, uid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUrls(db *sqlx.DB, uid int) error {
	deleteSQL := "DELETE FROM url WHERE `uid` = ?"
	_, err := db.Exec(deleteSQL, uid)
	if err != nil {
		return err
	}
	return nil
}

func GetUrlsViaUid(db *sqlx.DB, uid int) (*[]UrlVar, error) {
	searchSQL := "SELECT * FROM url WHERE `uid` = ?"
	rows, err := db.Queryx(searchSQL, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []UrlVar
	for rows.Next() { // Iterate and fetch the records from result cursor
		var item UrlVar
		err := rows.StructScan(&item)
        if err != nil {
            return nil, err
        }
		result = append(result, item)
	}
	return &result, nil
}

func GetUrls(db *sqlx.DB) (*[]UrlVar, error) {
	searchSQL := "SELECT * FROM url"
	rows, err := db.Queryx(searchSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []UrlVar
	for rows.Next() { // Iterate and fetch the records from result cursor
		var item UrlVar
		err := rows.StructScan(&item)
        if err != nil {
            return nil, err
        }
		result = append(result, item)
	}
	return &result, nil
}

func CreateTable(db *sqlx.DB, create_sql string) error {
	statement, err := db.Prepare(create_sql) // Prepare SQL Statement
	if err != nil {
		return err
	}
	statement.Exec() // Execute SQL Statements
	return nil
}