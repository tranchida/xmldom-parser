package xmldom_test

import (
	"fmt"

	"github.com/Rodion-Bozhenko/xmldom-parser"
)

const (
	ExampleXml = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE junit SYSTEM "junit-result.dtd">
<testsuites>
	<testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/xmldom-parser">
		<properties>
			<property name="go.version">go1.8.1</property>
		</properties>
		<testcase classname="xmldom-parser" id="ExampleParseXML" time="0.004"></testcase>
		<testcase classname="xmldom-parser" id="ExampleParse" time="0.005"></testcase>
    <testcase xmlns:test="mock" id="AttrNamespace"></testcase>
	</testsuite>
</testsuites>`
)

func ExampleParseXML() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	fmt.Printf("name = %v\n", node.Name)
	fmt.Printf("attributes.len = %v\n", len(node.Attributes))
	fmt.Printf("children.len = %v\n", len(node.Children))
	fmt.Printf("root = %v", node == node.Root())
	// Output:
	// name = testsuites
	// attributes.len = 0
	// children.len = 1
	// root = true
}

func ExampleNode_GetAttribute() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	attr := node.FirstChild().GetAttribute("name")
	fmt.Printf("%v = %v\n", attr.Name, attr.Value)
	// Output:
	// name = github.com/subchen/xmldom-parser
}

func ExampleNode_GetChildren() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	children := node.FirstChild().GetChildren("testcase")
	for _, c := range children {
		fmt.Printf("%v: id = %v\n", c.Name, c.GetAttributeValue("id"))
	}
	// Output:
	// testcase: id = ExampleParseXML
	// testcase: id = ExampleParse
	// testcase: id = AttrNamespace
}

func ExampleNode_FindByID() {
	root := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	node := root.FindByID("ExampleParseXML")
	fmt.Println(node.XML())
	// Output:
	// <testcase classname="xmldom-parser" id="ExampleParseXML" time="0.004" />
}

func ExampleNode_FindOneByName() {
	root := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	node := root.FindOneByName("property")
	fmt.Println(node.XML())
	// Output:
	// <property name="go.version">go1.8.1</property>
}

func ExampleNode_FindByName() {
	root := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	nodes := root.FindByName("testcase")
	for _, node := range nodes {
		fmt.Println(node.XML())
	}
	// Output:
	// <testcase classname="xmldom-parser" id="ExampleParseXML" time="0.004" />
	// <testcase classname="xmldom-parser" id="ExampleParse" time="0.005" />
	// <testcase xmlns:test="mock" id="AttrNamespace" />
}

func ExampleNode_Query() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	// xpath expr: https://github.com/antchfx/xpath

	// find all children
	fmt.Printf("children = %v\n", len(node.Query("//*")))

	// find node matched tag name
	nodeList := node.Query("//testcase")
	for _, c := range nodeList {
		fmt.Printf("%v: id = %v\n", c.Name, c.GetAttributeValue("id"))
	}
	// Output:
	// children = 6
	// testcase: id = ExampleParseXML
	// testcase: id = ExampleParse
	// testcase: id = AttrNamespace
}

func ExampleNode_QueryOne() {
	node := xmldom.Must(xmldom.ParseXML(ExampleXml)).Root
	// xpath expr: https://github.com/antchfx/xpath

	// find node matched attr name
	c := node.QueryOne("//testcase[@id='ExampleParseXML']")
	fmt.Printf("%v: id = %v\n", c.Name, c.GetAttributeValue("id"))
	// Output:
	// testcase: id = ExampleParseXML
}

func ExampleDocument_XML() {
	doc := xmldom.Must(xmldom.ParseXML(ExampleXml))
	fmt.Println(doc.XML())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?><!DOCTYPE junit SYSTEM "junit-result.dtd"><testsuites><testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/xmldom-parser"><properties><property name="go.version">go1.8.1</property></properties><testcase classname="xmldom-parser" id="ExampleParseXML" time="0.004" /><testcase classname="xmldom-parser" id="ExampleParse" time="0.005" /><testcase xmlns:test="mock" id="AttrNamespace" /></testsuite></testsuites>
}

func ExampleDocument_XMLPretty() {
	doc := xmldom.Must(xmldom.ParseXML(ExampleXml))
	fmt.Println(doc.XMLPretty())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <!DOCTYPE junit SYSTEM "junit-result.dtd">
	// <testsuites>
	//   <testsuite tests="2" failures="0" time="0.009" name="github.com/subchen/xmldom-parser">
	//     <properties>
	//       <property name="go.version">go1.8.1</property>
	//     </properties>
	//     <testcase classname="xmldom-parser" id="ExampleParseXML" time="0.004" />
	//     <testcase classname="xmldom-parser" id="ExampleParse" time="0.005" />
	//     <testcase xmlns:test="mock" id="AttrNamespace" />
	//   </testsuite>
	// </testsuites>
}

func ExampleNewDocument() {
	doc := xmldom.NewDocument("testsuites")

	testsuiteNode := doc.Root.CreateNode("testsuite").SetAttributeValue("name", "github.com/subchen/xmldom-parser")
	testsuiteNode.CreateNode("testcase").SetAttributeValue("name", "case 1").Text = "PASS"
	testsuiteNode.CreateNode("testcase").SetAttributeValue("name", "case 2").Text = "FAIL"

	fmt.Println(doc.XMLPretty())
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <testsuites>
	//   <testsuite name="github.com/subchen/xmldom-parser">
	//     <testcase name="case 1">PASS</testcase>
	//     <testcase name="case 2">FAIL</testcase>
	//   </testsuite>
	// </testsuites>
}
