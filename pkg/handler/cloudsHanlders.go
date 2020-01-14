// Copyright (c) 2020 Dat Vu Tuan <tuandatk25a@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/vCloud-DFTBA/faythe-ui/pkg/config"
	"github.com/vCloud-DFTBA/faythe-ui/pkg/model"
)

func listClouds(w http.ResponseWriter, r *http.Request) {
	cfg := config.Get()
	u := model.MakeURL(cfg.FaytheURL, "clouds")
	client := &http.Client{}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth(cfg.FaytheUsername, cfg.FaythePassword)
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyText)
}

func createCloud(w http.ResponseWriter, r *http.Request) {
	cfg := config.Get()
	u := model.MakeURL(cfg.FaytheURL, "clouds", "openstack")
	client := &http.Client{}
	buf, _ := ioutil.ReadAll(r.Body)

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(buf))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth(cfg.FaytheUsername, cfg.FaythePassword)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyText)
}

func deleteCloud(w http.ResponseWriter, r *http.Request) {
	cfg := config.Get()
	vars := mux.Vars(r)
	u := model.MakeURL(cfg.FaytheURL, "clouds", vars["id"])
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", u, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth(cfg.FaytheUsername, cfg.FaythePassword)
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyText)
}
