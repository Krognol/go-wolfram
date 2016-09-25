# go-wolfram
Single file library for the Wolfram Alpha API written in Golang without extra dependencies

# Installing
`go get github.com/Krognol/go-wolfram`

# Example
```go
package main

import "github.com/Krognol/go-wolfram"

func main() {
	//Initialize a new client
	c := &wolfram.Client{AppID:"your app id here"}

	//Get a result without additional parameters
	res, err := c.GetQueryResult("1+1", nil)

	if err != nil {
		panic(err)
	}

	//Iterate through the pods and subpods
	//and print out their title attributes
	for i := range res.Pods {
		println(res.Pods[i].Title)

		for j := range res.Pods[i].SubPods {
			println(res.Pods[i].SubPods[j].Title)
		}
	}
}
```
### Output

```
> go run main.go

< Input
< Result
< Number name
< Visual representation
< Number line
< Illustration
```

# Adding extra parameters

```go
package main

import "github.com/Krognol/go-wolfram"

func main() {
	//Initialize a new client
	c := &wolfram.Client{AppID:"your app id here"}
  
  	params := url.Values{}
	params.Set("assumption", "DateOrder_**Day.Month.Year--")
  
	//Get a result with additional parameters
	res, err := c.GetQueryResult("26-9-2016", params)

	if err != nil {
		panic(err)
	}

	//Iterate through the pods and subpods
	//and print out their title attributes
	for i := range res.Pods {
		println(res.Pods[i].Title)

		for j := range res.Pods[i].SubPods {
			println(res.Pods[i].SubPods[j].Title)
		}
	}
}
```

### Output

```
> go run main.go

< Input interpretation
< Date formats
< Time difference from today (Friday, September 23, 2016)
< Time in 2016
< Observances for September 26 (Sweden)
< Anniversaries for September 26, 2016
< Daylight information for September 26 in Stockholm, Sweden
< Phase of the Moon
```
