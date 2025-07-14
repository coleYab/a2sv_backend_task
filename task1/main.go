package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name     string
	Subjects map[string]float64
}

func New(name string) *Student {
	return &Student{
		Name:     name,
		Subjects: map[string]float64{},
	}
}

func (s *Student) addSubject() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter subject name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter your score: ")
	scoreS, _ := reader.ReadString('\n')
	score, err := strconv.ParseFloat(strings.TrimSpace(scoreS), 64)
	if err != nil {
		fmt.Println("ERROR: please enter a valid score")
		return
	}

	if score < 0 || score > 100 {
		fmt.Println("ERROR: your score is out of valid range(0 - 4)")
		return
	}

	s.Subjects[name] = score
	fmt.Println("Subject added succesfully")
}

func (s *Student) calculateAverge() {
	var tot float64 = float64(0)
	count := 0
	for _, score := range s.Subjects {
		tot += score
		count += 1
	}

	if count == 0 {
		fmt.Println("ERROR: you have no subject added to calculate the average")
		return
	}

	avg := tot / float64(count)
	fmt.Printf("The average score is: %v\n", avg)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("====    Welcome to lets calculate your average =====")
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	student := New(name)

	for {
		fmt.Println("Main Menu:")
		fmt.Println("1. Add Subject")
		fmt.Println("2. Calculate Average")
		fmt.Println("0. Exit")
		fmt.Print("Enter your choice: ")
		choiceS, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(choiceS))
		if err != nil {
			fmt.Println("ERROR: please enter a valid choice")
			continue
		}

		switch choice {
		case 1:
			student.addSubject()
		case 2:
			student.calculateAverge()
		case 0:
			{
				fmt.Println("GoodBye")
				os.Exit(0)
			}
		default:
			{
				fmt.Println("ERROR: invalid chioce")
				continue
			}
		}
	}
}
