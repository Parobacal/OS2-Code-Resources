package main

import (
	"fmt"
	//"time"
	"time"

	"github.com/gocolly/colly"
)

type trabajito struct {
	Url         string
	Referencias int
}

type resultado struct {
	ConteoPalabras int
	ConteoEnlaces  int
	Url            string
	Origen         string
	Sha            string
}

func main() {
	jobs := make(chan trabajito, 100)
	results := make(chan trabajito, 100)
	for i:= 0; i < 5; i ++ {

	//definimos nuestro scraper
		go worker(jobs, results)
	}
	//go worker(jobs, results)
	//go worker(jobs, results)
	//go worker(jobs, results)

	//for i := 0; i < 100; i++ {
	//    jobs <- i
	//    time.Sleep(1000)
	//}
	jobs <- trabajito{"https://es.wikipedia.org/wiki/Chuck_Norris", 2}
	//close(jobs)

	for r := range results {
		//fmt.Println(<-results)
		//t := <-results
		fmt.Println("visitando resultados")
		jobs <- r
	}
}

func worker(jobs <-chan trabajito, results chan<- trabajito) {

	for j := range jobs {
		Url := j.Url
		Nr := j.Referencias

		conteo := 0
		aux := ""

		c := colly.NewCollector()
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		c.OnHTML("div#mw-content-text p", func(e *colly.HTMLElement) {
			fmt.Println(e)
		})

		c.OnHTML("div#mw-content-text p a", func(e *colly.HTMLElement) {
			//fmt.Println(e.Attr("href"))
			if conteo < Nr {
				fmt.Println(e.Request.AbsoluteURL(e.Attr("href")))
				aux = e.Request.AbsoluteURL(e.Attr("href"))
				results <- trabajito{aux, Nr - 1}
				conteo = conteo + 1
			}
		})

		c.Visit(Url)
		time.Sleep(3000)
	}
}
