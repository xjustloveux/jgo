// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jconf

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jfile"
	"os"
	"reflect"
	"strings"
)

type Config interface {
	// Format returns format
	Format() jfile.Format
	// SetFormat set format
	SetFormat(jfile.Format)
	// FileName returns file name
	FileName() string
	// SetFileName set file name
	SetFileName(string)
	// Root returns root path
	Root() string
	// SetRoot set root path
	SetRoot(string)
	// EnvFileName returns env file name
	EnvFileName() string
	// SetEnvFileName set env file name
	SetEnvFileName(string)
	// EnvKey returns env key
	EnvKey() string
	// SetEnvKey set env key
	SetEnvKey(string)
	// EnvVal returns env value
	EnvVal() string
	// SetEnvVal set env value
	SetEnvVal(string)
	// EnableEnv enable env
	EnableEnv()
	// DisableEnv disable env
	DisableEnv()
	// Load read config file
	// have env file name, then after load file will load env file, and overwrite the value
	Load() error
	// Get returns interface
	Get(...interface{}) (interface{}, error)
	// String returns string
	String(...interface{}) (string, error)
	// Bool returns bool
	Bool(...interface{}) (bool, error)
	// Int returns int
	Int(...interface{}) (int, error)
	// Int8 returns int8
	Int8(...interface{}) (int8, error)
	// Int16 returns int16
	Int16(...interface{}) (int16, error)
	// Int32 returns int32
	Int32(...interface{}) (int32, error)
	// Int64 returns int64
	Int64(...interface{}) (int64, error)
	// Uint returns uint
	Uint(...interface{}) (uint, error)
	// Uint8 returns uint8
	Uint8(...interface{}) (uint8, error)
	// Uint16 returns uint16
	Uint16(...interface{}) (uint16, error)
	// Uint32 returns uint32
	Uint32(...interface{}) (uint32, error)
	// Uint64 returns uint64
	Uint64(...interface{}) (uint64, error)
	// Float32 returns float32
	Float32(...interface{}) (float32, error)
	// Float64 returns float64
	Float64(...interface{}) (float64, error)
	// InterfaceMapInterface returns map[interface{}]interface{}
	InterfaceMapInterface(...interface{}) (map[interface{}]interface{}, error)
	// StringMapInterface returns map[string]interface{}
	StringMapInterface(...interface{}) (map[string]interface{}, error)
	// SliceInterface returns []interface{}
	SliceInterface(...interface{}) ([]interface{}, error)
	// SliceString returns []string
	SliceString(...interface{}) ([]string, error)
	// SliceBool returns []bool
	SliceBool(...interface{}) ([]bool, error)
	// SliceInt returns []int
	SliceInt(...interface{}) ([]int, error)
	// SliceInt8 returns []int8
	SliceInt8(...interface{}) ([]int8, error)
	// SliceInt16 returns []int16
	SliceInt16(...interface{}) ([]int16, error)
	// SliceInt32 returns []int32
	SliceInt32(...interface{}) ([]int32, error)
	// SliceInt64 returns []int64
	SliceInt64(...interface{}) ([]int64, error)
	// SliceUint returns []uint
	SliceUint(...interface{}) ([]uint, error)
	// SliceUint8 returns []uint8
	SliceUint8(...interface{}) ([]uint8, error)
	// SliceUint16 returns []uint16
	SliceUint16(...interface{}) ([]uint16, error)
	// SliceUint32 returns []uint32
	SliceUint32(...interface{}) ([]uint32, error)
	// SliceUint64 returns []uint64
	SliceUint64(...interface{}) ([]uint64, error)
	// SliceFloat32 returns []float32
	SliceFloat32(...interface{}) ([]float32, error)
	// SliceFloat64 returns []float64
	SliceFloat64(...interface{}) ([]float64, error)
	// Convert args can be map[string]interface{} or config args
	Convert(interface{}, ...interface{}) error
}

type config struct {
	format      jfile.Format
	fileName    string
	root        string
	envFileName string
	data        map[string]interface{}
	env         bool
	envKey      string
	envVal      string
}

