package conf

import "github.com/BurntSushi/toml"

type Conf struct {
	MysqlConf mySqlConf `toml:"mysql"`
	ServerConf serverConf `toml:"server"`
}

type serverConf struct {
	Port int `toml:"port"`
}

type mySqlConf struct {
	UserName string `toml:"user_name"`
	Password string `toml:"password"`
	Address string `toml:"address"`
	DbName string `toml:"db_name"`
	MaxIdleConn int `toml:"max_idle_conn"`
	MaxOpenConn int `toml:"max_open_conn"`
}

var DefaultConf Conf


func (c *mySqlConf) GetDsn() string {
	return c.UserName + ":" + c.Password + "@(" + c.Address + ")/" + c.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func init() {
	_, err := toml.DecodeFile("conf.toml", &DefaultConf)
	if err != nil {
		panic(err)
	}
}