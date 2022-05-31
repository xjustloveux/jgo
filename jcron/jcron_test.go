// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jtime"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	testErr := "TEST ERROR:"
	SetFormat(jfile.Json)
	SetEnvFileName("")
	inEnvKey := ""
	SetEnvKey(inEnvKey)
	outEnvKey := EnvKey()
	assert.Equal(t, inEnvKey, outEnvKey, fmt.Sprintf("%v != %v", outEnvKey, inEnvKey))
	inEnvVal := ""
	SetEnvVal(inEnvVal)
	outEnvVal := EnvVal()
	assert.Equal(t, inEnvVal, outEnvVal, fmt.Sprintf("%v != %v", outEnvVal, inEnvVal))
	DisableEnv()
	EnableEnv()
	SetLogFunc(func(i ...interface{}) {})
	if err := Init(); err == nil {
		t.Error(fmt.Sprint(testErr, " Init must be return error"))
	}
	SetFileName("../files/test-jcron-error.json")
	if err := Init(); err == nil {
		t.Error(fmt.Sprint(testErr, " Init must be return error"))
	}
	SetFileName("../files/test-jcron.json")
	if err := Init(); err != nil {
		t.Error(err)
	}
	if err := AddJob("Job01", jobFunc(job01)); err != nil {
		t.Error(err)
	}
	if err := AddJobFunc("Job02", job02); err != nil {
		t.Error(err)
	}
	if err := AddJobFunc("Job04", job04); err != nil {
		t.Error(err)
	}
	if err := AddJobFunc("NilJob", nil); err != nil {
		t.Error(err)
	}
	if _, err := GetSchedule("NilSch"); err == nil {
		t.Error(fmt.Sprint(testErr, " GetSchedule must be return error"))
	}
	{
		var sch Schedule
		var err error
		name := "Sch01"
		if sch, err = GetSchedule(name); err != nil {
			t.Error(err)
		} else {
			var str string
			if str, err = sch.Name(); err != nil {
				t.Error(err)
			}
			assert.Equal(t, name, str, fmt.Sprintf("%v != %v", str, name))
			if str, err = sch.CronExpression(); err != nil {
				t.Error(err)
			}
			cron := "1/3,2-43/13 * * * * ? *"
			assert.Equal(t, cron, str, fmt.Sprintf("%v != %v", str, cron))
			if str, err = sch.Job(); err != nil {
				t.Error(err)
			}
			jobName := "Job01"
			assert.Equal(t, jobName, str, fmt.Sprintf("%v != %v", str, jobName))
			if str, err = sch.Desc(); err != nil {
				t.Error(err)
			}
			desc := "this is schedule 01-----------"
			assert.Equal(t, desc, str, fmt.Sprintf("%v != %v", str, desc))
			var testTime time.Time
			if testTime, err = sch.NextTime(); err != nil {
				t.Error(err)
			}
			assert.Equal(t, true, testTime.IsZero(), fmt.Sprintf("%v != %v", testTime.IsZero(), true))
			if testTime, err = sch.PrevTime(); err != nil {
				t.Error(err)
			}
			assert.Equal(t, true, testTime.IsZero(), fmt.Sprintf("%v != %v", testTime.IsZero(), true))
			if err = sch.Stop(); err != nil {
				t.Error(err)
			}
			if err = sch.Resume(); err != nil {
				t.Error(err)
			}
			var st Status
			if st, err = sch.Status(); err != nil {
				t.Error(err)
			}
			assert.Equal(t, Run, st, fmt.Sprintf("%v != %v", st, Run))
			var m01 map[string]interface{}
			if m01, err = sch.Data(); err != nil {
				t.Error(err)
			}
			m01["test"] = t
			m := make(map[string]interface{})
			m["event"] = "trigger"
			m["test"] = t
			if err = sch.Trigger(m); err != nil {
				t.Error(err)
			}
		}
	}
	m := make(map[string]interface{})
	m["event"] = "trigger"
	m["test"] = t
	if err := Trigger(jobFunc(job03), m); err != nil {
		t.Error(err)
	}
	m = make(map[string]interface{})
	m["event"] = "trigger"
	m["test"] = t
	if err := TriggerFunc(job02, m); err != nil {
		t.Error(err)
	}
	outStatus := GetStatus()
	assert.Equal(t, Stop, outStatus, fmt.Sprintf("%v != %v", outStatus, Stop))
	Interrupt()
	defer func() {
		Wait()
		WaitTrigger()
	}()
	Start()
	Interrupt()
	<-time.After(jtime.Second)
	Start()
	Start()
	if err := AddJob("Job03", jobFunc(job03)); err != nil {
		t.Error(err)
	}
	m = make(map[string]interface{})
	m["event"] = "schedule"
	m["test"] = t
	if err := AddSchedule(&SchInfo{Name: "Sch03", Cron: "*/2 * * * * ?", JobName: "Job03", JobData: m, Status: "Run"}); err != nil {
		t.Error(err)
	}
	if _, err := GetSchedule("Sch02"); err != nil {
		t.Error(err)
	}
	if err := Trigger(nil, nil); err == nil {
		t.Error(fmt.Sprint(testErr, " Trigger must be return error"))
	}
	if err := TriggerFunc(nil, nil); err == nil {
		t.Error(fmt.Sprint(testErr, " Trigger must be return error"))
	}
	m = make(map[string]interface{})
	m["event"] = "trigger"
	m["test"] = t
	if err := Trigger(jobFunc(job01), m); err != nil {
		t.Error(err)
	}
	m = make(map[string]interface{})
	m["event"] = "trigger"
	m["test"] = t
	if err := TriggerFunc(job02, m); err != nil {
		t.Error(err)
	}
	<-time.After(6 * jtime.Second)
}

