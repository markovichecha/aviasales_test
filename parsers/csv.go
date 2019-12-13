package parsers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

var reg = regexp.MustCompile(`(\w*){(\d*),(\d*)}`)

func fieldToIndex(csvFields []string) map[int]reflect.StructField {
	t := reflect.TypeOf(Hotel{})
	hFields := make(map[string]reflect.StructField)
	for i := 0; i < t.NumField(); i++ {
		hField := t.Field(i)
		if csvTag := hField.Tag.Get("csv"); csvTag != "" {
			switch hField.Type.Kind() {
			case reflect.Slice:
				match := reg.FindStringSubmatch(csvTag)
				if len(match) == 4 {
					from, err := strconv.ParseInt(match[2], 10, 32)
					if err != nil {
						log.Fatal(err)
					}
					to, err := strconv.ParseInt(match[3], 10, 32)
					if err != nil {
						log.Fatal(err)
					}
					for a := from; a <= to; a++ {
						baseTag := fmt.Sprintf("%s%d", match[1], a)
						hFields[baseTag] = hField
					}
				}
			default:
				hFields[csvTag] = hField
			}
		}
	}
	fMap := make(map[int]reflect.StructField)
	for i, csvField := range csvFields {
		if ftype, ok := hFields[csvField]; ok {
			fMap[i] = ftype
		}
	}
	return fMap
}

// CSVParse converts the CSV dump into a Hotel structure
func CSVParse(file *os.File, hotels chan Hotel) {
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

		for pos, fieldRef := range fMap {
			field := reflect.ValueOf(&hotel).Elem().FieldByName(fieldRef.Name)
			value := record[pos]
			switch fieldRef.Type.Kind() {
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
