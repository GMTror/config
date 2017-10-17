package config

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	desName = "description"
)

func Description(i interface{}) (string, error) {
	d, err := description(reflect.ValueOf(i), "", "", "")
	if err != nil {
		return "", err
	}

	return d, nil
}

func description(result reflect.Value, tag, des, defaultVal string) (string, error) {
	switch result.Kind() {
	case reflect.Int:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Int8:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Int16:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Int32:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Int64:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Uint:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Uint8:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Uint16:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Uint32:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Uint64:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Float32:
		return descriptionPrint(tag, tag, defaultVal), nil
	case reflect.Float64:
		return descriptionPrint(tag, tag, defaultVal), nil
	case reflect.String:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Bool:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Ptr:
		return descriptionPtr(result, tag, des, defaultVal)
	case reflect.Struct:
		return descriptionStruct(result, tag, des, defaultVal)
	case reflect.Slice:
		return descriptionPrint(tag, des, defaultVal), nil
	case reflect.Map:
		return descriptionPrint(tag, des, defaultVal), nil
	default:
		return "", errors.New("type error")
	}
}

func descriptionPrint(tag, des, defaultVal string) (description string) {
	if tag == "" {
		return ""
	}
	description = tag
	if des != "" {
		description = fmt.Sprintf("%s - %s", tag, des)
	}
	if defaultVal == "" {
		return
	}
	if des == "" {
		description = fmt.Sprintf("%s -", tag)
	}
	return fmt.Sprintf("%s (defauilt: %s)", description, defaultVal)
}

func descriptionPtr(result reflect.Value, tag, des, defaultVal string) (string, error) {
	if des != "" || defaultVal != "" {
		return descriptionPrint(tag, des, defaultVal), nil
	}

	if _, ok := result.Interface().(Unmarshaler); ok {
		return descriptionPrint(tag, des, defaultVal), nil
	}

	return description(reflect.Indirect(result), tag, des, defaultVal)
}

func descriptionStruct(result reflect.Value, tag, des, defaultVal string) (string, error) {
	resultType := result.Type()
	for i := 0; i < resultType.NumField(); i++ {
		fieldType := resultType.Field(i)
		bTag := fieldType.Tag.Get(tagName)
		if bTag != pasName {
			sTag := getTag(tag, bTag)
			sVal := fieldType.Tag.Get(valName)
			sDes := fieldType.Tag.Get(desName)
			d, err := description(result.Field(i), sTag, sDes, sVal)
			if err != nil {
				return "", err
			}
			if d != "" {
				if des == "" {
					des = d
				} else {
					des = fmt.Sprintf("%s\n%s", des, d)
				}
			}
		}
	}

	return des, nil
}
