package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

// AlexaSkillSchema is the JSON representation of the interaction model.
type AlexaSkillSchema struct {
	Intents []struct {
		Name    string        `json:"name"`
		Samples []interface{} `json:"samples"`
		Slots   []struct {
			Name    string   `json:"name"`
			Type    string   `json:"type"`
			Samples []string `json:"samples"`
		} `json:"slots,omitempty"`
	} `json:"intents"`
	Types []struct {
		Name   string `json:"name"`
		Values []struct {
			ID   interface{} `json:"id"`
			Name struct {
				Value    string   `json:"value"`
				Synonyms []string `json:"synonyms"`
			} `json:"name"`
		} `json:"values"`
	} `json:"types"`
	Prompts []struct {
		ID                string `json:"id"`
		PromptVersion     string `json:"promptVersion"`
		DefinitionVersion string `json:"definitionVersion"`
		Variations        []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"variations"`
	} `json:"prompts"`
	Dialog struct {
		Version string `json:"version"`
		Intents []struct {
			Name                 string `json:"name"`
			ConfirmationRequired bool   `json:"confirmationRequired"`
			Prompts              struct {
			} `json:"prompts"`
			Slots []struct {
				Name                 string `json:"name"`
				Type                 string `json:"type"`
				ElicitationRequired  bool   `json:"elicitationRequired"`
				ConfirmationRequired bool   `json:"confirmationRequired"`
				Prompts              struct {
					Elicit string `json:"elicit"`
				} `json:"prompts"`
			} `json:"slots"`
		} `json:"intents"`
	} `json:"dialog"`
}

// ParseSchemaJSON parses the interaction model JSON into the provided
// AlexaSkillSchema struct.
func (s *AlexaSkillSchema) ParseSchemaJSON(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, s)
}

// GetSynonyms returns all the synonyms for the requested slot value.
func (s *AlexaSkillSchema) GetSynonyms(slotValue string) (synonyms []string) {
	for _, typ := range s.Types {
		for _, value := range typ.Values {
			if strings.EqualFold(value.Name.Value, slotValue) {
				return append(synonyms, value.Name.Synonyms...)
			}
		}
	}
	return
}

// GetSlotValue returns the slot value of the requested intent.
func (s *AlexaSkillSchema) GetSlotValue(echoReq *alexa.EchoRequest) (string,
	error) {
	for _, intent := range s.Intents {
		if intent.Name == echoReq.GetIntentName() {
			return echoReq.GetSlotValue(intent.Slots[0].Name)
		}
	}
	return "", errors.New("cannot get slot value: no matching intent name found in local schema")
}
