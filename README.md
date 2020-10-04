# Gopher translator service

# Lint it
```
$ cd /where/go/src/is/github.com/mlinuxgada/gophertranslator
$ goling ./...
$
```

No lint msgs must appear, eg all is good

# Run

Start on port 1234:
```
$ go run main.go  --port=1234
INFO[0000] starting server on: localhost:1234
...
```

If no port defined, will start on 8080!

# Requests:

#### Word translator

Word translator translates in gopher lang defined words.
Example:
```
POST: localhost:1234/word
```

body:
```json
{
    "english-word": "apple"
}
```

response: 200 Ok

response body:
```json
{
    "gopher-word": "gapple"
}
```

If empty word is submitted, OR struct is not correct, we receive generic err response:

request body:
```json
{
    "english-word": ""
}
```

response: 422 Unprocessable entity
response body:
```json
{
    "status": 422,
    "detail": "Given word is empty string"
}
```

#### 
}


```json
{
    "status": 422,
    "detail": "Given word is empty string"
}
```

####  Sentence translator


Sentence translator translates in gopher lang defined sentence.
Example:
```
POST: localhost:1234/sentence
```

body:
```json
{
    "english-sentence": "Apple xray chair square?"
}
```

response: 200 Ok

response body:
```json
{
    "gopher-sentence": "Gapple gexray airchogo aresquogo?"
}
```

If empty sentence is submitted, OR struct is not correct, we receive generic err response:

request body:
```json
{
    "english-sentence": ""
}
```

response: 422 Unprocessable entity
response body:
```json
{
    "status": 422,
    "detail": "Given word is empty string"
}
```

#### History 

History controller is able to show all the history, eg all the words and sentences and their translated versions, ordered aphabetically.

Initial req/resp /no translation calls made/:

```
GET: localhost:1234/history
```

response: 200 Ok

response body:
```json
{
    "history": null
}
```


When several requests are made, eg /word : apple, gexray, chair, square, we call history again:

```
GET: localhost:1234/history
```

response: 200 Ok

response body:
```json
{
    "history": [
        {
            "apple": "gapple"
        },
        {
            "chair": "airchogo"
        },
        {
            "square": "aresquogo"
        },
        {
            "xray": "gexray"
        }
    ]
}
```

When we add sentence, like /sentence: "Apple xray chair square?", when call /history endpoint, we get:

```
GET: localhost:1234/history
```

response: 200 Ok

response body:
```json
{
    "history": [
        {
            "Apple xray chair square?": "Gapple gexray airchogo aresquogo?"
        },
        {
            "apple": "gapple"
        },
        {
            "chair": "airchogo"
        },
        {
            "square": "aresquogo"
        },
        {
            "xray": "gexray"
        }
    ]
}
```

Note: history logic is in memory, eg while the app is live. If we restart it, its resetted.
