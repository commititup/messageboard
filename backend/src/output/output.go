package output

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


//FUNCTION FOR SENDIG DATA BACK TO CLIENT IN JSON FORMAT WITH SOME HEADER
func Send(w http.ResponseWriter, data map[string]interface{}, errMsg error, statusCode ...interface{}) {
	response := make(map[string]interface{})
	if errMsg == nil {
		response["error"] = nil
		response["success"] = true
	} else {
		response["error"] = errMsg.Error()
		response["success"] = false
	}

	response["data"] = data

	code := 200
	for _, arg := range statusCode {
		switch t := arg.(type) {
		case int:
			code = t
		}
	}

	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Println("Error in encoding response, err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in encoding response")
	} else {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
}

func RawSend(w http.ResponseWriter, data interface{}, errMsg error) {
	response := make(map[string]interface{})
	if errMsg == nil {
		response["error"] = nil
		response["success"] = true
	} else {
		response["error"] = errMsg.Error()
		response["success"] = false
	}
	response["data"] = data

	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Println("Error in encoding response, err: ", err)
		fmt.Fprintln(w, "Error in encoding response")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
}
