package postcode

import (
    "database/sql"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    _ "github.com/go-sql-driver/mysql"
    "github.com/LindsayBradford/go-dbf/godbf"
    "github.com/julienschmidt/httprouter"
    "syscall"
    "time"
    ylog	"github.com/postcode/lib/ylog"
    yconfig	"github.com/postcode/lib/yconfig"
)

const level = 5
var table map[int]string
// Version - версия
var Version string

// Init - инициализация системы по входным параметрам
func Init(Ver string) string {
	var err error
	var fileName string
	var Nodes [level]Node
	Version				= Ver

	table = map[int]string{
	    0: "regions",
	    1: "autonoms",
	    2: "areas",
	    3: "citys",
	    4: "indexes",
	}

	for i := 0; i<level; i++ {
	    Nodes[i].ID			= make(map[string]int)
	    Nodes[i].Name		= make(map[string]string)
	    Nodes[i].Description	= make(map[string]string)
	}

	initdbPtr		:= flag.String("init", "", " - Path to file data, create database and init params")
	ConfigPtr		= flag.String("config", "", " - Path to config, if not found, then file is taken in /etc/postcode/"+os.Args[0])
	flag.Parse()
	if len(os.Args) == 1 {
		fmt.Printf("\n use: %s -h for help \n", os.Args[0])
		os.Exit(0)
	}
	yconfig.Conf(config, ConfigPtr)
	ylog.YLogInit(config.YLog, config.Interface)

	Db, err = sql.Open(config.Core.DatabaseType, Params(config.Core, config.Core.DatabaseType))
	if err != nil {
		ylog.YLog(3, Ident, err.Error())
		os.Exit(3)
	}
	defer Db.Close()
	Db.SetMaxIdleConns(0)
	if len(*initdbPtr) > 2 {
	    t0 := time.Now()
	    fmt.Printf("\nInit database\n")
	    if config.Core.DatabaseType == "mysql" {
		err = InitMysql()
		if err != nil {
		    fmt.Printf("%v\n", err)
		}
	    }
	    fmt.Printf("\nUnzip file %s\n", *initdbPtr)
	    files, err		:= Unzip(*initdbPtr, "./")
	    if err != nil {
		ylog.YLog(3, Ident, err.Error())
		os.Exit(3)
	    }
	    fileName		= files[0]
	    fmt.Printf("\nFile %v  read\n", fileName)
	
	    dbfTable, err := godbf.NewFromFile(fileName, "cp866")
	    fmt.Printf("Records found - %d\n", dbfTable.NumberOfRecords())
	    for i := 0; i < dbfTable.NumberOfRecords(); i++ {
		sl							:= dbfTable.GetRowAsSlice(i) 
		if len(sl[5])<1{
		    sl[5]						= sl[4]
		}
		if len(sl[6])<1{
		    sl[6]						= sl[4]
		}
		if len(sl[7])<1{
		    sl[7]						= sl[4]
		}
		Nodes[0].ID[sl[4]]					= i
		Nodes[1].Name[sl[5]]					= sl[4]
		Nodes[1].ID[sl[5]]					= i
		Nodes[2].Name[sl[6]]					= sl[5]
		Nodes[2].ID[sl[6]]					= i
		Nodes[3].Name[sl[7]]					= sl[6]
		Nodes[3].ID[sl[7]]					= i
		Nodes[3].Description[sl[7]]				= sl[8]
		Nodes[4].Name[sl[0]]					= sl[7]
		Nodes[4].ID[sl[0]]					= i
	    }
	    count, err							:= AddRegions(Nodes);
	    fmt.Printf("Records added regions  - %d\n", count-1)
	    count, err							= AddAll(Nodes);
	    if config.Core.DatabaseType == "mysql" {
		err = InitMysqlIndex()
		if err != nil {
		    fmt.Printf("%v\n", err)
		}
	    }
	    fmt.Printf("Time has passed: %.2f min\n", time.Since(t0).Minutes())
	    os.Exit(0)
	}
	InitAPI();
	ioutil.WriteFile(config.Pidfile, []byte(fmt.Sprintf("%d\n", os.Getpid())), 0644)
	defer os.Remove(config.Pidfile)
	return config.Pidfile
}

func InitAPI(){
    fmt.Print("Run\n")
    strInterface := fmt.Sprintf("%s:%d", config.Interface, config.Port)
    ylog.YLog(1, Ident, "Run as demon")
    router := httprouter.New()
	router.GET("/api/v1/stat", BasicAuth(HStatApi, config.UserName, config.Password))
	router.GET("/api/v1/getAll/:name", BasicAuth(HGetAll, config.UserName, config.Password))
	router.GET("/api/v1/getAllIndexesID/:name/:id", BasicAuth(HGetAllIndexes, config.UserName, config.Password))
	router.GET("/api/v1/getAddressByPostIndex/:id", BasicAuth(HGetAddress, config.UserName, config.Password))
	router.GET("/api/v1/getNamesByTopId/:name/:id", BasicAuth(HGetNamesByTopID, config.UserName, config.Password))
	router.POST("/api/v1/getAddressByPostCodes", BasicAuth(HGetAddresses, config.UserName, config.Password))

    syscall.Dup2(int(ylog.F[3].Fd()), 1)
    syscall.Dup2(int(ylog.F[3].Fd()), 2)

    if config.Https == "no" {
	ylog.YLog(1, Ident, "Listening HTTP on _"+strInterface+"_")
	http.ListenAndServe(strInterface, router)
    }else{
	ylog.YLog(1, Ident, "Listening HTTPS on "+strInterface)
	http.ListenAndServeTLS(strInterface, "cert.pem", "key.pem", router)
    }
}
