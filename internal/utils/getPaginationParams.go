package utils

import (
	"log"
	"net/url"
	"strconv"
)

func GetPaginationParams(URL url.Values) struct {
	Limit  int32
	Offset int32
} {
	Limit_str := URL.Get("perPage")
	Offset_str := URL.Get("page")
	limit_num, err := strconv.Atoi(Limit_str)
	if err != nil {
		limit_num = 8
		//log.Println(err)
	}
	if limit_num < 1 {
		limit_num = 1
	}
	Offset_num, err := strconv.Atoi(Offset_str)
	if err != nil {
		Offset_num = 1
		//log.Println(err)
	}
	if Offset_num < 0 {
		Offset_num = 0
	}
	Offset_num = (Offset_num - 1) * (limit_num)
	log.Println(limit_num, Offset_num)
	return struct {
		Limit  int32
		Offset int32
	}{Limit: int32(limit_num), Offset: int32(Offset_num)}
}
