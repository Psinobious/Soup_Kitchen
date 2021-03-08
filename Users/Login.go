package users

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type UserLogin struct{
	userID string
	password string
}
type UserLogInHandler struct{
	Path string
	UserRepository UserRepository
}
type LoggedInUser struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
}
type OAuthClient struct {

}
func (u *UserLogInHandler) Login(writer http.ResponseWriter, request *http.Request){
	requestBody, _ := ioutil.ReadAll(request.Body)
	userLogInRequest := UserLogin{}
	_ = json.Unmarshal(requestBody, &userLogInRequest)
}
func getToken(){

}