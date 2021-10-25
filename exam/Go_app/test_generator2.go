package main

import (
	
	"testing"
	"https://github.com/dailymotion/allure-go"
	"https://github.com/GabbyyLS/allure-go-common"
)

func TestHelloWorld(t *testing.T) {
    allure.Test(t,
                allure.Description("автоматизированный тест"),
                allure.Action(func() {
                    s := "Hello world"
                    if len(s) == 0 {
                      	t.Errorf("Expected 'hello world' string, but got %s ", s)
                    }
                }))
}
