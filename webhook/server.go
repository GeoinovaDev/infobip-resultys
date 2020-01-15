package webhook

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"git.resultys.com.br/lib/lower/convert/decode"
	"git.resultys.com.br/lib/lower/exception"
	"git.resultys.com.br/lib/lower/promise"
	"git.resultys.com.br/sdk/infobip-golang/log"
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
	defer s.mutex.Unlock()

	p := promise.New()

	if s.ExistHook(messageID) {
		p = s.hooks[messageID]
		s.RemoveHook(messageID)
		return p
	}

	s.hooks[messageID] = p

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

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	log.GetInstance().Add(body)
	go s.process(body)

	w.Write([]byte("ok"))
}

func (s *Server) process(body string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	defer func() {
		err := recover()
		msg := ""

		switch err.(type) {
		case string:
			msg = err.(string)
		case []string:
			msg = strings.Join(err.([]string), ". ")
		case error:
			msg = fmt.Sprint(err)
		default:
			msg = "erro de runtime"
		}

		if err != nil {
			exception.Raise(msg, exception.WARNING)
		}

		fmt.Println(err)
	}()

	fmt.Println(body)

	json := response.ResultsResponse{Messages: make([]message.Message, 1)}
	decode.JSON(body, &json)

	for i := 0; i < len(json.Messages); i++ {
		message := json.Messages[i]
		messageID := message.MessageID

		if s.ExistHook(messageID) {
			s.ResolveHook(messageID, message)
		} else {
			p := promise.New()
			s.hooks[messageID] = p
			s.hooks[messageID].Resolve(message)
		}
	}

}
