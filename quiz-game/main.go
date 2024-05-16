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
	csvfile := flag.String("csv", "problems.csv", "a csv file is in format of q/a type")
	timelimit:=flag.Int("limit",30,"the time limit for quiz")
	flag.Parse()

	file, err := os.Open(*csvfile)
	if err != nil {
		exit(fmt.Sprint("failed to open csv: %s\n", *csvfile))
	}
	r :=csv.NewReader(file)
	lines,err := r.ReadAll()
	if err !=nil{
		exit("failed to parse the provided csv file.")
	}
	
	problems :=parseline(lines)

	timer := time.NewTimer(time.Duration(*timelimit) *time.Second)
	
	correct := 0
	for i,p:= range problems{
		fmt.Printf("problem #%d: %s = \n", i+1,p.question)
		answerch:=make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerch<- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("you scored %d out of %d.\n",correct,len(problems))
			return
		case answer := <-answerch:
			if answer == p.answer{
				correct++
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
	return returns
}
type problem struct {
	question string
	answer string
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
