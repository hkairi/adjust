package fetcher

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Result struct {
	Url  string
	Hash string
}

type fetcher struct {
	Limit  int
	Urls   []string
	Client http.Client
	hasher hash.Hash
}

func New(limit int, urls []string) *fetcher {
	return &fetcher{
		Limit:  limit,
		Urls:   urls,
		hasher: md5.New(),
	}
}

func (f *fetcher) Start() {
	var wg sync.WaitGroup

	urlsCh := make(chan string, len(f.Urls))
	for _, url := range f.Urls {
		urlsCh <- url
	}
	close(urlsCh)

	wg.Add(len(f.Urls))

	for i := 0; i < f.Limit; i++ {
		go func() {
			for url := range urlsCh {
				go f.process(url, &wg)
			}
		}()
	}
	wg.Wait()
}

func (f *fetcher) HashText(s string) string {
	f.hasher.Write([]byte(s))
	hashedUrl := hex.EncodeToString(f.hasher.Sum(nil))

	return hashedUrl
}

func (f *fetcher) process(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	var full_url string = fmt.Sprintf("http://%s", url)
	resp, err := f.Client.Get(full_url)
	if err != nil {
		log.Fatalln(url, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(fmt.Errorf("Reading body! %v", err))
	}
	defer resp.Body.Close()

	hashedUrl := f.HashText(string(body))

	fmt.Printf("%-40s %s \n", url, hashedUrl)
}
