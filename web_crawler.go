package main


import (
    "fmt"
    "sync"
    )
    
func Serial(url string, fetcher Fetcher, fetched map[string]bool) {
    
    if fetched[url] {
        return
    }
    fetched[url]  = true
    urls, err := fetcher.Fetch(url)
    if err != nil {
        return
    }
    for _, u := range urls {
        Serial(u, fetcher, fetched)
    }
    
    return
}

func ConcurrentMutex(url string, fetcher Fetcher, f *fetchState) {
    
    f.mu.Lock()
    already := f.fetched[url]
    f.fetched[url] = true
    f.mu.Unlock()
    if already {
        return
    }
    urls, err := fetcher.Fetch(url)
    if err != nil {
        return 
    }
    var done sync.WaitGroup
    for _, u := range urls {
        done.Add(1)
        go func(u string) {
            defer done.Done()
            ConcurrentMutex(u, fetcher, f)
        }(u)
    done.Wait()
        
    }
    
    return
    
    
}

func makeState() *fetchState{
    f := &fetchState{}
    f.fetched = make(map[string]bool)
    return f
    
}


type fetchState struct {
    mu sync.Mutex
    fetched map[string]bool
}

func Worker(url string, ch chan []string, fetcher Fetcher) {
    urls, err := fetcher.Fetch(url)
    if err != nil {
        ch <- []string{}
    }else {
        ch <- urls
    }
    
}
func Coordinator(ch chan []string, fetchr Fetcher) {
    n := 1
    fetched := make(map[string]bool)
    for urls := range ch {
        for _, url := range urls {
            if fetched[url] == false {
                fetched[url] = true
                n += 1
                go Worker(url, ch, fetcher)
            }
        }
        n -= 1
        if n == 0 {
            break
        }
        
    }
    return
}

func ConcurrentChannel(u string, fecher Fetcher) {
    ch := make(chan []string)
    go func(){
        ch <- []string{u}
    }()
    Coordinator(ch, fetcher)
}

func main() {
    
    fmt.Printf("=== Serial===\n")
	Serial("http://golang.org/", fetcher, make(map[string]bool))
	
	fmt.Printf("===== Concurrent Mutex =====\n")
	ConcurrentMutex("http://golang.org/", fetcher, makeState())
	
	fmt.Printf("==== Concurrent Channel ====\n")
	ConcurrentChannel("http://golang.org/", fetcher)
	
    
}

type Fetcher interface {
    Fetch(url string) (urls []string, err error)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct{
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) ([]string, error) {
    if res, ok := f[url] ; ok {
        fmt.Printf("found %s \n", url)
        return res.urls, nil
    }
    fmt.Printf("Missing, Not Found %s \n", url)
    return nil, fmt.Errorf("not found %s", url)
}

var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

