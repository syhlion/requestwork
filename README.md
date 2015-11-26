# requestworker

a lib for go to batch processing send web request

## Install

`go get github.com/syhlion/requestwork`

### Usage

```

func main() {

    // Init Job
    j := &requestwork.Job{}
    resq, err := http.NewRequest("GET", "http://tw.yahoo.com", nil)
    if err != nil {
        panic(err)
    }
    j.Resq = resq
    result := make(chan string)
    j.Command = func(resp *http.Response, err error) {
        if err != nil {
            return
        }
        b, err := ioutil.ReadAll(resp.Body)
        result <- string(b)

    }

    // Init Worker
    w := &requestwork.Worker{
            JobQuene:   make(chan *requestwork.Job),
            Threads:    5,
            HttpClient: &http.Client{},
    }
    go w.Start()
    w.JobQuene <- j
    println(<-result)
    fmt.Println("end")
}

```
