package main

import(
	"net/http"
	"fmt"
	"encoding/xml"
)

/*
these structs will house the unmarshlled response.
they should be hierarchically shaped like the XML.
but can omit irrelevant data.
*/
type Status struct{
	Text string
}

type User struct{
	XMLName xml.Name
	Status Status
}
// var user User

func main(){
	// perform an HTTP request for the twitter status of user: Crazcalm
	response, _ := http.Get("http://twitter.com/users/crazcalm.xml")

	// initial the structure of the XML.response
	user := User{xml.Name{"", "user"}, Status{""}}
	// unmarshal the XML into out structures
	xml.Unmarshal(response.Body, &user)
	fmt.Printf("status: %s", user.Status.Text)
}