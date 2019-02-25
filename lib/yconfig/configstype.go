package yconfig

import (
    ylog	"github.com/postcode/lib/ylog"
)
// структуры для конфиг файла
type Config struct {
    YLog		ylog.YLogType	`json:"YLog"`
    Pidfile		string		`json:"pidfile"`
    WorkDir		string		`json:"workdir"`
    Interface		string		`json:"interface"`
    Port		int		`json:"port"`
    Service		string		`json:"service"`
    Https		string		`json:"https"`
    UserName		string		`json:"username"`
    Password		string		`json:"password"`
    Core		DBParams	`json:"core"`
}

type DBParams struct{
    DatabaseType	string		`json:"database_type"`
    Host		string		`json:"host"`
    Port		int		`json:"post"`
    Sid			string		`json:"sid"`
    DbName		string		`json:"dbname"`
    UserName		string		`json:"username"`
    Password		string		`json:"password"`
}

