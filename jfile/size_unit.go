// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jfile

// SizeUnit size unit
type SizeUnit int64

const (
	Byte    SizeUnit = 1
	Kb               = 128 * Byte
	KB               = 8 * Kb
	Mb               = 128 * KB
	MB               = 8 * Mb
	Gb               = 128 * MB
	GB               = 8 * Gb
	Tb               = 128 * GB
	TB               = 8 * Tb
	Pb               = 128 * TB
	PB               = 8 * Pb
	Eb               = 128 * PB
	EB               = 8 * Eb
	Unknown          = -1
)

// ToInt64 SizeUnit type to int64
func (u SizeUnit) ToInt64() int64 {
	return int64(u)
}

// ParseSizeUnit takes a string SideMode and returns the SizeUnit constant.
func ParseSizeUnit(u string) (SizeUnit, error) {
	switch u {
	case "Byte":
		return Byte, nil
	case "Kb":
		return Kb, nil
	case "KB":
		return KB, nil
	case "Mb":
		return Mb, nil
	case "MB":
		return MB, nil
	case "Gb":
		return Gb, nil
	case "GB":
		return GB, nil
	case "Tb":
		return Tb, nil
	case "TB":
		return TB, nil
	case "Pb":
		return Pb, nil
	case "PB":
		return PB, nil
	case "Eb":
		return Eb, nil
	case "EB":
		return EB, nil
	}
	return Unknown, errorFmt(errorNotValidSizeUnit, u)
}