func (c *config) Format() jfile.Format {
	return c.format
}

func (c *config) SetFormat(f jfile.Format) {
	c.format = f
}

func (c *config) FileName() string {
	return c.fileName
}

func (c *config) SetFileName(name string) {
	c.fileName = name
}

func (c *config) Root() string {
	return c.fileName
}

func (c *config) SetRoot(root string) {
	c.root = root
}

func (c *config) EnvFileName() string {
	return c.envFileName
}

func (c *config) SetEnvFileName(name string) {
	c.envFileName = name
}

func (c *config) EnvKey() string {
	return c.envKey
}

func (c *config) SetEnvKey(key string) {
	c.envKey = key
}

func (c *config) EnvVal() string {
	return c.envVal
}

func (c *config) SetEnvVal(val string) {
	c.envVal = val
}

func (c *config) EnableEnv() {
	c.env = true
}

func (c *config) DisableEnv() {
	c.env = false
}

func (c *config) Load() (err error) {
	if c.fileName == "" {
		return errors(errorFileNameEmpty)
	}
	var b []byte
	if b, err = jfile.Load(fmt.Sprint(c.root, c.fileName)); err != nil {
		return err
	}
	if err = jfile.Decode(c.format.String(), b, c.data); err != nil {
		return err
	}
	if c.env {
		if c.envFileName != "" {
			if b, err = jfile.Load(fmt.Sprint(c.root, c.envFileName)); err != nil {
				return err
			}
		} else {
			var val string
			if c.envVal != "" {
				val = c.envVal
			} else {
				var key string
				if c.envKey != "" {
					key = c.envKey
				} else {
					key = envKey
					c.envKey = key
				}
				if val = jcast.String(c.data[key]); val == "" {
					val = os.Getenv(key)
				}
				if val == "" {
					return nil
				}
				c.envVal = val
			}
			arr := strings.Split(c.fileName, ".")
			l := len(arr) - 1
			name := arr[:l]
			ext := arr[l:]
			path := ""
			for i, s := range name {
				if i > 0 {
					path = fmt.Sprint(path, ".")
				}
				path = fmt.Sprint(path, s)
			}
			path = fmt.Sprint(path, dash, val, ".")
			for _, s := range ext {
				path = fmt.Sprint(path, s)
			}
			if b, err = jfile.Load(fmt.Sprint(c.root, path)); err != nil {
				return err
			}
		}
		env := make(map[string]interface{})
		if err = jfile.Decode(c.format.String(), b, env); err != nil {
			return err
		}
		var m1, m2 map[interface{}]interface{}
		if m1, err = jcast.InterfaceMapInterface(c.data); err != nil {
			return err
		}
		if m2, err = jcast.InterfaceMapInterface(env); err != nil {
			return err
		}
		if c.data, err = jcast.StringMapInterface(overwriteMap(m1, m2)); err != nil {
			return err
		}
	}
	return nil
}

func (c *config) Get(args ...interface{}) (interface{}, error) {
	return c.get(c.data, args...)
}

func (c *config) String(args ...interface{}) (string, error) {
	if v, err := c.Get(args...); err != nil {
		return "", err
	} else {
		return jcast.String(v), nil
	}
}

func (c *config) Bool(args ...interface{}) (bool, error) {
	if v, err := c.Get(args...); err != nil {
		return false, err
	} else {
		return jcast.Bool(v)
	}
}

func (c *config) Int(args ...interface{}) (int, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Int(v)
	}
}

func (c *config) Int8(args ...interface{}) (int8, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Int8(v)
	}
}

func (c *config) Int16(args ...interface{}) (int16, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Int16(v)
	}
}

func (c *config) Int32(args ...interface{}) (int32, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Int32(v)
	}
}

func (c *config) Int64(args ...interface{}) (int64, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Int64(v)
	}
}

func (c *config) Uint(args ...interface{}) (uint, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Uint(v)
	}
}

func (c *config) Uint8(args ...interface{}) (uint8, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Uint8(v)
	}
}

func (c *config) Uint16(args ...interface{}) (uint16, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Uint16(v)
	}
}

