package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
    "encoding/json"
    "bytes"
)

func checkFail(t *testing.T, err error) {
    if err != nil {
        t.Fatal(err)
    }
}

func setup(t *testing.T) {
    db := ConnectDb()
    defer db.Close()

    sqlData, err := ioutil.ReadFile("schema.sql")
    checkFail(t, err)

    sqlStmts := string(sqlData)
    for _, stmt := range strings.Split(sqlStmts, ";") {
        _, err := db.Exec(stmt)
        checkFail(t, err)
    }

}

func TestList(test *testing.T) {
    setup(test)
    testServer := httptest.NewServer(Handlers())
    defer testServer.Close()

    res, err := http.Get(testServer.URL + "/list")
    checkFail(test, err)

    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode != 200 {
        test.Errorf("Expected status code 200 got %d", res.StatusCode)
    }

    var lists []List
    err = json.Unmarshal(body, &lists)
    checkFail(test, err)

    if len(lists) != 0 {
        test.Errorf("Expected [] got %s", lists)
    }

    b, err := json.Marshal(ListCreateRequest{Name: "Work"})
    checkFail(test, err)

    res, err = http.Post(testServer.URL + "/list", "application/json", bytes.NewReader(b))
    checkFail(test, err)

    if res.StatusCode != 200 {
        test.Errorf("Expected status code 200 got %d", res.StatusCode)
    }

    res, err = http.Get(testServer.URL + "/list")
    checkFail(test, err)

    body, err = ioutil.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode != 200 {
        test.Errorf("Expected status code 200 got %d", res.StatusCode)
    }

    err = json.Unmarshal(body, &lists)
    checkFail(test, err)

    if len(lists) != 1 {
        test.Errorf("Expected [] got %s", lists)
    }
}
