package ylog

import (
    "fmt"
    "errors"
    "log"
    "time"
    "reflect"
    "log/syslog"
    "os"
    "github.com/fluent/fluent-logger-golang/fluent"
)

var F [6] *os.File
var prgName string = os.Args[0]
var logItem  = make([]string, 0)
var sysYLog   [5]*syslog.Writer
var debugLog int
var logger *fluent.Fluent
var host string

func initSysyLog() {

    sysYLog[0],_ 		= syslog.New(syslog.LOG_DEBUG, prgName)
    defer sysYLog[0].Close()
    sysYLog[1], _		= syslog.New(syslog.LOG_INFO, prgName)
    defer sysYLog[1].Close()
    sysYLog[2], _		= syslog.New(syslog.LOG_WARNING, prgName)
    defer sysYLog[2].Close()
    sysYLog[3], _		= syslog.New(syslog.LOG_ERR, prgName)
    defer sysYLog[3].Close()
    sysYLog[4],_		= syslog.New(syslog.LOG_NOTICE, prgName)
    defer sysYLog[4].Close()
    log.SetFlags(0)
}

func initFluent(LogConfig YLogOutType) error{
    var err error
    
    if len(LogConfig.ServerIP)<3{
	return errors.New("IP server Fluend is empty")
    }
    if LogConfig.ServerPort<1{
	return errors.New("Port server Fluend is empty")
    }
    logger, err = fluent.New(fluent.Config{FluentHost:LogConfig.ServerIP, FluentPort:LogConfig.ServerPort})
    if err != nil{
	return err
    }
    defer logger.Close()
    return nil
}

func YLogInit(YLogConfig YLogType, interfaceName string) error {
    var err error
    CFile := map[string]int{"DebugFile":0, "InfoFile":1, "WarningFile":2, "ErrorFile":3, "NoticeFile":4, "LogFile":5}

    if YLogConfig.Debug == "yes"{
	debugLog = 1
    }else{
	debugLog = 0
    }
    bool := false
    host  = interfaceName
    for val := range YLogConfig.YLogOut{
	if YLogConfig.YLogOut[val].LogType == "file"{
	    for i :=1 ; i < reflect.ValueOf(YLogConfig.YLogOut[val]).NumField(); i++ {
		ifv := reflect.ValueOf(YLogConfig.YLogOut[val]).Field(i)
		if ((len(ifv.String())>5)&&(reflect.TypeOf(YLogConfig.YLogOut[val]).Field(i).Type.String() == "string")) {
		    F[CFile[reflect.TypeOf(YLogConfig.YLogOut[val]).Field(i).Name]], err = os.OpenFile(ifv.String(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		    if err != nil {
			fmt.Printf("INIT Error: %s file %s\n", err, ifv.String())
			os.Exit(-1)
			return err
		    }
		}
	    }
	    bool = true
	}
	if YLogConfig.YLogOut[val].LogType == "syslog"{
	    initSysyLog()
	    bool = true
	}
	if YLogConfig.YLogOut[val].LogType == "fluent"{
	    err  = initFluent(YLogConfig.YLogOut[val])
	    if err != nil {
		bool		= false
	    }else{
		bool		= true
	    }
	}
	if bool {
	    logItem = append(logItem, YLogConfig.YLogOut[val].LogType)
	    bool		= false
	}
    }
    return nil
}

func YLog(level int, ident string, msg string){
    var err error
    levels := []string{"Debug", "Info", "Warning", "Error", "Notice", "Log"}
    if len(msg) > 0 {
	if (level < 0 ) && ( level > 4){
	    level = 5
	}
	if (level == 0 )&&(debugLog == 0){ // не записывать логи , если в конфирурации нет переменной DEBUG
	    return
	}
	msg = host + " " + ident + " "+levels[level] + " " + msg
	    for val := range logItem {
		if logItem[val] == "file"{
		    if (level < 6) && (F[level] != nil){
			if _, err = F[level].WriteString(time.Now().Format("02-01-2006 15:04:05") + " " + prgName + " " + msg + "\n"); err != nil {
			    fmt.Printf("Error write log %s \n", err)
			}
		    }else{
			if _, err = F[5].WriteString(time.Now().Format("02-01-2006 15:04:05") + " " + prgName + " " + msg + "\n"); err != nil {
			    fmt.Printf("Error write log %s\n", err)
			}
		    }
		}
		if logItem[val] == "syslog"{
		    log.SetOutput(sysYLog[level])
		    log.Print(msg)
		}
		if logItem[val] == "fluent"{
		    strMap 		:= make(map[string]string)
		    strMap["level"]	= levels[level]
		    strMap["host"]	= host
		    strMap["ident"]	= ident
		    strMap["message"]	= msg
		    logger.Post("log.es", strMap)
		}
	    }
	}
}

func Destr(){
    for i:=0; i<6; i++{
	defer F[i].Close()
    }
}

func Debug(ident string, mesg string) {
    YLog(0,ident,mesg)
}

func Info(ident string, mesg string) {
    YLog(1,ident,mesg)
}

func Warning(ident string, mesg string) {
    YLog(2,ident,mesg)
}

func Error(ident string, mesg string) {
    YLog(3,ident,mesg)
}
