package requestwork

import (
	"net/http"
)

type job struct {
	req *http.Request
	end chan result
}

type result struct {
	resp *http.Response
	err  error
}

type Worker struct {
	JobQuene   chan *job
	Threads    int
	HttpClient *http.Client
}

func (w *Worker) Excuste(req *http.Request) (resp *http.Response, err error) {
	j := &job{req, make(chan result)}

	w.JobQuene <- j
	r := <-j.end
	return r.resp, r.err

}

func (w *Worker) run() {
	for j := range w.JobQuene {
		resp, err := w.HttpClient.Do(j.req)
		j.end <- result{resp, err}
		close(j.end)
	}

}

func (w *Worker) Start() {

	for i := 0; i < w.Threads; i++ {
		go w.run()
	}

}
