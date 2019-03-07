# postcode
Распаковать и установить(mysql локально без пароля)
```
tar -xvf ./postcode02.tar.gz
./install.sh
```
Установка сервера происходит в /opt/postcode. В эту директорию необходимо загрузить PIndx.zip с почты россии. И далее:
```
./installdb.sh
```
После для установки, как сервиса:
```
./installservice.sh
```

