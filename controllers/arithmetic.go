package controllers

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strings"
	m "teltechcalcweb/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Private Map of arithmetic actions. (Perhaps we can use for I18n)
var actionMap = map[string]func(int, int) (int, error){
	"add":      add,
	"sub":      sub,
	"subtract": sub,
	"div":      div,
	"divide":   div,
	"multiply": multi,
	"times":    multi,
	"X":        multi,
}

var resultCache = map[string]int{}

// ArithmeticController Web Controller that will handle all arithmetic expressions
type ArithmeticController struct {
	beego.Controller
}

// Add adds two integers x and y
// @router /:action [get]
func (c *ArithmeticController) Add() {

	var err error
	var x, y, answer int
	var action = c.Ctx.Input.Param(":action")
	var arithmetic func(int, int) (int, error)
	var ok bool
	cached := false
	defer func() {
		if err != nil {
			c.Data["json"] = m.Result{Action: action, Error: err.Error()}
			c.Ctx.Output.SetStatus(500)
			logs.Error("Error... %s", err.Error())
		} else {
			c.Data["json"] = m.Result{Action: action, Answer: answer, Cached: cached, X: x, Y: y}
		}
		c.ServeJSON(true)
	}()

	if x, y, err = c.validate(); err != nil {
		return
	} else if arithmetic, ok = actionMap[action]; !ok {
		err = fmt.Errorf("'%s' is not a recognizable action ", action)
		return
	} else if answer, ok = resultCache[createCacheKey(arithmetic, x, y)]; ok {
		cached = true
		return
	} else if answer, err = arithmetic(x, y); err != nil {
		return
	}

	resultCache[createCacheKey(arithmetic, x, y)] = answer
	logs.Info("%+v", resultCache)

}

// validate Validates our input
func (c *ArithmeticController) validate() (x, y int, err error) {
	if x, err = c.GetInt("x"); err != nil {
		err = fmt.Errorf("'%s' is not a valid integer", c.GetString("x"))
		return
	} else if y, err = c.GetInt("y"); err != nil {
		err = fmt.Errorf("'%s' is not a valid integer", c.GetString("y"))
		return
	}
	return
}

//add adds two numbers x and y
func add(x, y int) (answer int, err error) {
	answer = x + y
	return
}

//sub subtracts two numbers x and y (x - y)
func sub(x, y int) (answer int, err error) {
	answer = x - y
	return
}

//div divides two numbers x and y (x / y)
func div(x, y int) (answer int, err error) {
	if y == 0 {
		err = fmt.Errorf("Cannot Divid by zero")
	} else {
		answer = x / y
	}
	return
}

//multi multiplies two numbers x and y (x * y)
func multi(x, y int) (answer int, err error) {
	answer = x * y
	return
}

func createCacheKey(f func(int, int) (int, error), x, y int) string {
	ints := []int{x, y}
	sort.Ints(ints)
	return functionName(f) + "-" + strings.Trim(strings.Replace(fmt.Sprint(ints), " ", "", -1), "[]")
}

func functionName(i func(int, int) (int, error)) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
