package events

import driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"

func createEcosystem(EcoSystemID, title, description, status, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	_, err = db.ExecNeo(`CREATE(
		n:ECOSYSTEM{EcoSystemID:{EcoSystemID}, title:{title}, description:{description}, 
		status:{status}})`,
		map[string]interface{}{
			"EcoSystemID": EcoSystemID,
			"title": title,
			"description": description,
			"status": status,
		})
		if err != nil {
			return err
		}
	return nil
}
func destroyEcosystem(EcoSystemID, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	_, err = db.ExecNeo(`"MATCH(n:ECOSYSTEM{EcoSystemID:{EcoSystemID}}) 
							DETACH DELETE n"`,
		map[string]interface{}{
			"EcoSystemID": EcoSystemID,
		})
		if err != nil {
			return err
		}
	return nil
}
func addStudent(EcoSystemID, UserID, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	_, err = db.ExecNeo(`"MATCH (e:ECOSYSTEM),(u:USER)
							WHERE e.EcoSystemID = '{EcoSystemID}' AND u.UserID = '{UserID}'
							CREATE (u)-[r:STUDENT_OF]->(e)"`,
		map[string]interface{}{
			"EcoSystemID": EcoSystemID,
			"UserID": UserID,
		})
		if err != nil {
			return err
		}
	return nil
}
func removeStudent(EcoSystemID, UserID, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	_, err = db.ExecNeo(`"MATCH (u:USER{UserID:$UserID})-[r.STUDENT_OF]->(e:ECOSYSTEM{EcoSystemID=$EcoSystemID})
							DELETE r`,
		map[string]interface{}{
			"EcoSystemID": EcoSystemID,
			"UserID": UserID,
		})
		if err != nil {
			return err
		}
	return nil
}
func ApplyToEcoSystem(userToken, EcosystemID string){}
func LeaveEcoSystem(userToken, EcosystemID string){}
func AcceptApplicant(ApplicantToken, ModeratorToken string){}
func RejectApplicant(ApplicantToken, ModeratorToken string){}
