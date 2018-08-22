package task

import (
	"bufio"
	"io"
	"net/http"
	"strings"
)

type RequestTask struct {
	Method 	string 				`json:"method"`
	Url 	string				`json:"url"`
	Headers map[string][]string `json:"headers"`
	Body 	string 				`json:"body"`
}

type ResponseTask struct {
	Id 	uint64	`json:"id"`
	Url string	`json:"url"`
}

type Task struct {
	Id 				uint64				`json:"id"`
	StatusCode 		int					`json:"statusCode"`
	Headers 		map[string][]string	`json:"headers"`
	ContentLength	uint64				`json:"contentLength"`
}

func NewTask(requestTask RequestTask) (task Task, err error) {

	req, err := http.NewRequest(requestTask.Method, requestTask.Url, strings.NewReader(requestTask.Body))
	if err != nil {
		return
	}
	req.Header = requestTask.Headers

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	task.StatusCode = resp.StatusCode
	task.Headers = resp.Header

	reader := bufio.NewReader(resp.Body)
	p := make([]byte, 1024)
	for {
		n, errRead := reader.Read(p)
		if errRead == io.EOF {
			break
		}
		if errRead != nil {
			err = errRead
			return
		}
		task.ContentLength += uint64(n)
	}

	task.Id = StoreAdd(requestTask.Url)

	return

}
