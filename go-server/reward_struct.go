package main

import ("fmt"
	"net/url")


type Reward struct {
	Type string
	ID string
	Project_ID string
	Title string
	Description string
	Value float64
}

func getEmptyReward() *Reward {
	reward := Reward{
		Type: "Reward",
		ID: getUUID()}
	return &reward
}

func getRewardFromMap(m map[string]interface{}) *Reward {
	r, ok := m["value"].(map[string]interface{})
	if ok {
		fmt.Println(r)
		rew := Reward{}
		rew.Type = "Reward"
		id, _ := r["ID"].(string)
		rew.ID = id
		projectID, _ := r["Project_ID"].(string)
		rew.Project_ID = projectID
		title, _ := r["Title"].(string)
		rew.Title = title
		description, _ := r["Description"].(string)
		rew.Description = description
		value, _ := r["Value"].(float64)

		rew.Value = value
		return &rew
	}
	return nil
}

func getRewardByProject(projectID string) *[]Reward {
	db := getDB();

	result := ViewResponse{}
	parameters := url.Values{}
	parameters.Set("key", "\""+projectID+"\"")

	err := db.GetView("rewards","get_rewards", &result, &parameters)
	fmt.Println(err)

	rewards := []Reward{}
	rewards = rewards
	for _, value := range result.Rows {
		m, ok := value.(map[string]interface{})
		if ok {
			r := getRewardFromMap(m)
			if r != nil {
				rewards = append(rewards, *r)
			}
		}
	}
	return &rewards
}
