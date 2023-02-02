[![JGo Web](https://jgo.dev/assets/images/logo_300.svg)](https://jgo.dev/)

[![JGo release](https://img.shields.io/github/v/release/xjustloveux/jgo)](https://github.com/xjustloveux/jgo/releases)
[![codecov](https://codecov.io/gh/xjustloveux/jgo/branch/master/graph/badge.svg?token=RCO5VO2YU6)](https://codecov.io/gh/xjustloveux/jgo)
[![Build Status](https://github.com/xjustloveux/jgo/actions/workflows/go.yml/badge.svg)](https://github.com/xjustloveux/jgo/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/xjustloveux/jgo)](https://goreportcard.com/report/github.com/xjustloveux/jgo)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/xjustloveux/jgo)](https://pkg.go.dev/mod/github.com/xjustloveux/jgo)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/xjustloveux/jgo/blob/master/LICENSE)

---

* [Overview](#Overview)
* [Middlewares](#Middlewares)
* [Installation](#Installation)
* [Example](#Example)
    * [jsql](#jsql)
    * [jcron](#jcron)
    * [jlog](#jlog)
    * [jfile](#jfile)
    * [jtime](#jtime)
* [Environment](#Environment)
* [API](#API)

# Overview

---

JGo provides an easier configuration for writing sql, log, and cron jobs.

JGo Web：[https://jgo.dev](https://jgo.dev)

JGo Web Project：[https://github.com/xjustloveux/jgo.web](https://github.com/xjustloveux/jgo.web)

# Middlewares

---

**JGo minimizes dependencies on third-party middleware to avoid conflicts.**

jlog only import [logrus](https://github.com/sirupsen/logrus) middleware.

jsql only import [govaluate](https://github.com/Knetic/govaluate) middleware, but it is designed on the basis
of [mysql](https://github.com/go-sql-driver/mysql), [go-mssqldb](https://github.com/denisenkom/go-mssqldb), [godror](https://github.com/godror/godror) and [pq](https://github.com/lib/pq).

# Installation

---

```shell
go get github.com/xjustloveux/jgo
```

# Example

---

## jsql

### Configuration

Configuration file default `json` format, you can use `jsql.SetFormat` to set you want, but you must be used
[jfile.RegisterCodec](#RegisterCodec) register codec.

| Property Name                      | Required | Type                   | Default Value | Comment                                                                                                                                                                                                                                                                                                                                                |
|------------------------------------|----------|------------------------|---------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| DaoPath                            | true     | string                 | empty         | It your xml files folder path.                                                                                                                                                                                                                                                                                                                         |
| Default                            | true     | string                 | empty         | It your default DataSource name.                                                                                                                                                                                                                                                                                                                       |
| DataSource                         | true     | map[string]interface{} | empty         |                                                                                                                                                                                                                                                                                                                                                        |
| DataSource.Type                    | true     | string                 | empty         | You can set `MySql`, `MSSql`, `Oracle` or `PostgreSql`, others as long as the sql parameter supports '?'.                                                                                                                                                                                                                                              |
| DataSource.DSN                     | true     | string                 | empty         | DataSourceName. If you have information security considerations, you can encrypt the DataSource into a string, and set the decryption function.                                                                                                                                                                                                        |
| DataSource.DN                      | false    | string                 | empty         | DriverName. If your type is `MySql`, `MSSql`, `Oracle` or `PostgreSql`, the DriverName default use [mysql](https://github.com/go-sql-driver/mysql), [go-mssqldb](https://github.com/denisenkom/go-mssqldb), [godror](https://github.com/godror/godror) and [pq](https://github.com/lib/pq) driver. You also can set this value to use your DriverName. |
| DataSource.DbName                  | false    | string                 | empty         | It your db name.                                                                                                                                                                                                                                                                                                                                       |
| DataSource.ConnMaxLifetime         | false    | time.Duration          | 120           |                                                                                                                                                                                                                                                                                                                                                        |
| DataSource.ConnMaxLifetimeDuration | false    | string                 | Second        | Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day                                                                                                                                                                                                                                                                                        |
| DataSource.ConnMaxIdleTime         | false    | time.Duration          | 0             |                                                                                                                                                                                                                                                                                                                                                        |
| DataSource.ConnMaxIdleTimeDuration | false    | string                 | Hour          | Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day                                                                                                                                                                                                                                                                                        |
| DataSource.MaxOpenConns            | false    | int                    | 0             |                                                                                                                                                                                                                                                                                                                                                        |
| DataSource.MaxIdleConns            | false    | int                    | 0             |                                                                                                                                                                                                                                                                                                                                                        |
| DataSource.EncodeData              | false    | string                 | empty         | If you have information security considerations, you can encrypt the DataSource into a string, and set the decryption Format and function.                                                                                                                                                                                                             |
| DataSource.Format                  | false    | jfile.Format           | jfile.Json    | `DataSource.EncodeData` format. If you want use other format, you must be use [jfile.RegisterCodec](#RegisterCodec) register codec.                                                                                                                                                                                                                    |

### Usage

#### config.json

```json
{
  "db": {
    "daoPath": "dao/",
    "default": "exampleMySql",
    "dataSource": {
      "exampleMySql": {
        "type": "MySql",
        "dsn": "user:password@tcp(192.168.1.1:3306)/DBName?checkConnLiveness=false&maxAllowedPacket=0&charset=utf8mb4&parseTime=true",
        "dbName": "DBName"
      },
      "exampleMSSql": {
        "type": "MSSql",
        "dsn": "Data Source=192.168.1.1,1433;Initial Catalog=DBName;Integrated Security=False;User ID=user;Password=password;Connection Timeout=120;MultipleActiveResultSets=True",
        "dbName": "DBName"
      },
      "exampleOracle": {
        "type": "Oracle",
        "dsn": "user/password@192.168.1.1:1521/ORCLCDB",
        "dbName": "DBName"
      },
      "examplePostgreSql": {
        "type": "PostgreSql",
        "dsn": "postgresql://user:password@192.168.1.1:5432/DBName?sslmode=disable",
        "dbName": "DBName"
      }
    }
  }
}
```

#### dao/example.xml

For details, please refer to [XmlTag](#XmlTag).

⚠️**The tags ${} and #{} will directly become SQL statements at the end, which may cause SQL injection problems. Please
use them with caution.**

```xml
<?xml version="1.0" encoding="UTF-8"?>
<dao>
    <select id="example1">
        SELECT * FROM TABLE1
    </select>
    <select id="example2">
        SELECT * FROM TABLE2
        <where>
            COL1 = @{COL1}
            <if test="!nil(COL2) and COL2 != ''">
                AND COL2 = @{COL2}
            </if>
        </where>
    </select>
    <select id="example3">
        SELECT ${COL1}, ${COL2} FROM TABLE3
        <orderBy last="true">
            ${SORT} DESC
        </orderBy>
    </select>
    <select id="example4">
        SELECT
        <foreach params="list" open="" separator="," close="">
            #{val}
        </foreach>
        FROM TABLE4
    </select>
    <insert id="insertExample">
        INSERT INTO ${TABLE}
        <foreach params="list" open="(" separator="," close=")">
            #{val}
        </foreach>
        VALUES
        <foreach params="list" open="(" separator="," close=")">
            @{#{val}}
        </foreach>
    </insert>
</dao>
```

#### main.go

```go
package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	_ "github.com/lib/pq"
	"github.com/xjustloveux/jgo/jsql"
)

func main() {
	if err := jsql.Init(); err != nil {
		fmt.Println(err)
		return
	}
}
```

#### example1

```go
package main

func example1() {
	// SELECT * FROM TABLE1
	if agent, err := jsql.GetAgent(); err != nil {
		fmt.Println(err)
	} else {
		var res jsql.Result
		if res, err = agent.Query("example1"); err != nil {
			fmt.Println(err)
		} else {
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}
```

#### example2

```go
package main

func example2() {
	// SELECT * FROM TABLE2 WHERE COL1 = ? AND COL2 = ?
	param := make(map[string]interface{})
	param["COL1"] = "VAL1"
	param["COL2"] = "VAL2"
	if agent, err := jsql.GetAgent(); err != nil {
		fmt.Println(err)
	} else {
		var res jsql.Result
		if res, err = agent.QueryPage("example2", 6, 10, param); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}
```

#### example3

```go
package main

func example3() {
	// SELECT COL_NAME1, COL_NAME2 FROM TABLE3 ORDER BY SORT_COL_NAME DESC
	param := make(map[string]interface{})
	param["COL1"] = "COL_NAME1"
	param["COL2"] = "COL_NAME2"
	param["SORT"] = "SORT_COL_NAME"
	if agent, err := jsql.GetAgent("exampleMSSql"); err != nil {
		fmt.Println(err)
	} else {
		var res jsql.Result
		if res, err = agent.QueryPage("example3", 6, 10, param); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}

	// you also can use struct
	type Param struct {
		COL1 string
		COL2 string
		SORT string
	}
	param2 := Param{
		COL1: "COL_NAME1",
		COL2: "COL_NAME2",
		SORT: "SORT_COL_NAME",
	}
	type Data struct {
		COL_NAME1 string
		COL_NAME2 string
	}
	type List struct {
		Rows []Data
	}
	var list List
	if agent, err := jsql.GetAgent("exampleMSSql"); err != nil {
		fmt.Println(err)
	} else {
		if _, err = agent.QueryPage("example3", 6, 10, param2, &list); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(list)
		}
	}
}
```

#### example4

```go
package main

func example4() {
	// SELECT COL_NAME1, COL_NAME2 FROM TABLE4
	param := make(map[string]interface{})
	list := []string{"COL_NAME1", "COL_NAME2"}
	param["list"] = list
	if agent, err := jsql.GetAgent("exampleOracle"); err != nil {
		fmt.Println(err)
	} else {
		var res jsql.Result
		if res, err = agent.QueryPage("example4", 6, 10, param); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}
```

#### example5

```go
package main

func example5() {
	// INSERT INTO TABLE5 (COL_NAME1, COL_NAME2, COL_NAME3) VALUES (?, ?, ?)
	list := []string{"COL_NAME1", "COL_NAME2", "COL_NAME3"}
	param := make(map[string]interface{})
	param["list"] = list
	param["TABLE"] = "TABLE5"
	param["COL_NAME1"] = 123
	param["COL_NAME2"] = "Test"
	param["COL_NAME3"] = time.Now()
	if agent, err := jsql.GetAgent(); err != nil {
		fmt.Println(err)
	} else {
		if _, err = agent.Insert("insertExample", param); err != nil {
			fmt.Println(err)
		}
	}
}
```

#### example6

```go
package main

func example6() {
	// INSERT INTO TABLE5 (COL_NAME1, COL_NAME2, COL_NAME3) VALUES (?, ?, ?)
	list := []string{"COL_NAME1", "COL_NAME2", "COL_NAME3"}
	param := make(map[string]interface{})
	param["list"] = list
	param["TABLE"] = "TABLE5"
	param["COL_NAME1"] = 123
	param["COL_NAME2"] = "Test1"
	param["COL_NAME3"] = time.Now()
	if agent, err := jsql.GetAgent(); err != nil {
		fmt.Println(err)
	} else {
		// or you can use agent.UseTx easier
		if _, err = agent.Begin(); err != nil {
			fmt.Println(err)
		} else {
			defer func() {
				if err != nil {
					if e := agent.Rollback(); e != nil {
						fmt.Println(e)
					}
				}
			}()
			if _, err = agent.InsertTx("insertExample", param); err != nil {
				fmt.Println(err)
				return
			}
			param["COL_NAME1"] = 456
			param["COL_NAME2"] = "Test2"
			param["COL_NAME3"] = time.Now()
			if _, err = agent.InsertTx("insertExample", param); err != nil {
				fmt.Println(err)
				return
			}
			if err = agent.Commit(); err != nil {
				fmt.Println(err)
			}
		}
	}
}
```

#### example7

```go
package main

func example7() {
	// SELECT * FROM TABLE6 WHERE COL1 = ?
	var res jsql.Result
	ta := &jsql.TableAgent{Table: "TABLE6"}
	ta.Equal("COL1", "VAL1")
	if res, err = ta.QueryPage(6, 10); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.TotalRecord())
		for _, item := range res.Rows() {
			fmt.Println(item)
		}
	}
}
```

### XmlTag

| Tag Name | Layer | Attr Name | Required | Type   | Comment                                                                                                  |
|----------|-------|-----------|----------|--------|----------------------------------------------------------------------------------------------------------|
| dao      | 1     |           | true     |        |                                                                                                          |
| select   | 2     |           |          |        |                                                                                                          |
|          |       | id        | true     | string |                                                                                                          |
| insert   | 2     |           |          |        |                                                                                                          |
|          |       | id        | true     | string |                                                                                                          |
| update   | 2     |           |          |        |                                                                                                          |
|          |       | id        | true     | string |                                                                                                          |
| delete   | 2     |           |          | string |                                                                                                          |
|          |       | id        | true     | string |                                                                                                          |
| other    | 2     |           |          | string |                                                                                                          |
|          |       | id        | true     | string |                                                                                                          |
| if       | 3 up  |           |          |        |                                                                                                          |
|          |       | test      | true     | string | expression, support nil check use nil()<br/>middleware: [govaluate](https://github.com/Knetic/govaluate) |
| foreach  | 3 up  |           |          |        |                                                                                                          |
|          |       | params    | true     | string | param key, param type can be map or slice                                                                |
|          |       | open      | false    | string |                                                                                                          |
|          |       | separator | false    | string |                                                                                                          |
|          |       | close     | false    | string |                                                                                                          |
| where    | 3 up  |           |          |        |                                                                                                          |
| orderBy  | 3 up  |           |          |        |                                                                                                          |
|          |       | last      | false    | bool   | for QueryPage                                                                                            |

## jcron

### Configuration

Configuration file default `json` format, you can use `jcron.SetFormat` to set you want, but you must be used
[jfile.RegisterCodec](#RegisterCodec) register codec.

| Property Name    | Required | Type                   | Default Value | Comment                            |
|------------------|----------|------------------------|---------------|------------------------------------|
| Location         | false    | string                 | empty         | time location                      |
| Schedule         | true     | []*SchInfo             | empty         |                                    |
| Schedule.Name    | true     | string                 | empty         | schedule name.                     |
| Schedule.Cron    | true     | string                 | empty         | [CronExpression](#CronExpression). |
| Schedule.JobName | true     | string                 | empty         | schedule job name.                 |
| Schedule.JobData | false    | map[string]interface{} | empty         | schedule job data.                 |
| Schedule.Desc    | false    | string                 | empty         | schedule job description.          |
| Schedule.Status  | false    | string                 | run           | schedule status.                   |

### Usage

#### config.json

```json
{
  "cron": {
    "schedule": [
      {
        "Name": "Sch01",
        "Cron": "7-43/13 * * * * ? *",
        "JobName": "Job01",
        "JobData": {
          "data": "val"
        },
        "Desc": "this is schedule 01-----------"
      },
      {
        "Name": "Sch02",
        "Cron": "3,7,11,32-57/7 * * * * ? *",
        "JobName": "Job02",
        "JobData": {},
        "Desc": "this is schedule 02-----------"
      }
    ]
  }
}
```

#### main.go

```go
package main

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcron"
	"github.com/xjustloveux/jgo/jtime"
	"time"
)

func main() {
	if err := jcron.Init(); err != nil {
		fmt.Println(err)
		return
	}
	if err := jcron.AddJobFunc("Job01", Job01); err != nil {
		fmt.Println(err)
		return
	}
	if err := jcron.AddJobFunc("Job02", Job02); err != nil {
		fmt.Println(err)
		return
	}
	jcron.Start()
	select {
	case <-time.After(3 * jtime.Minute):
	}
	jcron.Wait()
}

func Job01(data map[string]interface{}) {
	// do something
}

func Job02(data map[string]interface{}) {
	// do something
}
```

### CronExpression

| Name         | Required | Allowed Values                      | Allowed Special Characters |
|--------------|----------|-------------------------------------|----------------------------|
| Seconds      | true     | 0-59                                | ,-*/                       |
| Minutes      | true     | 0-59                                | ,-*/                       |
| Hours        | true     | 0-23                                | ,-*/                       |
| Day of month | true     | 1-31                                | ,-*/?                      |
| Month        | true     | 1-12 or January-December or JAN-DEC | ,-*/                       |
| Day of week  | true     | 0-6 or Sunday-Saturday or SUN-SAT   | ,-*/?                      |
| Year         | false    | 2020-2080                           | ,-*/                       |

## jlog

### Configuration

Configuration file default `json` format, you can use `jlog.SetFormat` to set you want, but you must be used
[jfile.RegisterCodec](#RegisterCodec) register codec.

| Property Name                        | Required | Type                   | Default Value                                          | Comment                                                                                                              |
|--------------------------------------|----------|------------------------|--------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| Params                               | false    | map[string]interface{} | empty                                                  | `Appender.Output.P` and `Appender.Output.LinkName` replace params                                                    |
| Default                              | false    | []string               | empty                                                  | Appender name                                                                                                        |
| Appender                             | true     | map[string]interface{} | empty                                                  |                                                                                                                      |
| Appender.Level                       | false    | string                 | info                                                   | `logrus.Level`                                                                                                       |
| Appender.ReportCaller                | false    | bool                   | true                                                   | `logrus.ReportCaller`                                                                                                |
| Appender.Formatter                   | false    | *formatter             | default formatter                                      |                                                                                                                      |
| Appender.Formatter.Type              | false    | string                 | TEXT                                                   | `TEXT`, `JSON`                                                                                                       |
| Appender.Formatter.Location          | false    | string                 | empty                                                  | time location                                                                                                        |
| Appender.Formatter.Text              | false    | *logrus.TextFormatter  | &logrus.TextFormatter{TimestampFormat: jtime.DateTime} |                                                                                                                      |
| Appender.Formatter.Json              | false    | *logrus.JSONFormatter  | &logrus.JSONFormatter{TimestampFormat: jtime.DateTime} |                                                                                                                      |
| Appender.Output                      | true     | *output                | default output                                         |                                                                                                                      |
| Appender.Output.Name                 | false    | string                 | empty                                                  | writer name. you can use `.AddWriter` add writer before `.Init`                                                      |
| Appender.Output.P                    | true     | string                 | empty                                                  | log file path. default `${Program}` tag replace program name. time format refer to [FormatString](#FormatString)     |
| Appender.Output.Clock                | false    | string                 | Local                                                  | `*time.Location` string, use for file path time format.                                                              |
| Appender.Output.LinkName             | false    | string                 | empty                                                  | symlink file path. default `${Program}` tag replace program name. time format refer to [FormatString](#FormatString) |
| Appender.Output.MaxAge               | false    | time.Duration          | 365                                                    | MaxAge number                                                                                                        |
| Appender.Output.MaxAgeDuration       | false    | string                 | Day                                                    | Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day                                                      |
| Appender.Output.RotationTime         | false    | time.Duration          | 24                                                     | RotationTime number                                                                                                  |
| Appender.Output.RotationTimeDuration | false    | string                 | Hour                                                   | Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day                                                      |
| Appender.Output.RotationSize         | false    | int64                  | 10                                                     | RotationSize number                                                                                                  |
| Appender.Output.RotationSizeUnit     | false    | string                 | MB                                                     | Byte, Kb, KB, Mb, MB, Gb, GB, Tb, TB, Pb, PB, Eb, EB                                                                 |
| Appender.Output.RotationCount        | false    | int                    | 0                                                      | RotationCount number                                                                                                 |
| Appender.Output.Handler              | false    | string                 | empty                                                  | handler name. you can use `.AddHandler` add handler before `.Init`                                                   |
| Logs                                 | true     | []*logs                | empty                                                  |                                                                                                                      |
| Logs.Program                         | true     | []string               | empty                                                  | go program name. ex: `main` or `main.go` or `pkg:sample`                                                             |
| Logs.Appender                        | true     | []string               | empty                                                  | Appender name.                                                                                                       |

***Note:* `jlog.Init` default create 'default' program name and 'console' appender name.**

### Usage

#### config.json

```json
{
  "log": {
    "params": {
      "path": "/log"
    },
    "appender": {
      "sys": {
        "level": "Error",
        "formatter": {
          "type": "JSON"
        },
        "output": {
          "p": "${path}/sys/%yyyy-%MM-%dd/system.log",
          "utc": false,
          "linkName": "${path}/sys/system",
          "rotationSize": 100,
          "rotationSizeUnit": "KB"
        }
      },
      "web": {
        "formatter": {
          "text": {
            "timestampFormat": "2006-01-02"
          }
        },
        "output": {
          "p": "${path}/web/%yyyy-%MM-%dd/website.log",
          "linkName": "${path}/web/website"
        }
      }
    },
    "logs": [
      {
        "program": [
          "program1",
          "main"
        ],
        "appender": [
          "sys"
        ]
      },
      {
        "program": [
          "program2"
        ],
        "appender": [
          "console",
          "web"
        ]
      },
      {
        "program": [
          "pkg:packageName"
        ],
        "appender": [
          "console",
          "sys"
        ]
      }
    ]
  }
}
```

#### main.go

sys appender log

```go
package main

import (
	"fmt"
	"github.com/xjustloveux/jgo/jlog"
	"time"
)

func main() {
	if err := jlog.Init(); err != nil {
		fmt.Println(err)
		return
	}
	jlog.Info("info message")
	jlog.Warn(time.Now())
	jlog.Error(fmt.Errorf("error message"))
}
```

#### program2.go

console and web appender log

```go
package main

func example() {
	jlog.Info("this is program2")
}
```

## jfile

### Convert

```go
type Example struct{
str string
}

m := make(map[string]interface{})
m["str"] = "example"
ex := Example{}

jfile.Convert(m, &ex)
```

### RegisterCodec

```go
package main

import (
	"encoding/json"
	"github.com/xjustloveux/jgo/jfile"
)

type jsonCodec struct{}

func (jsonCodec) Encode(m map[string]interface{}) ([]byte, error) {
	return json.MarshalIndent(m, "", "  ")
}

func (jsonCodec) Decode(b []byte, m map[string]interface{}) error {
	return json.Unmarshal(b, &m)
}

func main() {
	jfile.RegisterCodec(jfile.Json.String(), jsonCodec{})
}
```

## jtime

### FormatString

E.g. 2022-06-20 14:44:55.012345678 -0700 MST IN America/Phoenix Location

| Pattern | Description                                | E.g.            |
|---------|--------------------------------------------|-----------------|
| %d      | day in month, 1-31                         | 20              |
| %dd     | day in month, 01-31                        | 20              |
| %ddd    | day in week, Sun-Mon                       | Mon             |
| %dddd   | day in week, Sunday-Monday                 | Monday          |
| %D      | day in year, 1-366                         | 171             |
| %DD     | day in year, 01-366                        | 171             |
| %DDD    | day in year, 001-366                       | 171             |
| %f      | nanosecond, 0-9999                         | 0               |
| %ff     | nanosecond, 00-9999                        | 01              |
| %fff    | nanosecond, 000-9999                       | 012             |
| %ffff   | nanosecond, 0000-9999                      | 0123            |
| %F      | nanosecond, -9999                          |                 |
| %FF     | nanosecond, -9999                          | 01              |
| %FFF    | nanosecond, -9999                          | 012             |
| %FFFF   | nanosecond, -9999                          | 0123            |
| %g      | Era designator, A.D. or B.C.               | A.D.            |
| %h      | hour in day, 0-11                          | 2               |
| %hh     | hour in day, 00-11                         | 02              |
| %H      | hour in day, 0-23                          | 14              |
| %HH     | hour in day, 00-23                         | 14              |
| %k      | hour in day, 1-12                          | 2               |
| %kk     | hour in day, 01-12                         | 02              |
| %K      | hour in day, 1-24                          | 14              |
| %KK     | hour in day, 01-24                         | 14              |
| %l      | time zone name                             | America/Phoenix |
| %m      | minute in hour, 0-59                       | 44              |
| %mm     | minute in hour, 00-59                      | 44              |
| %M      | month in year, 1-12                        | 6               |
| %MM     | month in year, 01-12                       | 06              |
| %MMM    | month in year, Jan-Dec                     | Jun             |
| %MMMM   | month in year, January-December            | June            |
| %s      | second in minute, 0-59                     | 55              |
| %ss     | second in minute, 00-59                    | 55              |
| %t      | AM/PM marker                               | P               |
| %tt     | AM/PM marker                               | PM              |
| %w      | week in year, 1-53                         | 25              |
| %W      | day in week, 0-6                           | 1               |
| %y      | year                                       | 22              |
| %yy     | year                                       | 22              |
| %yyy    | year                                       | 2022            |
| %yyyy   | year                                       | 2022            |
| %z      | time zone offset from UTC, hour            | -7              |
| %zz     | time zone offset from UTC, hour            | -07             |
| %zzz    | time zone offset from UTC, hour and minute | -07:00          |
| %zzzz   | time zone offset from UTC, second          | -25200          |
| %Z      | time zone name                             | MST             |

# Environment

---

jsql, jlog, jcron, jconf support environment variables.

Config required to be set file name, this file will be basic setting.

jsql, jlog, jcron default file name is 'config.json'.

You can override the basic setting by setting the env file name, env val or env key.

### Example1

#### default.json

```json
{
  "val": 123,
  "type": "default"
}
```

#### override.json

```json
{
  "type": "override",
  "name": "test"
}
```

#### main.go

```go
package main

import (
	"fmt"
	"github.com/xjustloveux/jgo/jconf"
)

func main() {
	conf := jconf.New()
	conf.SetFileName("default.json")
	conf.SetEnvFileName("override.json")
	if err := conf.Load(); err != nil {
		fmt.Println(err)
	}
}
```

#### last json data

```json
{
  "val": 123,
  "type": "override",
  "name": "test"
}
```

### Example2

#### main.go

```go
package main

import (
	"fmt"
	"github.com/xjustloveux/jgo/jconf"
)

func main() {
	conf := jconf.New()
	conf.SetFileName("default.json")
	conf.SetEnvVal("dev")
	// Override file name is 'default-dev.json'.
	if err := conf.Load(); err != nil {
		fmt.Println(err)
	}
}
```

### Example3

#### main.go

```go
package main

import (
	"fmt"
	"github.com/xjustloveux/jgo/jconf"
)

func main() {
	conf := jconf.New()
	conf.SetFileName("default.json")
	conf.SetEnvKey("goEnv")
	// If the 'default.json' file or os have 'goEnv' environment variables.
	// Override file name is 'default-$goEnv.json'.
	// E.g. goEnv value is 'dev', then the override file name is 'default-dev.json'.
	if err := conf.Load(); err != nil {
		fmt.Println(err)
	}
}
```

### Example4

#### main.go

```go
package main

import (
	"fmt"
	"github.com/xjustloveux/jgo/jconf"
)

func main() {
	conf := jconf.New()
	conf.SetFileName("default.json")
	// If the 'default.json' file or os have 'jEnv' environment variables.
	// Override file name is 'default-$jEnv.json'.
	// E.g. jEnv value is 'dev', then the override file name is 'default-dev.json'.
	if err := conf.Load(); err != nil {
		fmt.Println(err)
	}
}
```

# API

---

[jcast](https://pkg.go.dev/github.com/xjustloveux/jgo/jcast) ,
[jconf](https://pkg.go.dev/github.com/xjustloveux/jgo/jconf) ,
[jcron](https://pkg.go.dev/github.com/xjustloveux/jgo/jcron) ,
[jevent](https://pkg.go.dev/github.com/xjustloveux/jgo/jevent) ,
[jfile](https://pkg.go.dev/github.com/xjustloveux/jgo/jfile) ,
[jlog](https://pkg.go.dev/github.com/xjustloveux/jgo/jlog) ,
[jruntime](https://pkg.go.dev/github.com/xjustloveux/jgo/jruntime) ,
[jslice](https://pkg.go.dev/github.com/xjustloveux/jgo/jslice) ,
[jsql](https://pkg.go.dev/github.com/xjustloveux/jgo/jsql) ,
[jtime](https://pkg.go.dev/github.com/xjustloveux/jgo/jtime)