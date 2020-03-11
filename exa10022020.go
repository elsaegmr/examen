package main


import (
	"bufio"                                                
	"fmt"                              
	"os" 
	"io/ioutil"                       
)

type Student struct {
	name, note, note2, note3 string
}


func main() {
	menu := `what do you want to do?
[1] -- Add
[2] -- Show
[3] -- Exit
----->	`

	var eleccion int
	
	for eleccion != 3 {
		fmt.Print(menu)
		fmt.Scanln(&eleccion)
		scanner := bufio.NewScanner(os.Stdin)
		switch eleccion {
		case 1:
			var students []Student
			for i:= 0; i<5; i++ {
			var student Student
			fmt.Println("Write name:")
			if scanner.Scan() {
				student.name = scanner.Text()
			}
			fmt.Println("Write first note:")
			if scanner.Scan() {
				student.note = scanner.Text()
			}

			fmt.Println("Write second note:")
			if scanner.Scan() {
				student.note2 = scanner.Text()
			}

			fmt.Println("Write third note:")
			if scanner.Scan() {
				student.note3 = scanner.Text()
			}
		
			students = append(students,student)
		}
			
		for i:= 0; i<len(students); i++ {
			Adddata(students[i].name, students[i].note, students[i].note2, students[i].note3)
		}
		
		case 2:
			
			data, err := ioutil.ReadFile("text.txt")
    		if err != nil {
        		fmt.Println(err)
    		}

    		fmt.Println("The file contains \n " + string(data))	
		}
	}
}

func Adddata (name string, note string, note2 string, note3 string) {
	f, err := os.OpenFile("text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) 
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(name+" "+note+" "+ note2+" "+note3+"\n"); err != nil {
		fmt.Println(err)
	}

}