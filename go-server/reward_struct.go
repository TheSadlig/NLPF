package main

import ("net/url")


type Reward struct {
	Type string
	ID string
	Project_ID string
	Title string
	Description string
	Value float64
	InvestorNumber float64
}

func getEmptyReward() *Reward {
	reward := Reward{
		Type: "Reward",
		ID: getUUID()}
	return &reward
}

func getInvestmentNumberByReward(rewardID string) float64 {
	db := getDB();

	result := ViewResponse{}
	parameters := url.Values{}
	parameters.Set("key", "\""+rewardID+"\"")

	db.GetView("investment","get_investments_number_by_reward", &result, &parameters)
	if len(result.Rows) > 0 {
		r, _ := result.Rows[0].(map[string]interface{})
		return r["value"].(float64)
	}
	return 0
}

func getRewardFromMap(m map[string]interface{}) *Reward {
	r, ok := m["value"].(map[string]interface{})
	if ok {
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

func getRewardByProject(projectID string) (*[]Reward, float64) {
	db := getDB();

	result := ViewResponse{}
	parameters := url.Values{}
	parameters.Set("key", "\""+projectID+"\"")

	db.GetView("rewards","get_rewards", &result, &parameters)
	
	rewards := []Reward{}
	rewards = rewards
	projectIncome := float64(0)
	for _, value := range result.Rows {
		m, ok := value.(map[string]interface{})
		if ok {
			r := getRewardFromMap(m)
			r.InvestorNumber = getInvestmentNumberByReward(r.ID)
			projectIncome += r.InvestorNumber * r.Value
			if r != nil {
				rewards = append(rewards, *r)
			}
		}
	}
	return &rewards, projectIncome
}
