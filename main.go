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

var tasksToDo []task
var tasksDone []task

func newTask(name string, status bool) task {
	var t task
	t.name = name
	t.status = status
	return t
}

func createTask (taskname string){
	t := newTask(taskname, false)
	taskWriter(t)
	taskLoader(todoParser())
}


func taskWriter(task task){
	i, err := os.Open("./todos")
	defer i.Close()
	if err != nil {fmt.Println(err)}
	
	i.Write([]byte(fmt.Sprintf("%v", tasklist)) )
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
		t.status, _ = strconv.ParseBool(fields[2])

		tasks = append(tasks, t)
	}
	return tasks
	
}



func main(){
	taskLoader(todoParser())
	createTask("task1")
	// fmt.Println(tasks)
	fmt.Println("these are the done tasks: ", tasksDone)
	fmt.Println("thse are the tasks to do: ", tasksToDo)
}

