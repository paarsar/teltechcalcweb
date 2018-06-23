package controllers

import (
	"strings"
	m "teltechcalcweb/models"
	"testing"

	"github.com/astaxie/beego/context"
)

func TestActionCaching(t *testing.T) {

	tables := []struct {
		action string
		xstr   string
		ystr   string
		result int
		cached bool
	}{
		{"add", "2", "1", 3, false},
		{"add", "2", "1", 3, true},
		{"multiply", "2", "1", 2, false},
		{"multiply", "2", "1", 2, true},
	}

	for _, table := range tables {

		a := ArithmeticController{}
		a.Ctx = &context.Context{
			Input:  context.NewInput(),
			Output: context.NewOutput(),
		}

		a.Ctx.Input.SetParam(":action", table.action)
		a.Ctx.Input.SetParam("x", table.xstr)
		a.Ctx.Input.SetParam("y", table.ystr)
		a.Data = map[interface{}]interface{}{}

		a.Action()
		result := a.Data["json"].(m.Result)

		if table.result != result.Answer || table.cached != result.Cached {
			t.Errorf(" x: (%s == %d) y: (%s == %d) result: ((expected)%d == (actual)%d)  Cached: ((exp)%t == (act)%+v)", table.xstr, result.X, table.ystr, result.Y, table.result, result.Answer, table.cached, result.Cached)
		}

	}
}

func TestAction(t *testing.T) {

	tables := []struct {
		action string
		xstr   string
		ystr   string
		result int
		hasErr bool
	}{
		{"add", "2", "1", 3, false},
		{"sub", "2", "1", 1, false},
		{"multiply", "2", "1", 2, false},
		{"divide", "2", "1", 2, false},
		{"", "2", "1", 0, true},
		{"divide", "2", "0", 0, true},
	}

	for _, table := range tables {

		a := ArithmeticController{}
		a.Ctx = &context.Context{
			Input:  context.NewInput(),
			Output: context.NewOutput(),
		}

		a.Ctx.Input.SetParam(":action", table.action)
		a.Ctx.Input.SetParam("x", table.xstr)
		a.Ctx.Input.SetParam("y", table.ystr)
		a.Data = map[interface{}]interface{}{}

		a.Action()
		result := a.Data["json"].(m.Result)

		if table.result != result.Answer || table.hasErr == (len(result.Error) == 0) {
			t.Errorf(" x: (%s == %d) y: (%s == %d) result: ((expected)%d == (actual)%d)  has err: (%t == %+v)", table.xstr, result.X, table.ystr, result.Y, table.result, result.Answer, table.hasErr, result.Error)
		}

	}
}

func TestValidate(t *testing.T) {

	tables := []struct {
		xstr   string
		ystr   string
		x, y   int
		hasErr bool
	}{
		{"2", "1", 2, 1, false},
		{"2", "b", 2, 0, true},
	}

	for _, table := range tables {

		a := ArithmeticController{}
		a.Ctx = &context.Context{
			Input: context.NewInput(),
		}

		a.Ctx.Input.SetParam("x", table.xstr)
		a.Ctx.Input.SetParam("y", table.ystr)

		x, y, err := a.validate()

		if x != table.x || y != table.y || (err != nil) != table.hasErr {
			t.Errorf("x: (%s == %d) y: (%s == %d) has err: (%t == %+v)", table.xstr, table.x, table.ystr, table.y, table.hasErr, err)
		}
	}

}

func TestDiv(t *testing.T) {
	tables := []struct {
		x     int
		y     int
		check int
		err   bool
	}{
		{2, 1, 2, false},
		{1, 1, 1, false},
		{10, 5, 2, false},
		{-10, 5, -2, false},
		{-10, 0, -2, true},
	}

	for _, table := range tables {
		value, err := div(table.x, table.y)
		if (err != nil) != table.err {
			t.Errorf("err:%b want(%b)", err)
			continue
		}

		if err == nil && value != table.check {
			t.Errorf("x:%d y:%d value:%d want(%d)", table.x, table.y, value, table.check)
		}
	}
}
func TestSub(t *testing.T) {
	tables := []struct {
		x     int
		y     int
		check int
	}{
		{2, 1, 1},
		{1, 1, 0},
		{5, 4, 1},
		{-5, 4, -9},
	}

	for _, table := range tables {
		value, _ := sub(table.x, table.y)
		if value != table.check {
			t.Errorf("x:%d y:%d value:%d want(%d)", table.x, table.y, value, table.check)
		}
	}
}
func TestAdd(t *testing.T) {
	tables := []struct {
		x     int
		y     int
		check int
	}{
		{2, 1, 3},
		{1, 1, 2},
		{5, 4, 9},
		{-5, 4, -1},
	}

	for _, table := range tables {
		value, _ := add(table.x, table.y)
		if value != table.check {
			t.Errorf("x:%d y:%d value:%d want(%d)", table.x, table.y, value, table.check)
		}
	}
}
func TestMulti(t *testing.T) {
	tables := []struct {
		x     int
		y     int
		check int
	}{
		{2, 1, 2},
		{1, 1, 1},
		{5, 4, 20},
		{-5, 4, -20},
	}

	for _, table := range tables {
		value, _ := multi(table.x, table.y)
		if value != table.check {
			t.Errorf("x:%d y:%d value:%d want(%d)", table.x, table.y, value, table.check)
		}
	}
}

func TestFunctionName(t *testing.T) {
	var f = func(int, int) (int, error) {
		return 1, nil
	}

	name := functionName(f)

	if !strings.Contains(name, ".TestFunctionName.func1") {
		t.Error("naming incorrect")
	}
	return
}

func TestCreateCacheKeyNumbersInOrder(t *testing.T) {
	tables := []struct {
		x     int
		y     int
		check string
	}{
		{2, 1, "1:2"},
		{1, 1, "1:1"},
		{5, 4, "4:5"},
		{-5, 4, "-5:4"},
	}

	var f = func(int, int) (int, error) {
		return 1, nil
	}

	for _, table := range tables {
		value := createCacheKey(f, table.x, table.y)

		if strings.Contains(value, table.check) {
			t.Errorf("x:%d y:%d value:%s want(%s)", table.x, table.y, value, table.check)
		}
	}
}
