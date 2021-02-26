package Events

import (
	"reflect"
	"time"
)
type Publication struct{
	Title string
	publishers []*Publisher
	Date_Created string
	Description string
	status string
	date_created time.Time
	update_time time.Time
}
type Publisher struct{
	first_name string
	last_name string
	userID string
}
type User struct {
	userID string
	firstName string
	lastName string
	email string	
	password string
	date_created string
	update_time string
}
type Privilege struct {
	Title string
	Administrative_Distance int
	Description string
}
type Ecosystem struct {
	EcosystemID string
	title string
	description string 
	status string
	date_created string
	update_time string
}
type tag struct {
	Title string
	Description string
}
type Item struct{
	OwnerID string
	ProjectID string
	project_name string
	Filename string
	file_size float64
	location string
	date_created string
	update_time string
}

func (u *User) GetUserFields(field string, value string) string {
	r := reflect.ValueOf(u)
	f := reflect.Indirect(r).FieldByName(field)
	return f.FieldByName(value).String()
}

func (i *Item) GetItemFields(field string, value string) string {
	r := reflect.ValueOf(i)
	f := reflect.Indirect(r).FieldByName(field)
	return f.FieldByName(value).String()
}
func (e *Ecosystem) GetEcosystemFields(field string, value string) string {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.FieldByName(value).String()
}