# v1.3.16 (2023.02.15)
New Features：
* Add jsql.Agent function
* (*SqlAndArgs, Count, Exists, InsertWithLastInsertId)

# v1.3.15 (2023.02.02)
New Features：
* Add jlog ReportCaller config setting
* Add jruntime GetCallerProgramLine function

# v1.3.14 (2023.01.03)
Bug Fixes：
* Fixed jsql query page bug

# v1.3.13 (2022.12.11)
Changes：
* Change jcron load location timings

# v1.3.12 (2022.12.10)
Bug Fixes：
* Fixed jcron next time bug

# v1.3.11 (2022.12.10)
Bug Fixes：
* Fixed jlog rotate bug

# v1.3.10 (2022.12.05)
Bug Fixes：
* Remove jsql query page redundant subject event

# v1.3.9 (2022.12.06)
New Features：
* Add jcast.TimeLocTransform function

# v1.3.8 (2022.12.05)
New Features：
* jlog default appender setting custom

# v1.3.7 (2022.12.05)
Bug Fixes：
* Fixed jsql query page in postgresql bug

# v1.3.6 (2022.12.05)
New Features：
* jcron and jlog add time location setting

# v1.3.5 (2022.12.01)
Changes：
* Add jsql query page convert data RowStart, RowEnd and TotalRecord

# v1.3.4 (2022.12.01)
Bug Fixes：
* Fixed getDefault function warning

# v1.3.3 (2022.11.30)
Bug Fixes：
* Fixed jsql element order tag last attr ignore error

# v1.3.2 (2022.11.28)
Bug Fixes：
* Fixed jsql table schema bug

# v1.3.1 (2022.11.24)
New Features：
* Json value support

Bug Fixes：
* Fixed ignore dash character '_' error

