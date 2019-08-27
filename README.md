# festivo-go
Festivo API client for GO

Example usage : 
```
package main

import (
  "fmt"
  "github.com/tcheronneau/festivo-go"
)

func main() {
  hapi,_ := festigo.New("YOUR_API")
  query := map[string]interface{}{
    "country": "LU",
    "year": "2019",
    //"month": "08",
    //"day": "26",
  }
  h,err := hapi.GetHolidays(query)
  if err != nil {
    fmt.Println(err)
  }
  for _,holi := range h {
    fmt.Printf("%s\n", holi)
  }
}
```
Will result in : 
```
{New Year's Day LU 2019-01-01 %!s(bool=true) 2019-01-01T00:00:00.000Z 2019-01-02T00:00:00.000Z public}
{Good Friday LU 2019-04-19 %!s(bool=false) 2019-04-19T00:00:00.000Z 2019-04-20T00:00:00.000Z observance}
{Easter Sunday LU 2019-04-21 %!s(bool=false) 2019-04-21T00:00:00.000Z 2019-04-22T00:00:00.000Z observance}
{Easter Monday LU 2019-04-22 %!s(bool=true) 2019-04-22T00:00:00.000Z 2019-04-23T00:00:00.000Z public}
{Labour Day LU 2019-05-01 %!s(bool=true) 2019-05-01T00:00:00.000Z 2019-05-02T00:00:00.000Z public}
{Ascension Day LU 2019-05-30 %!s(bool=true) 2019-05-30T00:00:00.000Z 2019-05-31T00:00:00.000Z public}
{Whit Monday LU 2019-06-10 %!s(bool=true) 2019-06-10T00:00:00.000Z 2019-06-11T00:00:00.000Z public}
{L’anniversaire du Grand-Duc LU 2019-06-23 %!s(bool=true) 2019-06-23T00:00:00.000Z 2019-06-24T00:00:00.000Z public}
{Assumption LU 2019-08-15 %!s(bool=true) 2019-08-15T00:00:00.000Z 2019-08-16T00:00:00.000Z public}
{All Saints' Day LU 2019-11-01 %!s(bool=true) 2019-11-01T00:00:00.000Z 2019-11-02T00:00:00.000Z public}
{Armistice 1918 LU 2019-11-11 %!s(bool=true) 2019-11-11T00:00:00.000Z 2019-11-12T00:00:00.000Z public}
{Christmas Day LU 2019-12-25 %!s(bool=true) 2019-12-25T00:00:00.000Z 2019-12-26T00:00:00.000Z public}
{Boxing Day LU 2019-12-26 %!s(bool=true) 2019-12-26T00:00:00.000Z 2019-12-27T00:00:00.000Z public}
```

Object Holiday has several accessible fields : 
```
  Name string
  Country string
  Date string
  Public bool
  Start string
  End string
  Type string
```
So if you want to access only date or name you will change the loop : 
```
  for _,holi := range h {
    fmt.Printf("%s\n", holi.Date)
  }
```
