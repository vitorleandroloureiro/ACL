package models

type Role struct {
	Name                 string
	Allowed_resources    []Resource
	Disallowed_resources []Resource
}
