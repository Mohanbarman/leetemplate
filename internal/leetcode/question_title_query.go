package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type questionTitleRes struct {
	Data struct {
		Question struct {
			QuestionID         string `json:"questionId"`
			QuestionFrontendID string `json:"questionFrontendId"`
			Title              string `json:"title"`
			TitleSlug          string `json:"titleSlug"`
			IsPaidOnly         bool   `json:"isPaidOnly"`
			Difficulty         string `json:"difficulty"`
			Likes              int    `json:"likes"`
			Dislikes           int    `json:"dislikes"`
			CategoryTitle      string `json:"categoryTitle"`
		} `json:"question"`
	} `json:"data"`
}

func getQuestionTitle(slug string) (questionTitleRes, error) {
	titleQuery := fmt.Sprintf(`
    {
     "query": "query questionTitle($titleSlug: String!) { question(titleSlug: $titleSlug) {questionId questionFrontendId title titleSlug isPaidOnly difficulty likes dislikes categoryTitle }}",
     "variables": {"titleSlug":"%s"},
     "operationName":"questionTitle"
    }
    `,
		slug,
	)
	titleResp, err := http.Post(leetcodeAPI, "application/json", bytes.NewBufferString(titleQuery))
	if err != nil {
		return questionTitleRes{}, err
	}

	if titleResp.StatusCode >= 400 {
		fmt.Fprintf(os.Stderr, "leetcode getTitle query failed: %s\n", titleResp.Status)
		errBody, err := io.ReadAll(titleResp.Body)
		fmt.Fprintf(os.Stderr, string(errBody))
		return questionTitleRes{}, err
	}
	defer titleResp.Body.Close()

	var titleResponse questionTitleRes
	titleBody, err := io.ReadAll(titleResp.Body)
	if err != nil {
		return questionTitleRes{}, err
	}

	err = json.Unmarshal(titleBody, &titleResponse)
	if err != nil {
		return questionTitleRes{}, err
	}

	return titleResponse, nil
}
