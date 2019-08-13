package modulesEmail

import "feifancron/db"

type AhgFfEmailSubscribePoint struct {
	PointId int `orm:"point_id" json:"point_id"`
}

func (*AhgFfEmailSubscribePoint) GetTable() string {
	return "ahg_ff_email_subscribe_point"
}

func GetPoint() int {
	aesp := new(AhgFfEmailSubscribePoint)
	sql := "SELECT point_id FROM " + aesp.GetTable()
	Dbs.GetDb().QueryRow(sql).Scan(&aesp.PointId)
	return aesp.PointId
}

func SetPoint(id int) error {
	sql := "UPDATE ahg_ff_email_subscribe_point SET point_id = ?"
	stmt, err := Dbs.GetDb().Prepare(sql)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	return nil
}