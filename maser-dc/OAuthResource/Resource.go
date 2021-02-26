package OAuth
import (
    "golang.org/x/crypto/bcrypt"
)

type AuthorizationGrant struct {
	Username string `json:"username"`
	Password string `json:"password"`
	GrantType string `json:"grantType"`
}
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
func processAuthorizationGrant(){

}
func requestAccessTokens(){

}
func processAccessToken(){

}
func invalidateAccessTokens(){

}