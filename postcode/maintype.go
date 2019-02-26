package postcode

import (
    "database/sql"
    yconfig	"github.com/postcode/lib/yconfig"
)

// Ident - строка идентификатор внешний системы (ключ сессии), название программы
var Ident string

// Db - переменная работы с БД.
var Db *sql.DB

var config		= new(yconfig.Config)

// ConfigPtr - указатель на строку параметров запуска платежной системы.
var ConfigPtr *string

// Rec - структура записи старистики системы.
var Rec Stat

// Stat - структура для статистики статистика
type Stat struct{
    Version		string		`json:"version_k"`		// версия ПО формат 1.0.1
    Regions		int		`json:"regions"`		// кол-во записей в таблице регионы
    Autonoms		int		`json:"autonoms"`		// кол-во записей в таблице автономии
    Areas		int		`json:"areas"`			// кол-во записей в таблице области
    Citys		int		`json:"city"`			// кол-во записей в таблице города
    Indexes		int		`json:"indexes"`		// кол-во записей в таблице индексы
    WorkTime		int64		`json:"work_time_k"`		// время обработки запроса статистики в наносендах
}

// общая структура для инициализации базы
type Node struct{
    ID			map[string]int					// id в bd
    Name		map[string]string				// имя сущности (регион, город, область...)
    Description		map[string]string				// описание, дополнительное поле sity1 в sity
}

// OutGist - структура выдачи обектов
type OutGist struct{
    Count		int		`json:"count"`			// количество
    Gist		[]Gist		`json:"gist"`			// масив объектов выдачи
}

// Gist - структура сущности название региона + id, города или области
type Gist struct{
    ID			int		`json:"id"`			// id в базе
    Name		string		`json:"name"`			// наименование сущности
    Description		sql.NullString	`json:"description"`		//
}

type Indexes struct{
    Count		int		`json:"count"`			// количество
    Indexes		[]string	`json:"indexes"`		// почтовый индекс
}

type Address struct{
    Region		string		`json:"region"`			// регион
    Autonom		string		`json:"autonom"`		// автономя
    Area		string		`json:"area"`			// область
    City		string		`json:"city"`			// город
    City1		sql.NullString	`json:"city1"`			// город1
    Index		int		`json:"indexes"`		// индекс
}

type Addresses struct{
    Count		int		`json:"count"`			// количество
    Address		[]Address	`json:"adress"`		// адреса
}