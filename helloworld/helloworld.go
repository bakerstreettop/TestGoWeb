// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package helloworld

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><meta charset='gb2312'/><title> ☎ BakerStreet </title><body>你好，旅行者。。!</body></html>")
}
