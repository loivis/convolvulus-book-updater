package main

import (
	convvls "github.com/loivis/convolvulus-update"
)

func main() {
	m := convvls.Message{
		Data: []byte(`[{"title": "foo"}]`),
	}
	_ = convvls.Update(nil, m)
}
