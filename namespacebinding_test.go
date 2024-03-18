// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"bytes"
	"testing"
)

func TestAdd(t *testing.T) {
	url := "http://binding-test-example.com/abc"
	err := NameSpaceBinding.Add(url, "xyz")
	if err != nil {
		t.Error(err)
	}

	type Person struct {
		XMLName Name `xml:"http://binding-test-example.com/abc person"`
		Id      int  `xml:"id,attr"`
	}

	v := &Person{Id: 13}
	want := `<xyz:person xmlns:xyz="http://binding-test-example.com/abc" id="13"></xyz:person>`
	got, err := Marshal(v)
	if err != nil {
		t.Error(err)
	}
	if string(got) != want {
		t.Errorf("got `%s`, want `%s`", got, want)
	}

}

func TestClear(t *testing.T) {
	url := "http://binding-test-example.com/abc"
	err := NameSpaceBinding.Add(url, "xyz")
	if err != nil {
		t.Error(err)
	}
	if NameSpaceBinding.get(url) != "xyz" {
		t.Error("binding was not set")
	}
	NameSpaceBinding.Clear()
	if NameSpaceBinding.get(url) == "xyz" {
		t.Error("binding was not cleared")
	}
}

func TestSkipNamespaceAttrForPrefix(t *testing.T) {
	var b bytes.Buffer
	enc := NewEncoder(&b)

	prefix1, url1 := "ns1", "binding-test-example.com:ns1"
	enc.AddNamespaceBinding(url1, prefix1)
	enc.AddSkipNamespaceAttrForPrefix(url1, prefix1)

	prefix2, url2 := "ns2", "binding-test-example.com:ns2"
	enc.AddNamespaceBinding(url2, prefix2)
	enc.AddSkipNamespaceAttrForPrefix(url2, prefix2)

	type Person struct {
		XMLName Name `xml:"binding-test-example.com:ns1 person"`
		Id      int  `xml:"binding-test-example.com:ns2 id"`
	}

	v := &Person{Id: 42}

	if err := enc.Encode(v); err != nil {
		t.Error(err)
	}
	if err := enc.Close(); err != nil {
		t.Error(err)
	}

	got := string(b.Bytes())
	want := `<ns1:person><ns2:id>42</ns2:id></ns1:person>`
	if string(got) != want {
		t.Errorf("got `%s`, want `%s`", got, want)
	}
}
