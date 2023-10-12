package main

type User struct {
	id            int    `json:"id" binding:"required"`
	name          string `json:"name"`
	userCondition Conditions
}

type Conditions struct {
	conditions []int `json:"conditions"`
	status     []bool
}
