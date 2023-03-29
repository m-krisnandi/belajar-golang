package main

import "fmt"

type person struct {
	name string
	age int
}

type student struct {
	name string
	grade int
	age int
	// embedded struct
	person
}

func main() {
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	// inisisalisasi object struct
	var s2 = student{"ethan hunt", 2, 22, person{"ethan", 21}}
	
	var s3 = student{name: "jason"}

	var s4 = student{}
	s4.name = "wick"
	s4.grade = 2

	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)
	fmt.Println("student 4 :", s4.name)

	// variable object pointer
	var s5 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 5, name :", s5.name)

	s5.name = "ethan hunt"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 5, name :", s5.name)

	// embedded struct
	var s6 = student{}
	s6.name = "wick"
	s6.age = 21
	s6.person.age = 22
	s6.grade = 2

	fmt.Println("name :", s6.name)
	fmt.Println("age :", s6.age)
	fmt.Println("person age :", s6.person.age)
	fmt.Println("grade :", s6.grade)

	// pengisisan sub struct
	var p1 = person{name: "wick", age: 21}
	var s7 = student{person: p1, grade: 2}

	fmt.Println("name :", s7.person.name)
	fmt.Println("age :", s7.person.age)
	fmt.Println("person age :", s7.grade)

	// anonymous struct
	var s8 = struct {
		person
		grade int
	}{}

	// anonymous struct dengan pengisian property
	// {
	// 	person: person{name: "wick", age: 21},
	// 	grade: 2,
	// }

	s8.person = person{"wick", 21}
	s8.grade = 2

	fmt.Println("name :", s8.person.name)
	fmt.Println("age :", s8.person.age)
	fmt.Println("grade :", s8.grade)

	// kombinasi slice dan struct
	var allStudents = []person {
		{name: "wick", age: 21},
		{name: "ethan", age: 22},
		{name: "jason", age: 23},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	// slice anonymous struct
	var allStudents2 = []struct {
		person
		grade int
	} {
		{person: person{name: "wick", age: 21}, grade: 2},
		{person: person{name: "ethan", age: 22}, grade: 2},
		{person: person{name: "jason", age: 23}, grade: 2},
	}

	for _, student := range allStudents2 {
		fmt.Println(student.name, "age is", student.age, "and grade is", student.grade)
	}

	// deklarasi struct dengan keyword var
	var student struct {
		person
		grade int
	}

	// deklarasi sekaligus inisialisasi
	var student2 = struct {
		grade int
	} {
		grade: 2,
	}

	student.person = person{"wick", 21}
	student.grade = 2

	fmt.Println("name :", student.person.name)
	fmt.Println("age :", student.person.age)
	fmt.Println("grade :", student.grade)

	fmt.Println("grade :", student2.grade)

	// nested struct
	type student3 struct {
		person struct {
			name string
			age int
		}
		grade int
		hobbies []string
	}

	var s9 = student3{}
	s9.person.name = "wick"
	s9.person.age = 21
	s9.grade = 2
	s9.hobbies = []string {"Playing video games", "Watching movies", "Playing football"}

	fmt.Println("name :", s9.person.name)
	fmt.Println("age :", s9.person.age)
	fmt.Println("grade :", s9.grade)
	fmt.Println("hobbies :", s9.hobbies)

	// tag pada struct
	type person2 struct {
		name string `required max:"10"` 
		age int `required min:"18" max:"30"`
	}

	// type alias
	type Person struct {
		name string
		age int
	}
	type People = Person // type People adalah alias dari Person

	var p2 = Person{"wick", 21}
	fmt.Println(p2)

	var p3 = People{"wick", 21}
	fmt.Println(p3)

	people := People{"wick", 21}
	fmt.Println(Person(people))

	people2 := Person{"wick", 21}
	fmt.Println(People(people2))

	// struct literal
	type People1 = struct {
		name string
		age int
	}

	type People2 struct {
		name string
		age int
	}

	// contoh alias pada type data
	type Number = int
	var num Number = 12
	fmt.Println(num)

}