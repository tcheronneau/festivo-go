package festigo

import (
  "net/url"
  "encoding/json"
  "io"
  "path"
  "io/ioutil"
  "errors"
  "net/http"
//  "reflect"
  //"fmt"
)

type Client struct {
  key string
  baseURL url.URL
  *http.Client
}

type Response struct {
  Status int
  Envelope map[string]string
  Requests map[string]interface{}
  Message string
  Errors string
  Holidays Holidays
}

type Holidays struct {
  Query map[string]string
  Holidays []Holiday
}

type Holiday struct {
  Name string
  Country string
  Date string
  Public bool
  Start string
  End string
  Type string
}

type Query struct {
  Country string
  Year int
  Month int
  Day int
  Language string
  Previous bool
  Upcoming bool
  Public bool
  Format string
  Pretty bool
  State bool
  Region string
  Type string
}

func New(key string) (*Client, error) {
  baseURL := "https://getfestivo.com/v1/holidays"
  u, err := url.Parse(baseURL)
  if err != nil {
    return nil, err
  }
  return &Client{key,*u,&http.Client{}}, nil
}

func (c *Client) newRequest(method, requestPath string, query url.Values, body io.Reader) (*http.Request, error) {
  url := c.baseURL
  url.Path = path.Join(url.Path, requestPath)
  key := c.key
  query.Add("api_key", key)
  url.RawQuery = query.Encode()
  req, err := http.NewRequest(method, url.String(), body)
  if err != nil {
    return req, err
  }

  req.Header.Add("Content-Type", "application/json")
  return req, err
}

func (c *Client) GetHolidays(query map[string]interface{}) ([]Holiday, error) {
  holidays := make([]Holiday,0)
  q := url.Values{}
  for k,v := range query {
    q.Add(k,v.(string))
  }
  req,err := c.newRequest("GET","",q,nil)
  if err != nil {
    return holidays,err
  }
  resp, err := c.Do(req)
  if err != nil {
    return holidays,err
  }
  if resp.StatusCode != 200 {
    return holidays,errors.New(resp.Status)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return holidays,err
  }
  response := Response{}
  err = json.Unmarshal(data,&response)
  holdays := response.Holidays
  holidays = holdays.Holidays
  return holidays, err
}
