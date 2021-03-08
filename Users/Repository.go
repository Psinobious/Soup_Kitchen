package users

import (
	"time"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)
type UserRepository interface {
	changeFirstName(UserID, FirstName string)(error)
	changeLastName(UserID, LastName string)(error)
}
type UserNeo4jRepository struct {
	Driver neo4j.Driver
}
type User struct {
	userID string
	firstName string
	lastName string
	email string	
	password string
	status string
	date_created time.Time
	update_time time.Time
}

func (u *UserNeo4jRepository) findUser(UserID string, tx neo4j.Transaction) (*User, error){

	query := "Match (u:USER{UserID:'$UserID') return n"
	parameters := map[string] interface{}{
		"UserID": UserID,
	}
	result, err := tx.Run(query, parameters)
	if err != nil {
		panic(err)
	}
	
	record, err := result.Single()
	FirstName, _ := record.Get("FirstName")
	LastName, _ := record.Get("LastName")
	Email, _ := record.Get("Email")
	Status, _ := record.Get("Status")
	Date_created, _ := record.Get("DateCreated")
	Update_time, _ := record.Get("UpdateTime")
	
	return &User{
		userID: UserID,
		firstName: FirstName.(string),
		lastName: LastName.(string),
		email: Email.(string),
		status: Status.(string),
		date_created: Date_created.(time.Time),
		update_time: Update_time.(time.Time),
		}, nil	
}
func (u *UserNeo4jRepository) persistUser(UserID, FirstName, LastName, Email, Password string, tx neo4j.Transaction) error{
	query := `CREATE(n:USER{UserID:{UserID}, FirstName:{FirstName}, 
				LastName:{LastName}, Email:{Email}, Password:{Password}, 
				Date_Created:{TIMESTAMP()}, Date_Last_Modified: {TIMESTAMP()}
				})`
				
	parameters := map[string]interface{}{
		"UserID": UserID,
		"FirstName": FirstName,
		"LastName": LastName,
		"Email": Email,
		"Password": Password,
	}
	_, err := tx.Run(query, parameters)
	if err != nil {
		panic(err)
	}
	return nil
}
func (u *UserNeo4jRepository) deleteUser(UserID string, tx neo4j.Transaction) {
	query := "MATCH(n:USER{UserID:$UserID}) DELETE n"
	parameters := map[string]interface{}{
			"UserID": UserID,
		}
	_, err := tx.Run(query, parameters)
	if err != nil {
		panic(err)
	}
}
func (u *UserNeo4jRepository) changeFirstName(UserID, FirstName string) (error){
	session := u.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})

	defer session.Close()

	query := "MATCH (n:USER {UserID: $UserID}) SET n.FirstName = $FirstName"
	parameters := map[string]interface{}{
		"FirstName": FirstName,
		"UserID": UserID,
	}
	_, err := session.WriteTransaction(func(tx neo4j.Transaction)(interface{}, error){
		_, err := tx.Run(query, parameters)
		return nil, err
	})
	if err != nil {
		panic(err)
	}
	return nil
}
func (u *UserNeo4jRepository) changeLastName(LastName, UserID string) error{
	session := u.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})

	defer session.Close()
	
	query := "MATCH (n:USER {UserID: $UserID}) SET n.LastName = $LastName"
	parameters := map[string]interface{}{
		"LastName": LastName,
		"UserID": UserID,
	}
	_, err := session.WriteTransaction(func(tx neo4j.Transaction)(interface{}, error){
		_, err := tx.Run(query, parameters)
		return nil, err
	})
	if err != nil {
		panic(err)
	}
	return nil
}
func (u *UserNeo4jRepository) changePassword(Password string, tx neo4j.Transaction) error{
	query :="MATCH(n:USER{UserID:$UserID}) SET n.Password = $Password"
	parameters := map[string]interface{}{
			"Password": Password,
	}
	_, err := tx.Run(query, parameters)
	if err != nil {
		panic(err)
	}
	return nil
}

