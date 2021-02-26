package main

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"log"
	"fmt"
)
func main() {
    dbUri := "neo4j://localhost:7687"
    driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("username", "password", ""))
    if err != nil {
		panic(err)
	}
	defer driver.Close()

	r := mux.NewRouter()
	fmt.Println("Starting server on port 8080.....")
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
func helloWorld(uri, username, password string) (string, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return "", err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]interface{}{"message": "hello, world"})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return greeting.(string), nil
}