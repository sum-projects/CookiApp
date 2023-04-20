package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	req, err := http.NewRequest("POST", s.config.MailService, bytes.NewBuffer(jsonData))
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
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error calling mail service: %v", err)
		}
		return fmt.Errorf("error calling mail service: %s", resBody)
	}

	return err
}
