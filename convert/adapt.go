package convert

import (
	"errors"
	"reflect"
	"time"
)

const timeType string = "time.Time"

// SourceToDestination get fields from "src" and set there values
// to fields of "dest" where the names and types of fields between
// two structures are equals.
//
//  src := struct{
//  	ID string
//      Name string
//		Age	int64
//  }{
//		ID: "123",
//		Name: "test_convert",
//		Age: 34,
//  }

//  dest := struct{
// 		ID string
//		Age int64
//      Address string
// 	}{}
//
// SourceToDestination(src, &dest)
//
// Here dest will be {ID: "123", Age: 34, Address: ""}
// The "dest" must be pointer otherwise function will fail.
// Fields must be exported otherwise function will fail. Function also
// also support convert from int64 to time.Time and vice versa.
func SourceToDestination(src interface{}, dest interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case *reflect.ValueError:
				err = r.(*reflect.ValueError)
			case string:
				err = errors.New(r.(string))
			}
		}
	}()

	srcValue := reflect.ValueOf(src)
	destValue := reflect.ValueOf(dest).Elem()

	if srcValue.Kind() != reflect.Slice {
		convertSourceValueToDestinationValue(srcValue, destValue)
		return nil
	}

	numberOfElements := srcValue.Len()
	for i := 0; i < numberOfElements; i++ {
		srcVal := srcValue.Index(i)
		typ := destValue.Type().Elem()
		destVal := reflect.New(typ).Elem()
		convertSourceValueToDestinationValue(srcVal, destVal)
		dd := reflect.Append(destValue, destVal)
		destValue.Set(dd)

	}

	return nil
}

func convertSourceValueToDestinationValue(srcValue reflect.Value, destValue reflect.Value) {
	srcType := srcValue.Type()
	destType := destValue.Type()
	numberOfFields := srcType.NumField()

	for i := 0; i < numberOfFields; i++ {
		srcFieldType := srcType.Field(i)
		srcFieldValue := srcValue.Field(i)
		destFieldValue := destValue.FieldByName(srcFieldType.Name)
		destField, ok := destType.FieldByName(srcFieldType.Name)
		if !ok {
			continue
		}

		srcFieldKind := srcFieldType.Type.Kind()
		fieldValue := destValue.FieldByName(srcFieldType.Name)

		// When the names of the fields are equals, but kinds are different, than
		// function suppose that fields are of type time.Time and int64.
		if srcFieldValue.Type().String() == timeType || destFieldValue.Type().String() == timeType {
			convertTime(srcFieldValue, fieldValue)
			continue
		}

		// if the two fields are of kind "struct",
		// we must invoke recursive method convertSourceValueToDestinationValue
		if srcFieldKind == reflect.Struct && destField.Type.Kind() == reflect.Struct {
			convertSourceValueToDestinationValue(srcFieldValue, fieldValue)
			continue
		}

		// when the names of the fields are equal, kinds are equal and fields
		// are slice, function will foreach over each element of array
		if srcFieldKind == reflect.Slice {
			if srcFieldType.Type.Elem().Kind() == reflect.Struct {
				for k := 0; k < srcFieldValue.Len(); k++ {
					srcVal := srcFieldValue.Index(k)
					typ := fieldValue.Type().Elem()
					destVal := reflect.New(typ).Elem()
					convertSourceValueToDestinationValue(srcVal, destVal)
					dd := reflect.Append(fieldValue, destVal)
					fieldValue.Set(dd)
				}
				continue
			}
		}

		fieldValue.Set(srcValue.Field(i))
	}
}

func convertTime(t1 reflect.Value, t2 reflect.Value) {
	seconds := t1.Interface()
	if t, ok := t1.Interface().(time.Time); ok {
		seconds = t.Unix()
	}

	if t, ok := t2.Interface().(time.Time); ok {
		t = time.Unix(seconds.(int64), 0).UTC()
		t2.Set(reflect.ValueOf(t))
		return
	}

	t2.Set(reflect.ValueOf(seconds))
}
