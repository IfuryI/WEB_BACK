# КиноПоиск
### Команда
* [Теняев Олег](https://github.com/grillow)
* [Афимин Илья](https://github.com/IfuryI)
* [Бахметьев Глеб](https://github.com/polyanimal)
* [Зотов Алексей](https://github.com/let-robots-reign)

### Менторы
* [Климова Наталия](https://github.com/Tataklim) - ментор по фронтенду
* [Рыбаков Дмитрий](https://github.com/bulletmys) - ментор по бекенду

### Swagger
https://app.swaggerhub.com/apis/let-robots-reign/Kinopoisk/1.0.0

### Деплой
http://89.208.198.186:4000/

### Ссылка на репозиторий фронтенда
https://github.com/frontend-park-mail-ru/2021_1_kekEnd

ieafimin@MSK-C02D5445MD6R WEB_BACK % ab -c 10 -n 6000 -r https://cinemedia.ru/api/v1/movies/26
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking cinemedia.ru (be patient)
Completed 600 requests
Completed 1200 requests
Completed 1800 requests
Completed 2400 requests
Completed 3000 requests
Completed 3600 requests
Completed 4200 requests
Completed 4800 requests
Completed 5400 requests
Completed 6000 requests
Finished 6000 requests


Server Software:        nginx/1.18.0
Server Hostname:        cinemedia.ru
Server Port:            443
SSL/TLS Protocol:       TLSv1.2,ECDHE-RSA-AES256-GCM-SHA384,2048,256
Server Temp Key:        ECDH X25519 253 bits
TLS Server Name:        cinemedia.ru

Document Path:          /api/v1/movies/26
Document Length:        1787 bytes

Concurrency Level:      10
Time taken for tests:   47.947 seconds
Complete requests:      6000
Failed requests:        0
Total transferred:      11772000 bytes
HTML transferred:       10722000 bytes
Requests per second:    125.14 [#/sec] (mean)
Time per request:       79.912 [ms] (mean)
Time per request:       7.991 [ms] (mean, across all concurrent requests)
Transfer rate:          239.77 [Kbytes/sec] received

Connection Times (ms)
min  mean[+/-sd] median   max
Connect:       22   47  14.1     44     291
Processing:     9   33  14.5     32     198
Waiting:        9   32  14.1     30     197
Total:         41   80  21.7     76     354

Percentage of the requests served within a certain time (ms)
50%     76
66%     82
75%     87
80%     91
90%    105
95%    117
98%    133
99%    145
100%    354 (longest request)



asdad


ieafimin@MSK-C02D5445MD6R WEB_BACK % ab -c 10 -n 6000 -r https://cinemedia.ru/api/v1/movies/26     
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking cinemedia.ru (be patient)
Completed 600 requests
Completed 1200 requests
Completed 1800 requests
Completed 2400 requests
Completed 3000 requests
Completed 3600 requests
Completed 4200 requests
Completed 4800 requests
Completed 5400 requests
Completed 6000 requests
Finished 6000 requests


Server Software:        nginx/1.18.0
Server Hostname:        cinemedia.ru
Server Port:            443
SSL/TLS Protocol:       TLSv1.2,ECDHE-RSA-AES256-GCM-SHA384,2048,256
Server Temp Key:        ECDH X25519 253 bits
TLS Server Name:        cinemedia.ru

Document Path:          /api/v1/movies/26
Document Length:        1787 bytes

Concurrency Level:      10
Time taken for tests:   48.183 seconds
Complete requests:      6000
Failed requests:        0
Total transferred:      11772000 bytes
HTML transferred:       10722000 bytes
Requests per second:    124.53 [#/sec] (mean)
Time per request:       80.305 [ms] (mean)
Time per request:       8.031 [ms] (mean, across all concurrent requests)
Transfer rate:          238.59 [Kbytes/sec] received

Connection Times (ms)
min  mean[+/-sd] median   max
Connect:       19   47  15.4     45     267
Processing:     8   33  17.7     31     386
Waiting:        8   31  16.6     29     386
Total:         37   80  24.5     76     439

Percentage of the requests served within a certain time (ms)
50%     76
66%     83
75%     88
80%     92
90%    105
95%    118
98%    134
99%    153
100%    439 (longest request)
