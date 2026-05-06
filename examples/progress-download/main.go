package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"charm.land/bubbles/v2/progress"
	boba "github.com/btwiuse/boba"
)

var p *boba.Program

type progressWriter struct {
	total      int
	downloaded int
	reader     io.Reader
	onProgress func(float64)
}

func (pw *progressWriter) Start() {
	// TeeReader calls pw.Write() each time a new response is received
	_, err := io.Copy(io.Discard, io.TeeReader(pw.reader, pw))
	if err != nil {
		p.Send(progressErrMsg{err})
	}
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	pw.downloaded += len(p)
	if pw.total > 0 && pw.onProgress != nil {
		pw.onProgress(float64(pw.downloaded) / float64(pw.total))
	}
	return len(p), nil
}

func getResponse(url string) (*http.Response, error) {
	resp, err := http.Get(url) // nolint:gosec
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("receiving status of %d for url: %s", resp.StatusCode, url)
	}
	return resp, nil
}

var DEFAULT_URL = "https://no-cors.deno.dev/https://download.blender.org/demo/color_vortex.blend"

func main() {
	url := flag.String("url", DEFAULT_URL, "url for the file to download")
	flag.Parse()

	if *url == "" {
		flag.Usage()
		os.Exit(1)
	}

	resp, err := getResponse(*url)
	if err != nil {
		fmt.Println("could not get response", err)
		os.Exit(1)
	}
	defer resp.Body.Close() // nolint:errcheck

	// Don't add TUI if the header doesn't include content size
	// it's impossible see progress without total
	if resp.ContentLength <= 0 {
		fmt.Println("can't parse content length, aborting download")
		os.Exit(1)
	}

	pw := &progressWriter{
		total:  int(resp.ContentLength),
		reader: resp.Body,
		onProgress: func(ratio float64) {
			p.Send(progressMsg(ratio))
		},
	}

	m := model{
		pw:       pw,
		progress: progress.New(progress.WithDefaultBlend()),
	}
	// Start Bubble Tea
	p = boba.NewProgram(m)

	// Start the download
	go pw.Start()

	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
