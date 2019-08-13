package main

import (
	"feifanEmailCron/email"
	"feifanEmailCron/modules"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	data []modulesEmail.AhgEcmsFfExpertThought
	EmailContent string
)

func main() {
	//生成邮件内容
	lastId := emailContent()
	//有新文章的时候发送订阅邮件
	if lastId > 0 {
		fmt.Println("正在批量发送邮件...")
		//记录上一次的发送文章IDlastId
		modulesEmail.SetPoint(lastId)
		//批量发送邮件
		batch(0)
		fmt.Println("邮件发送完成！")
	} else {
		fmt.Println("没有新邮件需要发送！")
	}
}

//生成邮件内容
func emailContent() int {
	//获得订阅的数据
	data = modulesEmail.GetNewsTop()
	if len(data) == 0 {
		return 0
	}
	//打开邮件模板
	templateText := OpenFile("email")
	//获得模板内容
	text := templateText[:]
	//设置funcMap
	funcMap := template.FuncMap{
		"title": strings.Title,
	}
	//实例化一个 html/template
	tmpl, err := template.New("email").Funcs(funcMap).Parse(string(text))
	//生成一个参数容器
	tmp := make(map[string]interface{})
	tmp["list"] = data
	//进行渲染
	fileName := "news" + time.Now().Format("2006-01-02")
	file := "./template/" + fileName + ".tpl"
	//删除原有文件
	os.Remove(file)
	newEmail, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0766);
	if err != nil {
		panic(err)
	}
	defer newEmail.Close()
	err = tmpl.Execute(newEmail, tmp)
	if err != nil {
		panic(err)
	}
	//生成的文件内容
	email := OpenFile(fileName)
	//获得模板内容
	EmailContent = string(email[:])
	if len(data) > 0{
		return data[0].Id
	}
	return 0
}

//批量发送邮件
func batch(last int) {
	where := ""
	if last > 0 {
		where = "id<" + strconv.Itoa(last)
	}
	res := modulesEmail.GetBatch(where, "id DESC", 100)
	if len(res) == 0 {
		return
	}
	fmt.Printf("----------------last %d---------------------\n", last)
	//发送邮件
	for _, v := range res {
		to := make([]string, 0)
		to = append(to, v.Email)
		feifanEmail.SendMailTLS(to, EmailContent)
	}
	//获得最后面的元素
	maxNumber := len(res)
	lastE := res[maxNumber-1]
	batch(lastE.Id)
}

//打开文件
func OpenFile(name string) []byte{
	path := "./template/" + name + ".tpl"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil{
		panic(err)
	}
	return content
}