# v1.3.0 (2022.11.23)
New Features：
* Add `PostgreSql` sql type, it is designed on basic of [pq](https://github.com/lib/pq)
* Add jsql test

Changes：
* `Oracle` sql type design reference middleware changed from [go-oci8](https://github.com/mattn/go-oci8) to [godror](https://github.com/godror/godror)

Bug Fixes：
* Fixed jsql.TableAgent.Drop no agent bug
* Fixed jcast.strToTime bug

# v1.2.14 (2022.11.05)
Changes：
* jsql.TableSchema PrimaryKey column type from string to interface{}

# v1.2.13 (2022.11.05)
Changes：
* jsql.TableSchema PrimaryKey column type from int8 to string

# v1.2.12 (2022.11.05)
Changes：
* jsql.TableSchema PrimaryKey column type from *int8 to int8

# v1.2.11 (2022.11.05)
New Features：
* Add jcron.GetScheduleInfo function

# v1.2.10 (2022.11.04)
New Features：
* Add jsql.TableSchema IsIdentity column

Bug Fixes：
* Fixed jsql sql default value column name bug

# v1.2.9 (2022.11.03)
Bug Fixes：
* Fixed jcast time format bug

# v1.2.8 (2022.11.03)
Bug Fixes：
* Fixed get record scan type nil no value bug

# v1.2.7 (2022.11.03)
Bug Fixes：
* Fixed get record scan type nil bug

# v1.2.6 (2022.11.03)
Bug Fixes：
* Fixed table schema convert bug
* Fixed scan type nil bug

# v1.2.5 (2022.11.03)
Bug Fixes：
* Fixed Agent.Tables bug

# v1.2.4 (2022.11.03)
Bug Fixes：
* Fixed table schema mssql and oracle sql param bug

# v1.2.3 (2022.10.31)
New Features：
* Add jfile.Exist function

Changes：
* errorf and errors function change to errorFmt and errorStr

# v1.2.2 (2022.10.30)
New Features：
* Add jsql.GetDataSource function
* Add jsql config data 'dbName'
* Add jsql Agent.UseTx, Agent.Tables, Agent.TablesTx, Agent.TableSchema and Agent.TableSchemaTx function

# v1.2.1 (2022.10.22)
Changes：
* jconf add root path and default value './config/'
* jsql, jcron and jlog can use SetRoot change root path value

# v1.2.0 (2022.10.22)
Changes：
* jsql, jcron and jlog change default load filename to 'config.json'
* jsql config data need under the 'db' tag name
* jcron config data need under the 'cron' tag name
* jlog config data need under the 'log' tag name
* Update golang version to 1.19
* Update sirupsen/logrus version to v1.9.0

# v1.1.4 (2022.06.26)
Bug Fixes：
* Fixed jlog replace program name have '.go' string

# v1.1.3 (2022.06.26)
Bug Fixes：
* Fixed jlog rotate time bug

# v1.1.2 (2022.06.26)
Remove：
* Remove doc.go

Bug Fixes：
* Fixed jlog rotate time bug

# v1.1.1 (2022.06.24)
Changes：
* jsql, jcron and jlog change log function to subscribe event

New Features：
* Add CHANGELOG.md
* Add jevent feature

Remove：
* Remove jsql, jcron, jlog config data debug name

Bug Fixes：
* Fixed jlog logger bug
* Fixed jtime.FormatString bug

# v1.1.0 (2022.06.21)
Remove：
* Remove package github.com/lestrrat-go/file-rotatelogs

# v1.0.13 (2022.06.17)
Bug Fixes：
* Fixed jcast bug for 32bit platform

# v1.0.12 (2022.06.17)
Bug Fixes：
* Fixed jcast deploy different platform max size type check bug

# v1.0.11 (2022.06.16)
Changes：
* jsql result change time string to time

# v1.0.10 (2022.06.14)
New Features：
* Create go ci action and change travis to github action 

Bug Fixes：
* Fixed jsql xml tag bug

# v1.0.9 (2022.06.10)
Bug Fixes：
* Fixed jfile convert bug

# v1.0.8 (2022.05.31)
Changes：
* Update go.mod
* Change jsql function from ParseTag to parseTag

New Features：
* Add jcron schedule default status features

# v1.0.7 (2022.05.30)
New Features：
* Add codecov ci action
* Add jcast deploy different platform max size type check
* Add jsql auto convert to struct features

Bug Fixes：
* Fixed jcast and jcron.CronExpression security build problem

# v1.0.6 (2022.05.27)
New Features：
* Add jsql xml where tag

# v1.0.5 (2022.05.26)
Changes：
* Change jsql load xml use from struct to xml tag

New Features：
* Add jlog package name judgment
* Add jruntime.GetPkgName and jruntime.GetCallerPkgName
* Add jcast.String xml.CharData and xml.Comment type

Bug Fixes：
* Fixed jfile.Convert bug

# v1.0.4 (2022.05.25)
Bug Fixes：
* Fixed jsql xml remove comment bug

# v1.0.3 (2022.05.25)
New Features：
* Add jcast.StringMapString

Bug Fixes：
* Fixed jcron lock unlock bug
* Fixed jsql xml and TableAgent bug

# v1.0.2 (2022.05.19)
Changes：
* CronExpression time type and operation mode change
* jlog most common function argument change

New Features：
* Add jlog.GetLogger, jlog.Fields, jlog.Level and jlog.LogFunction

Remove：
* Remove jcron.status unused ParseStatus function
* Remove jlog unused function

Bug Fixes：
* Fixed jcast.Value bug
* Fixed jcron chan bug
* Fixed jlog.appender write function use file-rotatelogs package bug

# v1.0.1 (2022.05.16)
New Features：
* Add travis ci
* jcast.Value add []bool type

Remove：
* Remove jtime and jslice unused function

Bug Fixes：
* Fixed jcast.strToTime and jcast.Float64 bug
* Fixed jconf.overwriteMap, jconf.overwriteSlice and config.Convert bug

# v1.0.0 (2022.05.13)
Initial Release