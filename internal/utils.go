package lucy

import (
	"fmt"
	"github.com/supercmmetry/lucy/types"
	"reflect"
	"strconv"
)


func Marshal(v interface{}) types.Exp {
	vtype := reflect.TypeOf(v)
	vvalue := reflect.ValueOf(v)

	for vtype.Kind() != reflect.Struct {
		if vtype.Kind() == reflect.Ptr {
			vvalue = vvalue.Elem()
		} else if vtype.Kind() == reflect.Slice {
			vvalue = reflect.New(vtype.Elem()).Elem()
		}
		vtype = vtype.Elem()
	}

	tagMap := make(map[string]interface{})

	for i := 0; i < vtype.NumField(); i++ {
		if tagName, ok := vtype.Field(i).Tag.Lookup("lucy"); ok {
			tagMap[tagName] = vvalue.Field(i).Interface()
		}
	}

	return tagMap
}

func Unmarshal(src map[string]interface{}, dest interface{}) {
	vtype := reflect.TypeOf(dest)
	vvalue := reflect.ValueOf(dest)

	if vtype.Kind() != reflect.Struct {
		vtype = reflect.TypeOf(dest).Elem()
		vvalue = reflect.ValueOf(dest).Elem()
	}

	for i := 0; i < vtype.NumField(); i++ {
		if tagName, ok := vtype.Field(i).Tag.Lookup("lucy"); ok {
			obj := src[tagName]

			switch reflect.TypeOf(vvalue.Field(i).Interface()).Kind() {
			case reflect.String:
				vvalue.Field(i).SetString(obj.(string))
			case reflect.Bool:
				vvalue.Field(i).SetBool(obj.(bool))
			case reflect.Int:
				vvalue.Field(i).SetInt(obj.(int64))
			case reflect.Int8:
				vvalue.Field(i).SetInt(obj.(int64))
			case reflect.Int16:
				vvalue.Field(i).SetInt(obj.(int64))
			case reflect.Int32:
				vvalue.Field(i).SetInt(obj.(int64))
			case reflect.Int64:
				vvalue.Field(i).SetInt(obj.(int64))
			case reflect.Uint:
				vvalue.Field(i).SetUint(obj.(uint64))
			case reflect.Uint8:
				vvalue.Field(i).SetUint(obj.(uint64))
			case reflect.Uint16:
				vvalue.Field(i).SetUint(obj.(uint64))
			case reflect.Uint32:
				vvalue.Field(i).SetUint(obj.(uint64))
			case reflect.Uint64:
				vvalue.Field(i).SetUint(obj.(uint64))
			case reflect.Float32:
				vvalue.Field(i).SetFloat(obj.(float64))
			case reflect.Float64:
				vvalue.Field(i).SetFloat(obj.(float64))
			}
		}
	}
}

func SFormat(format string, I []interface{}) string {
	newStr := ""
	index := 0
	for _, chr := range format {
		if chr == '?' {
			index += 1
			i := index - 1

			switch reflect.TypeOf(I[i]).Kind() {
			case reflect.String:
				{
					subStr := ""
					targStr := I[i].(string)
					for _, c := range targStr {
						if c == '\'' || c == '"' || c == '\\' {
							subStr += "\\"
						}
						subStr += string(c)
					}
					newStr += "'" + subStr + "'"
				}
			case reflect.Int:
				{
					newStr += strconv.Itoa(I[i].(int))
				}
			case reflect.Int64:
				{
					newStr += strconv.FormatInt(I[i].(int64), 10)
				}
			case reflect.Int32:
				{
					newStr += strconv.FormatInt(int64(I[i].(int32)), 10)
				}
			case reflect.Int16:
				{
					newStr += strconv.FormatInt(int64(I[i].(int16)), 10)
				}
			case reflect.Int8:
				{
					newStr += strconv.FormatInt(int64(I[i].(int8)), 10)
				}
			case reflect.Uint:
				{
					newStr += strconv.FormatUint(uint64(I[i].(uint)), 10)
				}
			case reflect.Uint8:
				{
					newStr += strconv.FormatUint(uint64(I[i].(uint8)), 10)
				}
			case reflect.Uint16:
				{
					newStr += strconv.FormatUint(uint64(I[i].(uint16)), 10)
				}
			case reflect.Uint32:
				{
					newStr += strconv.FormatUint(uint64(I[i].(uint32)), 10)
				}
			case reflect.Uint64:
				{
					newStr += strconv.FormatUint(I[i].(uint64), 10)
				}
			case reflect.Float32:
				{
					newStr += fmt.Sprintf("%f", I[i].(float32))
				}
			case reflect.Float64:
				{
					newStr += fmt.Sprintf("%f", I[i].(float64))
				}
			case reflect.Bool:
				{
					if I[i].(bool) {
						newStr += "true"
					} else {
						newStr += "false"
					}
				}
			}
		} else {
			newStr += string(chr)
		}
	}
	return newStr
}

func Format(format string, I ...interface{}) string {
	return SFormat(format, I)
}

func SanitizeExp(exp types.Exp) {
	for k, v := range exp {
		exp[k] = Format("?", v)
	}
}

func GetTypeName(i interface{}) string {
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.TypeOf(i).Elem().Kind() == reflect.Struct {
			return reflect.TypeOf(i).Elem().Name()
		} else if reflect.TypeOf(i).Elem().Kind() == reflect.Slice &&
			reflect.TypeOf(i).Elem().Elem().Kind() == reflect.Struct {
			return reflect.TypeOf(i).Elem().Elem().Name()
		}
	} else if reflect.TypeOf(i).Kind() == reflect.Struct {
		return reflect.TypeOf(i).Name()
	}
	return ""
}

type Queue struct {
	elements *[]interface{}
}

func (q *Queue) Init() {
	elements := make([]interface{}, 0)
	q.elements = &elements
}

func (q *Queue) GetAll() *[]interface{} {
	return q.elements
}

func (q *Queue) Push(elem interface{}) {
	*q.elements = append(*q.elements, elem)
}

func (q *Queue) Pop() (interface{}, error) {
	if (len(*q.elements)) == 0 {
		return nil, Error(EmptyQueue)
	}
	elems := q.elements

	lastElem := (*elems)[len(*elems)-1]
	*elems = (*elems)[:len(*elems)-1]

	return lastElem, nil
}

func (q *Queue) Get() (interface{}, error) {
	if (len(*q.elements)) == 0 {
		return nil, Error(EmptyQueue)
	}
	elem := (*q.elements)[0]
	*q.elements = (*q.elements)[1:]
	return elem, nil
}

func (q *Queue) IsEmpty() bool {
	return len(*q.elements) == 0
}
