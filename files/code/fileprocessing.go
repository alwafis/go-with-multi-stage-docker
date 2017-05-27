package main

import (
	"fmt"
	"io/ioutil"
  "strings"
)

type Simplefile struct {
    Title string
    Description string
    Extra string
}

func (s *Simplefile) save() error {
    filename := s.Title + ".txt"
    return ioutil.WriteFile(filename, []byte(s.Extra + ";" + s.Description), 0600)
}

func loadSimpleFile(title string) *Simplefile {
    filename := title + ".txt"
    body, _ := ioutil.ReadFile(filename)
    s := strings.Split(string(body), ";")
    return &Simplefile{Title: title, Description: s[1], Extra: s[0]}
}

func main() {
    obj1 := &Simplefile{Title: "TestSimpleFile", Description: "Simple Description", Extra: "Simple Extra"}
    obj1.save()
    obj2 := loadSimpleFile("TestSimpleFile")
    fmt.Println("This is title: " + string(obj2.Title))
    fmt.Println("This is description: " + string(obj2.Description))
    fmt.Println("This is extra: " + string(obj2.Extra))
}
