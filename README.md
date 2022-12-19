# Exchange-REST-tt

#### Rest API поддерживает 2-а Get запроса:
    
        http://localhost:8080/v1/exchange/USD
    
        http://localhost:8080/v1/exchange
    

### Развертывание 
  1. Скачать публичный репозиторий.
  2. Запустить Docker desktop
  3. Выполнить в терминале комманду 
  ````
    docker-compose up
  ````
  В процессе работы контейнера используются следующийе порты: 8080. 

### Тестирование 
    В заголовке запроса необходимо указать ключ авторизации.
        X-Api-Key: 123321

### TODO 
  1. Дописать интеграционные тесты
  2. Добавить склонения Рубля
  3. Имплементировать единый sync.Map, хронящий данные от https://www.cbr-xml-daily.ru/daily_json.js
  4. Доработать логирование и ошибки  
  5. Пофиксить swagger