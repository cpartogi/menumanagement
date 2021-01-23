package helper

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func RandOTP() int64 {
	rand.Seed(time.Now().UnixNano())

	var s string
	for i := 0; i < 5; i++ {
		s += fmt.Sprintf("%d", rand.Intn(9))
	}
	otp, _ := strconv.Atoi(s)
	return int64(otp)
}

func IsValidMD5(s string) bool {
	r, _ := regexp.Compile("^[a-fA-F0-9]{32}$")
	return r.MatchString(s)
}

func ValidateOTPExpired(createTime time.Time) bool {
	now := time.Now().Local()
	expiredAt := createTime.Add(time.Second * 180).Local()
	valid := now.Before(expiredAt)

	if valid {
		return true
	}

	return false
}

func ProcessReflectDataStrict(dataReflect reflect.Value) map[string]interface{} {
	if !dataReflect.IsValid() {
		return nil
	}

	data := make(map[string]interface{}, dataReflect.NumField())
	for i := 0; i < dataReflect.NumField(); i++ {
		value := dataReflect.Field(i).Interface()

		if value != "" && value != 0 {
			key := dataReflect.Type().Field(i).Name
			data[strcase.ToSnake(key)] = value
		}
	}

	return data
}

func ProcessReflectData(dataReflect reflect.Value) map[string]interface{} {
	if !dataReflect.IsValid() {
		return nil
	}

	data := make(map[string]interface{}, dataReflect.NumField())
	for i := 0; i < dataReflect.NumField(); i++ {
		value := dataReflect.Field(i).Interface()
		key := dataReflect.Type().Field(i).Name
		data[strcase.ToSnake(key)] = value
	}

	return data
}
