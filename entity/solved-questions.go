package entity

// All of this are capatilised so that we can access them at other locations and is serialized
type QuestionsDetails struct {
	ProblemCategory               string `json:"problem_category" bson:"problem_category"` //binding:"required" can be used for validation as well
	ProblemLevel                  string `json:"problem_level" bson:"problem_level"`
	QuestionDescription           string `json:"question_description" bson:"question_description"`
	QuestionLink                  string `json:"question_link" bson:"question_link"`
	BruteForceSolutionDescription string `json:"brute_force_solution_description" bson:"brute_force_solution_description"`
	BruteForceSolution            string `json:"brute_force_solution" bson:"brute_force_solution"`
	BruteForceSolutionTc          string `json:"brute_force_solution_tc" bson:"brute_force_solution_tc"`
	BruteForceSolutionSc          string `json:"brute_force_solution_sc" bson:"brute_force_solution_sc"`
	OptimalSolutionDescription    string `json:"optimal_solution_description" bson:"optimal_solution_description"`
	OptimalSolution               string `json:"optimal_solution" bson:"optimal_solution"`
	OptimalSolutionTc             string `json:"optimal_solution_tc" bson:"optimal_solution_tc"`
	OptimalSolutionSc             string `json:"optimal_solution_sc" bson:"optimal_solution_sc"`
	IsDeleted                     bool   `json:"is_deleted" bson:"is_deleted"`
}
