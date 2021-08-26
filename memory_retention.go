package memory_retention

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

// This global variable used save client answers.
// - Key is a hash value that identifies an individual.
var data = make(map[string][]string)

// Topics specified by the presenter.
var topics = make(map[string]string)

// Save whether the key is in use.
// If the key does not exist in this map, it is not used.
var keyMap = make(map[string]bool)

// Create a key.
//
// Arguments:
// hash {string} - hash value.
func CreateKey(hash string) {
	mutex.Lock()
	keyMap[hash] = true
	mutex.Unlock()
}

// Deletes the specified key.
//
// Arguments:
// key {string} - Key value.
func DeleteKey(key string) error {
	mutex.Lock()
	if err := exist(key); err != nil {
		return err
	}

	delete(keyMap, key)
	delete(data, key)
	delete(topics, key)
	mutex.Unlock()
	return nil
}

// Clear the answer.
//
// Arguments:
// key {string} - Key value.
func DeleteAnswer(key string) error {
	mutex.Lock()
	if err := exist(key); err != nil {
		return err
	}
	data[key] = []string{}
	mutex.Unlock()
	return nil
}

// All clear.
func DeleteAll() {
	data = make(map[string][]string)
	keyMap = make(map[string]bool)
	topics = make(map[string]string)
}

// Check if the key exists.
// If the key does not exist, an error will be thrown.
//
// Arguments:
// key {string} - Key value.
func exist(key string) error {
	_, exist := keyMap[key]
	if !exist {
		return fmt.Errorf("Key is not exist.")
	}
	return nil
}

// Add an answer.
//
// Arguments:
// key {string} - Key value.
// value {string} - Answer.
func AddAnswer(key string, value string) error {
	mutex.Lock()
	if err := exist(key); err != nil {
		return err
	}
	if data[key] == nil {

	}
	data[key] = append(data[key], value)
	mutex.Unlock()
	return nil
}

// Returns a list of answers.
//
// Arguments:
// key {string} - Key value.
//
// Returns:
// {[]string} - Answers.
func GetAnswer(key string) ([]string, error) {
	mutex.Lock()
	if err := exist(key); err != nil {
		return nil, err
	}
	answer := data[key]
	mutex.Unlock()

	return answer, nil
}

// Set a topic.
//
// Arguments:
// key {string} - Key value.
// value {string} - topic
func SetTopic(key string, value string) error {
	mutex.Lock()
	if err := exist(key); err != nil {
		return err
	}
	topics[key] = value
	mutex.Unlock()

	return nil
}

// Get topic
//
// Arguments:
// key {string} - key value.
//
// Returns:
// {string} - topic.
func GetTopic(key string) (string, error) {
	mutex.Lock()
	if err := exist(key); err != nil {
		return "", err
	}
	topic := topics[key]
	mutex.Unlock()

	return topic, nil
}
