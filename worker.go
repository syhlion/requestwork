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

func New(httpClient *http.Client, threads int) *Worker {

	w := &Worker{
		jobQuene:   make(chan *job),
		threads:    threads,
		httpClient: httpClient,
	}
	go w.start()
	return w

}

type Worker struct {
	jobQuene   chan *job
	threads    int
	httpClient *http.Client
}

func (w *Worker) Execute(req *http.Request) (resp *http.Response, err error) {
	j := &job{req, make(chan result)}

	w.jobQuene <- j
	r := <-j.end
	return r.resp, r.err

}

func (w *Worker) run() {
	for j := range w.jobQuene {
		resp, err := w.httpClient.Do(j.req)
		j.end <- result{resp, err}
		close(j.end)
	}

}

func (w *Worker) start() {

	for i := 0; i < w.threads; i++ {
		go w.run()
	}

}
