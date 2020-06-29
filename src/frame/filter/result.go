package filter

import "net/http"

func result200(response http.ResponseWriter, result []byte) {
	response.WriteHeader(200)
	response.Write(result)
}

func result500(response http.ResponseWriter, result []byte) {
	response.WriteHeader(500)
	response.Write(result)
}
