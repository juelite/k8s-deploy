package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"k8s-deploy/k8s-client"
)

const (
	listen_addr = ":1711"
	err_params = `{"code": 201, "message": "参数解析失败或不正确", "date": []}`
	request_suc = `{"code": 200, "message": "success", "date": []}`
)

func main()  {
	http.HandleFunc("/create", create)
	http.HandleFunc("/update", update)

	fmt.Println("listen on ", listen_addr)
	err := http.ListenAndServe(listen_addr, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var (
		name, app, image string
		pods, port int32
	)
	err := r.ParseForm()
	params := r.Form
	if err != nil {
		io.WriteString(w, err_params)
		return
	}
	// name
	if val, ok := params["name"]; !ok || val[0] == "" {
		io.WriteString(w, err_params)
		return
	} else {
		name = val[0]
	}
	// app
	if val, ok := params["app"]; !ok || val[0] == "" {
		io.WriteString(w, err_params)
		return
	} else {
		app = val[0]
	}
	// image
	if val, ok := params["image"]; !ok || val[0] == "" {
		io.WriteString(w, err_params)
		return
	} else {
		image = val[0]
	}
	// pods
	if val, ok := params["pods"]; !ok || val[0] == "" {
		io.WriteString(w, err_params)
		return
	} else {
		v, _ := strconv.Atoi(val[0])
		pods = int32(v)
	}
	// port
	if val, ok := params["port"]; !ok || val[0] == "" {
		io.WriteString(w, err_params)
		return
	} else {
		v, _ := strconv.Atoi(val[0])
		port = int32(v)
	}

	err = k8s_client.Create(name, app, image, pods, port)
	if err != nil {
		request_err := `{"code": 202, "message": "`+err.Error()+`", "date": []}`
		io.WriteString(w, request_err)
	}
	io.WriteString(w, request_suc)
	return
}

func update(w http.ResponseWriter, r *http.Request) {
	var (
		name, image string
		pods int32
	)
	err := r.ParseForm()
	params := r.Form
	if err != nil {
		io.WriteString(w, err_params)
		return
	}
	// name
	if val, ok := params["name"]; !ok || val[0] == "" {
		io.WriteString(w, err_params)
		return
	} else {
		name = val[0]
	}
	// image
	if val, ok := params["image"]; !ok || val[0] == "" {
		io.WriteString(w, err_params)
		return
	} else {
		image = val[0]
	}
	// pods
	if val, ok := params["pods"]; !ok || val[0] == "" {
		io.WriteString(w, err_params)
		return
	} else {
		v, _ := strconv.Atoi(val[0])
		pods = int32(v)
	}
	err = k8s_client.Update(name, image, pods)
	if err != nil {
		request_err := `{"code": 202, "message": "`+err.Error()+`", "date": []}`
		io.WriteString(w, request_err)
	}
	io.WriteString(w, request_suc)
	return
}
