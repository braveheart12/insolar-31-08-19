//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// +build apitests

package introspection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Create and initialize an HTTP client for connection re-use:
var introspectionClient *http.Client

func init() {
	introspectionClient = &http.Client{}
}

type Counter struct {
	Name  string `json:"Name"`
	Count string `json:"Count""`
}

type Counters struct {
	Counters []Counter `json:"Counters"`
}

type Filters struct {
	TypeActivate                           Filter `json:"TypeActivate "`
	TypeAdditionalCallFromPreviousExecutor Filter `json:"TypeAdditionalCallFromPreviousExecutor "`
	TypeCallMethod                         Filter `json:"TypeCallMethod "`
	TypeCode                               Filter `json:"TypeCode "`
	TypeDeactivate                         Filter `json:"TypeDeactivate "`
	TypeError                              Filter `json:"TypeError "`
	TypeExecutorResults                    Filter `json:"TypeExecutorResults "`
	TypeFilamentSegment                    Filter `json:"TypeFilamentSegment "`
	TypeGetCode                            Filter `json:"TypeGetCode "`
	TypeGetFilament                        Filter `json:"TypeGetFilament "`
	TypeGetJet                             Filter `json:"TypeGetJet "`
	TypeGetObject                          Filter `json:"TypeGetObject "`
	TypeGetPendings                        Filter `json:"TypeGetPendings "`
	TypeGetRequest                         Filter `json:"TypeGetRequest "`
	TypeHasPendings                        Filter `json:"TypeHasPendings "`
	TypeHotObjects                         Filter `json:"TypeHotObjects "`
	TypeID                                 Filter `json:"TypeID "`
	TypeIDs                                Filter `json:"TypeIDs "`
	TypeIndex                              Filter `json:"TypeIndex "`
	TypeJet                                Filter `json:"TypeJet "`
	TypeMeta                               Filter `json:"TypeMeta "`
	TypeObjIndex                           Filter `json:"TypeObjIndex "`
	TypeObjState                           Filter `json:"TypeObjState "`
	TypePass                               Filter `json:"TypePass "`
	TypePassState                          Filter `json:"TypePassState "`
	TypePendingFinished                    Filter `json:"TypePendingFinished "`
	TypePendingsInfo                       Filter `json:"TypePendingsInfo "`
	TypeReplication                        Filter `json:"TypeReplication "`
	TypeRequest                            Filter `json:"TypeRequest "`
	TypeRequestInfo                        Filter `json:"TypeRequestInfo "`
	TypeResultInfo                         Filter `json:"TypeResultInfo "`
	TypeReturnResults                      Filter `json:"TypeReturnResults "`
	TypeSagaCallAcceptNotification         Filter `json:"TypeSagaCallAcceptNotification "`
	TypeSetCode                            Filter `json:"TypeSetCode "`
	TypeSetIncomingRequest                 Filter `json:"TypeSetIncomingRequest "`
	TypeSetOutgoingRequest                 Filter `json:"TypeSetOutgoingRequest "`
	TypeSetResult                          Filter `json:"TypeSetResult "`
	TypeState                              Filter `json:"TypeState "`
	TypeStillExecuting                     Filter `json:"TypeStillExecuting "`
	TypeUnknown                            Filter `json:"TypeUnknown "`
	TypeUpdate                             Filter `json:"TypeUpdate "`
}

type Filter struct {
	Enable   bool   `json:"Enable"`
	Filtered string `json:"Filtered"`
}

type SetFilterJSON struct {
	Name   string `json:Name`
	Enable bool   `json:"Enable"`
}

func SendEmptyPostRequest(url string) []byte {
	responseBody := SendRequest(nil, url, "POST", "")
	return responseBody

}

func SendRequest(payload []byte, intUrl string, method string, contentType string) []byte {
	// Create a new HTTP request and send it:
	request, err := http.NewRequest(method, intUrl, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("ContentType", contentType)
	response, err := introspectionClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer request.Body.Close()

	// Receive and return the response body:
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(responseBody))

	return responseBody
}

func GetMessageNodeCounters(introspectionPortAddress string) Counters {
	var messageCounters Counters

	var err error
	var responce []byte

	responce = SendEmptyPostRequest("http://" + introspectionPortAddress + "/getMessagesStat")
	err = json.Unmarshal(responce, &messageCounters)
	if err != nil {
		log.Fatalln(err)
	}
	return messageCounters
}

func SetMessageNodeCounters(introspectionPortAddress string, messageType string, enableFilter bool) SetFilterJSON {
	setMessageTypeFilterEnable := SetFilterJSON{
		Name:   messageType,
		Enable: enableFilter,
	}

	jsonPayload, err := json.Marshal(setMessageTypeFilterEnable)
	if err != nil {
		log.Fatalln(err)
	}

	var response []byte
	response = SendRequest(jsonPayload, "http://"+introspectionPortAddress+"/setMessagesFilter", "POST", "application/json")
	var result SetFilterJSON
	// Unmarshal the response:
	err = json.Unmarshal(response, &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}
