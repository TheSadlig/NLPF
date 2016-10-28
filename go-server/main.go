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
	db := getDB()
	
	rawData := c.PostForm("data")

	var parsed map[string]interface{}
	data :=[]byte(rawData)

	json.Unmarshal(data, &parsed)
	p := getEmptyProject()


	val, ok := parsed["data"].(map[string]interface{})

	if ok {
		p.Description = val["desc"].(string)
		p.Name = val["name"].(string)
		p.Date = val["date"].(string)
		p.User_ID = val["userID"].(string)
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
				db.Save(&reward, reward.ID, "")
			}
		}
	}
	db.Save(&p, p.ID, "")
	
	c.JSON(200, gin.H{"success": true})

}


/*
{"data": {
"lastname": "Gildas",
"firstname": "lebel",
"mail" : "gildaslebel@hezze.fr", "password": "efzeiluh"}
}
*/
func create_user(c *gin.Context) {
	db := getDB()
	
	rawData := c.PostForm("data")

	var parsed map[string]interface{}
	data :=[]byte(rawData)

	json.Unmarshal(data, &parsed)
	u := getEmptyUser()


	val, ok := parsed["data"].(map[string]interface{})

	if ok {
		u.Lastname = val["lastname"].(string)
		u.Firstname = val["firstname"].(string)
		u.Mail = val["mail"].(string)
		u.Password = val["password"].(string)

	}
	db.Save(&u, u.ID, "")
	
	c.JSON(200, gin.H{"success": true})

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
	r.POST("/api/createUser", create_user)
	r.Run(":9090")
}
