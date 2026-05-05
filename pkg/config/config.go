package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var cfg any

func Load[T any]() (*T, error) {
	if cfg != nil {
		return nil, errors.New("config already loaded")
	}

	var _cfg T
	cfg = &_cfg
	loadEnv(".env")

	var fields = reflect.TypeOf(_cfg)
	v := reflect.ValueOf(&_cfg).Elem()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		tag := field.Tag.Get("env")

		if tag == "" {
			continue
		}

		val := os.Getenv(tag)
		if val == "" {
			continue
		}

		switch field.Type.Kind() {
		case reflect.String:
			v.Field(i).SetString(val)

		case reflect.Int, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(val, 10, 64)

			if err != nil {
				return nil, errors.New("invalid int for %s", tag)
			}

			v.Field(i).SetInt(n)

		case reflect.Bool:
			b, err := strconv.ParseBool(val)
			if err != nil {
				return nil, fmt.Errorf("invalid bool for %s", tag)
			}

			v.Field(i).SetBool(b)
		}
	}

	cfg = &_cfg
	return cfg.(*T), nil
}
