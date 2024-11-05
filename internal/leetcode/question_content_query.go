package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type questionContentRes struct {
	Data struct {
		Question struct {
			Content string `json:"content"`
		} `json:"question"`
	} `json:"data"`
}

func getQuestionContent(slug string) (questionContentRes, error) {
	query := fmt.Sprintf(`
    {
      "query":"query questionContent($titleSlug: String!) { question(titleSlug: $titleSlug) { content mysqlSchemas dataSchemas } }",
      "variables":{"titleSlug":"%s"},
      "operationName":"questionContent"
    }
    `,
		slug,
	)
	titleResp, err := http.Post(leetcodeAPI, "application/json", bytes.NewBufferString(query))
	if err != nil {
		return questionContentRes{}, err
	}

	if titleResp.StatusCode >= 400 {
		fmt.Fprintf(os.Stderr, "leetcode getContent query failed: %s\n", titleResp.Status)
		errBody, err := io.ReadAll(titleResp.Body)
		fmt.Fprintf(os.Stderr, string(errBody))
		return questionContentRes{}, err
	}
	defer titleResp.Body.Close()

	var response questionContentRes
	body, err := io.ReadAll(titleResp.Body)
	if err != nil {
		return questionContentRes{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return questionContentRes{}, err
	}

	return response, nil
}
