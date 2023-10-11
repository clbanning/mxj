// https://www.reddit.com/r/golang/comments/9eclgy/xml_unmarshaling_internal_references/

package main

import (
	"fmt"

	"github.com/clbanning/mxj/v2"
)

var data = []byte(`<app>
   <users>
     <user id="1" name="Jeff" />
     <user id="2" name="Sally" />
   </users>
   <messages>
     <message id="1" from_user="1" to_user="2">Hello!!</message>
   </messages>
</app>`)

func main() {
	m, err := mxj.NewMapXml(data)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%v\n", m)

	type mystruct struct {
		FromUser string
		ToUser   string
		Message  string
	}
	myStruct := mystruct{}
	val, err := m.ValueForKey("user", "-id:1")
	if val != nil {
		myStruct.FromUser = val.(map[string]interface{})["-name"].(string)
	} else {
		// if there no val, then err is at least KeyNotExistError
		fmt.Println("err:", err)
		return
	}
	val, err = m.ValueForKey("user", "-id:2")
	if val != nil {
		myStruct.ToUser = val.(map[string]interface{})["-name"].(string)
	} else {
		// if there no val, then err is at least KeyNotExistError
		fmt.Println("err:", err)
		return
	}
	val, err = m.ValueForKey("#text")
	if val != nil {
		myStruct.Message = val.(string)
	} else {
		// if there no val, then err is at least KeyNotExistError
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("%#v\n", myStruct)
}
