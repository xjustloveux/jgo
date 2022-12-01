// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"database/sql"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jtime"
	"reflect"
	"time"
)

type dataSource struct {
	Type                    string
	DSN                     string
	DN                      string
	DbName                  string
	ConnMaxLifetime         time.Duration
	ConnMaxLifetimeDuration string
	ConnMaxIdleTime         time.Duration
	ConnMaxIdleTimeDuration string
	MaxOpenConns            int
	MaxIdleConns            int
	EncodeData              string
	Format                  jfile.Format
	db                      *sql.DB
}

func (*dataSource) getDefault() *dataSource {
	return &dataSource{
		ConnMaxLifetime:         120,
		ConnMaxLifetimeDuration: "Second",
		ConnMaxIdleTime:         0,
		ConnMaxIdleTimeDuration: "Hour",
		MaxOpenConns:            0,
		MaxIdleConns:            0,
		EncodeData:              "",
		Format:                  jfile.Json,
	}
}

func (ds *dataSource) open() error {
	if ds.db != nil {
		return errorStr(errorDbAlreadyOpen)
	}
	var f reflect.Value
	nds := ds
	if decodeFunc != nil {
		f = reflect.ValueOf(decodeFunc)
		if f.Type().NumIn() == 1 && f.Type().NumOut() == 2 {
			if ds.EncodeData != "" {
				var err error
				var encodeStr string
				dsm := make(map[string]interface{})
				tds := (&dataSource{}).getDefault()
				in := make([]reflect.Value, 1)
				in[0] = reflect.ValueOf(ds.EncodeData)
				out := f.Call(in)
				if o := out[1]; !o.IsNil() {
					if o.CanInterface() {
						var ok bool
						if err, ok = o.Interface().(error); ok {
							return err
						} else {
							return errorStr(errorDecodeFuncOut2NotErrorType)
						}
					} else {
						return errorStr(errorDecodeFuncOut2NotErrorType)
					}
				}
				if o := out[0]; o.Type().Kind() == reflect.String {
					encodeStr = o.String()
				} else {
					return errorStr(errorDecodeFuncOut1NotStringType)
				}
				if err = jfile.Decode(ds.Format.String(), []byte(encodeStr), dsm); err != nil {
					return err
				}
				if err = jfile.Convert(dsm, tds); err != nil {
					return err
				}
				nds = tds
			}
		} else {
			return errorStr(errorDecodeFuncType)
		}
	}
	driverName := ""
	dataSourceName := ""
	if nds.DN != "" {
		driverName = nds.DN
	} else {
		if t, err := ParseDBType(nds.Type); err != nil {
			return err
		} else {
			driverName = t.DriverName()
		}
	}
	if decodeFunc != nil && ds.EncodeData == "" {
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(nds.DSN)
		out := f.Call(in)
		if o := out[1]; !o.IsNil() {
			if o.CanInterface() {
				if err, ok := o.Interface().(error); ok {
					return err
				} else {
					return errorStr(errorDecodeFuncOut2NotErrorType)
				}
			} else {
				return errorStr(errorDecodeFuncOut2NotErrorType)
			}
		}
		if o := out[0]; o.Type().Kind() == reflect.String {
			dataSourceName = o.String()
		} else {
			return errorStr(errorDecodeFuncOut1NotStringType)
		}
	} else {
		dataSourceName = nds.DSN
	}
	if db, err := sql.Open(driverName, dataSourceName); err != nil {
		return err
	} else {
		if d, timeErr := jtime.ParseTimeDuration(nds.ConnMaxLifetimeDuration); nds.ConnMaxLifetime > 0 && timeErr == nil {
			db.SetConnMaxLifetime(nds.ConnMaxLifetime * d)
		}
		if d, timeErr := jtime.ParseTimeDuration(nds.ConnMaxIdleTimeDuration); nds.ConnMaxIdleTime > 0 && timeErr == nil {
			db.SetConnMaxIdleTime(nds.ConnMaxIdleTime * d)
		}
		if nds.MaxOpenConns > 0 {
			db.SetMaxOpenConns(nds.MaxOpenConns)
		}
		if nds.MaxIdleConns > 0 {
			db.SetMaxIdleConns(nds.MaxIdleConns)
		}
		ds.db = db
	}
	return nil
}

func (ds *dataSource) close() error {
	if ds.db == nil {
		return errorStr(errorDbNotOpen)
	}
	if err := ds.db.Close(); err != nil {
		return err
	} else {
		ds.db = nil
	}
	return nil
}
