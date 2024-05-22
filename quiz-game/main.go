package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/csv"
	"strings"
	"time"

)

func main() {
	csvfile := flag.String("csv", "problems.csv", "a csv file is in format of q/a type") //initializing csvfile
	timelimit:=flag.Int("limit",10,"the time limit for quiz")
	fmt.Println("time limit 10 sec")
	flag.Parse()

	file, err := os.Open(*csvfile) //here 
	if err != nil {   //if error is not nil it will exit 
		exit(fmt.Sprint("failed to open csv: %s\n", *csvfile))
	} //if error is nill it continues the program excution
	r :=csv.NewReader(file) 
	lines,err := r.ReadAll()
	if err !=nil{
		exit("failed to parse the provided csv file.")
	}
	
	problems :=parseline(lines)

	timer := time.NewTimer(time.Duration(*timelimit) *time.Second) //time limit in game
	
	correct := 0
	for a,p:= range problems{
		fmt.Printf("problem #%d: %s = \n", a+1,p.question) // p is initializing as problem 
		answerch:=make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerch<- answer //assign a value to a variable through channel
		}()
		select {
		case <-timer.C:
			fmt.Printf("you scored %d out of %d.\n",correct,len(problems))
			return
		case answer := <-answerch:
			if answer == p.answer{
				correct++ //if coorect answer given by user it 
			}
		}
	}
	fmt.Printf("you scored %d out of %d.\n",correct,len(problems))

}
func parseline(lines [][]string) []problem {
	returns:= make([]problem, len(lines))
	for i,lines:= range lines{
		returns[i] = problem{
			question:lines[0],
			answer: strings.TrimSpace(lines[1]),
		}
	}
	return returns //return the value to returns(variable)
}
type problem struct { 
	question string
	answer string
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
