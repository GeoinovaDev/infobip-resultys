package call

import (
	"git.resultys.com.br/sdk/infobip-golang/ivr"
)

// Call struct
type Call struct {
	IVR *ivr.IVR
}

// Wait esperar ate a ligação completar
func (call *Call) Wait() {

}

// FirstResponse retorna a resposta da primeira request
func (ivr *IVR) FirstResponse() string {
	if len(ivr.responses) == 0 {
		return ""
	}

	return ivr.responses[0]
}

// LastResponse retorna a resposta da ultima request
func (ivr *IVR) LastResponse() string {
	if len(ivr.responses) == 0 {
		return ""
	}

	return ivr.responses[len(ivr.responses)-1]
}

// Responses retorna todas as respostas das requests realizadas
func (ivr *IVR) Responses() []string {
	return ivr.responses
}
