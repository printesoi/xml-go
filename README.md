# xml [![GoDoc](https://godoc.org/github.com/printesoi/xml-go?status.svg)](https://godoc.org/github.com/printesoi/xml-go) [![Tests](https://github.com/printesoi/xml-go/actions/workflows/test.yml/badge.svg)](https://github.com/printesoi/xml-go/actions/workflows/test.yml) [![Coverage Status](https://coveralls.io/repos/github/printesoi/xml-go/badge.svg)](https://coveralls.io/github/printesoi/xml-go) [![Go Report Card](https://goreportcard.com/badge/github.com/printesoi/xml-go)](https://goreportcard.com/report/github.com/printesoi/xml-go)

Package `github.com/printesoi/xml-go` implements a XML 1.0 marshaler that is a drop in `encoding/xml` replacement maintaining full interface compatibility to it but aims to producing [C14N Exclusive XML Canonicalization](https://www.w3.org/TR/xml-exc-c14n/) compatible byte sequences. This includes sorting of all rendered element attributes as per the C14N-XML spec.

## Example
A Struct defined in the normal way as known from `encoding/xml`

```go
type Person struct {
	XMLName xml.Name `xml:"http://example.com/ns1 person"`
	Name    string   `xml:"name"`
	Phone   string   `xml:"http://example.com/ns2 phone,omitempty"`
}
```
is marshaled to
```xml
<ns1:person xmlns:ns1="http://example.com/ns1">
  <ns1:name>Foo</ns1:name>
  <ns2:phone xmlns:ns2="http://example.com/ns2">123</ns2:phone>
</ns1:person>
```

## Example 2

Skip adding xmlns:prefix="url":

```go
type Ambassador struct {
	XMLName xml.Name `xml:"example.com:ns1 ambassador"`
	Name    string   `xml:"name"`
	Origin  string   `xml:"example.com:ns2 origin,omitempty"`
}

var b bytes.Buffer
enc := NewEncoder(&b)

prefix1, url1 := "ns1", "example.com:ns1"
enc.AddNamespaceBinding(url1, prefix1)
enc.AddSkipNamespaceAttrForPrefix(url1, prefix1)

prefix2, url2 := "ns2", "example.com:ns2"
enc.AddNamespaceBinding(url2, prefix2)
enc.AddSkipNamespaceAttrForPrefix(url2, prefix2)

v := &Ambassador{
    Name: "Mahit Dzmare",
    Origin: "Lsel Station",
}

if err := enc.Encode(v); err != nil {
    // error
}
if err := enc.Close(); err != nil {
    // error
}

ambassadorXML := b.Bytes()
```
is marshaled to
```xml
<ns1:ambassador>
  <ns1:name>Mahit Dzmare</ns1:name>
  <ns2:origin>Lsel Station</ns2:origin>
</ns1:ambassador>
```
