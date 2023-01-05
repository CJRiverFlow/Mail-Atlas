package constants

//Request params
const (
	ResultsFrom = 0
	ResultsMax  = 100
	UserAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"
)

//Query strings
const (
	QueryString = `{
        "search_type": "match",
        "query":
        {
            "term": "%s"
        },
        "from": %v,
        "max_results": %v,
        "sort_fields": ["-@timestamp"],
        "_source": []
    }`
)
