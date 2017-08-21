package util

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/goinggo/mapstructure"
)

func GetProgramName() string {
	base := path.Base(os.Args[0])
	ext := path.Ext(base)
	return strings.TrimSuffix(base, ext)
}

// GetCurFilename
// Get current file name, without suffix
func GetCurFilename() string {
	_, fulleFilename, _, _ := runtime.Caller(0)
	//fmt.Println(fulleFilename)
	filenameWithSuffix := path.Base(fulleFilename)
	//fmt.Println("filenameWithSuffix=", filenameWithSuffix)
	fileSuffix := path.Ext(filenameWithSuffix)
	//fmt.Println("fileSuffix=", fileSuffix)

	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	//fmt.Println("filenameOnly=", filenameOnly)

	return filenameOnly
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Map2Struct(obj interface{}, result interface{}) error {
	return mapstructure.Decode(obj, result)
}

func MD5Encode(in string) string {
	hash := md5.New()
	hash.Write([]byte(in))
	b := hash.Sum(nil)
	return hex.EncodeToString(b)
}
func Set(list interface{}) interface{} {
	m := make(map[interface{}]bool)
	result := []interface{}{}
	switch list.(type) {
	case []int:
		//fmt.Println("List is []int: ", list.([]int))
		for _, v := range list.([]int) {
			m[v] = true
		}
	case []string:
		//fmt.Println("List is []string: ", list.([]string))
		for _, v := range list.([]string) {
			m[v] = true
		}
	}
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}
