package Events

import (
	"github.com/elastic/go-elasticsearch/v8"
	driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"reflect"
	"encoding/json"
	"log"
	"bytes"
	"context"
)
type ProjectResult struct {
	Publication
}
func (p *Publication) GetProjectFields(field string, value string) string {
	r := reflect.ValueOf(p)
	f := reflect.Indirect(r).FieldByName(field)
	return f.FieldByName(value).String()
}
func getPublication(publiicationID, neo4jURL string) {
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	cypher := `MATCH (p:PUBLICATION{ProjectID:{ProjectID}})
			USING INDEX p:PUBLICATION(ProjectID)`

	
	_, _, _, err = db.QueryNeoAll(cypher, map[string]interface{}{
		"ProjectID": publiicationID,
	})
	if err != nil{
		return
	}
}
func search(request, index string) error{
	var r  map[string]interface{}
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": request,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
	return nil
}
func createProject(projectID, title, description, date_created, time_updated, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	cypher := `CREATE(
		n:PROJECT{ProjectID:{ProjectID}, Title:{Title}, 
		description:{description}, status:{status}})`
	_, err = db.ExecNeo(cypher, map[string]interface{}{
		"ProjectID": projectID,
		"Title": title,
		"description": description,
		"date_created": date_created,
		"time_updated": time_updated,
	})
	if err != nil {
		return err
	}
	return nil
}
func deleteProject(projectID, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()

	cypher := `MATCH(n:PROJECT{ProjectID:{ProjectID}}) DELETE n`

	_, err = db.ExecNeo(cypher, map[string]interface{}{
			"ProjectID": projectID,
		})
		if err != nil {
			return err
		}
	return nil
}
func addFile() error{
	return nil
}
func DeleteItem() error{
	return nil
}
func changeTitle(projectID, title string) error{
	return nil
}
func changeDesription(projectID, desription string) error{
	return nil
}
