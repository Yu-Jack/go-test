package repository

import (
	"encoding/json"
	"io/ioutil"
	"jack-test/internal/dataservice"
	"sync"
)

var mu sync.Mutex
var fileName string = "data.json"

func (repository *repository) GetList() (dataservice.UserTodoList, error) {
	mu.Lock()
	defer mu.Unlock()

	var userList dataservice.UserTodoList
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		return userList, err
	}

	err = json.Unmarshal(content, &userList.Users)

	if err != nil {
		return userList, err
	}

	return userList, nil
}

func (repository *repository) SaveList(userList dataservice.UserTodoList) error {
	mu.Lock()
	defer mu.Unlock()

	jsonData, err := json.MarshalIndent(userList.Users, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, jsonData, 0644)

	if err != nil {
		return err
	}

	return nil
}
