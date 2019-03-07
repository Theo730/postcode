package postcode

import (
    "fmt"
    "time"
    ylog	"github.com/postcode/lib/ylog"
    yconfig	"github.com/postcode/lib/yconfig"
)

// Params - в зависимости от входных параметров выбирает параметры для соединения с БД
func Params(cfg yconfig.DBParams, str string) string {
    var info string

    if str == "mysql" {
	if cfg.Host == "localhost" {
	    info		= fmt.Sprintf("%s:%s@/%s", cfg.UserName, cfg.Password, cfg.DbName)
	}else{
	    info		= fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	}
    }
// для oracle 12, возможно будет использоваться
    if str == "oci8" {
	info			= fmt.Sprintf("%s:%s@%s/%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Sid)
    }
    return info
}

// Добавление всех регионов в базу данных
func AddRegions(Nodes [level]Node)(int, error){
    var  i int
    stmt, err 			:= Db.Prepare("INSERT INTO regions SET NAME = ?")
    if err != nil {
	panic(err.Error())
    }
    defer stmt.Close()
    i 				= 0
    for k := range Nodes[0].ID {
	if len(k)>0{
    	    result, _ 		:= stmt.Exec(k)
	    j, _ 		:= result.LastInsertId()
	    Nodes[0].ID[k]	= int(j)
	} 
	i++
	if err != nil {
	    panic(err.Error())
	}
    }
    return i, nil
}

// Добавление всех сущностей, кроме регионов, в базу данных
func AddAll(Nodes [level]Node)(int, error){
    var j int64 
    var str string
    
    for i:=1; i<level ;i++ {
	j 				= 0
	if table[i] == "city"{
	    str				= fmt.Sprintf("INSERT INTO %s SET NAME = ?, TOP_ID =?, DESCRIPTION = ?", table[i])
	}else{
	    str				= fmt.Sprintf("INSERT INTO %s SET NAME = ?, TOP_ID =?", table[i])
	}
	stmt, err 			:= Db.Prepare(str)
	if err != nil {
	    panic(err.Error())
	}
	defer stmt.Close()
	for k := range Nodes[i].ID {
	    if len(k)>0{
		if table[i] == "city" {
		    result, _ 		:= stmt.Exec(k, Nodes[i-1].ID[Nodes[i].Name[k]], Nodes[i].Description[k])
		    j, _ 		= result.LastInsertId()
		}else{
		    result, _ 		:= stmt.Exec(k, Nodes[i-1].ID[Nodes[i].Name[k]])
		    j, _ 		= result.LastInsertId()
		}
		Nodes[i].ID[k]		= int(j)
	    } 
	    j++
	    if err != nil {
		panic(err.Error())
	    }
	}
	fmt.Printf("Records added %s - %d\n", table[i], j-1)
    }
    return int(j), nil
}

// GetStat - статистика по системе.
func GetStat() (rec Stat, err error){

    t0 			:= time.Now()

    err			= Db.QueryRow("SELECT count(*) FROM regions").Scan(&rec.Regions)
    if err != nil {
	ylog.YLog(3, Ident, "Stat select param error, Error: "+err.Error())
	rec.Regions = 0
	return rec, err
    }
    err			= Db.QueryRow("SELECT count(*) FROM autonoms").Scan(&rec.Autonoms)
    if err != nil {
	ylog.YLog(3, Ident, "Stat select param error, Error: "+err.Error())
	rec.Autonoms = 0
	return rec, err
    }
    err 		= Db.QueryRow("SELECT count(*) FROM areas").Scan(&rec.Areas)
    if err != nil {
	ylog.YLog(3, Ident, "Stat select param error, Error: "+err.Error())
	rec.Areas = 0
	return rec, err
    }
    err			= Db.QueryRow("SELECT count(*) FROM citys").Scan(&rec.Citys)
    if err != nil {
	ylog.YLog(3, Ident, "Stat select param error, Error: "+err.Error())
	rec.Citys = 0
	return rec, err
    }
    err			= Db.QueryRow("SELECT count(*) FROM indexes ").Scan(&rec.Indexes)
    if err != nil {
	ylog.YLog(3, Ident, "Stat select param error, Error: "+err.Error())
	rec.Indexes = 0
	return rec, err
    }
    rec.WorkTime = int64(time.Since(t0))
    return rec, nil
}

// GetGist - запрос всех сущностей в базе данных, входящие параметры tableName - имя сущности, id - если надо, id родителя
func GetGist(tableName string, id int = 0) (out OutGist, err error){
    var str string
    if id == 0{
	str			= fmt.Sprintf("SELECT ID, NAME, DESCRIPTION FROM %s", tableName)
    }else{
	str			= fmt.Sprintf("SELECT ID, NAME, DESCRIPTION FROM %s WHERE TOP_ID = %d", tableName, id)
    }
    results, err 	:= Db.Query(str)
    if err != nil {
	ylog.YLog(3, Ident, "DataBase Error "+err.Error())
	return out, err
    }
    j			:= 0
    for results.Next() {
	var gist Gist

	err = results.Scan(&gist.ID, &gist.Name, &gist.Description)
	if err != nil {
	    panic(err.Error()) 
	}
	out.Gist		= append(out.Gist, gist)
	j++
    }
    out.Count			= j
    return out, nil
}

func GetAllIndexes(tableName string, tLevel int, id string ) (out Indexes, err error){
    str 		:= "(SELECT NAME FROM indexes WHERE TOP_ID in" 
    for i:= level-2; i>tLevel; i-- {
	str		= fmt.Sprintf("%s (SELECT ID FROM %s WHERE TOP_ID in ", str, table[i])
    
    } 
    str			= fmt.Sprintf("%s (%s)", str, id)
    for i:= level-1; i>tLevel; i-- {
	str		= fmt.Sprintf("%s )", str)
    } 
    fmt.Printf("%s\n",str)
    results, err 	:= Db.Query(str)
    if err != nil {
	ylog.YLog(3, Ident, "DataBase Error "+err.Error())
	return out, err
    }
    j			:= 0
    for results.Next() {
	var index string

	err = results.Scan(&index)
	if err != nil {
	    panic(err.Error()) 
	}
	out.Indexes	= append(out.Indexes, index)
	j++
    }
    out.Count		= j
    return out, nil
}

// GetAddress - получание адреса по почтовому индексу
func GetAddress(id string) (out Address, err error){
    err		= Db.QueryRow("SELECT indexes.name, citys.name, citys.description, areas.name, autonoms.name, regions.name "+ 
	"FROM indexes, citys, areas, autonoms, regions WHERE indexes.name = ? and indexes.top_id = citys.id and citys.top_id = areas.id "+
	"and areas.top_id = autonoms.id and autonoms.top_id = regions.id", id).Scan(&out.Index, &out.City, &out.City1, &out.Area, &out.Autonom, &out.Region)
    if err != nil {
	ylog.YLog(3, Ident, "Stat select param error, Error: "+err.Error())
	return out, err
    }
    return out, nil
}

// GetAddresses - получение всех адресов по массиву почтовых индексов
func GetAddresses(indexes Indexes) (out Addresses, err error){
    str 		:=  fmt.Sprintf("'%s'", indexes.Indexes[0])
    for i:=1; i<len(indexes.Indexes);i++ {
	str		= fmt.Sprintf("%s, '%s'", str, indexes.Indexes[i])
    }
    str			= fmt.Sprintf("SELECT indexes.name, citys.name, citys.description, areas.name, autonoms.name, regions.name "+ 
	"FROM indexes, citys, areas, autonoms, regions WHERE indexes.name in ( %s ) and indexes.top_id = citys.id and citys.top_id = areas.id "+
	"and areas.top_id = autonoms.id and autonoms.top_id = regions.id", str)
    results, err 	:= Db.Query(str)
    j			:=0
    for results.Next() {
	var address Address
	err 		= results.Scan(&address.Index, &address.City, &address.City1, &address.Area, &address.Autonom, &address.Region)
	if err != nil {
	    ylog.YLog(3, Ident, "Stat select param error, Error: "+err.Error())
	    return out, err
	}
	out.Address	= append(out.Address, address)
	j++
    }
    out.Count		= j
    return out, nil
}
