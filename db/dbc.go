package db

import (
	"fmt"
	"math/rand"
	"shareMe/util"
	"strconv"
	"strings"
	"time"
)

func generateShareCode() string {
	shareCode := ""
	rand.Seed(time.Now().UnixNano())
	min := 55
	max := 90
	for i := 0; i < 4; i++ {
		randomInt := rand.Intn(max-min+1) + min
		if randomInt < 65 {
			shareCode += strconv.Itoa(randomInt - 55)
		} else {
			shareCode += string(rune(randomInt))
		}
	}
	return shareCode
}

func GenerateShareCode() string {
	shareCode := generateShareCode()
	for {
		if !QueryShareCodeExist(shareCode) {
			break
		}
		shareCode = generateShareCode()
	}

	return shareCode
}

func CreateFileRecord(fileNames []string, shareCode string) bool {
	valToStore := ""
	for i := 0; i < len(fileNames)-1; i++ {
		valToStore += fileNames[i] + util.Delim // delim
	}
	if len(fileNames) > 0 {
		valToStore += fileNames[len(fileNames)-1]
	}
	err := DB.Put([]byte(shareCode), []byte(valToStore), nil)
	err = DB.Put([]byte(shareCode+"_T"), []byte("File"), nil)
	if err != nil {
		return false
	}
	return true
}

func PoolByShareCode(shareCode string) bool {
	err := DB.Delete([]byte(shareCode), nil)
	err = DB.Delete([]byte(shareCode+"_T"), nil)
	if err != nil {
		return false
	}
	return true
}

func QueryShareCodeExist(shareCode string) bool {
	ok, err := DB.Has([]byte(shareCode), nil)
	if err != nil {
		return false
	}
	if ok {
		return true
	}
	return false
}

func QueryShareCodeType(shareCode string) string {
	data, err := DB.Get([]byte(shareCode+"_T"), nil)
	if err != nil {
		return "Error"
	}
	return string(data[:])
}

func QueryFileByShareCode(shareCode string) []string {
	var returnList []string
	data, err := DB.Get([]byte(shareCode), nil)
	if err == nil {
		fileNamesStr := string(data[:])
		arr := strings.Split(fileNamesStr, util.Delim)
		for i := 0; i < len(arr); i++ {
			returnList = append(returnList, arr[i])
		}
	}
	return returnList
}

func CreateOtherRecord(shareCode string, info string, recordType string) bool {
	err := DB.Put([]byte(shareCode), []byte(info), nil)
	err = DB.Put([]byte(shareCode+"_T"), []byte(recordType), nil)
	if err != nil {
		return false
	}
	return true
}

func QueryOtherByShareCode(shareCode string) string {
	var returnStr = ""
	data, err := DB.Get([]byte(shareCode), nil)
	if err == nil {
		returnStr = string(data[:])
	}
	return returnStr
}

func GetAll() {
	iter := DB.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println("key: ", key, "value: ", value)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println("Error Occurred during Destroy Iter Release!")
	}
}
