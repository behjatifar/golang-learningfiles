package main

import "fmt"

var globalscope = 124

func main() {
	job := "Senior Software Engineer"
	var firstName string = "amirmahdi"
	var lastName = "Behjatifar"
	const Country = "Iran"
	const City = "Neyshabur"
	const myWebsite = "https://Behjatifar.com/"
	// var i5, f3, x2, f6, h4 = 55, 65, "Arya", "AmirMahdi", 23
	var (
		Experince  = "5years"
		Skills     = "Golang,Docker,Nodejs,Nestjs"
		University = "Oxford UK Master Degreee"
	)
	{
		// Points to Global Scope but after that we declare it again so it makes missunderstanding
		// in go we behave to variables with scope that we have access to them
		fmt.Println(globalscope)
		globalscope = 125
		fmt.Println(globalscope)
	}
	{
		localvariable := "hello neighbors"
		fmt.Println(localvariable)
	}
	// thats why we don't access to localvariable here because it declare in scope and out of it we can't having access

	age := 20

	fmt.Printf("Firstname: %s Lastname: %s age: %d job: %s \n", firstName, lastName, age, job)
	fmt.Printf("Experince: %s Skills: %s Academic Edu: %s \n", Experince, Skills, University)
	fmt.Printf("borned Country: %s MotherCity: %s  Website: %s ", Country, City, myWebsite)
}
