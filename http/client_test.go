package http

import (
	"io/ioutil"
	"testing"
)

var testData []byte
var testError error

func TestOverview(t *testing.T) {
	testData, _ = ioutil.ReadFile("testdata/overview_1.json")

	client := createClient()
	overview, _ := client.Overview()

	expectedOverview := Overview{
		Node:        "rabbit@host",
		ClusterName: "rabbit@host",
		Version:     "3.6.10",
	}

	compareOverview(expectedOverview, overview, t)
}

func createClient() *Client {
	return newClientWithDataRet("http://test.com", "user", "password", mockRetriever)
}

func mockRetriever(url string, m string, us string, p string) ([]byte, error) {
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
}
