package parsers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"reflect"
	"regexp"
	"strconv"
)

var rangeTagReg = regexp.MustCompile(`(\w*){(\d*),(\d*)}`)

type fieldInfo struct {
	name string
	kind reflect.Kind
}

func fieldToIndex(csvFields []string) map[int]fieldInfo {
	t := reflect.TypeOf(Hotel{})
	hFields := make(map[string]fieldInfo)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fi := fieldInfo{field.Name, field.Type.Kind()}
		if tag := field.Tag.Get("csv"); tag != "" {
			switch field.Type.Kind() {
			case reflect.Slice:
				match := rangeTagReg.FindStringSubmatch(tag)
				if len(match) != 4 {
					hFields[tag] = fi
					break
				}
				from, err := strconv.ParseUint(match[2], 10, 32)
				if err != nil {
					log.Fatal("A non-numeric character encountered")
				}
				to, err := strconv.ParseUint(match[3], 10, 32)
				if err != nil {
					log.Fatal("A non-numeric character encountered")
				}
				baseTag := match[1]
				for a := from; a <= to; a++ {
					tag = fmt.Sprintf("%s%d", baseTag, a)
					hFields[tag] = fi
				}
			default:
				if tag == "-" {
					tag = ""
				} else {
					hFields[tag] = fi
				}
			}
		}
	}
	fMap := make(map[int]fieldInfo)
	for i, csvField := range csvFields {
		if ftype, ok := hFields[csvField]; ok {
			fMap[i] = ftype
		}
	}
	return fMap
}

// CSVParse converts the CSV dump into a Hotel structure
func CSVParse(file io.Reader, hotels chan Hotel) {
	r := csv.NewReader(file)

	fields, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	if err == io.EOF {
		return
	}
	fMap := fieldToIndex(fields)

	for {
		var hotel Hotel
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for pos, fieldinfo := range fMap {
			field := reflect.ValueOf(&hotel).Elem().FieldByName(fieldinfo.name)
			value := record[pos]
			switch fieldinfo.kind {
			case reflect.String:
				field.SetString(value)
			case reflect.Float64:
				if s, err := strconv.ParseFloat(value, 64); err == nil {
					field.SetFloat(s)
				}
			case reflect.Uint8:
				if s, err := strconv.ParseUint(value, 10, 8); err == nil {
					field.SetUint(s)
				}
			case reflect.Slice:
				field.Set(reflect.Append(field, reflect.ValueOf(value)))
			}
		}
		hotels <- hotel
	}
}
