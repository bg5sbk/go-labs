package labs19

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func Verify(data interface{}) error {
	typeOfData := reflect.TypeOf(data)
	valueOfData := reflect.ValueOf(data)

	if typeOfData.Kind() == reflect.Ptr {
		typeOfData = typeOfData.Elem()
		valueOfData = valueOfData.Elem()
	}

	numField := typeOfData.NumField()

	for i := 0; i < numField; i++ {
		fieldInfo := typeOfData.Field(i)

		if fieldInfo.Type.Kind() == reflect.String {
			maxLength := fieldInfo.Tag.Get("max-length")

			if maxLength != "" {
				maxLength2, _ := strconv.Atoi(maxLength)
				if valueOfData.Field(i).Len() > maxLength2 {
					return errors.New(fmt.Sprintf("len(%s.%s) > %s", typeOfData.Name(), fieldInfo.Name, maxLength))
				}
			}
		}
	}

	return nil
}

var cache = make(map[reflect.Type][]func(reflect.Value) error)

func FastVerify(data interface{}) error {
	typeOfData := reflect.TypeOf(data)
	valueOfData := reflect.ValueOf(data)

	if typeOfData.Kind() == reflect.Ptr {
		typeOfData = typeOfData.Elem()
		valueOfData = valueOfData.Elem()
	}

	verifyCallbacks, cached := cache[typeOfData]

	if !cached {
		numField := typeOfData.NumField()

		for i := 0; i < numField; i++ {
			fieldInfo := typeOfData.Field(i)

			if fieldInfo.Type.Kind() == reflect.String {
				maxLength := fieldInfo.Tag.Get("max-length")

				if maxLength != "" {
					var (
						maxLength2, _ = strconv.Atoi(maxLength)
						errorMessage  = fmt.Sprintf("len(%s.%s) > %s", typeOfData.Name(), fieldInfo.Name, maxLength)
						fieldIndex    = i
					)
					verifyCallbacks = append(verifyCallbacks, func(v reflect.Value) error {
						if v.Field(fieldIndex).Len() > maxLength2 {
							return errors.New(errorMessage)
						}
						return nil
					})
				}
			}
		}

		cache[typeOfData] = verifyCallbacks
	}

	for _, callback := range verifyCallbacks {
		if err := callback(valueOfData); err != nil {
			return err
		}
	}

	return nil
}

var cache2 = make(map[reflect.Type][]func(uintptr) error)

func VeryFastVerify(data interface{}) error {
	typeOfData := reflect.TypeOf(data)
	valueOfData := reflect.ValueOf(data)

	if typeOfData.Kind() == reflect.Ptr {
		typeOfData = typeOfData.Elem()
		valueOfData = valueOfData.Elem()
	}

	verifyCallbacks, cached := cache2[typeOfData]

	if !cached {
		beginAddr := valueOfData.UnsafeAddr()
		numField := typeOfData.NumField()

		for i := 0; i < numField; i++ {
			fieldInfo := typeOfData.Field(i)

			if fieldInfo.Type.Kind() == reflect.String {
				maxLength := fieldInfo.Tag.Get("max-length")

				if maxLength != "" {
					var (
						maxLength2, _ = strconv.Atoi(maxLength)
						errorMessage  = fmt.Sprintf("len(%s.%s) > %s", typeOfData.Name(), fieldInfo.Name, maxLength)
						fieldOffset   = valueOfData.Field(i).UnsafeAddr() - beginAddr
					)
					verifyCallbacks = append(verifyCallbacks, func(ptr uintptr) error {
						if (*reflect.StringHeader)(unsafe.Pointer(ptr+fieldOffset)).Len > maxLength2 {
							return errors.New(errorMessage)
						}
						return nil
					})
				}
			}
		}

		cache2[typeOfData] = verifyCallbacks
	}

	addr := valueOfData.UnsafeAddr()

	for _, callback := range verifyCallbacks {
		if err := callback(addr); err != nil {
			return err
		}
	}

	return nil
}
