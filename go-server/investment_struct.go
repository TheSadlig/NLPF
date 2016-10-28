package main

type Investment struct {
	Type string
	ID string
	Reward_ID string
	User_ID string
}

func getEmptyInvestment() *Investment {
	project := Investment{
		Type: "Investment",
		ID: getUUID()}
	return &project
}
