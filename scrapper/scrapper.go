package scrapper

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type extractedJob struct {
	id      string
	name    string
	title   string
	endDate string
}

// Scrape saramin by a term
func Scrape(term string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?searchType=search&recruitSort=relation&recruitPageCount=40&searchword=" + term

	var jobs []extractedJob
	channel := make(chan []extractedJob)
	totalPages := getPages(baseURL)

	for i := 1; i <= totalPages; i++ {
		go getPage(i, baseURL, channel)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-channel
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int, url string, mainChannel chan<- []extractedJob) {
	var jobs []extractedJob
	channel := make(chan extractedJob)

	pageURL := url + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting page:", page, pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			checkErr(err)
		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	itemRecruit := doc.Find(".item_recruit")
	itemRecruit.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, channel)
	})

	for i := 0; i < itemRecruit.Length(); i++ {
		job := <-channel
		jobs = append(jobs, job)
	}

	mainChannel <- jobs
}

func extractJob(card *goquery.Selection, channel chan<- extractedJob) {
	id, _ := card.Attr("value")
	name := CleanString(card.Find(".corp_name").Text())
	title := CleanString(card.Find(".job_tit>a").Text())
	endDate := CleanString(card.Find(".job_date>span").Text())
	channel <- extractedJob{id: id, name: name, title: title, endDate: endDate}
}

func getPages(url string) int {
	pages := 0

	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			checkErr(err)
		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"URL", "Name", "Title", "End Date"}

	err = w.Write(headers)
	checkErr(err)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?rec_idx=" + job.id, job.name, job.title, job.endDate}
		err := w.Write(jobSlice)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatal("Request failed with Status:", res.StatusCode)
	}

}
