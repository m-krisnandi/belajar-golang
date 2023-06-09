package main

var students = []*Student{}

type Student struct {
	Id string
	Name string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, student := range students {
		if student.Id == id {
			return student
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "John", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "Paul", Grade: 3})
	students = append(students, &Student{Id: "s003", Name: "George", Grade: 3})
}