# jgo
[![codecov](https://codecov.io/gh/xjustloveux/jgo/branch/master/graph/badge.svg?token=RCO5VO2YU6)](https://codecov.io/gh/xjustloveux/jgo)
[![Build Status](https://app.travis-ci.com/xjustloveux/jgo.svg?branch=master)](https://app.travis-ci.com/xjustloveux/jgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/xjustloveux/jgo)](https://goreportcard.com/report/github.com/xjustloveux/jgo)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/xjustloveux/jgo)](https://pkg.go.dev/mod/github.com/xjustloveux/jgo)

---

<a href="https://dream-ja.com/">
  <img alt="JaJa" align="right" width="30%" height="150px" src="https://dream-ja.com/assets/svg/logo-white.svg">
</a>

* [Overview](#Overview)
* [Middlewares](#Middlewares)
* [Installation](#Installation)
* [Example](#Example)
    * [jsql](#jsql)
    * [jlog](#jlog)
    * [jcron](#jcron)
* [Environment](#Environment)
* [Api](#Api)
* [License](#License)

# Overview

---

Jgo provides an easier configuration for writing sql, log, and cron jobs.

# Middlewares

---

**Jgo minimizes dependencies on third-party middleware to avoid conflicts.**

jlog import two third-party middlewares that [logrus](https://github.com/sirupsen/logrus)
and [file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs).

jsql only import [govaluate](https://github.com/Knetic/govaluate) middleware, but it is designed on the basis
of [mysql](https://github.com/go-sql-driver/mysql), [go-mssqldb](https://github.com/denisenkom/go-mssqldb)
and [go-oci8](https://github.com/mattn/go-oci8).

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

| Property Name                      | Required | Type                   | Default Value | Comment                                                                                                                                                                                                                                                                                                  |
|------------------------------------|----------|------------------------|---------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Debug                              | false    | bool                   | false         | When set true that will print the error sql.                                                                                                                                                                                                                                                             |
| DaoPath                            | true     | string                 | empty         | It your xml files folder path.                                                                                                                                                                                                                                                                           |
| Default                            | true     | string                 | empty         | It your default DataSource name.                                                                                                                                                                                                                                                                         |
| DataSource                         | true     | map[string]interface{} | empty         |                                                                                                                                                                                                                                                                                                          |
| DataSource.Type                    | true     | string                 | empty         | You can set `MySql`, `MSSql` or `Oracle`, others as long as the sql parameter supports '?'.                                                                                                                                                                                                              |
| DataSource.DSN                     | true     | string                 | empty         | DataSourceName. If you have information security considerations, you can encrypt the DataSource into a string, and set the decryption function.                                                                                                                                                          |
| DataSource.DN                      | false    | string                 | empty         | DriverName. If your type is `MySql`, `MSSql` or `Oracle`, the DriverName default use [mysql](https://github.com/go-sql-driver/mysql), [go-mssqldb](https://github.com/denisenkom/go-mssqldb) and [go-oci8](https://github.com/mattn/go-oci8) driver. You also can set this value to use your DriverName. |
| DataSource.ConnMaxLifetime         | false    | time.Duration          | 120           |                                                                                                                                                                                                                                                                                                          |
| DataSource.ConnMaxLifetimeDuration | false    | string                 | Second        | Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day                                                                                                                                                                                                                                          |
| DataSource.ConnMaxIdleTime         | false    | time.Duration          | 0             |                                                                                                                                                                                                                                                                                                          |
| DataSource.ConnMaxIdleTimeDuration | false    | string                 | Hour          | Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day                                                                                                                                                                                                                                          |
| DataSource.MaxOpenConns            | false    | int                    | 0             |                                                                                                                                                                                                                                                                                                          |
| DataSource.MaxIdleConns            | false    | int                    | 0             |                                                                                                                                                                                                                                                                                                          |
| DataSource.EncodeData              | false    | string                 | empty         | If you have information security considerations, you can encrypt the DataSource into a string, and set the decryption Format and function.                                                                                                                                                               |
| DataSource.Format                  | false    | jfile.Format           | jfile.Json    | `DataSource.EncodeData` format. If you want use other format, you must be use [jfile.RegisterCodec](#RegisterCodec) register codec.                                                                                                                                                                      |

### Usage

#### sql.json

```json
{
  "daoPath": "dao/",
  "default": "exampleMySql",
  "dataSource": {
    "exampleMySql": {
      "type": "MySql",
      "dsn": "user:password@tcp(192.168.1.1:3306)/DBName?checkConnLiveness=false&maxAllowedPacket=0&charset=utf8mb4&parseTime=true"
    },
    "exampleMSSql": {
      "type": "MSSql",
      "dsn": "Data Source=192.168.1.1,1433;Initial Catalog=DBName;Integrated Security=False;User ID=user;Password=password;Connection Timeout=120;MultipleActiveResultSets=True"
    },
    "exampleOracle": {
      "type": "Oracle",
      "dsn": "user/password@192.168.1.1:1521/ORCLCDB"
    }
  }
}
```

#### dao/example.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<dao>
    <select id="example1">
        SELECT * FROM TABLE1
    </select>
    <select id="example2">
        SELECT * FROM TABLE2
        WHERE COL1 = @{COL1}
        <if test="!Nil(COL2) and COL2 != ''">
            AND COL2 = @{COL2}
        </if>
    </select>
    <select id="example3">
        SELECT ${COL1}, ${COL2} FROM TABLE3
        <orderBy lastTag="true">
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
	_ "github.com/mattn/go-oci8"
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

## jlog

### Configuration

Configuration file default `json` format, you can use `jlog.SetFormat` to set you want, but you must be used
[jfile.RegisterCodec](#RegisterCodec) register codec.

| Property Name                        | Required | Type                   | Default Value                                          | Comment                                                          |
|--------------------------------------|----------|------------------------|--------------------------------------------------------|------------------------------------------------------------------|
| Debug                                | false    | bool                   | false                                                  | When set true that will print the jlog error.                    |
| Params                               | false    | map[string]interface{} | empty                                                  | `Appender.Output.P` or `Appender.Output.LinkName` replace params |
| Appender                             | true     | map[string]interface{} | empty                                                  |                                                                  |
| Appender.Level                       | false    | string                 | info                                                   | `logrus.Level`                                                   |
| Appender.Formatter                   | false    | *formatter             | default formatter                                      |                                                                  |
| Appender.Formatter.Type              | false    | string                 | TEXT                                                   | `TEXT`, `JSON`                                                   |
| Appender.Formatter.Text              | false    | *logrus.TextFormatter  | &logrus.TextFormatter{TimestampFormat: jtime.DateTime} |                                                                  |
| Appender.Formatter.Json              | false    | *logrus.JSONFormatter  | &logrus.JSONFormatter{TimestampFormat: jtime.DateTime} |                                                                  |
| Appender.Output                      | true     | *output                | default output                                         |                                                                  |
| Appender.Output.P                    | true     | string                 | empty                                                  | file-rotatelogs `.New(P, ...)`                                   |
| Appender.Output.UTC                  | false    | bool                   | false                                                  | file-rotatelogs `.WithClock`                                     |
| Appender.Output.LinkName             | false    | string                 | empty                                                  | file-rotatelogs `.WithLinkName`                                  |
| Appender.Output.MaxAge               | false    | time.Duration          | 365                                                    | file-rotatelogs `.WithMaxAge`                                    |
| Appender.Output.MaxAgeDuration       | false    | string                 | Day                                                    | Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day  |
| Appender.Output.RotationTime         | false    | time.Duration          | 24                                                     | file-rotatelogs .WithRotationTime                                |
| Appender.Output.RotationTimeDuration | false    | string                 | Hour                                                   | Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day  |
| Appender.Output.RotationSize         | false    | int64                  | 10                                                     | file-rotatelogs .WithRotationSize                                |
| Appender.Output.RotationSizeUnit     | false    | string                 | MB                                                     | Byte, Kb, KB, Mb, MB, Gb, GB, Tb, TB, Pb, PB, Eb, EB             |
| Appender.Output.RotationCount        | false    | uint                   | 0                                                      | file-rotatelogs `.WithRotationCount`                             |
| Logs                                 | true     | []*logs                | empty                                                  |                                                                  |
| Logs.Program                         | true     | []string               | empty                                                  | go program name. ex: main or main.go                             |
| Logs.Appender                        | true     | []string               | empty                                                  | Appender name.                                                   |

**note: `jlog.Init` default create 'default' program name and 'console' appender name.**

### Usage

#### log.json

```json
{
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
        "p": "${path}/sys/%Y-%m-%d/system.log",
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
        "p": "${path}/web/%Y-%m-%d/website.log",
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
    }
  ]
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

## jcron

### Configuration

Configuration file default `json` format, you can use `jcron.SetFormat` to set you want, but you must be used
[jfile.RegisterCodec](#RegisterCodec) register codec.

| Property Name    | Required | Type                   | Default Value | Comment                                        |
|------------------|----------|------------------------|---------------|------------------------------------------------|
| Debug            | false    | bool                   | false         | When set true that will print the jcron error. |
| Schedule         | true     | []*SchInfo             | empty         |                                                |
| Schedule.Name    | true     | string                 | empty         | schedule name.                                 |
| Schedule.Cron    | true     | string                 | empty         | [CronExpression](#CronExpression).             |
| Schedule.JobName | true     | string                 | empty         | schedule job name.                             |
| Schedule.JobData | false    | map[string]interface{} | empty         | schedule job data.                             |
| Schedule.Desc    | false    | string                 | empty         | schedule job description.                      |

### Usage

#### log.json

```json
{
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

# Environment

---

jsql, jlog, jcron, jconf support environment variables.

Config required to be set file name, this file will be basic setting.

jsql default file name is 'sql.json'.

jlog default file name is 'log.json'.

jcron default file name is 'cron.json'.

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

# Api

---

## jfile

### Load

```go
jfile.Load("C:/example.txt")
```

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

### GetCodec

```go
jfile.GetCodec(jfile.Json.String())
```

### Encode

```go
m := make(map[string]interface{})
jfile.Encode(jfile.Json.String(), m)
```

### Decode

```go
package main

import (
	"github.com/xjustloveux/jgo/jfile"
)

func main() {
	var err error
	var b []byte
	if b, err = jfile.Load("C:/example.json"); err != nil {
		return
	}
	m := make(map[string]interface{})
	if err = jfile.Decode(jfile.Json.String(), b, m); err != nil {
		return
	}
}
```

### ParseSizeUnit

```go
jfile.ParseSizeUnit("MB")
```

## jtime

### ParseTimeDuration

```go
jtime.ParseTimeDuration("Day")
```

## jslice

### Filter

```go
package main

import (
	"github.com/xjustloveux/jgo/jslice"
)

func main() {
	s := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	jslice.Filter(s, func(i interface{}) bool {
		if v, ok := i.(int); ok {
			return v > 10
		}
		return false
	})
}
```

### Insert

```go
package main

import (
	"github.com/xjustloveux/jgo/jslice"
)

func main() {
	s := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	jslice.Insert(3, s, 6)
}
```

### InsertAll

```go
package main

import (
	"github.com/xjustloveux/jgo/jslice"
)

func main() {
	s1 := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	s2 := []int{2, 4, 6}
	jslice.InsertAll(3, s1, s2)
}
```

## jconf

### New

## jruntime
### GetFuncName
### GetCallerName
### GetCallerProgramName

## jcast

### VerifyPtr
### Value
### Time
### TimeLoc
### TimeString
### TimeFormatString
### String
### Bool
### Int
### Int8
### Int16
### Int32
### Int64
### Uint
### Uint8
### Uint16
### Uint32
### Uint64
### Float32
### Float64
### InterfaceMapInterface
### StringMapInterface
### SliceInterface
### SliceString
### SliceBool
### SliceInt
### SliceInt8
### SliceInt16
### SliceInt32
### SliceInt64
### SliceUint
### SliceUint8
### SliceUint16
### SliceUint32
### SliceUint64
### SliceFloat32
### SliceFloat64

# License

---

[MIT](https://github.com/xjustloveux/jgo/blob/master/LICENSE)