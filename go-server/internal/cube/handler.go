package cube

import (
	"encoding/json"
	"go-server/internal/logger"
	"net/http"
	"strconv"
)

func CalculateCubeHandler(w http.ResponseWriter, r *http.Request){
 numberString := r.URL.Query().Get("number");
 number,err := strconv.Atoi(numberString);
 if(err !=nil){
	logger.Log.Error("calculation fail: "+err.Error())
	return
 }

 result:= Cube(number)
 response:= map[string] interface{}{
	"number":number,
	"cube":result,
 }

 json.NewEncoder(w).Encode(response)

 


}