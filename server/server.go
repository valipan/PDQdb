/*
 * Copyright (C) 2015-present Adam Hanna <ahanna@alumni.mines.edu>
 * Copyright (C) 2015-present Jonathan Barronville <jonathan@belairlabs.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package server

import (
	"github.com/adam-hanna/PDQdb/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func StartServer(hostname string, port uint16) error {
	mx := mux.NewRouter()
	mx.HandleFunc("/key/{key}", ProcessKey)
	mx.HandleFunc("/", serveMainRoute)

	var host string = hostname + strconv.Itoa(int(port))

	return http.ListenAndServe(host, mx)

}

func serveMainRoute(res http.ResponseWriter, req *http.Request) {
	// write the headers
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "text/plain")

	// send back the response
	res.Write([]byte("Pretty Damn Quick!\n"))
}

func ProcessKey(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		// extract the key from the url of the request
		vars := mux.Vars(req)
		key := vars["name"]

		//grab the key that the user is looking for
		var val []byte = data.DataSet[key]

		// write the headers
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "text/plain")

		// send back the response
		res.Write(val)
	} else {
		// do POST / PUT stuff

		// write the headers
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "text/plain")

		// send back the response
		res.Write([]byte("Pretty Damn Quick!\n"))
	}
}