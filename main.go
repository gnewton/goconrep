package main

import (
	"archive/tar"
	"bufio"
	"encoding/json"
	"fmt"
	core "github.com/gnewton/goconrep/corelib"
	"github.com/xi2/xz"
	"io"
	"log"
	"strings"
	"sync"
)

const numWorkers = 16

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var sizeTotal int64 = 0

	jsonChannel := make(chan string, 100)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(jsonChannel, &wg)
	}

	if true {
		//r, f, err := genericReader("data/core_small.tar.gz")
		r, f, err := genericReader("/home/gnewton/data/CORE/2018/core_2018-03-01_metadata.tar.gz")
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}

		tr := tar.NewReader(r)

		for {
			hdr, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			sizeTotal += hdr.Size / 1024 / 1024
			fmt.Println(hdr.Name, hdr.Size/1024/1024, "MB", "   ", sizeTotal, "MB")

			// Is a compressed json xz
			if strings.HasSuffix(hdr.Name, ".json.xz") {
				reader, err := xz.NewReader(tr, 0)
				if err != nil {
					log.Fatal(err)
				}

				scanner := bufio.NewScanner(reader)
				buf := make([]byte, 0, 4*1024*1024)
				scanner.Buffer(buf, 1024*1024*64)
				var n int64 = 0
				for scanner.Scan() {
					if true {
						//decodeJson(scanner.Text(), n)
						jsonChannel <- scanner.Text()
						n++
					}
				}

				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	close(jsonChannel)
	wg.Wait()
}

func worker(c chan string, wg *sync.WaitGroup) {
	go func() {
		for s := range c {
			decodeJson(s)
		}
		wg.Done()
	}()
}

func decodeJson(s string) {
	var doc core.Doc2018
	dec := json.NewDecoder(strings.NewReader(s))

	if err := dec.Decode(&doc); err == io.EOF {
		return
	} else if err != nil {
		log.Println(s)
		log.Fatal(err)
	}
	doc.RawRecordXML = ""

}
