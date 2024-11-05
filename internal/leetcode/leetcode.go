package leetcode

const leetcodeAPI = "https://leetcode.com/graphql/"

type Difficulty int

const (
	Easy Difficulty = iota + 1
	Medium
	Hard
)

type Question struct {
	Id         string // The leetcode integer ID
	Title      string
	TitleSlug  string
	Difficulty Difficulty
	Content    string
}

func GetQuestion(slug string) (Question, error) {
	titleResponse, err := getQuestionTitle(slug)

	if err != nil {
		return Question{}, err
	}
	contentResponse, err := getQuestionContent(slug)
	if err != nil {
		return Question{}, err
	}

	question := Question{
		Id:        titleResponse.Data.Question.QuestionFrontendID,
		Title:     titleResponse.Data.Question.Title,
		TitleSlug: titleResponse.Data.Question.TitleSlug,
		Content:   contentResponse.Data.Question.Content,
	}

	switch titleResponse.Data.Question.Difficulty {
	case "Easy":
		question.Difficulty = Easy
		break
	case "Medium":
		question.Difficulty = Medium
		break
	case "Hard":
		question.Difficulty = Hard
		break
	}

	return question, nil
}
