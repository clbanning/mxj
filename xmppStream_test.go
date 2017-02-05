package mxj

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestXMPPStreamTag(t *testing.T) {
	fmt.Println("----------- TestXMPPStreamTag ...")
	var data = `
<stream:stream
    from='example.com'
    xmlns="jabber:client"
    xmlns:stream="http://etherx.jabber.org/streams"
    version="1.0">
<stream:features>
  <bind xmlns="urn:ietf:params:xml:ns:xmpp-bind"/>
  <sm xmlns="urn:xmpp:sm:3"/>
</stream:features>
<stream:stream>`

	HandleXMPPSteamTag()
	buf := bytes.NewBufferString(data)
	for {
		m, raw, err := NewMapXmlReaderRaw(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal("err:", err)
		}
		fmt.Println(string(raw))
		fmt.Println(m)
	}
}
