package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (s *Server) sendMail(msg MailPayload) error {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	mailServiceURL := "http://mailer-service/send"
	req, err := http.NewRequest("POST", mailServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		return errors.New("error calling mail service")
	}

	return err
}
