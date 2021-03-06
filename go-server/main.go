package main

import( "github.com/gin-gonic/gin"
	"encoding/json")

func get_projects(c *gin.Context) {

	projects := getProjects()
	
	c.JSON(200, gin.H{"success": true, "data": makeTransmittableProjectList(projects)})
}

func get_projects_by_id(c *gin.Context) {

	rawData := c.PostForm("data")

	var parsed map[string]interface{}
	data :=[]byte(rawData)

	json.Unmarshal(data, &parsed)

	val, ok := parsed["data"].(map[string]interface{})
	project := Project{}
	if ok {
		ID := val["ID"].(string)
		project = *getProjectById(ID)
	}

	c.JSON(200, gin.H{"success": true, "data": getTransmittableProject(project)})
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

func connect_user(c *gin.Context) {
	rawData := c.PostForm("data")

	var parsed map[string]interface{}
	data :=[]byte(rawData)

	json.Unmarshal(data, &parsed)

	val, ok := parsed["data"].(map[string]interface{})
	mail := ""
	password := ""
	
	if ok {
		mail = val["mail"].(string)
		password = val["password"].(string)
	}

	existant := getUserByMail(mail)
	
	if existant.ID == "" || existant.Password != password {
		c.JSON(200, gin.H{"success": false})
	} else {
		c.JSON(200, gin.H{"success": true, "data": existant})
	}
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

	existant := getUserByMail(u.Mail)
	
	if existant.ID == "" {
		db.Save(&u, u.ID, "")
		c.JSON(200, gin.H{"success": true})
	} else {
		c.JSON(200, gin.H{"success": false})
	}
}

func invest(c *gin.Context) {
	db := getDB()
	
	rawData := c.PostForm("data")

	var parsed map[string]interface{}
	data :=[]byte(rawData)

	json.Unmarshal(data, &parsed)

	invest := getEmptyInvestment()
	
	val, ok := parsed["data"].(map[string]interface{})

	rewardID := ""
	userID := ""
	
	if ok {
		rewardID = val["rewardID"].(string)
		userID = val["userID"].(string)
	}
	invest.Reward_ID = rewardID
	invest.User_ID = userID

	db.Save(&invest, invest.ID, "")
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
	
	r.GET("/api/getProjects", get_projects)
	
	r.POST("/api/getProjectById", get_projects_by_id)

	r.POST("/api/connectUser", connect_user)

	r.POST("/api/invest", invest)
	
	r.POST("/api/createProject", create_project)
	r.POST("/api/createUser", create_user)
	r.Run(":9090")
}
