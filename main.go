// Copyright 2014 Daiki Kohara
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	. "github.com/daikikohara/enotify-slack/event"
)

var (
	config = NewConfig("conf.yml")
	logger = NewLogger(config.Logfile)
	boltdb = NewBolt(config.Boltdb.Dbfile, config.Boltdb.Bucketname)
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	payload := make(chan *Payload)
	for provider, _ := range config.Provider {
		go getEvent(provider, payload)
	}
	go sendEvent(payload)
	shutdownHook()
}

// getEvent gets events from event providers.
// API depends on the provider, so the actual implementation to get events is obtained via GetApi(factory method).
func getEvent(provider string, payload chan *Payload) {
	api := GetApi(provider)
	for {
		events, err := api.Get(config.Provider[provider].Url, config.Keyword, config.Nickname)
		handleError("Error on getting event from "+provider, err, payload)
		for _, event := range events {
			key := []byte(provider + ":" + fmt.Sprintf("%v", event.Id))
			if event.IsValid(config.Place) && !boltdb.Exists(key) {
				value, _ := json.Marshal(event)
				boltdb.Put(key, value)
				payload <- ConstructSlackMessage(provider, &event, config)
			}
		}
		time.Sleep(time.Duration(config.Provider[provider].Interval) * time.Second)
	}
}

// handleError handles error.
func handleError(pretext string, err error, errorMsg chan *Payload) {
	if err != nil {
		msg := pretext + "\n" + err.Error()
		logger.Println(msg)
		if config.ErrorToSlack {
			errorMsg <- ConstructSlackError(msg, config.Slack.Channel)
		}
	}
}

// sendEvent sends Payload which contains Event to Slack.
// this function should be called by only one goroutine to avoid the Slack rate limit(1 payload/sec)
func sendEvent(payload chan *Payload) {
	for {
		p := <-payload
		data, err := json.Marshal(p)
		resp, err := http.Post(config.Slack.Url, "application/json", strings.NewReader(string(data)))
		if err != nil {
			logger.Println(err.Error())
		}
		defer resp.Body.Close()
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Println(err.Error())
		}
		time.Sleep(time.Duration(config.Slack.Interval) * time.Second)
	}
}

// shutdownHook handles OS signals.
// It catches SIGINT and SIGTERM to close resources and then exits the process.
func shutdownHook() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	boltdb.Db.Close()
	logger.Close()
	fmt.Println("bye")
	os.Exit(0)
}
