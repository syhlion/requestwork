package requestworker

import (
	"net/http"
)

type Job struct {
	Resq    *http.Request
	Command func(resp *http.Response, err error)
}

type Worker struct {
	JobQuene   chan *Job
	Threads    int
	HttpClient *http.Client
}

func (w *Worker) run() {
	for j := range w.JobQuene {
		res, err := w.HttpClient.Do(j.Resq)
		j.Command(res, err)
		res.Body.Close()
	}

}

func (w *Worker) Start() {

	for i := 0; i < w.Threads; i++ {
		go w.run()
	}

}
