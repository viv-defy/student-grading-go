package student_grading

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	students := []student{}
	scanner := bufio.NewScanner(f)
	isHeader := true
	for scanner.Scan() {
		studentRecord := scanner.Text()
		if isHeader {
			isHeader = false
			continue
		}

		tokens := strings.Split(studentRecord, ",")

		testScores := []int{}
		for i, v := range tokens {
			if i > 2 {
				score, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				testScores = append(testScores, score)
			}
		}

		students = append(students, student{
			firstName:  tokens[0],
			lastName:   tokens[1],
			university: tokens[2],
			test1Score: testScores[0],
			test2Score: testScores[1],
			test3Score: testScores[2],
			test4Score: testScores[3],
		})
	}

	return students
}

func calculateGrade(students []student) []studentStat {
	studentStats := []studentStat{}
	for _, student := range students {
		score := float32(student.test1Score+student.test2Score+student.test3Score+student.test4Score) / float32(4)
		var grade Grade
		switch {
		case score < 35:
			grade = F
		case score >= 35 && score < 50:
			grade = C
		case score >= 50 && score < 70:
			grade = B
		case score >= 70:
			grade = A
		}

		studentStats = append(studentStats, studentStat{
			student:    student,
			finalScore: float32(score),
			grade:      grade,
		})
	}

	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var topper studentStat
	for _, gradedStudent := range gradedStudents {
		if gradedStudent.finalScore > topper.finalScore {
			topper = gradedStudent
		}
	}

	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	universityTopper := make(map[string]studentStat)
	for _, gradedStudent := range gs {
		if topper, ok := universityTopper[gradedStudent.university]; ok {
			if gradedStudent.finalScore > topper.finalScore {
				universityTopper[gradedStudent.university] = gradedStudent
			}
		} else {
			universityTopper[gradedStudent.university] = gradedStudent
		}
	}
	return universityTopper
}
