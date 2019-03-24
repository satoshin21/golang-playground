package lessons

import (
	"fmt"
	"github.com/go-ini/ini"
)

type ConfigList struct {

	Port int
	DbName string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		cfg.Section("web").Key("port").MustInt(),
		cfg.Section("db").Key("name").MustString("example.sql"),
		cfg.Section("db").Key("driver").String()}
}

func main() {
	fmt.Println(Config.Port, Config.DbName, Config.SQLDriver)
}