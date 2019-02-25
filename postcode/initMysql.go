package postcode

import (
    "fmt"
    ylog	"github.com/postcode/lib/ylog"
)

// InitMysql - создание структуры таблиц Mysql
func InitMysql() error {
    var err error

    _, err = Db.Exec("DROP TABLE IF EXISTS citys ")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }
    _, err = Db.Exec("CREATE TABLE citys ( "+
	"id int(10) unsigned NOT NULL AUTO_INCREMENT, "+
	"name varchar(255) DEFAULT NULL, "+
	"description varchar(255) DEFAULT NULL, "+
	"top_id int(10) UNSIGNED DEFAULT 0, "+
	"PRIMARY KEY (id)) ENGINE=InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }

    _, err = Db.Exec("DROP TABLE IF EXISTS areas ")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }
    _, err = Db.Exec("CREATE TABLE areas ( "+
	"id int(10) unsigned NOT NULL AUTO_INCREMENT, "+
	"name varchar(255) DEFAULT NULL, "+
	"description varchar(255) DEFAULT NULL, "+
	"top_id int(10) UNSIGNED DEFAULT 0, "+
	"PRIMARY KEY (id)) ENGINE=InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }

    _, err = Db.Exec("DROP TABLE IF EXISTS autonoms ")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }
    _, err = Db.Exec("CREATE TABLE autonoms ( "+
	"id int(10) unsigned NOT NULL AUTO_INCREMENT, "+
	"name varchar(255) DEFAULT NULL, "+
	"description varchar(255) DEFAULT NULL, "+
	"top_id int(10) UNSIGNED DEFAULT 0, "+
	"PRIMARY KEY (id)) ENGINE=InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }

    _, err = Db.Exec("DROP TABLE IF EXISTS indexes ")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }
    _, err = Db.Exec("CREATE TABLE indexes ( "+
	"id int(10) unsigned NOT NULL AUTO_INCREMENT, "+
	"name varchar(255) DEFAULT NULL, "+
	"description varchar(255) DEFAULT NULL, "+
	"top_id int(10) UNSIGNED DEFAULT 0, "+
	"PRIMARY KEY (id)) ENGINE=InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }
    _, err = Db.Exec("DROP TABLE IF EXISTS regions ")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }
    _, err = Db.Exec("CREATE TABLE regions ( "+
	"id int(10) unsigned NOT NULL AUTO_INCREMENT, "+
	"name varchar(255) DEFAULT NULL, "+
	"description varchar(255) DEFAULT NULL, "+
	"PRIMARY KEY (id)) ENGINE=InnoDB CHARACTER SET utf8 COLLATE utf8_general_ci")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
    }
    return err
}

func InitMysqlIndex() error {
	var err error

    _, err = Db.Exec("CREATE INDEX index_city_top ON citys (top_id);")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }
    _, err = Db.Exec("CREATE INDEX index_city_top ON areas (top_id);")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }
    _, err = Db.Exec("CREATE INDEX index_city_top ON autonoms (top_id);")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	return err
    }
    _, err = Db.Exec("CREATE INDEX index_top ON indexes (top_id);")
    if err != nil {
	ylog.YLog(3, fmt.Sprintf("%s",err), err.Error())
	}
    return err
}