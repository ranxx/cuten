curl "http://localhost:9999/login/zhangsan" -X POST -d '{"user":"张三","passwd":"123456"}' --header "Content-Type: application/json"
curl "http://localhost:9999/login/lisi" -X POST -d '{"user":"张三","passwd":"123456"}' --header "Content-Type: application/json"
curl "http://localhost:9999/login" -X POST -d '{"user":"李四","passwd":"qwert"}' --header "Content-Type: application/json"
curl "http://localhost:9999/login?user=axing&passwd=123456"
curl "http://localhost:9999/v1?user=axing&passwd=123456"
curl "http://localhost:9999/v1/user?user=axing&passwd=123456"

