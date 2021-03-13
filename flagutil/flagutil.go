package flagutil

import (
	"os"
	"reflect"
	"errors"
	"strconv"
)

const TAG_NAME = "flag"
const TAG_PREFIX = "--"

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func Parse(obj interface{}) error {
	t := reflect.TypeOf(obj)
	//fmt.Println(t.Kind())
	if t.Kind() != reflect.Ptr {
		return errors.New("parse error: object is not ptr")
	}
	el := reflect.ValueOf(obj).Elem()
	if el.NumField() < 1 {
		return errors.New("parse error: field empty")
	}

	args := os.Args[1:]
	elT := el.Type()
	for i := 0; i < el.NumField(); i++ {
		fT := elT.Field(i)
		f := el.Field(i)
		tag := fT.Tag.Get(TAG_NAME)
		//fmt.Println("tag:", tag)
		k, found := Find(args, TAG_PREFIX + tag)
		if found == true {
			//fmt.Println(k)
			switch f.Kind() {
			case reflect.Bool:
				f.SetBool(true)
			case reflect.Int:
				v, err := strconv.Atoi(args[k+1])
				if err != nil {
					return errors.New("parse error: convert int64 fail")
				}
				f.SetInt(int64(v))
			case reflect.String:
				f.SetString(args[k+1])
			default:
				return errors.New("parse error: switched default")
			}
		}
	}
	return nil
}