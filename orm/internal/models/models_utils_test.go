package models

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type NotApplicableModel struct {
	Id int
}

func (n *NotApplicableModel) IsApplicableTableForDB(db string) bool {
	return db == "default"
}

func TestIsApplicableTableForDB(t *testing.T) {
	assert.False(t, IsApplicableTableForDB(reflect.ValueOf(&NotApplicableModel{}), "defa"))
	assert.True(t, IsApplicableTableForDB(reflect.ValueOf(&NotApplicableModel{}), "default"))
}

func TestSnakeString(t *testing.T) {
	camel := []string{"PicUrl", "HelloWorld", "HelloWorld", "HelLOWord", "PicUrl1", "XyXX"}
	snake := []string{"pic_url", "hello_world", "hello_world", "hel_l_o_word", "pic_url1", "xy_x_x"}

	answer := make(map[string]string)
	for i, v := range camel {
		answer[v] = snake[i]
	}

	for _, v := range camel {
		res := SnakeString(v)
		if res != answer[v] {
			t.Error("Unit Test Fail:", v, res, answer[v])
		}
	}
}

func TestSnakeStringWithAcronym(t *testing.T) {
	camel := []string{"ID", "PicURL", "HelloWorld", "HelloWorld", "HelLOWord", "PicUrl1", "XyXX"}
	snake := []string{"id", "pic_url", "hello_world", "hello_world", "hel_lo_word", "pic_url1", "xy_xx"}

	answer := make(map[string]string)
	for i, v := range camel {
		answer[v] = snake[i]
	}

	for _, v := range camel {
		res := SnakeStringWithAcronym(v)
		if res != answer[v] {
			t.Error("Unit Test Fail:", v, res, answer[v])
		}
	}
}

func TestCamelString(t *testing.T) {
	snake := []string{"pic_url", "hello_world_", "hello__World", "_HelLO_Word", "pic_url_1", "pic_url__1"}
	camel := []string{"PicUrl", "HelloWorld", "HelloWorld", "HelLOWord", "PicUrl1", "PicUrl1"}

	answer := make(map[string]string)
	for i, v := range snake {
		answer[v] = camel[i]
	}

	for _, v := range snake {
		res := CamelString(v)
		if res != answer[v] {
			t.Error("Unit Test Fail:", v, res, answer[v])
		}
	}
}
