package main

import (
	"fmt"
	"time"
	"strconv"
	"os"
	"strings"
	"bufio"
)

type task struct{
	name string
	deadline time.Time
	status bool //true == done false == not done
	ood bool //is not done and past the deadline
}

var tasksToDo []task
var tasksDone []task

func newTask(name string, deadline time.Time, status bool) task {
	var t task
	t.name = name
	t.deadline = deadline
	t.status = status
	t.ood = false
	return t
}

func taskLoader(tasklist []task){
	for _, task := range tasklist{
		if task.status {tasksDone = append(tasksDone, task)} 
		if !task.status && task.deadline.Compare(time.Now()) < 1 {task.ood = true}
	 	if !task.status {tasksToDo = append(tasksToDo, task)}
	}
}

func todoParser() []task {
	i, _ := os.Open("./todos")
	defer i.Close()
	
	var tasks []task
	
	scanner := bufio.NewScanner(i)
	var t task
	for scanner.Scan(){
		line := scanner.Text()
		fields := strings.Split(line, ",")
		for i := range fields {
					fields[i] = strings.TrimSpace(fields[i])
		}
		
		t.name = fields[0]
		t.deadline, _ = time.Parse("2006-01-02", "2023-02-03")
		t.status, _ = strconv.ParseBool(fields[2])
		t.ood, _ = strconv.ParseBool(fields[3])

		tasks = append(tasks, t)
	}
	return tasks
	
}

func createTask (taskname string, deadline string){
	dl, err := time.Parse("2006-01-02", deadline)
	if err != nil {fmt.Println(err)}

	newTask(taskname, dl, false)
	//writetofile
	taskLoader(todoParser())
}

func main(){
	taskLoader(todoParser())
	
	// fmt.Println(tasks)
	fmt.Println("these are the done tasks: ", tasksDone)
	fmt.Println("thse are the tasks to do: ", tasksToDo)
}

