package requestwork

import (
	"net/http"
)

type Job struct {
	Req     *http.Request
	Command func(resp *http.Response, err error)
}

type Worker struct {
	JobQuene   chan *Job
	Threads    int
	HttpClient *http.Client
}

func (w *Worker) run() {
	for j := range w.JobQuene {
		res, err := w.HttpClient.Do(j.Req)
		j.Command(res, err)
		if err != nil {
			continue
		}
		res.Body.Close()
	}

}

func (w *Worker) Start() {

	for i := 0; i < w.Threads; i++ {
		go w.run()
	}

}
