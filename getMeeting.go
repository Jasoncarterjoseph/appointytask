package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client Client) GetMeeting(meetingId int) (getMeetingResponse GetMeetingResponse, err error) {

	endpoint := fmt.Sprintf("/meetings/%d", meetingId)
	httpMethod := http.MethodGet

	var b []byte
	b, err = client.executeRequest(endpoint, httpMethod)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &getMeetingResponse)
	if err != nil {
		return
	}

	return
}
