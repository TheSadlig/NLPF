package main

import( "github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
        "os"
	"time")


type Data struct {
    name string
    data string
}
type TestDocument struct {
    Title string
    Note string
}        

func v1_groupname(c *gin.Context) {


     c.Writer.Header().Set("Content-Type","application/json")
     c.Writer.WriteHeader(200)
     value, success := os.LookupEnv("GroupName")
     
     if (success) {
          c.Writer.Write([]byte("{\"status\": \"success\", \"data\": {\"groupname\" : \""+value+"\"}}"))
     } else {
          c.Writer.Write([]byte("{\"status\": \"success\", \"data\": {\"groupname\" : \"radon\"}}"))
     }

//     c.JSON(200, gin.H{"status": "success", "data": {"groupname" : "radon"} })
                                                       
}

func main() {
// Creates a gin router with default middleware:
// logger and recovery (crash-free) middleware
	router := gin.Default()

	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("127.0.0.1",5984,timeout)
	conn, err = conn, err

	auth := couchdb.BasicAuth{Username: "user", Password: "password" }
	db := conn.SelectDB("myDatabase", &auth)

	theDoc := TestDocument{
		Title: "My Document",
		Note: "This is a note",
	}

	theId := "zzz" //use whatever method you like to generate a uuid
	rev, err := db.Save(theDoc, theId, "")
	rev = rev
	
	router.GET("/v1/groupname", v1_groupname)
	router.Run(":9090")
}
