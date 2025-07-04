package models

import "time"

type PersonalInfo struct {
	FullName string `json:"fullName" form:"fullName"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Location string `json:"location" form:"location"`
	LinkedIn string `json:"linkedin" form:"linkedin"`
	GitHub   string `json:"github" form:"github"`
	Website  string `json:"website" form:"website"`
	Summary  string `json:"summary" form:"summary"`
}

type Education struct {
	Institution string `json:"institution" form:"institution"`
	Degree      string `json:"degree" form:"degree"`
	StartDate   string `json:"startDate" form:"startDate"`
	EndDate     string `json:"endDate" form:"endDate"`
	Description string `json:"description" form:"description"`
}

type Experience struct {
	Company     string `json:"company" form:"company"`
	Position    string `json:"position" form:"position"`
	StartDate   string `json:"startDate" form:"startDate"`
	EndDate     string `json:"endDate" form:"endDate"`
	Description string `json:"description" form:"description"`
}

type Skill struct {
	Name  string `json:"name" form:"name"`
	Level string `json:"level" form:"level"`
}

type CV struct {
	PersonalInfo PersonalInfo `json:"personalInfo"`
	Education    []Education  `json:"education"`
	Experience   []Experience `json:"experience"`
	Skills       []Skill      `json:"skills"`
	Languages    []string     `json:"languages"`
	Language     string       `json:"language" form:"language"` // UI language (en/es)
	CreatedAt    time.Time    `json:"createdAt"`
}
