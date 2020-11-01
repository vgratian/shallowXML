package main

import (
    "os"
    "io/ioutil"
    "fmt"
    "strings"
)

var RED string = "\033[31m"
var PURPLE string = "\033[35m"
var CYAN string = "\033[36m"
var BOLD string = "\033[1m"
var END string = "\033[0m"


type Element struct {
    xml string
    i, j, m, n int
}

func NewElement(xml string) Element {
    /*
    Creates a new XML element. All we do here is just the boundaries
    of the XML tags and store as indices
    */
    var e Element
    e.xml = xml           // raw string XML
    e.i = 0               // index of opening bracket of opening tag
    e.j = 1               // index of closing bracket of opening tag
    e.n = len(xml) - 1    // index of opening bracket of closing tag
    e.m = e.n - 1         // index of closing bracket of closing tag

    for xml[e.j] != '>' {
        e.j++
    }
    for xml[e.m:e.m+2] != "</" {
        e.m--
    }
    return e
}


func GetLabel(e Element) string {
    /*
    Returns label of the XML element

    Examle:
    if e.xml = "<version>1.0</version>", 
    will return: "version"

    */
    return e.xml[e.i+1:e.j]
}


func GetContent(e Element) string {
    /*
    Returns content of the XML element

    Examle:
    if e.xml = "<version>1.0</version>", 
    will return: "1.0"

    */
    var c string
    if e.xml[e.j+1] != '<' {
        c = e.xml[e.j+1:e.m]
    }
    return c
}


func GetChildren(e Element) []Element {
    /*
    Returns list of direct children of element
    */   
    var children []Element
    var k, p, q int
    var open_tag string


    children = []Element{}

    if e.xml[e.j+1] == '<' {
        p = e.j + 1
	    open_tag = ""
        
        for p < e.m {

            if e.xml[p] == '>' {

                if e.xml[k+1] != '/' {

                    tag := e.xml[k+1:p]

                    if open_tag == "" {
                        open_tag = tag
                        q = k
                    } 

                } else {
                    tag := e.xml[k+2:p]

                    if open_tag == tag {
                        s := e.xml[q:p+1]
                        child := NewElement(s)
                        children = append(children, child)
                        q = p
                        open_tag = ""
                    } 
                }

            } else if e.xml[p] == '<' {
                k = p
            } 

            p++
        }
    }

    return children
}


func PrintRecursively(e Element, depth int) {
    /*
    Print element and its children recursively. Indentation is used
    to highlight depth of the element. On each line 3 values are printed:
      - label of element    (in pink and between angle brackets)
      - content of element (in cyan and between round brackets)
      - number of children (in red)
    */

    children := GetChildren(e)  
    // Print indentation and tag name
    fmt.Printf("%s<%s%s%s%s> ", strings.Repeat(" | ", depth), BOLD, PURPLE, GetLabel(e), END)
    // Print content in brackets
    fmt.Printf("(%s%s%s)  ", CYAN, GetContent(e), END)
    // Print num children
    fmt.Printf("%s%d%s\n", RED, len(children), END)

    for i := range(children) {
        PrintRecursively(children[i], depth+1)
    }
}

func main() {

    var xml []byte
    var root_element Element

    if len(os.Args) != 2 {
        fmt.Println("Expected path to XML file")
        os.Exit(1)
    }

    xml, err := ioutil.ReadFile(os.Args[1])

    if err != nil {
        fmt.Printf("Reading XML file [%s] failed:\n", os.Args[1])
        fmt.Println(err)
        os.Exit(1)
    }

    root_element = NewElement(string(xml))

    PrintRecursively(root_element, 0)
}




