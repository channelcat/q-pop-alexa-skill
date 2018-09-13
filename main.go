package main

import (
	"log"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

const (

	// EchoAppID from Amazon Dashboard.
	EchoAppID  = "amzn1.ask.skill.746568a7-460a-4021-925d-33c05672b21a"
	sslCert    = "/etc/letsencrypt/live/skillserver.qpop.services/fullchain.pem"
	sslKey     = "/etc/letsencrypt/live/skillserver.qpop.services/privkey.pem"
	schemaFile = "schema.json"
)

var (
	// Applications defines the alexa skill's meta information.
	Applications = map[string]interface{}{
		"/echo/qpop": alexa.EchoApplication{ // Route
			AppID:          EchoAppID,
			OnLaunch:       EchoLaunchHandler,
			OnIntent:       EchoIntentHandler,
			OnSessionEnded: EchoSessionEndedHandler,
		},
	}
	schema AlexaSkillSchema
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	alexa.RunSSL(Applications, ":443", sslCert, sslKey)
}

// EchoLaunchHandler handles the LaunchRequest from the echo service.
func EchoLaunchHandler(echoReq *alexa.EchoRequest,
	echoResp *alexa.EchoResponse) {
	msg := "Welcome to q. Pop. A League of Legends utility that grants voice control capabilities to matchmaking and champion select."
	echoResp.OutputSpeech(msg).EndSession(false)
}

// EchoIntentHandler handles all IntentRequests from the echo service.
func EchoIntentHandler(echoReq *alexa.EchoRequest,
	echoResp *alexa.EchoResponse) {

	// Load the local interaction model JSON.
	err := schema.ParseSchemaJSON(schemaFile)
	if err != nil {
		ErrorHandler(err, echoResp)
		return
	}
}

// EchoSessionEndedHandler handles the SessionEndedRequest from the echo
// service.
func EchoSessionEndedHandler(echoReq *alexa.EchoRequest,
	echoResp *alexa.EchoResponse) {
	echoResp.EndSession(true)
}

// ErrorHandler logs the provided error then returns it in the EchoResponse.
func ErrorHandler(err error, echoResp *alexa.EchoResponse) {
	log.Println(err)
	echoResp.OutputSpeech(err.Error()).EndSession(true)
}
