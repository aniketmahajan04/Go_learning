package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

/*
 A function for which failure is an expected behavior returns an additional result, conventonally the lastone.
 If the failure has only one possible cause, the result is a boolean, usually
 called ok,as in this example of a cache look up that always succeeds unless there was no entry
 for that key :
 value, ok := cache.Lookup(key)
 	if !ok {
	 // ... cahe[key] does not exist...
	}
*/

func main() {
	// 1.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//  Create a new error with combining two errors
	doc, err := html.Parse(resp.Body)
	resp.body.Closed()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// 3 only resevered for main function
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
		// more convenient way than Fprintf
		log.Fatalf("Site is down: %v\n", err)
	}

	// .5 in rare cases
	dir, err := ioutil.TempDir("", "scratch")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}

	// .4 in some cases
	if err := Ping(); err != nil {
		log.Printf("ping failed: %v; networking disabled", err)
		// or
		fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
	}
}

// 2
// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all atempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil //success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
