package main

import (
	"bytes"
	"net/http"
  "os/exec"
  "os"
  "encoding/json"
  "io/ioutil"
  "strings"
)

const (
	PUSH_URL   = "https://api.pushbullet.com/v2/pushes"
  DEVICE_URL = "https://api.pushbullet.com/v2/devices"
)

type Push struct {
  Type string `json:"type"`
  Title string `json:"title"`
  Message string `json:"body"`
  Device string `json:"device"`
}

type Device struct {
  Active bool `json:"active"`
  Iden string `json:"iden"`
}

type Container struct {
	Devices []Device `json:"devices"`
}

func GetDevices(token string) []Device {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", DEVICE_URL, nil)
	req.Header.Set("Authorization", "Bearer " + token)
	res, _ := client.Do(req)

	decoder := json.NewDecoder(res.Body)

	var container Container

	err := decoder.Decode(&container)
	if err != nil {
		panic(err)
	}
	return container.Devices
}

func DoPush(push *Push, token string) {
	client := &http.Client{}
  body, _ := json.Marshal(*push)
	req, _ := http.NewRequest("POST", PUSH_URL, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", "Bearer " + token)
	req.Header.Set("Content-Type", "application/json")
	client.Do(req)
}

func GetToken() (string, error) {
  dir, err := exec.Command("bash", "-c", "echo -n $HOME").Output()
  if err != nil {
    return "", err
  }
  res, err := ioutil.ReadFile(string(dir) + "/.longrun-token")
  return string(res), err
}


func main() {
  token, err := GetToken()
  if err != nil {
    panic(err)
  }
  
  devices := GetDevices(token)
  
  cmd := strings.Join(os.Args[1:], " ")
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
  for d := 0; d < len(devices); d++ {
    if devices[d].Active {
      push.Device = devices[d].Iden
      DoPush(&push, token)
    }
  }
}