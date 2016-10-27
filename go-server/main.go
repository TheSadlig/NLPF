package main

import( "github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"time"
	"encoding/json"
	"reflect"
	"fmt")


type Data struct {
    name string
    data string
}

func get_projects(c *gin.Context) {

	projects := getProjects()
	c.JSON(200, getTransmittableProjectList(projects))

}

func getAuth() couchdb.BasicAuth {
	return couchdb.BasicAuth{"admin", "admin"}
}

func getConn () *couchdb.Connection {
	var timeout = time.Duration(50000 * time.Millisecond)
	conn, err := couchdb.NewConnection("127.0.0.1",5984,timeout)
	fmt.Println(err)
	return conn

}
func createDB () {
	conn := getConn()
	auth := getAuth()
	err := conn.CreateDB("nlpf", &auth);
	fmt.Println(err)
}

func getDB () *couchdb.Database {
	conn := getConn()
	auth := getAuth()
	db := conn.SelectDB("nlpf", &auth)
	
	return db
}

func main() {
// Creates a gin router with default middleware:
// logger and recovery (crash-free) middleware
	router := gin.Default()

	createDB()
	
	var parsed map[string]interface{}

	data := []byte(`
    {
        "success": true,
        "data": {
"errors" : ["Coucou !", "Hello"], "user":  "Roman"
}
    }`)


	err := json.Unmarshal(data, &parsed)
	fmt.Println(err);
	fmt.Println("JSON Content:");
	fmt.Println(parsed["success"]);

	for key, value := range parsed {
		fmt.Println(reflect.TypeOf(value))
		fmt.Println("1Key:", key, "Value:", value)
		val, ok := value.(map[string]interface{})
		if ok {
			for k2, v2 := range val {
				fmt.Println(reflect.TypeOf(v2))
				fmt.Println("2Key:", k2, "Value:", v2)
				
				fmt.Println(reflect.TypeOf(v2))
				val2, ok := v2.([]interface{})
				if ok {
					for k3, v3 := range val2 {
						fmt.Println(reflect.TypeOf(v3))
						fmt.Println("3Key:", k3, "Value:", v3)
					}
				}
			}
		}
	}
	/*
	var parsedData map[string]interface{}
	err = json.Unmarshal(, &parsedData)
	fmt.Println(err);
	fmt.Println(parsedData);
	fmt.Println(parsedData["errors"]);
	*/

	createViews()

	fmt.Println("Liste des projets: ", getProjects())
	
	router.GET("/api/getProjects", get_projects)
	router.Run(":9090")
}
