package events

import (
	driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func createUser(UserID, FirstName, LastName, Email, Password, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	_, _, _, err = db.QueryNeoAll(`CREATE(n:USER{UserID:{UserID}, FirstName:{FirstName}, 
									LastName:{LastName}, Email:{Email}, Password:{Password}, 
									Date_Created:{TIMESTAMP()}, Date_Last_Modified: {TIMESTAMP()}
									})`,
		map[string]interface{}{
			"UserID": UserID,
			"FirstName": FirstName,
			"LastName": LastName,
			"Email": Email,
			"Password": Password,
		})
		if err != nil {
			return err
		}

	return nil
}
func deleteUser(UserID, adminToken, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	_, _, _, err = db.QueryNeoAll(`"MATCH(n:USER{UserID:{UserID}}) DELETE n"`,
		map[string]interface{}{
			"UserID": UserID,
		})
		if err != nil {
			return err
		}
	return nil
}
func changeFirstName(FirstName, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)	
	defer db.Close()
	_, _, _, err = db.QueryNeoAll(`"MATCH(n:USER{UserID:{UserID}}) 
							SET n.FirstName = {FirstName}"`,
		map[string]interface{}{
			"FirstName": FirstName,
		})
		if err != nil {
			return err
		}
	return nil
}
func changeLastName(LastName, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)	
	defer db.Close()
	_, _, _, err = db.QueryNeoAll(`"MATCH(n:USER{UserID:{UserID}}) 
							SET n.LastName = '{LastName}'`,
		map[string]interface{}{
			"LastName": LastName,
		})
		if err != nil {
			return err
		}
	return nil
}
func changePassword(Password, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)	
	defer db.Close()
	_, _, _, err = db.QueryNeoAll(`"MATCH(n:USER{UserID:$UserID}) 
							SET n.Password = $Password"`,
		map[string]interface{}{
			"Password": Password,
		})
		if err != nil {
			return err
		}
	return nil
}

