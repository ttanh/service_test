# service_test
The test for interview

## Installation
```
    $ git clone https://github.com/ttanh/service_test.git
```

#### Setup for local environment with Docker
```
$ docker-compose up
```

## Configration 
### Edit docker-compose.yml file
```
MYSQL_ROOT_PASSWORD: secret pass
MYSQL_USER: root                  # user
MYSQL_PASSWORD: secret pass
MYSQL_DATABASE: service_db        # database name

DB_USER: root           #user
DB_PASS: secret pass    #pass
DB_HOST: mysql          #mysql service on docker
DB_PORT: 3306           #port
DB_NAME: service_db     # database name
DB_LOG_MODE: 0          #config for mysql log mode, 0: false, 1: true
```
## APIS
Url base: http://localhost:9000

Error code: 
```
SUCCESS         = 0
ErrorInvalid    = -1
ErrorToken      = -2
ErrorFailed     = -3
ErrorStatusDone = -4
```

### 1. POST /token : Create token
#### Response
```
SPi7s4I4Rq79zwxpy8hv
```

### 2. POST /service1
#### Request
```
{
	"customer_id": "123456789",
	"customer_name": "Tuan Anh",
	"transaction_id": "pOA9e3Pgg64mbLgx4Le4b9qjyW5V3dH7CXmlC9K4CsZlFJbaB5v1TPXb8xRd3Hi3nBbFFOPyJTV4JwJR1fpfS6XE5dnurkdOQ8xLEPe58j2IDQgb1zAWd2gjjExbEtOO",
	"token": "SPi7s4I4Rq79zwxpy8hv",
	"date_time": "2017-12-12T15:34:05Z"
}
```
#### Response 
```
{
    "date_time": "2017-12-12T15:34:05Z",
    "transaction_id": "pOA9e3Pgg64mbLgx4Le4b9qjyW5V3dH7CXmlC9K4CsZlFJbaB5v1TPXb8xRd3Hi3nBbFFOPyJTV4JwJR1fpfS6XE5dnurkdOQ8xLEPe58j2IDQgb1zAWd2gjjExbEtOO",
    "code": 0,
    "description": "success"
}
```

### 3. POST /service2
#### Request
```
{
	"transaction_id": "pOA9e3Pgg64mbLgx4Le4b9qjyW5V3dH7CXmlC9K4CsZlFJbaB5v1TPXb8xRd3Hi3nBbFFOPyJTV4JwJR1fpfS6XE5dnurkdOQ8xLEPe58j2IDQgb1zAWd2gjjExbEtOO",
	"token": "Bw2YU4BJqV6pdVSSoEZx"
}
```
#### Response
```
{
    "date_time": "0001-01-01T00:00:00Z",
    "transaction_id": "pOA9e3Pgg64mbLgx4Le4b9qjyW5V3dH7CXmlC9K4CsZlFJbaB5v1TPXb8xRd3Hi3nBbFFOPyJTV4JwJR1fpfS6XE5dnurkdOQ8xLEPe58j2IDQgb1zAWd2gjjExbEtOO",
    "code": -4,
    "description": "record not found"
}
```

### 4. POST /process_done: Set status to done (status = 2) 
#### Request
```
{
	"transaction_id": "pOA9e3Pgg64mbLgx4Le4b9qjyW5V3dH7CXmlC9K4CsZlFJbaB5v1TPXb8xRd3Hi3nBbFFOPyJTV4JwJR1fpfS6XE5dnurkdOQ8xLEPe58j2IDQgb1zAWd2gjjExbEtOO",
	"token": "Bw2YU4BJqV6pdVSSoEZx"
}
```
#### Response
```
{
    "date_time": "2017-12-13T11:42:13.819808671Z",
    "transaction_id": "pOA9e3Pgg64mbLgx4Le4b9qjyW5V3dH7CXmlC9K4CsZlFJbaB5v1TPXb8xRd3Hi3nBbFFOPyJTV4JwJR1fpfS6XE5dnurkdOQ8xLEPe58j2IDQgb1zAWd2gjjExbEtOO",
    "code": 0,
    "description": "success"
}
```
