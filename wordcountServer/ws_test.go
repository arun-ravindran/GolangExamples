package wordcountServer

import (
	"net/rpc"
	"reflect"
	"testing"
)

func TestServer(t *testing.T) {
	servAddr := "localhost:8888"
	server := Server{addr: servAddr}
	server.Listen()
	input1 := "hello I am good hello bye bye bye bye good night hello"
	got, err := makeRequest(input1, servAddr)
	checkError(err)
	want := map[string]int{"hello": 3, "I": 1, "am": 1, "good": 2, "bye": 4, "night": 1}
	if !reflect.DeepEqual(got, want) {
		t.Error("Wrong word count")
	}
}

// Client
func makeRequest(input string, servAddr string) (map[string]int, error) {
	client, err := rpc.Dial("tcp", servAddr)
	checkError(err)
	args := Request{Input: input}
	resp := Response{Counts: make(map[string]int)}
	err = client.Call("Server.Compute", args, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Counts, nil
}
