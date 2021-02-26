package middleware

import (
	"encoding/json"
	"net/http"
)
type credentials struct{
	username string `json:"username"`
	password string `json:"password"`
}
type WebServiceHandler struct{

}
func (handler *WebServiceHandler) SignIn(w http.ResponseWriter, r *http.Request){
	var creds credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
}
func (handler *WebServiceHandler) SignOut(w http.ResponseWriter, r *http.Request){

}