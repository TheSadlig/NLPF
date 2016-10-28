package main

import( "github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"time"
	"encoding/json"
	"fmt")

func get_projects(c *gin.Context) {

	projects := getProjects()
		
	c.JSON(200, gin.H{"success": true, "data": makeTransmittableProjectList(projects)})

}

func create_project(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	db := getDB()
	
	rawData := c.PostForm("data")

	fmt.Println("raw:", rawData)
	var parsed map[string]interface{}
 /*
 {"data": {
"name": "Coucou",
"desc": "yolo",
"rewards" : [
{ "name": "rew1", "value": "100", "desc": "rewdesc1"},
{ "name": "rew1", "value" :"100", "desc": "rewdesc1"}
],
"date" : "03/12/1993",
"userID" : "jyfjytfhytfjytfjytf"
}                                                                                                               
}
*/
	data :=[]byte(rawData)

	err := json.Unmarshal(data, &parsed)
	fmt.Println(err);
	fmt.Println("JSON Content:");
	fmt.Println("parsed:", parsed);
	fmt.Println("parsed:", parsed["data"]);
	p := getEmptyProject()


	val, ok := parsed["data"].(map[string]interface{})
	fmt.Println(val)
	if ok {
		p.Description = val["desc"].(string)
		p.Name = val["name"].(string)
		p.Date = val["date"].(string)
		p.User_ID = val["userID"].(string)
		fmt.Println(p)
		rawRewards := val["rewards"]
		jsonRewards, ok := rawRewards.([]interface{})
		if ok {
			for _, v3 := range jsonRewards {
				r, _ := v3.(map[string]interface{})
				reward := getEmptyReward()
				reward.Project_ID = p.ID
				reward.Title = r["name"].(string)
				reward.Description = r["desc"].(string)
				value, _ := r["value"].(float64)
				reward.Value = value
				_, err = db.Save(&reward, reward.ID, "")
			}
		}
	}
	_, err = db.Save(&p, p.ID, "")
	
	c.JSON(200, gin.H{"success": true})

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
	//	router := gin.Default()
	r := gin.New()
	r.Use(func (c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
				c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
			}
			c.AbortWithStatus(200)
		}
	})
	createDB()
	


	createViews()

	fmt.Println("Liste des projets: ", getProjects())
	
	r.GET("/api/getProjects", get_projects)
	r.POST("/api/createProject", create_project)
	r.Run(":9090")
}
