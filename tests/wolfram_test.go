package tests

import (
	"testing"

	"github.com/Krognol/go-wolfram/wolfram"
)

func TestGetQueryResult(t *testing.T) {
	c := &wolfram.Client{AppID: "some appid"}

	_, err := c.GetQueryResult("What is the price of gold?", nil)
	if err != nil {
		t.Failed()
		t.Log(err.Error())
	}
}

func TestGetSimpleQueryResult(t *testing.T) {
	c := &wolfram.Client{AppID: "some appid"}

	_, _, err := c.GetSimpleQuery("What is the price of gold?", nil)
	if err != nil {
		t.Failed()
		t.Log(err.Error())
	}
}

func TestGetFastQueryRecognizerResult(t *testing.T) {
	c := &wolfram.Client{AppID: "some appid"}

	_, err := c.GetFastQueryRecognizer("Gold price", wolfram.Default)
	if err != nil {
		t.Failed()
		t.Log(err.Error())
	}
}

func TestGetShortAnswerQueryResult(t *testing.T) {
	c := &wolfram.Client{AppID: "some appid"}

	_, err := c.GetShortAnswerQuery("Price of gold", wolfram.Metric, 0)
	if err != nil {
		t.Failed()
		t.Log(err.Error())
	}
}

func TestGetSpokenAnswerResult(t *testing.T) {
	c := &wolfram.Client{AppID: "some appid"}

	_, err := c.GetSpokentAnswerQuery("Price of gold", wolfram.Metric, 0)
	if err != nil {
		t.Failed()
		t.Log(err.Error())
	}
}
