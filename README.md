# rzreversescheme

Из вашего трафика генерит swagger-схему

(тест) -> rzreversescheme(выступает как прокси) ->(приложение)

# Запуск

go run main.go 
Параметры:
 * -cmd-port - порт для api (default 9500)
 * -port - порт прокси сервера (default 9501)


# API
## Host matching

`
curl -i -X  POST http://localhost:9500/configure/host -H "Accept: application/json" -H "Content-Type: application/json"  -d '{ "host": "myhost", "hostfilter": "localhost", "postfilter": "*"}'
`

 * host - host name
 * hostfilter - фильтр по вхождению строки
 * portfilter - фильр по порту, число или "*"
 * pathfilter - фильтр пути запроса, вхождение строки
 
 ## Get scheme
 
 `http://localhost:9500/schema/<host>`

 * host - имя хоста, для которого нужна схема, если было обращение с портом (без матчинга) - указывать с портом, например: myhost:8090

# ENG
Reverse swagger scheme from http-trafic


