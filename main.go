package main

import (
	"bytes"
	"net/http"
  "os/exec"
  "os"
  "encoding/json"
  "io/ioutil"
  "os/user"
)

const (
	PUSH_URL  = "https://api.pushbullet.com/v2/pushes"
)

type Push struct {
  Type string `json:"type"`
  Title string `json:"title"`
  Message string `json:"body"`
  Device string `json:"device"`
}


func DoPush(push *Push, token string) {
	client := &http.Client{}
  body, _ := json.Marshal(*push)
	req, _ := http.NewRequest("POST", PUSH_URL, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	client.Do(req)
}

func GetToken() (string, error) {
  usr, err := user.Current()
  if err != nil {
    return "", err  
  }
  res, err := ioutil.ReadFile(usr.HomeDir + "/.longrun-token")
  return string(res), err
}


func main() {
  token, err := GetToken()
  if err != nil {
    panic(err)
  }
  
  for i := 1; i < len(os.Args); i++ {
    cmd := os.Args[1]
    out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
  
    failed := ""
    if err != nil {
      failed = "Failed: "
    }
    push := Push {
      Type: "note",
      Title: failed + cmd,
      Message: string(out),
      Device: "",
    }
    DoPush(&push, token)
  }
}