package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tri/maths"
	"tri/triangle"
)

type trian struct{
	Size int `json:"size"`
}

func TriangleHandle(rw http.ResponseWriter, r *http.Request) {
	(rw).Header().Set("Access-Control-Allow-Origin", "*")
	
	//fmt.Println(r)
	 newItem := trian{}
	 json.NewDecoder(r.Body).Decode(&newItem)
	if maths.Par(newItem.Size){
		if maths.Digit(newItem.Size){
			
			Conv, _ := json.MarshalIndent("Number too long", "", " ")
			fmt.Fprintf(rw, "%s", string(Conv))
			
		}else{
			response := triangle.TriangleBuildAPI(newItem.Size)
			Conv, _ := json.MarshalIndent(response, "", " ")
		fmt.Fprintf(rw, "%s", string(Conv))
		}
	
	}



}
