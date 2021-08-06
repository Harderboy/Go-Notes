package controller

import (
	"net/http"
	"regexp"
)

func registerCompanyRouter() {
	http.HandleFunc("/companies", handleCompanies)
	http.HandleFunc("/companies/", handleCompany)
}

func handleCompanies(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello Companies"))
}

func handleCompany(writer http.ResponseWriter, request *http.Request) {
	pattern, _ := regexp.Compile(`/companies/(\d+)`)
	matches := pattern.FindStringSubmatch(request.URL.Path)
	if len(matches) > 0 {
		// companyID,_:=strconv.Atoi(matches[1])
		writer.Write([]byte(matches[1]))
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}
