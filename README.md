# go-xivapi #

go-xivapi is a Go client library for accessing [xivapi](https://xivapi.com/docs)

## Usage ##

Construct a new XIVApi client, then use the various services on the client to
access different parts of the API. For example:

```go
client := xivapi.NewClient(nil)

// list all items containing the term "aiming"
items, _, err := client.Search.Items(context.Background(), "aiming")
```