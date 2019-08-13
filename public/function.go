package publicEmail

import (
	"feifanEmailCron/configs"
	"github.com/gohouse/converter"
	"gopkg.in/yaml.v2"
	"math/rand"
	"os"
	"time"
)

func RandInt(min, max int) int {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return r.Intn(max-min) + min
}

//database table convert to model
func RunModule(table string)  {
	cvt := converter.NewTable2Struct()
	cvt.Config(&converter.T2tConfig{
		RmTagIfUcFirsted: false,
		TagToLower: false,
		UcFirstOnly: false,
	})

	fileName := CamelString(table)

	err := cvt.Table(table).
		EnableJsonTag(true).
		PackageName("modules").
		TagKey("orm").
		RealNameMethod(fileName).
		SavePath("./modules/" + fileName + ".go").
		Dsn("").
		// 执行
		Run()
	if err != nil {
		panic(err)
	}
}

func ReadYamlConfig(path string)  (*configsEmail.Config,error){
	conf := &configsEmail.Config{}
	if f, err := os.Open(path); err != nil {
		return nil,err
	} else {
		defer f.Close()
		err = yaml.NewDecoder(f).Decode(conf)
		if err != nil {
			panic(err)
		}
	}
	return conf,nil
}

func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
