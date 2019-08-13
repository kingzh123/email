package modulesEmail

import (
	"feifancron/db"
	"strconv"
	"time"
)

type AhgFfEmailSubscribe struct {
	Id        int       `orm:"id" json:"id"`
	Email     string    `orm:"email" json:"email"`
	CreatedAt time.Time `orm:"created_at" json:"created_at"`
}

func (*AhgFfEmailSubscribe) GetTable() string {
	return "ahg_ff_email_subscribe"
}

//批量获得数据
func GetBatch(where string, sort string, limit int) []AhgFfEmailSubscribe {
	arr := make([]AhgFfEmailSubscribe, 0)
	ahg := new(AhgFfEmailSubscribe)
	sql := ""
	if len(where) > 0 {
		sql = `SELECT id,email FROM ` + ahg.GetTable() + ` WHERE ` + where + ` ORDER BY ` + sort + ` LIMIT ` + strconv.Itoa(limit)
	} else {
		sql = `SELECT id,email FROM ` + ahg.GetTable() + ` ORDER BY ` + sort + ` LIMIT ` + strconv.Itoa(limit)
	}
	rows, err := Dbs.GetDb().Query(sql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		a := new(AhgFfEmailSubscribe)
		rows.Scan(&a.Id, &a.Email)
		arr = append(arr, *a)
	}
	return arr
}