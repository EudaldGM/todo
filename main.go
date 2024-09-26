package main

import (
	"fmt"
	"strconv"
	"os"
	"strings"
	"bufio"
)

type task struct{
	name string
	status bool //true == done false == not done
}

var tasks []task
var tasksToDo []task
var tasksDone []task

func newTask (taskname string){
	var t = task{name: taskname, status: false}
	tasks = append(tasks, t) 
	taskLoader(tasks)
}

func taskWriter(tasks []task){
	f, err := os.Create("./todos")
	if err != nil{fmt.Println(err)}
	defer f.Close()
	for _, t := range tasks{
		f.Write([]byte(fmt.Sprintf("%v, %v\n", t.name, t.status)))
	}
}

func taskLoader(tasklist []task){
	for _, task := range tasklist{
		if task.status {tasksDone = append(tasksDone, task)} 
	 	if !task.status {tasksToDo = append(tasksToDo, task)}
	}
}

func todoParser() []task {
	i, _ := os.Open("./todos")
	defer i.Close()
	
	scanner := bufio.NewScanner(i)
	var t task
	for scanner.Scan(){
		line := scanner.Text()
		fields := strings.Split(line, ",")
		for i := range fields {
					fields[i] = strings.TrimSpace(fields[i])
		}
		
		t.name = fields[0]
		t.status, _ = strconv.ParseBool(fields[1])

		tasks = append(tasks, t)
	}
	return tasks
	
}



func main(){
	tasks = todoParser() //parse file
	taskLoader(tasks) //organize lists/
	newTask("atante")
	fmt.Println("these are the done tasks: ", tasksDone)
	fmt.Println("thse are the tasks to do: ", tasksToDo)
	taskWriter(tasks)
}

