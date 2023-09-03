package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 88,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 88,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Jack",
			LastName:  "Kevin",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 100,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 22,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 90,
				},
			},
		},
	}
}
