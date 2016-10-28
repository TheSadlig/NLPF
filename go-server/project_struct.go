package main

import "fmt"
type Project struct {
	Type string
	ID string
	Name string
	Description string
	User_ID string
	Date string
}

/*
name
income
date
desc
author
mail
rewards
*/

type TransmittableProject struct {
	ID string
	Name string
	Description string
	Rewards []Reward
	Author string
	Date string
	Mail string
	Income float64	
}

func getEmptyProject() *Project {
	project := Project{
		Type: "Project",
		ID: getUUID()}
	return &project
}

func getProjectById(projectID string) *Project {
	db := getDB();
	
	result := Project{}
	_, _ = db.Read(projectID, &result, nil)

	return &result
}

func getProjects() *[]Project {
	db := getDB();

	result := ViewResponse{}

	db.GetView("projects","get_projects", &result, nil)
	
	projects := []Project{}
	
	for _, value := range result.Rows {
		m, ok := value.(map[string]interface{})
		if ok {
			p := getProjectFromMap(m)
			if p != nil {
				projects = append(projects, *p)
			}
		}
	}
	return &projects
}

func makeTransmittableProjectList(project *[]Project) *[]TransmittableProject {
	transmittableProjects := []TransmittableProject{}
	
	for _, p := range *project {
		transmittableProjects = append(transmittableProjects, *getTransmittableProject(p))
	}
	return &transmittableProjects
}


func getProjectFromMap(m map[string]interface{}) *Project {
	r, ok := m["value"].(map[string]interface{})
	if ok {
		project := Project{}
		project.Type = "Project"
		id, _ := r["ID"].(string)
		project.ID = id
		user_id, _ := r["User_ID"].(string)
		project.User_ID = user_id
		name, _ := r["Name"].(string)
		project.Name = name
		desc, _ := r["Description"].(string)
		project.Description = desc
		date, _ := r["Date"].(string)
		project.Date = date
		
		return &project
	}
	return nil
}


func getTransmittableProject(project Project) *TransmittableProject {
	transmitProject := TransmittableProject{}

	transmitProject.Name = project.Name
	transmitProject.Description = project.Description
	transmitProject.Date = project.Date
	transmitProject.ID = project.ID
	a, b := getRewardByProject(project.ID)
	transmitProject.Rewards = *a
	transmitProject.Income = b
	fmt.Println("transmit:", transmitProject)
	user := getUserById(project.User_ID)
	fmt.Println("user:", user)
	
	transmitProject.Author = user.Firstname + " "+ user.Lastname
	transmitProject.Mail = user.Mail
	
	return &transmitProject
}
