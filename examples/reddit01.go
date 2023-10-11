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
	vals, err := m.ValuesForPath("app.users.user", "-id:1")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	if len(vals) == 1 {
		myStruct.FromUser = vals[0].(map[string]interface{})["-name"].(string)
	}
	vals, err = m.ValuesForPath("app.users.user", "-id:2")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	if len(vals) == 1 {
		myStruct.ToUser = vals[0].(map[string]interface{})["-name"].(string)
	}
	vals, err = m.ValuesForPath("app.messages.message.#text")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	if len(vals) == 1 {
		myStruct.Message = vals[0].(string)
	}

	fmt.Printf("%#v\n", myStruct)
}
