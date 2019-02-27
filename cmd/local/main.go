package main

import (
	convvls "github.com/loivis/convolvulus-update"
)

func main() {
	m := convvls.Message{
		Data: []byte(`[{"title": "我真的不无敌", "site": "起点中文网", "id": "1013723616", "author": "习仁"}]`),
	}
	_ = convvls.Update(nil, m)
}
