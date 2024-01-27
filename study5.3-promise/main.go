package main

import (
	"errors"
	"fmt"
	"time"
)

type Promise struct {
	callbacks []func(interface{})
	error     error
	data      interface{}
}

func (p *Promise) Then(callback func(interface{})) *Promise {
	p.callbacks = append(p.callbacks, callback)
	return p
}

func (p *Promise) Catch(callback func(error)) *Promise {
	if p.error != nil {
		callback(p.error)
	}
	return p
}

func (p *Promise) Resolve(data interface{}) {
	p.data = data
	go func() {
		for _, callback := range p.callbacks {
			callback(p.data)
		}
	}()
}

func (p *Promise) Reject(err error) {
	p.error = err
}

func fetchData() *Promise {
	p := new(Promise)

	go func() {
		data := "Data fetched successfully!"
		err := errors.New("Something went wrong")
		//var err error
		if err != nil {
			p.Reject(err)
		} else {
			p.Resolve(data)
		}
	}()
	//fmt.Println("end")

	return p
}

func main() {
	p := fetchData()

	p.Then(func(data interface{}) {
		fmt.Println("Data:", data)
	}).Catch(func(err error) {
		fmt.Println("Error:", err)
	})

	time.Sleep(7 * time.Second)
}
