package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const appName = "stdin-spinner"

var httpClient = http.DefaultClient

var config struct {
	Append       bool
	Frames       string
	RefreshEvery time.Duration
	Tee          string
	Width        int
}

func init() {
	config.RefreshEvery = time.Millisecond * 100
	config.Frames = "░▒▓█▓▒░"
	config.Width = 12
	log.SetFlags(0)
	log.SetOutput(os.Stderr)
	flag.StringVar(&config.Frames, "frames", config.Frames, "spinner animation frames")
	flag.StringVar(&config.Tee, "tee", config.Tee, "copy stdin to this file")
	flag.BoolVar(&config.Append, "a", config.Append, "(alias for -tee-append)")
	flag.BoolVar(&config.Append, "tee-append", config.Append, "append to the tee file (if specified)")
	flag.IntVar(&config.Width, "w", config.Width, "(alias for -width)")
	flag.IntVar(&config.Width, "width", config.Width, "spinner animation width")
	flag.DurationVar(&config.RefreshEvery, "r", config.RefreshEvery, "(alias for -refresh)")
	flag.DurationVar(&config.RefreshEvery, "refresh", config.RefreshEvery, "")
	flag.Parse()
}

func main() {
	var wg sync.WaitGroup
	var out io.Writer
	var count countWriter
	out = &count
	if config.Tee != "" {
		flags := os.O_CREATE | os.O_WRONLY
		if config.Append {
			flags = os.O_APPEND | os.O_WRONLY
		}
		f, err := os.OpenFile(config.Tee, flags, 0600)
		if err != nil {
			log.Fatalf("tee %s: %v", config.Tee, err)
		}
		defer f.Close()
		out = io.MultiWriter(&count, f)
	}
	done := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		fprintSpinner(os.Stderr, &count, done)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		io.Copy(out, os.Stdin)
		done <- struct{}{}
	}()
	wg.Wait()
}

func fprintSpinner(w io.Writer, c *countWriter, done <-chan struct{}) {
	previousCount := c.bytesWritten
	s := &spinner{
		frames: []rune(config.Frames),
		width:  config.Width,
	}
	ticker := time.NewTicker(config.RefreshEvery)
	first := make(chan struct{}, 1)
	first <- struct{}{}
	for {
		select {
		case <-done:
			fmt.Fprintf(w, "\r%s\r", strings.Repeat(" ", config.Width))
			return
		case <-first:
			fmt.Fprintf(w, "%s", s.bump())
		case <-ticker.C:
			if c.bytesWritten != previousCount {
				fmt.Fprintf(w, "\r%s", s.bump())
				previousCount = c.bytesWritten
			}
		}
	}
}
