package model

import (
	"fmt"
	"time"
)

type Person struct {
	Id         int
	Name       string
	Sex        string
	Age        int
	Phoneone   string
	Phonetwo   string
	Phonethree string
	IdCard     string
	Content    string
	Job        int
	InDate     string
	OutDate    string
}

func (person *Person) PrintInfo() {
	fmt.Println("ID:", person.Id)
	fmt.Println("Name:", person.Name)
	fmt.Println("Sex:", person.Sex)
	fmt.Println("Age:", person.Age)
	fmt.Println("Phoneonw:", person.Phoneone)
	fmt.Println("Phonetwo:", person.Phonetwo)
	fmt.Println("Phonethree:", person.Phonethree)
	time.Now()
}

func NewPerson(name string, sex string, age int, phoneone string, phonetwo string, phonethree string, idcard string, content string, job int) *Person {
	return &Person{Name: name, Sex: sex, Age: age, Phoneone: phoneone, Phonetwo: phonetwo, Phonethree: phonethree,
		IdCard: idcard, Content: content}
}
