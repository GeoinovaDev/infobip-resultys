package webhook

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"

	"git.resultys.com.br/lib/lower/convert/decode"
	"git.resultys.com.br/lib/lower/log"
	"git.resultys.com.br/lib/lower/net/loopback"
	"git.resultys.com.br/lib/lower/promise"
	"git.resultys.com.br/sdk/infobip-golang/message"
	"git.resultys.com.br/sdk/infobip-golang/response"
)

// Server struct
type Server struct {
	Port string

	hooks map[string]*promise.Promise
	mutex *sync.Mutex
}

// New ...
func New(port string) *Server {
	s := &Server{
		Port:  port,
		mutex: &sync.Mutex{},
		hooks: make(map[string]*promise.Promise),
	}

	s.Start()

	return s
}

// AddHook ...
func (s *Server) AddHook(messageID string) *promise.Promise {
	s.mutex.Lock()
	p := &promise.Promise{}
	s.hooks[messageID] = p
	s.mutex.Unlock()

	return p
}

// RemoveHook ...
func (s *Server) RemoveHook(messageID string) {
	delete(s.hooks, messageID)
}

// ResolveHook ...
func (s *Server) ResolveHook(messageID string, response interface{}) {
	s.hooks[messageID].Resolve(response)
	s.RemoveHook(messageID)
}

// ExistHook ...
func (s *Server) ExistHook(messageID string) bool {
	if _, ok := s.hooks[messageID]; ok {
		return true
	}

	return false
}

// Start ...
func (s *Server) Start() {
	go http.ListenAndServe(s.Port, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer func() {
		err := recover()
		if err != nil {
			log.Logger.Save(fmt.Sprint(err), log.WARNING, loopback.IP())
		}
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	json := response.ResultsResponse{Messages: make([]message.Message, 1)}
	decode.JSON(body, &json)

	s.mutex.Lock()
	messageID := json.Messages[0].MessageID
	if s.ExistHook(messageID) {
		s.ResolveHook(messageID, json)
	}
	s.mutex.Unlock()

	w.Write([]byte("ok"))
}
