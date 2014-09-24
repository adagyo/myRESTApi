package utils

import (
	"log"
	"net/http"
	"strconv"
)

type QueryRequestParameters struct {
	Limit  int
	Offset int
	Sort   string
}

func get(val string, def string) string {
	if val != "" {
		return val
	}
	return def
}

func ParseRequestParameters(request *http.Request) QueryRequestParameters {
	queryParams := request.URL.Query()
	var limit, offset int
	var err error

	limit, err = strconv.Atoi(get(queryParams.Get("limit"), "20"))
	if err != nil {
		limit = 20
		log.Println("[WARNING] Can't parse 'limit': " + err.Error())
	}
	if limit > 100 {
		log.Println("[WARNING] limit is > 100: " + strconv.Itoa(limit))
		limit = 20
	}

	offset, err = strconv.Atoi(get(queryParams.Get("offset"), "0"))
	if err != nil {
		offset = 0
		log.Println("[WARNING] Can't parse 'offset': " + err.Error())
	}
	if offset < 0 {
		log.Println("[WARNING] offset is < 0: " + strconv.Itoa(offset))
		offset = 0
	}

	// TODO: Filtrer les colonnes où le sort est autorisé
	sort := get(queryParams.Get("sort"), "userid")

	return QueryRequestParameters{
		Limit:  limit,
		Offset: offset,
		Sort:   sort,
	}
}
