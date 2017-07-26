package config

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	tagName = "env"
	valName = "default"
)

type Unmarshaler interface {
	UnmarshalConfig([]byte) error
}

func getTag(t1, t2 string) string {
	if t1 != "" {
		if t2 != "" {
			return fmt.Sprintf("%s_%s", t1, t2)
		} else {
			return t1
		}
	} else {
		return t2
	}
}

func ReadENV(i interface{}) error {
	if err := decode(reflect.ValueOf(i), "", ""); err != nil {
		return err
	}

	return nil
}

func decode(result reflect.Value, tag, defaultVal string) error {
	log.Print(result.Kind())

	switch result.Kind() {
	case reflect.Int:
		return decodeInt(result, tag, defaultVal)
	case reflect.Int8:
		return decodeInt8(result, tag, defaultVal)
	case reflect.Int16:
		return decodeInt16(result, tag, defaultVal)
	case reflect.Int32:
		return decodeInt32(result, tag, defaultVal)
	case reflect.Int64:
		return decodeInt64(result, tag, defaultVal)
	case reflect.Uint:
		return decodeUint(result, tag, defaultVal)
	case reflect.Uint8:
		return decodeUint8(result, tag, defaultVal)
	case reflect.Uint16:
		return decodeUint16(result, tag, defaultVal)
	case reflect.Uint32:
		return decodeUint32(result, tag, defaultVal)
	case reflect.Uint64:
		return decodeUint64(result, tag, defaultVal)
	case reflect.Float32:
		return decodeFloat32(result, tag, defaultVal)
	case reflect.Float64:
		return decodeFloat64(result, tag, defaultVal)
	case reflect.String:
		return decodeString(result, tag, defaultVal)
	case reflect.Bool:
		return decodeBool(result, tag, defaultVal)
	// case reflect.Complex64:
	// 	return decodeComplex64(result, tag, defaultVal)
	// case reflect.Complex128:
	// 	return decodeComplex128(result, tag, defaultVal)
	// case reflect.Interface:
	// 	return decodeInterface(result, tag, defaultVal)
	case reflect.Ptr:
		return decodePtr(result, tag, defaultVal)
	case reflect.Struct:
		return decodeStruct(result, tag, defaultVal)
	case reflect.Slice:
		return decodeSlice(result, tag, defaultVal)
	case reflect.Map:
		return decodeMap(result, tag, defaultVal)
	default:
		return errors.New("type error")
	}

	return nil
}

func decodeInt(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.Atoi(tVal)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeInt8(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseInt(tVal, 10, 8)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeInt16(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseInt(tVal, 10, 16)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeInt32(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseInt(tVal, 10, 32)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeInt64(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseInt(tVal, 10, 64)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeUint(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseUint(tVal, 10, 32)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeUint8(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseUint(tVal, 10, 8)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeUint16(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseUint(tVal, 10, 16)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeUint32(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseUint(tVal, 10, 32)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeUint64(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseUint(tVal, 10, 64)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeFloat32(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseFloat(tVal, 32)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeFloat64(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(0).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseFloat(tVal, 64)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

// func decodeComplex64(result reflect.Value, tag, defaultVal string) error {}
// func decodeComplex128(result reflect.Value, tag, defaultVal string) error {}

func decodeString(result reflect.Value, tag, defaultVal string) error {
	val := os.Getenv(tag)
	if val == "" {
		val = defaultVal
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

func decodeBool(result reflect.Value, tag, defaultVal string) error {
	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		} else {
			result.Set(reflect.ValueOf(false).Convert(result.Type()))
			return nil
		}
	}

	val, err := strconv.ParseBool(tVal)
	if err != nil {
		return err
	}

	result.Set(reflect.ValueOf(val).Convert(result.Type()))

	return nil
}

// func decodeInterface(result reflect.Value, tag, defaultVal string) error {}

func decodePtr(result reflect.Value, tag, defaultVal string) error {

	if result.IsNil() {
		resultType := result.Type()
		resultElemType := resultType.Elem()
		resultNewType := reflect.New(resultElemType)

		if u, ok := resultNewType.Interface().(Unmarshaler); ok {
			val := os.Getenv(tag)
			if val == "" {
				val = defaultVal
			}
			log.Print(tag)

			if err := u.UnmarshalConfig(bytes.NewBufferString(val).Bytes()); err != nil {
				return err
			}
		} else {
			if err := decode(reflect.Indirect(resultNewType), tag, defaultVal); err != nil {
				return err
			}
		}

		result.Set(resultNewType)
	} else {
		if u, ok := result.Interface().(Unmarshaler); ok {
			val := os.Getenv(tag)
			if val == "" {
				val = defaultVal
			}

			if err := u.UnmarshalConfig(bytes.NewBufferString(val).Bytes()); err != nil {
				return err
			}
		} else {
			if err := decode(reflect.Indirect(result), tag, defaultVal); err != nil {
				return err
			}
		}
	}

	return nil
}

func decodeStruct(result reflect.Value, tag, defaultVal string) error {
	resultType := result.Type()
	for i := 0; i < resultType.NumField(); i++ {
		fieldType := resultType.Field(i)
		sTag := getTag(tag, fieldType.Tag.Get(tagName))
		sVal := fieldType.Tag.Get(valName)
		if err := decode(result.Field(i), sTag, sVal); err != nil {
			return err
		}
	}

	return nil
}

func decodeSlice(result reflect.Value, tag, defaultVal string) error {
	resultType := result.Type()
	resultElemType := resultType.Elem()
	resultSliceType := reflect.SliceOf(resultElemType)

	rs := reflect.MakeSlice(resultSliceType, 0, 0)

	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		}
	}

	for _, val := range strings.Split(tVal, ",") {
		r := reflect.Indirect(reflect.New(resultElemType))
		decode(r, "", val)
		rs = reflect.Append(rs, r)
	}

	result.Set(rs)

	return nil
}

func decodeMap(result reflect.Value, tag, defaultVal string) error {
	resultType := result.Type()
	resultElemType := resultType.Elem()
	resultKeyType := resultType.Key()

	rm := reflect.MakeMap(reflect.MapOf(resultKeyType, resultElemType))

	tVal := os.Getenv(tag)

	if tVal == "" {
		if defaultVal != "" {
			tVal = defaultVal
		}
	}

	for _, kv := range strings.Split(tVal, ",") {
		vs := strings.SplitN(kv, ":", 2)
		key := reflect.ValueOf(vs[0])
		val := reflect.Indirect(reflect.New(resultElemType))
		decode(val, "", vs[1])
		rm.SetMapIndex(key, val)
	}

	result.Set(rm)

	return nil
}
