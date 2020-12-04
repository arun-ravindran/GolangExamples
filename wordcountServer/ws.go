// Demo of Go net/rpc
// Server computes the wordccount of an input string
package wordcountServer

import "net"
import "net/rpc"
import "strings"

type Server struct {
	addr string
}

type Request struct {
	Input string
}

type Response struct {
	Counts map[string]int
}

func (s *Server) Listen() {
	rpc.Register(s)
	listner, err := net.Listen("tcp", s.addr)
	checkError(err)
	go func() {
		rpc.Accept(listner)
	}()
}

func (s *Server) Compute(req Request, resp *Response) error {
	counts := make(map[string]int)
	input := req.Input
	tokens := strings.Fields(input)
	for _, tok := range tokens {
		counts[tok] += 1
	}
	resp.Counts = counts
	return nil
}
