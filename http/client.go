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

// A Client is a representation of a connection with the RabbitMQ HTTP API
type Client struct {
	url  string
	user string
	pass string
}

// NewClient creates a new Client instance with the given arguments
func NewClient(u string, us string, p string) *Client {
	return &Client{
		url:  u,
		user: us,
		pass: p,
	}
}

// Overview returns general information about the broker
func (c *Client) Overview() (Overview, error) {
	data, err := c.callAPI("/overview", "GET")
	if err != nil {
		return Overview{}, err
	}
	var o Overview
	json.Unmarshal(data, &o)
	return o, nil
}

func (c *Client) Connections() ([]Connection, error) {
	data, err := c.callAPI("/connections", "GET")
	if err != nil {
		return nil, err
	}
	var conns []Connection
	json.Unmarshal(data, &conns)

	return conns, nil
}

func (c *Client) callAPI(url string, m string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(m, c.url+url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.SetBasicAuth(c.user, c.pass)
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
