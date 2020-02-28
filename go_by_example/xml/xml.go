package main

import (
	"encoding/xml"
	"fmt"
)

// Fields tags contain directives for encoder/decoder
// "id,attr" means that id is an attribute, not a node
type Plant struct {
	XMLName	xml.Name `xml:"plant"`
	Id		int		 `xml:"id,attr"`
	Name	string	 `xml:"name"`
	Origin	[]string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{
		Id:   27,
		Name: "Coffee",
		Origin: []string{"Ethiopia", "Brazil"},
	}

	// Using MarshalIndent for more human-readable output
	out, _ := xml.MarshalIndent(coffee, "", "  ")
	// Take a look that we added generic XML header!
	fmt.Println(xml.Header + string(out) + "\n")

	// List of elements
	lst := []*Plant{coffee, &Plant{
		Id:     28,
		Name:   "Tea",
		Origin: []string{"India", "Africa"},
	}}
	outL, _ := xml.MarshalIndent(lst, "", "  ")
	fmt.Println(xml.Header + string(outL))

	// Now we can Unmarshal these records back to structs
	// Be aware that Marshal an Unmarshal can be one-way functions, just like here
	var p []*Plant
	if err := xml.Unmarshal(outL, &p); err != nil {
		panic(err)
	}
	// Marshaled 2 plants, unmarshaled 1...
	fmt.Println(p)
	
	// Nesting with tags
	tomato := &Plant{
		Id:     81,
		Name:   "Tomato",
		Origin: []string{"Mexico", "California"},
	}
	
	type Nesting struct {
		XMLName	xml.Name `xml:"nesting"`
		Plants	[]*Plant `xml:"parent>child>plant"`
	}
	
	nesting := &Nesting{
		Plants: []*Plant{coffee, tomato},
	}

	out, _ = xml.MarshalIndent(nesting, "", "  ")
	fmt.Println(string(out))

	// Simple type - if node name is not set, then type name is used
	out1, _ := xml.Marshal(true)
	fmt.Println(string(out1))
}
