// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCronExpression(t *testing.T) {
	testErr := "TEST ERROR:"
	cron := &CronExpression{0, 0, 0, 0, 0, 0, 0}
	for i := 0; i <= maxYear-minYear; i++ {
		cron.year |= 1 << i
	}
	for i := 1; i <= 12; i++ {
		cron.month |= 1 << i
	}
	for i := 1; i <= 31; i++ {
		cron.day |= 1 << i
	}
	for i := 0; i <= 23; i++ {
		cron.hour |= 1 << i
	}
	for i := 0; i <= 59; i++ {
		cron.minute |= 1 << i
	}
	for i := 0; i <= 59; i++ {
		cron.second |= 1 << i
	}
	if _, err := ParseCronExpression(""); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression(", , , , , ,"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* * * * * * *"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("E * * * * * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("E/7 * * * * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("E-7/7 * * * * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("7-E/7 * * * * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("7/77/777 * * * * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("60/77 * * * * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* 77 * * * * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* 60/77 * * * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* * 7-77 * * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* * 24/77 * * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* * * / * ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* * * * 7/E ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* * * * 32/77 ?"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* * * ? * E"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* * * * * ? 55"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if _, err := ParseCronExpression("* * * * * ? 2022-2077/A"); err == nil {
		t.Error(fmt.Sprint(testErr, " ParseCronExpression must be return error"))
	}
	if v, err := ParseCronExpression("* * * * * ? *"); err != nil {
		t.Error(err)
	} else {
		assert.Equal(t, *cron, *v, fmt.Sprintf("%v != %v", *v, *cron))
	}
	if v, err := ParseCronExpression("* * * ? * * *"); err != nil {
		t.Error(err)
	} else {
		cron.day = 0
		for i := 0; i <= 6; i++ {
			cron.weekday |= 1 << i
		}
		assert.Equal(t, *cron, *v, fmt.Sprintf("%v != %v", *v, *cron))
	}
	if v, err := ParseCronExpression("*/3 * * * * ?"); err != nil {
		t.Error(err)
	} else {
		cron.weekday = 0
		for i := 1; i <= 31; i++ {
			cron.day |= 1 << i
		}
		cron.second = 0
		for i := 0; i <= 59; i += 3 {
			cron.second |= 1 << i
		}
		assert.Equal(t, *cron, *v, fmt.Sprintf("%v != %v", *v, *cron))
	}
	if v, err := ParseCronExpression("1-29/3,32,44/7 * * * * ?"); err != nil {
		t.Error(err)
	} else {
		cron.second = 0
		for i := 1; i <= 29; i += 3 {
			cron.second |= 1 << i
		}
		cron.second |= 1 << 32
		for i := 44; i <= 59; i += 7 {
			cron.second |= 1 << i
		}
		assert.Equal(t, *cron, *v, fmt.Sprintf("%v != %v", *v, *cron))
	}
	if v, err := ParseCronExpression("1-29/3,32,44/7 3-40,5/9 * * * ?"); err != nil {
		t.Error(err)
	} else {
		cron.minute = 0
		for i := 3; i <= 40; i++ {
			cron.minute |= 1 << i
		}
		for i := 5; i <= 59; i += 9 {
			cron.minute |= 1 << i
		}
		assert.Equal(t, *cron, *v, fmt.Sprintf("%v != %v", *v, *cron))
	}
	if v, err := ParseCronExpression("1-29/3,32,44/7 3-40,5/9 * * * ? 2022-2077/7,2055/3"); err != nil {
		t.Error(err)
	} else {
		cron.year = 0
		for i := 2022 - minYear; i <= 2077-minYear; i += 7 {
			cron.year |= 1 << i
		}
		for i := 2055 - minYear; i <= maxYear-minYear; i += 3 {
			cron.year |= 1 << i
		}
		assert.Equal(t, *cron, *v, fmt.Sprintf("%v != %v", *v, *cron))
	}
	if v, err := ParseCronExpression("1-29/3,32,44/7 3-40,5/9 * * January ? 2022-2077/7,2055/3"); err != nil {
		t.Error(err)
	} else {
		cron.month = 1 << 1
		assert.Equal(t, *cron, *v, fmt.Sprintf("%v != %v", *v, *cron))
	}
	if v, err := ParseCronExpression("1-29/3,32,44/7 3-40,5/9 * * January-February,march/3,april,may,june,july,august,september,october,november,december ? 2022-2077/7,2055/3"); err != nil {
		t.Error(err)
	} else {
		cron.month |= 1 << 2
		for i := 3; i <= 12; i += 3 {
			cron.month |= 1 << i
		}
		cron.month |= 1 << 4
		cron.month |= 1 << 5
		cron.month |= 1 << 6
		cron.month |= 1 << 7
		cron.month |= 1 << 8
		cron.month |= 1 << 9
		cron.month |= 1 << 10
		cron.month |= 1 << 11
		cron.month |= 1 << 12
		assert.Equal(t, *cron, *v, fmt.Sprintf("%v != %v", *v, *cron))
	}
	if v, err := ParseCronExpression("1-29/3,32,44/7 3-40,5/9 * ? January-February,march/3,april,may,june,july,august,september,october,november,december sunday,monday,tuesday,wednesday,thursday,friday,saturday 2022-2077/7,2055/3"); err != nil {
		t.Error(err)
	} else {
		cron.day = 0
		cron.weekday |= 1 << 0
		cron.weekday |= 1 << 1
		cron.weekday |= 1 << 2
		cron.weekday |= 1 << 3
		cron.weekday |= 1 << 4
		cron.weekday |= 1 << 5
		cron.weekday |= 1 << 6
		assert.Equal(t, *cron, *v, fmt.Sprintf("%v != %v", *v, *cron))
	}
}
