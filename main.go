package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
) 


var db *sql.DB


type Student struct {
	ID int64
	name string
	level int64
}

type Course struct {
	code string
	title string
}


func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:  "root",
        Passwd: "engmso14789",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "GoSchool",
    }
    // Get a database handle.
    var err error

    db, err = sql.Open("mysql", cfg.FormatDSN())
	
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    // fmt.Println("Connected!")

	// Create Students table --- 

	db.Exec("CREATE TABLE IF NOT EXISTS students(id int auto_increment primary key,name varchar(255) NOT NULL,level int)")
	db.Exec("CREATE TABLE IF NOT EXISTS courses(code varchar(50) primary key, title varchar(255))")
	
	// UI

	println("Welcome!")
	println("1: Display All Students\n2: Display All Courses\n3: Add new student\n4: Add new course")

	var choice string

	fmt.Scanln(&choice)

	switch choice {
	case "1" :
		s,err := allStudents()
		if err != nil {
			println("Something went wrong, try again")
		}
		fmt.Println(s)
		fmt.Println("------------")
		break
	case "2" :
		c,err := allCourses()
		if err != nil {
			println("Something went wrong, try again")
		}
		fmt.Println(c)
		fmt.Println("------------")
		break

	case "3" : 
		var name string 
		var level int

		println("Enter Student's name : ")
		fmt.Scanln(&name)
		println("Enter Student's Level : ")
		fmt.Scanln(&level)

		addStudent(name,level)
		break
	case "4" : 
		var code string 
		var title string

		println("Enter Course Code : ")
		fmt.Scanln(&code)
		println("Enter Course Title : ")
		fmt.Scanln(&title)

		addCourse(code,title)
		break	
	
	default : 
		println("Invalid Input :(")
	}
	
	

}




func addStudent(name string, level int)(int64,error){
	 result, err := db.Exec("INSERT INTO students (name, level) VALUES (?,?)",name,level)

	 if err != nil {
		return 0,fmt.Errorf("Student : %v",err)
	 }

	 id , err := result.LastInsertId()

	if err != nil {
		return 0,fmt.Errorf("Student : %v",err)
	}

	return id, nil

}

func addCourse(code string, title string)(int64,error){
	 result, err := db.Exec("INSERT INTO courses (code, title) VALUES (?,?)",code,title)

	 if err != nil {
		return 0,fmt.Errorf("course : %v",err)
	 }

	 id , err := result.LastInsertId()

	if err != nil {
		return 0,fmt.Errorf("course : %v",err)
	}

	return id, nil

}

func allStudents()([]Student, error){

	var students []Student

	rows, err := db.Query("select * from students")

	if err!=nil{
		return students, fmt.Errorf("Students : %v", err)
	}

	defer rows.Close()

	for rows.Next(){
		var std Student

		rows.Scan(&std.ID, &std.name, &std.level)

		students = append(students, std)
	}

	return students, nil
}

func allCourses()([]Course, error){

	var courses []Course

	rows, err := db.Query("select * from courses")

	if err!=nil{
		return courses, fmt.Errorf("courses : %v", err)
	}

	defer rows.Close()

	for rows.Next(){
		var crs Course

		rows.Scan(&crs.code, &crs.title)

		courses = append(courses, crs)
	}

	return courses, nil
}
