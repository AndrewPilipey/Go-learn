package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name             string
	age              uint
	money            int16
	grade, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is %s. He is %d and he has money "+
		"equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	bob.setNewName("Alex")
	bob.getAllInfo()
	fmt.Fprintf(w, bob.getAllInfo()+"BeMonitoringHome")

}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "BeMonitoringContacts")

}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
