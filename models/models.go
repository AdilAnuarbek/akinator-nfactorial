package models

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

type Akinator struct {
	SessionID string `json:"sessionId"`
	Theme     string `json:"theme"`
	Region    string `json:"region"`
}

// Starts the Akinator game and immediately returns the session ID, question, and answers
// Example:
// "question": "Is your character real?",
// "answers": ["Yes","No","Don't know","Probably","Probably not"],
// "progress": 0
func (a *Akinator) StartAkinatorGame() (string, error) {
	a.SessionID = generateSessionID()
	reqBody, err := json.Marshal(map[string]string{
		"sessionId": a.SessionID,
		"theme":     a.Theme,
		"region":    a.Region,
	})
	if err != nil {
		return "", err
	}
	resp, err := http.Post("http://localhost:3000/start", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Decode the response
	var result struct {
		Question string `json:"question"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Question, nil
}

func (a *Akinator) AnswerAkinatorGame(answer string) (string, bool, string, error) {
	reqBody, err := json.Marshal(map[string]string{
		"sessionId": a.SessionID,
		"answer":    answer,
	})
	if err != nil {
		return "", false, "", err
	}
	resp, err := http.Post("http://localhost:3000/answer", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", false, "", err
	}
	defer resp.Body.Close()
	var result struct {
		Question string `json:"question"`
		Win      bool   `json:"win"`
		Name     string `json:"name"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	if result.Win {
		return result.Question, true, result.Name, nil
	}
	return result.Question, false, "", nil
}

func (a *Akinator) GuessAkinatorGame(answer string) (string, error) {
	reqBody, _ := json.Marshal(map[string]string{
		"sessionId": a.SessionID,
		"answer":    answer,
	})
	resp, err := http.Post("http://localhost:3000/guess", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result struct {
		Question string `json:"question"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Question, nil
}

func generateSessionID() string {
	b := make([]byte, 32) // 32 bytes = 256 bits of entropy
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("failed to generate session id")
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
