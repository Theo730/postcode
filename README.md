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
Для использования программы необходим установленный Mysql сервер. Перед запуском сервера необходимо инициализировать базу данных и данные.
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

## Компиляция

Необходим Go 1.8+ и все зависимые библиотеки, сервер проверялся на CentOS 7.1 и Debian 8.8. Устновка из исходников:
```
 go get https://github.com/Theo730/postcode.git
или
 cd ./go/src/github.com/
 wget https://github.com/Theo730/postcode.git
cd ./postcode/
make
```
## Установка компилированной (Linux CentOS 7.1)

## Директории
* test - директория с тестами
* lib - директория с дополнительными библиотеками 
* postcode - директория программы, выполненной в виде библиотеки
* programme - дирестория с архивом скомпелированной программы под Linux CentOS 7.0 c 