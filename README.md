### Vanitas-Go

Golang wrapper for the Vanitas antispam API for telegram

```
package test

import (
	"testing"

	"github.com/moezilla/Vanitas-Go"
)

func syl(w *testing.W) {
	_, err := Vanitas-Go.vanitas("1821151467")
	if err != nil {
		w.Error(err)
	}
}

```
