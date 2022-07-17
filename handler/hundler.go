package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetParamHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	param := vars["PARAM"]

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Hello, %s!", param)
}

func GetStatusBadHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(writer, "%d", http.StatusInternalServerError)
}

func PostMessageHandler(writer http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	str := string(body)

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "I got message:\n%v", str)
}

func SumHandler(writer http.ResponseWriter, request *http.Request) {
	a := request.Header.Get("a")
	b := request.Header.Get("b")

	resA, _ := strconv.Atoi(a)
	resB, _ := strconv.Atoi(b)
	writer.Header().Set("a+b", strconv.Itoa(resA+resB))

}
