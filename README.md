# Text Search App

This project implements a simple user interface with a go backend that performs a full text search against a Zinc database similar to Elasticsearch.

## Backend and Scripts

### Local testing
* Download and unzip the mail file. 

Using docker-compose
* Create a `/data` folder and add the write permissions
```
mkdir data
chmod a+rwx ./data
```

Run the following command to create the containers with docker compose
```
docker-compose up -d
```

To upload the data in Zinc database:
1. Go to `backend/indexer` and run `go build`
2. Run the indexer to push the data to the database:
```
$ ./indexer --path /abs/path/to/data
```
Access to the UI with the url `http://localhost:8080`  
You will see the main page to test the mail search:



## Techologies used
* Go with chi router
* Vue.js
* Zinc DB
* Docker

# References
* Running zinc with docker: [Documentation](https://docs.zincsearch.com/installation/)
* Zinc official docker image: [AWS link](https://gallery.ecr.aws/zinclabs/zinc)
