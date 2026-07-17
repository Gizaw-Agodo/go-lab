package main

import "fmt"

func main() {

	// =========================================
	// Create University
	// =========================================

	university, err := newUniversity("Addis Ababa University")
	if err != nil {
		fmt.Println(err)
		return
	}

	// =========================================
	// Create Instructors
	// =========================================

	instructor1, err := newInstructor(
		1,
		"Dr. John",
		"Computer Science",
		"john@aau.edu",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	instructor2, err := newInstructor(
		2,
		"Dr. Sarah",
		"Software Engineering",
		"sarah@aau.edu",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	// =========================================
	// Add Instructors
	// =========================================

	err = university.addInstructor(*instructor1)
	if err != nil {
		fmt.Println(err)
	}

	err = university.addInstructor(*instructor2)
	if err != nil {
		fmt.Println(err)
	}

	// =========================================
	// Create Courses
	// =========================================

	course1, err := newCourse(
		"CS101",
		"Introduction to Programming",
		3,
		50,
		*instructor1,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	course2, err := newCourse(
		"SE201",
		"Software Engineering",
		4,
		40,
		*instructor2,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	// =========================================
	// Add Courses
	// =========================================

	university.addCourse(*course1)
	university.addCourse(*course2)

	// =========================================
	// Create Students
	// =========================================

	student1, err := newStudent(
		1,
		"Alice",
		20,
		"alice@example.com",
		"Computer Science",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	student2, err := newStudent(
		2,
		"Bob",
		22,
		"bob@example.com",
		"Software Engineering",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	// =========================================
	// Add Students
	// =========================================

	university.addStudent(*student1)
	university.addStudent(*student2)

	// =========================================
	// Display Students
	// =========================================

	university.displayStudents()

	// =========================================
	// Find Student
	// =========================================

	foundStudent, err := university.findStudent(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("Student Found:")
	foundStudent.Display()

	// =========================================
	// Update Student
	// =========================================

	err = foundStudent.updateEmail("alice2026@example.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("After Email Update:")
	foundStudent.Display()

	// =========================================
	// Create Enrollments
	// =========================================

	enrollment1, _ := newEnrollement("CS101", 1)
	enrollment2, _ := newEnrollement( "SE201", 2)

	university.Enrollments = append(
		university.Enrollments,
		*enrollment1,
		*enrollment2,
	)

	// =========================================
	// Assign Grades
	// =========================================

	enrollment1.assignGrade(92)
	enrollment2.assignGrade(78)

	fmt.Println()
	fmt.Println("Enrollments")

	enrollment1.Display()
	enrollment2.Display()

	// =========================================
	// Interfaces Demo
	// =========================================

	fmt.Println()
	fmt.Println("Using Displayable Interface")

	DisplayInformation(*course1)
	DisplayInformation(*instructor1)

	fmt.Println()
	fmt.Println("Program Finished Successfully")
}