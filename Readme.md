
## shallowXML
I was too lazy to read about available libs, so I started my own XML parser. I started learning Go only yesteray, so I couldn't really do everything the way I wanted (e.g. I don't konw how to declare methods, so I used functions).

The end goal is to write a parser that is linear to the size of the input string and also stores nothing more than that single string in memory (the way I'll do it is to store the string in the root element and use pointers in all of it's children.)... maybe even less.

I will try to get similar API as that of [lxml](https://lxml.de/) / [elementree](http://effbot.org/zone/element.htm).

## How to use

First argument should be path of an XML file, e.g.:

```shell
$ go run parser.go example.xml
```

and output is:

```bash
<staff> ()  2
 | <person> ()  3
 |  | <firstname> (Jack)  0
 |  | <lastname> (McGreen)  0
 |  | <address> ()  2
 |  |  | <city> (London)  0
 |  |  | <country> (UK)  0
 | <person> ()  2
 |  | <firstname> (Anna)  0
 |  | <lastname> (Krikorian)  0
```

## Disclaimer

Of course this is just a toy-project / draft. No checks are done to detect invalid XML syntax, no error handling. XML elements with attributes will not be recognized.