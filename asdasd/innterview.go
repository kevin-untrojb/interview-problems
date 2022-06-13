package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Person struct {
	FirstName string `json:firstName`
	Lastname string `json:lastName`
}

func main() {
	//Notify everyone in the most efficient way possible

	wg := sync.WaitGroup{}
	errChan := make(chan error)
	defer close(errChan)
	finish := make(chan bool)
	errString := ""

	messages := []string{`{"firstName":"John", "lastName":"Doe"}`,
		`{"firstName":"Luke", "lastName":"Williams"}`,
		`{"firstName":"Tom", "lastName":"Collins"}`}
	for _, msg := range messages{
		wg.Add(1)
		go func(msg string) {
			defer func(){
				wg.Done()
				if err := recover(); err != nil{
					fmt.Println(err)
				}
			}()

			p := Person{}
			if err := json.Unmarshal([]byte(msg), &p); err != nil {
				errChan<- err
			}
			if err := notifyPerson(p); err != nil {
				errChan <-fmt.Errorf(fmt.Sprintf("error: person %v %v", p,err))

			}
		}(msg)
	}
	go func(){
		wg.Wait()
		close(finish)
	}()

	for {
		select {
			case <- finish:
				if errString != ""{
					fmt.Println(errString)
				}
				fmt.Println("everyone notified")
				return
		case err :=<- errChan:
			errString =errString + fmt.Sprintf("%v", err)
		}
	}



}

//Please do not change my signature
func notifyPerson(person Person) error {
	//Do not implement
	time.Sleep(time.Second * 5)
	//Do not uncomment (yet)
	 if time.Now().Unix()%2 == 0 {
	 	return fmt.Errorf("random error notifying %v", person)
	}
	if person.FirstName == "Tom" {
		panic(fmt.Sprintf("random panic notifying %v", person))
	}
	fmt.Println(person)
	return nil
}