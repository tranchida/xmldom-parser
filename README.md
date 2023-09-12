# xmldom-parser

It is a fork from subchen/go-xmldom.

## Motivation for fork

Since the original repo seems to be inactive, this repo is made to add the updates needed for my needs.
Many thanks for subchen's work!

[Original Repo](https://github.com/subchen/go-xmldom)

XML DOM processing for Golang, supports xpath query

- Parse XML into dom
- Query node using xpath
- Update XML using dom

## Installation

```bash
$ go get github.com/Rodion-Bozhenko/xmldom-parser
```

## Basic Usage

```go
xml := `<testsuite tests="2" failures="0" time="0.009" name="github.com/Rodion-Bozhenko/xmldom-parser">
    <testcase classname="xmldom-parser" name="ExampleParseXML" time="0.004"></testcase>
    <testcase classname="xmldom-parser" name="ExampleParse" time="0.005"></testcase>
</testsuite>`

doc := xmldom.Must(xmldom.ParseXML(xml))
root := doc.Root

name := root.GetAttributeValue("name")
time := root.GetAttributeValue("time")
fmt.Printf("testsuite: name=%v, time=%v\n", name, time)

for _, node := range root.GetChildren("testcase") {
    name := node.GetAttributeValue("name")
    time := node.GetAttributeValue("time")
    fmt.Printf("testcase: name=%v, time=%v\n", name, time)
}
```

## Xpath Query

```go
// find all children
fmt.Printf("children = %v\n", len(node.Query("//*")))

// find node matched tag name
nodeList := node.Query("//testcase")
for _, c := range nodeList {
    fmt.Printf("%v: name = %v\n", c.Name, c.GetAttributeValue("name"))
}

// find node matched attr name
c := node.QueryOne("//testcase[@name='ExampleParseXML']")
fmt.Printf("%v: name = %v\n", c.Name, c.GetAttributeValue("name"))
```

## Create XML

```go
doc := xmldom.NewDocument("testsuites")

suiteNode := doc.Root.CreateNode("testsuite").SetAttributeValue("name", "github.com/Rodion-Bozhenko/xmldom-parser")
suiteNode.CreateNode("testcase").SetAttributeValue("name", "case 1")
suiteNode.CreateNode("testcase").SetAttributeValue("name", "case 2")

fmt.Println(doc.XML())
```

## License

`xmldom-parser` is released under the Apache 2.0 license. See [LICENSE](https://github.com/Rodion-Bozhenko/xmldom-parser/blob/master/LICENSE)
