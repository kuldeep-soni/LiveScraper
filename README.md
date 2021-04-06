# LiveScraper
Service that will make a background request, fetch the respective web site, parse it and give back a valid result to the client

To run this service on your local system follow the below steps in command line:
1. Clone this repo in your current working directory
```
cd LiveScraper
go get -v -t ./...
go run app/website-scraping/cmd/main.go 
```
A local server should successfully spawn up

Hit the following url in your browser with valid amazon_id : http://localhost:8080/movie/amazon/{amazon_id}
