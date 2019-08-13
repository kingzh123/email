package modulesEmail

import (
	"feifanEmailCron/db"
	"feifancron/db"
	"time"
)

type AhgEcmsFfExpertThought struct {
	Id            int       `orm:"id" json:"id"`
	Classid       int       `orm:"classid" json:"classid"`
	Ttid          int       `orm:"ttid" json:"ttid"`
	Onclick       int       `orm:"onclick" json:"onclick"`
	Plnum         int       `orm:"plnum" json:"plnum"`
	Totaldown     int       `orm:"totaldown" json:"totaldown"`
	Newspath      string    `orm:"newspath" json:"newspath"`
	Filename      string    `orm:"filename" json:"filename"`
	Userid        int       `orm:"userid" json:"userid"`
	Username      string    `orm:"username" json:"username"`
	Firsttitle    int       `orm:"firsttitle" json:"firsttitle"`
	Isgood        int       `orm:"isgood" json:"isgood"`
	Ispic         int       `orm:"ispic" json:"ispic"`
	Istop         int       `orm:"istop" json:"istop"`
	Isqf          int       `orm:"isqf" json:"isqf"`
	Ismember      int       `orm:"ismember" json:"ismember"`
	Isurl         int       `orm:"isurl" json:"isurl"`
	Truetime      int       `orm:"truetime" json:"truetime"`
	Lastdotime    int       `orm:"lastdotime" json:"lastdotime"`
	Havehtml      int       `orm:"havehtml" json:"havehtml"`
	Groupid       int       `orm:"groupid" json:"groupid"`
	Userfen       int       `orm:"userfen" json:"userfen"`
	Titlefont     string    `orm:"titlefont" json:"titlefont"`
	Titleurl      string    `orm:"titleurl" json:"titleurl"`
	Stb           int       `orm:"stb" json:"stb"`
	Fstb          int       `orm:"fstb" json:"fstb"`
	Restb         int       `orm:"restb" json:"restb"`
	Keyboard      string    `orm:"keyboard" json:"keyboard"`
	Title         string    `orm:"title" json:"title"`
	Newstime      int       `orm:"newstime" json:"newstime"`
	Titlepic      string    `orm:"titlepic" json:"titlepic"`
	Eckuid        int       `orm:"eckuid" json:"eckuid"`
	Description   string    `orm:"description" json:"description"`
	Content       string    `orm:"content" json:"content"`
	Expert        string    `orm:"expert" json:"expert"`
	Image         string    `orm:"image" json:"image"`
	Theme         string    `orm:"theme" json:"theme"`
	PublishDate   time.Time `orm:"publish_date" json:"publish_date"`
	Source        string    `orm:"source" json:"source"`
	IsExpertIndex int       `orm:"is_expert_index" json:"is_expert_index"`
}

func (*AhgEcmsFfExpertThought) GetTable() string {
	return "ahg_ecms_ff_expert_thought"
}

//获得每日定时发送的最新5条记录
func GetNewsTop() []AhgEcmsFfExpertThought {
	arr := make([]AhgEcmsFfExpertThought, 0)
	aet := new(AhgEcmsFfExpertThought)
	sql := "SELECT id,title,image,description FROM " + aet.GetTable() + " WHERE id>? ORDER BY id DESC LIMIT 5"
	stmt, err := Dbs.GetDb().Prepare(sql)
	if err != nil {
		panic(err)
	}
	//获得上次发送的节点
	point := GetPoint()
	rows, err := stmt.Query(point)
	if err != nil {
		panic(err)
	}
	ossDomain := DbsEmail.GetConfig().Conf.Oss.Domain
	for rows.Next(){
		aeet := new(AhgEcmsFfExpertThought)
		rows.Scan(&aeet.Id, &aeet.Title, &aeet.Image, &aeet.Description)
		aeet.Image = ossDomain + aeet.Image
		arr = append(arr, *aeet)
	}
	return arr
}
