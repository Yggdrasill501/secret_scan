package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type BuildRequest struct {
    Parameter []Parameter `json:"parameter"`
}

type Parameter struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}

func main() {
    jenkinsURL := "http://your-jenkins-url/job/your-job-name/build" // replace with url to jenkins workflow
    username := os.Getenv("JENKINS_USER")
    apiToken := os.Getenv("JENKINS_API_TOKEN")

    buildRequest := BuildRequest{
        Parameter: []Parameter{
            {Name: "BITBUCKET_REPO", Value: "repo-url"}, //repository url
            {Name: "BITBUCKET_CREDENTIALS_ID", Value: "credentials-id"}, // replace with crendetials
        },
    }

    requestBody, err := json.Marshal(buildRequest)
    if err != nil {
        fmt.Println("Error marshalling JSON:", err)
        return
    }

    req, err := http.NewRequest("POST", jenkinsURL, bytes.NewBuffer(requestBody))
    if err != nil {
        fmt.Println("Error creating HTTP request:", err)
        return
    }
    req.SetBasicAuth(username, apiToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making HTTP request:", err)
        return
    }
    defer resp.Body.Close()

    fmt.Println("Triggered Jenkins build with status:", resp.Status)
}
