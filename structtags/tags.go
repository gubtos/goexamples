package structtags

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ClientConfig struct {
	URL                 string        `key:"url" default:""`
	Timeout             time.Duration `key:"timeout" default:"2s"`
	Retry               int           `key:"retry" default:"0"`
	MaxIdleConnsPerHost int           `key:"maxIdleConsPerHost" default:"20"`
	ExampleFloat        float64       `key:"exampleFloat" default:"20.2"`
}

var errUnimplementedType = errors.New("unknown type, handling type not implemented")

func DefaultClientConfig() ClientConfig {
	return FromDefaultTags[ClientConfig]()
}

func FromDefaultTags[T any]() T {
	var defaultClientConfig T

	cfgPtr := reflect.ValueOf(&defaultClientConfig)
	cfgElem := cfgPtr.Elem()
	cfgElemType := cfgElem.Type()

	for i := 0; i < cfgElem.NumField(); i++ {
		f := cfgElemType.Field(i)
		if defaultValue, ok := f.Tag.Lookup("default"); ok {
			elemType := cfgElem.Field(i).Type().String()
			field := cfgElem.Field(i)
			switch elemType {
			case "string":
				field.SetString(defaultValue)
			case "int32", "int", "int64":
				v, err := strconv.ParseInt(defaultValue, 10, 0)
				if err != nil {
					panic(err)
				}
				field.SetInt(v)
			case "time.Duration":
				v, err := time.ParseDuration(defaultValue)
				if err != nil {
					panic(err)
				}
				field.SetInt(int64(v.Abs()))
			case "float32", "float64":
				v, err := strconv.ParseFloat(defaultValue, 64)
				if err != nil {
					panic(err)
				}
				field.SetFloat(v)
			default:
				panic(errUnimplementedType)
			}
		}
	}

	return defaultClientConfig
}

func FromEnv[T any]() T {
	defaultClientConfig := FromDefaultTags[T]()

	cfgPtr := reflect.ValueOf(&defaultClientConfig)
	cfgElem := cfgPtr.Elem()
	cfgElemType := cfgElem.Type()

	for i := 0; i < cfgElem.NumField(); i++ {
		f := cfgElemType.Field(i)
		if keyValue, ok := f.Tag.Lookup("key"); ok {
			envValue := os.Getenv(strings.ToUpper(keyValue))
			if envValue == "" {
				continue
			}

			elemType := cfgElem.Field(i).Type().String()
			field := cfgElem.Field(i)
			switch elemType {
			case "string":
				field.SetString(envValue)
			case "int32", "int", "int64":
				v, err := strconv.ParseInt(envValue, 10, 0)
				if err != nil {
					panic(err)
				}
				field.SetInt(v)
			case "time.Duration":
				v, err := time.ParseDuration(envValue)
				if err != nil {
					panic(err)
				}
				field.SetInt(int64(v.Abs()))
			case "float32", "float64":
				v, err := strconv.ParseFloat(envValue, 64)
				if err != nil {
					panic(err)
				}
				field.SetFloat(v)
			default:
				panic(errUnimplementedType)
			}
		}
	}

	return defaultClientConfig
}