func (c *config) Uint32(args ...interface{}) (uint32, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Uint32(v)
	}
}

func (c *config) Uint64(args ...interface{}) (uint64, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Uint64(v)
	}
}

func (c *config) Float32(args ...interface{}) (float32, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Float32(v)
	}
}

func (c *config) Float64(args ...interface{}) (float64, error) {
	if v, err := c.Get(args...); err != nil {
		return 0, err
	} else {
		return jcast.Float64(v)
	}
}

func (c *config) InterfaceMapInterface(args ...interface{}) (map[interface{}]interface{}, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.InterfaceMapInterface(v)
	}
}

func (c *config) StringMapInterface(args ...interface{}) (map[string]interface{}, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.StringMapInterface(v)
	}
}

func (c *config) SliceInterface(args ...interface{}) ([]interface{}, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceInterface(v)
	}
}

func (c *config) SliceString(args ...interface{}) ([]string, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceString(v)
	}
}

func (c *config) SliceBool(args ...interface{}) ([]bool, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceBool(v)
	}
}

func (c *config) SliceInt(args ...interface{}) ([]int, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceInt(v)
	}
}

func (c *config) SliceInt8(args ...interface{}) ([]int8, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceInt8(v)
	}
}

func (c *config) SliceInt16(args ...interface{}) ([]int16, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceInt16(v)
	}
}

func (c *config) SliceInt32(args ...interface{}) ([]int32, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceInt32(v)
	}
}

func (c *config) SliceInt64(args ...interface{}) ([]int64, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceInt64(v)
	}
}

func (c *config) SliceUint(args ...interface{}) ([]uint, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceUint(v)
	}
}

func (c *config) SliceUint8(args ...interface{}) ([]uint8, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceUint8(v)
	}
}

func (c *config) SliceUint16(args ...interface{}) ([]uint16, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceUint16(v)
	}
}

func (c *config) SliceUint32(args ...interface{}) ([]uint32, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceUint32(v)
	}
}

func (c *config) SliceUint64(args ...interface{}) ([]uint64, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceUint64(v)
	}
}

func (c *config) SliceFloat32(args ...interface{}) ([]float32, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceFloat32(v)
	}
}

func (c *config) SliceFloat64(args ...interface{}) ([]float64, error) {
	if v, err := c.Get(args...); err != nil {
		return nil, err
	} else {
		return jcast.SliceFloat64(v)
	}
}

func (c *config) Convert(v interface{}, args ...interface{}) error {
	var err error
	var m map[string]interface{}
	if len(args) == 1 && args[0] != nil && reflect.TypeOf(args[0]).Kind() == reflect.Map {
		if m, err = jcast.StringMapInterface(args[0]); err != nil {
			return err
		}
	} else {
		if m, err = c.StringMapInterface(args...); err != nil {
			return err
		}
	}
	return jfile.Convert(m, v)
}

func (c *config) get(data interface{}, args ...interface{}) (interface{}, error) {
	if len(args) > 0 {
		a := args[0]
		if a == nil {
			return nil, errors(errorArgNil)
		}
		if data == nil {
			return nil, errorf(errorValNilOfArg, a)
		}
		na := args[1:]
		switch reflect.TypeOf(data).Kind() {
		case reflect.Map:
			if m, err := jcast.InterfaceMapInterface(data); err != nil {
				return nil, err
			} else {
				v := m[a]
				if len(na) > 0 {
					return c.get(v, na...)
				} else {
					return v, nil
				}
			}
		case reflect.Slice:
			var err error
			var s []interface{}
			if s, err = jcast.SliceInterface(data); err != nil {
				return nil, err
			}
			var idx int
			if idx, err = jcast.Int(a); err != nil {
				return nil, err
			}
			if l := len(s); idx < l {
				v := s[idx]
				if len(na) > 0 {
					return c.get(v, na...)
				} else {
					return v, nil
				}
			} else {
				return nil, errorf(errorIndexOut, idx, l)
			}
		default:
			return nil, errorf(errorValNotFoundOfArg, a)
		}
	} else {
		return data, nil
	}
}
