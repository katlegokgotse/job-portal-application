package models

type UserProfile struct {
	userId       int
	firstName    string
	lastName     string
	emailAddress string
	password     string
	phone        string
	role         UserRole
}
