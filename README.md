# Go Jobs Scraper

> **This document was generated with the help of GitHub Copilot.**  
> **This project was built as a practice exercise following a course from Nomad Coders.**

This project is a web application that scrapes job listings from `saramin.co.kr` based on a search term provided by the user.
The application is built using Go and the Echo web framework.

<br>

## Features

- Scrapes job listings from `saramin.co.kr`
- Saves the scraped job listings to a CSV file
- Provides a simple web interface for users to input search terms

<br>

## Prerequisites

- Go 1.16 or later
- Go modules enabled`

<br>

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/zzzzseong/go-jobs-scrapper.git
    ```

2. Install dependencies:
    ```sh
   go get github.com/PuerkitoBio/goquery
   go get github.com/labstack/echo
    ```

<br>

## Usage

1. Run the application:
    ```sh
    go run main.go
    ```

2. Open your web browser and navigate to `http://localhost:8080`.

3. Enter a job search term and click "Search".

4. The application will scrape job listings and prompt you to download a CSV file with the results.

<br>

## Project Structure

- `main.go`: Entry point of the application.
- `home.html`: HTML template for the web interface.
- `scrapper/scrapper.go`: Contains the logic for scraping job listings from `saramin.co.kr`.

<br>

## Reference

- [nomad-coders](https://nomadcoders.co/go-for-beginners): 쉽고 빠른 Go 시작하기
