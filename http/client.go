/*
Package http provides high-level access to the RabbitMQ HTTP API with all objects defined as structs.
*/
package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type dataRetriever func(url string, method string, user string, pass string) ([]byte, error)

// A Client is a representation of a connection with the RabbitMQ HTTP API
type Client struct {
	url     string
	user    string
	pass    string
	dataRet dataRetriever
}

// NewClient creates a new Client instance with the given arguments
func NewClient(u string, us string, p string) *Client {
	return &Client{
		url:     u,
		user:    us,
		pass:    p,
		dataRet: call,
	}
}

func newClientWithDataRet(u string, us string, p string, dr dataRetriever) *Client {
	return &Client{
		url:     u,
		user:    us,
		pass:    p,
		dataRet: dr,
	}
}

// Overview returns general information about the broker
func (c *Client) Overview() (Overview, error) {
	data, err := c.dataRet(c.url+"/overview", "GET", c.user, c.pass)
	if err != nil {
		return Overview{}, err
	}
	var o Overview
	json.Unmarshal(data, &o)
	return o, nil
}

// Connections returns info about all the connections in the broker
func (c *Client) Connections() ([]Connection, error) {
	data, err := c.dataRet(c.url+"/connections", "GET", c.user, c.pass)
	if err != nil {
		return nil, err
	}
	var conns []Connection
	json.Unmarshal(data, &conns)

	return conns, nil
}

func call(url string, m string, user string, pass string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(m, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Não foi possível ler a lista de deputados")
		return nil, err
	}
	return body, nil
}