func job01(m map[string]interface{}) {
	event := m["event"]
	fmt.Println("01: ", event, " - ", time.Now())
	t := m["test"].(*testing.T)
	var sch01 Schedule
	var err error
	if sch01, err = GetSchedule("Sch02"); err != nil {
		t.Error(err)
	} else {
		var j string
		if j, err = sch01.Job(); err != nil {
			t.Error(err)
		}
		if j == "Job02" && event.(string) == "schedule" {
			data := make(map[string]interface{})
			data["event"] = "schedule"
			if err = AddSchedule(&SchInfo{Name: "Sch02", Cron: "3,7,11/2,32-57/7 * * * * ? *", JobName: "Job04", JobData: data, Desc: "this is schedule 02-----------", Status: "Run"}); err != nil {
				t.Error(err)
			}
		}
	}
}

func job02(m map[string]interface{}) {
	fmt.Println("02: ", m["event"], " - ", time.Now())
}

func job03(m map[string]interface{}) {
	fmt.Println("03: ", m["event"], " - ", time.Now())
	t := m["test"].(*testing.T)
	testErr := "TEST ERROR:"
	var err error
	var sch Schedule
	if sch, err = GetSchedule("NilSch"); err == nil {
		t.Error(fmt.Sprint(testErr, " GetSchedule must be return error"))
	}
	name := "Sch02"
	if sch, err = GetSchedule(name); err != nil {
		t.Error(err)
	}
	if sch != nil {
		var inStr string
		var outStr string
		outStr = name
		if inStr, err = sch.Name(); err != nil {
			t.Error(err)
		}
		assert.Equal(t, outStr, inStr, fmt.Sprintf("%v != %v", inStr, outStr))
		outStr = "3,7,11/2,32-57/7 * * * * ? *"
		if inStr, err = sch.CronExpression(); err != nil {
			t.Error(err)
		}
		assert.Equal(t, outStr, inStr, fmt.Sprintf("%v != %v", inStr, outStr))
		outStr = "Job02 or Job04"
		if inStr, err = sch.Job(); err != nil {
			t.Error(err)
		}
		if inStr != "Job02" && inStr != "Job04" {
			t.Error(fmt.Sprintf("%v != %v", inStr, outStr))
		}
		jobName := inStr
		var sd map[string]interface{}
		data := make(map[string]interface{})
		data["event"] = "schedule"
		if sd, err = sch.Data(); err != nil {
			t.Error(err)
		}
		assert.Equal(t, data, sd, fmt.Sprintf("%v != %v", sd, data))
		outStr = "this is schedule 02-----------"
		if inStr, err = sch.Desc(); err != nil {
			t.Error(err)
		}
		assert.Equal(t, outStr, inStr, fmt.Sprintf("%v != %v", inStr, outStr))
		var st Status
		if st, err = sch.Status(); err != nil {
			t.Error(err)
		} else {
			if jobName == "Job04" {
				if st == Run {
					if _, err = sch.NextTime(); err != nil {
						t.Error(err)
					}
					if _, err = sch.PrevTime(); err != nil {
						t.Error(err)
					}
					if err = sch.Stop(); err != nil {
						t.Error(err)
					}
					tm := make(map[string]interface{})
					tm["event"] = "trigger"
					tm["test"] = t
					if err = sch.Trigger(tm); err != nil {
						t.Error(err)
					}
				} else {
					if err = sch.Resume(); err != nil {
						t.Error(err)
					}
				}
			}
		}
	}
	if sch, err = GetSchedule("Sch99"); err != nil {
		t.Error(err)
	} else {
		var st Status
		if st, err = sch.Status(); err != nil {
			t.Error(err)
		} else {
			if st == Stop {
				if err = AddJobFunc("Job99", job99); err != nil {
					t.Error(err)
				}
				if err = sch.Resume(); err != nil {
					t.Error(err)
				}
			} else {
				if err = AddJobFunc("Job99", nil); err != nil {
					t.Error(err)
				}
			}
		}
	}
}

func job04(m map[string]interface{}) {
	fmt.Println("04: ", m["event"], " - ", time.Now())
}

func job99(m map[string]interface{}) {
	fmt.Println("99: ", m["event"], " - ", time.Now())
}
