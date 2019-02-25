package ylog
// структуры для конфигурационного файла при использовании пакета
type YLogType struct{
    Debug		string		`json:"debug"`
    YLogOut		[]YLogOutType	`json:"ylogout"`

}
type YLogOutType struct{
    LogType		string		`json:"logtype"`  // возможные значения - file, syslog, fluent
    LogFile		string		`json:"logfile"`
    DebugFile		string		`json:"debugfile"`
    InfoFile		string		`json:"infofile"`
    WarningFile		string		`json:"warningfile"`
    ErrorFile		string		`json:"errorfile"`
    NoticeFile		string		`json:"noticefile"`
    ServerIP		string		`json:"server_ip"`
    ServerPort		int		`json:"server_port"`
}
