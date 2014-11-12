//Gets all the dead servers in your newrelic account and then removes them
//Great for ephemeral spot instance type EC2 architectures which can clog up
//Newrelic with a load of dead servers, removing them via the web interface is painful
//TODO run deleteions as parralel jobs using channels to sync
package main

import (
	"encoding/json"
	"flag"      //cli parsing
	"io/ioutil" //some convenience iostream methods
	"log"
	"net/http"
	"strconv"
)

const API_ENDPOINT = "https://api.newrelic.com/v2/servers"

var apiKey *string = flag.String("api-key", "", "Your newrelic api key")

//our minimum data that we'll marshall out of the NR json
type nrJson struct {
	Servers []struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		Reporting bool   `json:"reporting"`
	} `json:"servers"`
}

func main() {
	//get all the cli args
	flag.Parse()

	if *apiKey == "" {
		log.Fatal("You must supply your NewRelic api key use --api-key")
	}
	log.Println("Getting server list from NR")
	//create a http client for custom headers, rather than use the convenience GET method
	client := &http.Client{}
	req, err := http.NewRequest("GET", API_ENDPOINT+".json", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Api-Key", *apiKey)

	// GET server data from the api
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	//read body stream
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal("Didn't get a 200 statuscode ", res.StatusCode)
	}

	//marshal the json output into our struct
	unmarshaledJson := nrJson{}
	err = json.Unmarshal(body, &unmarshaledJson)
	if err != nil {
		log.Fatal(err)
	}

	var serverCount int = 0
	//foreach over servers and remove the duds
	servers := unmarshaledJson.Servers
	for _, s := range servers {
		if s.Reporting == false {
			if RemoveServer(s.Id) == true {
				serverCount++
			} else {
				log.Panicln("Couldn't remove server ", s.Id)
			}
		}
	}
	log.Println("Removed", serverCount, "dead servers from your NR account")
}

//Remove a dead server by id
func RemoveServer(serverId int) bool {
	log.Println("Removing dead server id", serverId)
	//create a manual http object to pass in the API key
	client := &http.Client{}
	var url string = API_ENDPOINT + "/" + strconv.Itoa(serverId) + ".json"
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Api-Key", *apiKey)

	// DELETE server data from the api
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Panic("Didn't get a 200 statuscode ", res.StatusCode)
		return false
	}

	return true
}
