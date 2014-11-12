//Gets all the dead servers in your newrelic account and then removes them
//Great for ephemeral spot instance type EC2 architectures which can clog up
//Newrelic with a load of dead servers
package main

import (
	"encoding/json"
	"flag" //cli parsing
	//"fmt"
	"io/ioutil" //some convenience iostream methods
	"log"
	"net/http"
)

const API_ENDPOINT = "https://api.newrelic.com/v2/servers.json"

var apiKey *string = flag.String("api-key", "", "Your newrelic api key")

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

	//create a http client for custom headers, rather than use the convenience GET method
	client := &http.Client{}
	req, err := http.NewRequest("GET", API_ENDPOINT, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Api-Key", *apiKey)

	// GET server data from the api
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	//read body
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal("Didn't get a 200 statuscode ", res.StatusCode)
	}

	log.Print("raw_json", string(body))

	unmarshaledJson := nrJson{}
	err = json.Unmarshal(body, &unmarshaledJson)
	if err != nil {
		log.Fatal(err)
	}

	//foreach over servers and remove dud ones
	servers := unmarshaledJson.Servers
	for _, s := range servers {
		if s.Reporting == false {
			RemoveServer(s.Id)
		}
	}
	//if server["reporting"] == false {
	//		s := server.(map[string]interface{})
	//		if s["reporting"] == false {
	//log.Printf("Removing dead server ", s)
	//			log.Printf("Removing dead server ", s["name"])
	//			fmt.Print(s["id"])
	//			server_id, _ := js.Int(s["id"])
	//			RemoveServer(server_id)
	//		}
	//		if server["id"] {
	//			log.Printf("index", server.id)
	//		}
	//	}
	//}
	//log.Printf("response", nrResponse.servers[0].id)

	//2. foreach over results and delete dud servers
}

func RemoveServer(serverId int) bool {
	log.Printf("server id ", serverId)
	return true
}
