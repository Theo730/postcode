# postcode
[eng]
Small programme for validating postal codes (data from the Russian post). Additionally, the program serves as a reference for:
* Regions
* Autonomous regions
* Areas
* City
* Zip codes in the above

[rus]
Проект работы с "ЭТАЛОННЫЙ СПРАВОЧНИК ПОЧТОВЫХ ИНДЕКСОВ ОБЪЕКТОВ ПОЧТОВОЙ СВЯЗИ", файл справочника расположен по адресу https://vinfo.russianpost.ru/database/ops.html. На момент написания был взят файл PIndx03.dbf, который содержит:
```
Unzip file PIndx.zip
File PIndx03.dbf  read
Records found - 48791
Records added regions  - 84
Records added autonoms - 90
Records added areas - 1817
Records added citys - 23302
Records added indexes - 48791
Time has passed: 2.57 min
```
Для использования программы необходим установленный Mysql сервер.
Перед запуском сервера необходимо инициализировать базу данных и данные.
```
./postCode -config ./config.json -init PIndx.zip
```
 После инициализации и запуска будут доступны все данные. Программа не совсем базаданных КЛАДР, точность данных до города.
 Выполняются запросы:
 * регионы
 * области
 * города
 * Почтовые индексы
 
 по спецификации элемента:
 * все почтовые индексы перечисленых сущностей
 
 по почтовому индексу:
 * адрес до города
 * запрос массива адресов
 
 Ключи запуска
    `./postcode -h`
    вызов краткой справки по аргументам коммандной строки

    `./postcode -init=<filename>`
    инициализация(создание) структуры таблиц в базе данных и записываются сами данные

    `./postcode -config=./config.json &`
    запуск программы как демона, штатный запуск
