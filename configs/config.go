package configsEmail


//备注 具体定义是根据 json 定义的说明取值的
type Config struct {
	Conf Conf  `yaml:"conf"`
}

type Conf struct {
	Mysql Mysql `yaml:"mysql"`
	Email Email `yaml:"email"`
	Oss Oss `yaml:"oss"`
}

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Database string `yaml:"database"`
	Port string `yaml:"port"`
	Charset string `yaml:"charset"`
	Net string `yaml:"net"`
}

type Email struct {
	Account string `yaml:"account"`
	Password string `yaml:"password"`
	Smtp string `yaml:"smtp"`
	Port int `yaml:"port"`
}

type Oss struct {
	Domain string `yaml:"domain"`
}