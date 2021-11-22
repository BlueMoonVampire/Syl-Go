### Syl-Go

Golang wrapper for the Sylviorus antispam API for telegram

```
package test

import (
	"testing"

	"github.com/moezilla/Syl-Go"
)

func syl(w *testing.W) {
	_, err := Syl-Go.syl("1821151467")
	if err != nil {
		w.Error(err)
	}
}

```
