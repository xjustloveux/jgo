// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

type configData struct {
	Cron *configPack
}

type configPack struct {
	Schedule []*SchInfo
}
