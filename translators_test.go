package translators

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bep/workers"
	qt "github.com/frankban/quicktest"
)

func TestGetTranslator(t *testing.T) {
	c := qt.New(t)

	d, _ := time.Parse("2006-Jan-02", "2018-Jan-06")

	c.Run("Basic", func(c *qt.C) {
		tnn := Get("nn_NO")
		c.Assert(tnn, qt.Not(qt.IsNil))
		c.Assert(tnn.MonthWide(d.Month()), qt.Equals, "januar")
	})

	c.Run("Para", func(c *qt.C) {
		p := workers.New(4)
		r, _ := p.Start(context.Background())

		for i := 0; i < 10; i++ {
			for _, locale := range []string{"nn_NO", "nn", "nyn", "sg", "se", "rwk", "mas"} {
				locale := locale
				r.Run(func() error {
					tnn := Get(locale)
					if tnn == nil {
						return errors.New("translator is nil")
					}

					if tnn.MonthWide(d.Month()) == "" {
						return errors.New("translator is invalid")
					}

					return nil
				})
			}
		}
	})

}
