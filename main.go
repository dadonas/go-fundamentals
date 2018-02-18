package main

//Test feature/new-go-feature
//feature2
//feature3
//Hotfix

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/go-fundamentals/model"
)

var (
	myCoursesArray = []string{"Java", "Golang", "Python", "NodeJS"}
)

func main() {

	var waitGroup sync.WaitGroup
	//goroutine expects 2 works to get done
	waitGroup.Add(2)

	printVarByValueAndByReference()
	printArray()
	printSlices()
	printAppendCourse()
	printMapOfCourses()
	printStructExample()

	performObjectOrientedExample()

	performVariaticFunc("Hello ", "variatic ", "function!!!", "\n")

	if valsCount, sum, err := performMultipleReturns(1, 2, 3); err == nil {
		fmt.Println("Values given:", valsCount, "Sum:", sum)
	}

	if _, _, err := performMultipleReturns(); err != nil {
		fmt.Println(err)
	}

	//Passing params as reference
	referenceMessage := "Reference message before change"
	fmt.Println(referenceMessage)
	func(message *string) {
		*message = "Reference message after change"
	}(&referenceMessage)
	fmt.Println(referenceMessage)

	//Goroutines
	go func() {
		//tells goroutine that work is done
		defer waitGroup.Done()

		fmt.Println("Executing routine 1 - before sleep...")
		time.Sleep(5 * time.Second)
		fmt.Println("Executing routine 1 - after sleep.")
	}()

	go func() {
		//tells goroutine that work is done
		defer waitGroup.Done()

		fmt.Println("Executing routine 2")
	}()

	//Go routine will force main to wait for 2 routines to get done
	waitGroup.Wait()

}

func performVariaticFunc(messages ...string) {
	for _, message := range messages {
		fmt.Print(message)
	}
}

func performMultipleReturns(values ...int) (valuesGiven int, sum int, err error) {
	if values == nil {
		err = errors.New("Informe algum valor")

		return
	}

	valuesGiven = len(values)

	for _, value := range values {
		sum += value
	}

	return
}

func printVarByValueAndByReference() {
	name := "Dadonas"
	refName := &name

	fmt.Println("Hello", *refName, "from memory address", refName, "and OS", runtime.GOOS)
}

func printArray() {
	theArray := []int{10, 20, 30, 40, 50}

	fmt.Println("This is the entire array:", theArray)
}

func printSlices() {
	myFirstCoursesSlice := myCoursesArray[:2]
	myUltimateCoursesSlice := myCoursesArray[2:]
	myLastCourseSlice := myCoursesArray[len(myCoursesArray)-1]

	fmt.Println("First Courses ->", myFirstCoursesSlice)
	fmt.Println("Ultimate Courses ->", myUltimateCoursesSlice)
	fmt.Println("Last Course ->", myLastCourseSlice)
}

func printAppendCourse() {
	myCoursesArray = append(myCoursesArray, "TDD")
	myLastCourseSlice := myCoursesArray[len(myCoursesArray)-1]

	fmt.Println("Last Course ->", myLastCourseSlice)
}

func printMapOfCourses() {
	/*
		coursesAndGrades := make(map[string]int)
		coursesAndGrades["Java"] = 100
		coursesAndGrades["Golang"] = 25
	*/

	coursesAndGrades := map[string]int{
		"Java":   100,
		"Golang": 25,
		"Python": 0,
		"NodeJS": 0,
	}

	//Delete Item
	//delete(coursesAndGrades, "NodeJS")

	for course, grade := range coursesAndGrades {
		fmt.Println(course, grade)
	}

}

func printStructExample() {
	type Course struct {
		Name       string
		University string
	}

	course := Course{
		Name:       "Java Programming",
		University: "Yale",
	}

	fmt.Println(course)
}

func performObjectOrientedExample() {
	var option model.PaymentOption

	option = model.CreateCash("", 10000.)
	if _, err := option.ProcessPayment(11000.); err != nil {
		fmt.Println(err)
	}

	option = model.CreateCreditCard("", 20000.)
	if _, err := option.ProcessPayment(21000.); err != nil {
		fmt.Println(err)
	}
}
