# postcode
Пример валидатора для bitrix. Все параметры лучше разместить в административной части в настройках <module> модуля.
## theo.postcode
theo.postcode.tar.gz - готовый модуль. Внимательнее на init.php , тут не проверяется пустой ли и после установки модуля /bitrix/php_interface/init.php не будет работать. Приоритет папки local, берется только 1 файл из одноименных.