package http

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var testData []byte
var testURL string
var testError error

func TestOverview(t *testing.T) {
	testData, _ = ioutil.ReadFile("testdata/overview_1.json")

	client := createClient("/overview")
	overview, err := client.Overview()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedOverview := Overview{
		Node:          "rabbit@host",
		ClusterName:   "rabbit@host",
		Version:       "3.6.10",
		MgtVersion:    "3.6.10",
		ErlangVersion: "18.3",
	}

	compareOverview(expectedOverview, overview, t)
}

func TestNodes(t *testing.T) {
	testData, _ = ioutil.ReadFile("testdata/nodes_1.json")

	client := createClient("/nodes")
	nodes, err := client.Nodes()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedNodes := []Node{}
	expectedNodes = append(expectedNodes, Node{
		Name:           "rabbit@8a288d0bd4f9",
		Type:           "disc",
		Running:        true,
		FilesDescUsed:  58,
		FilesDescTotal: 65536,
	})

	compareNodes(expectedNodes, nodes, t)
}

func TestNode(t *testing.T) {
	testData, _ = ioutil.ReadFile("testdata/node_1.json")

	client := createClient("/nodes/rabbit@8a288d0bd4f9")

	node, err := client.Node("rabbit@8a288d0bd4f9")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedNode := Node{
		Name:           "rabbit@8a288d0bd4f9",
		Type:           "disc",
		Running:        true,
		FilesDescUsed:  58,
		FilesDescTotal: 65536,
	}

	compareNode(expectedNode, node, t)
}

func createClient(apiURL string) *Client {
	testURL = "http://test.com" + apiURL
	return newClientWithDataRet("http://test.com", "user", "password", mockRetriever)
}

func mockRetriever(url string, m string, us string, p string) ([]byte, error) {
	if url != testURL {
		return nil, fmt.Errorf("Wrong url: expected %s but got %s", testURL, url)
	}
	return testData, testError
}

func compareOverview(o1 Overview, o2 Overview, t *testing.T) {
	if o1.Node != o2.Node {
		t.Fatalf("Expected node '%s' but got '%s'", o1.Node, o2.Node)
	}
	if o1.ClusterName != o2.ClusterName {
		t.Fatalf("Expected cluster name '%s' but got '%s'", o1.ClusterName, o2.ClusterName)
	}
	if o1.Version != o2.Version {
		t.Fatalf("Expected version '%s' but got '%s'", o1.Version, o2.Version)
	}
	if o1.MgtVersion != o2.MgtVersion {
		t.Fatalf("Expected management version '%s' but got '%s'", o1.MgtVersion, o2.MgtVersion)
	}
	if o1.ErlangVersion != o2.ErlangVersion {
		t.Fatalf("Expected erlang version '%s' but got '%s'", o1.ErlangVersion, o2.ErlangVersion)
	}
}

func compareNodes(ns1 []Node, ns2 []Node, t *testing.T) {
	if len(ns1) != len(ns2) {
		t.Fatalf("Expected '%d' nodes but got '%d'", len(ns1), len(ns2))
	}
	for i, n1 := range ns1 {
		n2 := ns2[i]
		compareNode(n1, n2, t)
	}
}

func compareNode(n1 Node, n2 Node, t *testing.T) {
	if n1.Name != n2.Name {
		t.Fatalf("Expected name '%s' but got '%s'", n1.Name, n2.Name)
	}
	if n1.Type != n2.Type {
		t.Fatalf("Expected type '%s' but got '%s'", n1.Type, n2.Type)
	}
	if n1.Running != n2.Running {
		t.Fatalf("Expected node state %t but got %t", n1.Running, n2.Running)
	}
	if n1.FilesDescUsed != n2.FilesDescUsed {
		t.Fatalf("Expected files descriptors used %d but got %d", n1.FilesDescUsed, n2.FilesDescUsed)
	}
	if n1.FilesDescTotal != n2.FilesDescTotal {
		t.Fatalf("Expected files descriptors total %d but got %d", n1.FilesDescTotal, n2.FilesDescTotal)
	}
}
