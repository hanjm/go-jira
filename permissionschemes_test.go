package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPermissionSchemeService_GetList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/3/permissionscheme"

	raw, err := ioutil.ReadFile("./mocks/all_permissionschemes.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprintf(w, string(raw))
	})

	permissionScheme, _, err := testClient.PermissionScheme.GetList()
	if permissionScheme == nil {
		t.Error("Expected role list. Role list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %v", err)
	}
}

func TestPermissionSchemeService_Get(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/rest/api/3/permissionscheme/10100"
	raw, err := ioutil.ReadFile("./mocks/permissionscheme.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEdpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		testRequestURL(t, request, testAPIEdpoint)
		fmt.Fprintf(writer, string(raw))
	})

	permissionScheme, _, err := testClient.PermissionScheme.Get(10100)
	if permissionScheme == nil {
		t.Errorf("Expected Role, got nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}

}