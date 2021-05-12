package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptrace"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

type StringService interface {
  Uppercase(string) (string, error)
  Count (string) int
}

type stringServise struct{}

func (stringServise) Uppercase(s string) (string, error){
  if s == "" {
    return "", ErrEmpty
  }
  return strings.ToUpper(s), nil
}

func (stringServise) Count(s string) int {
  return len(s)
}

var ErrEmpty = errors.New("Empty string")

type UppercaseRequest struct {
  S string `json:"s"`
}

type uppercaseResponse struct {
  V string `json:"v"`
  Err string `json:"error,omitempty"`
}

type countRequest struct {
  S string `json:"s"`
}

type countResponse struct {
  V int `json:"v"`
}


func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
  return func(_ context.Context, request interface{}) (interface{}, error){
    req := request.(countRequest)
    v := svc.Count(req.S)
    return countResponse{v}, nil
  }
}


func main(){
  svc := stringService{}
  uppercaseHandler := httptransport.NewServer(
    makeUppercaseEndpoint(svc),
    decodeUppercaseRequest,
    encodeResponse,
  )
  countHandler := httptransport.NewServer(
    makeCountEndopoint(svc),
    decodeCountRequest,
    encodeResponse,
  )
  
  http.Handle("/uppercase", uppercaseHandler)
  http.Handle("/count", countHandler)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error){
  var request UppercaseRequest
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
    return nil, err
  }
  return request, nil
}

func decodeCountRequest(_ context.Context, w http.ResponseWriter, response interface{}) error {
  return json.NewEncoder(w).Encode(response)
}
