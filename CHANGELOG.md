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