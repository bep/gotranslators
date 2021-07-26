// Package wraps all translators in github.com/go-playground/locales.
// The translators are not created until asked for in Get.
package translators

import (
	"strings"
	"sync"

	"github.com/go-playground/locales"
)

var (
	// One normally only need a small subset of all the languages,
	// so delay creation until needed.
	mu              sync.RWMutex
	translatorFuncs = make(map[string]func() locales.Translator)
	translators     = make(map[string]locales.Translator)
)

// Get gets the Translator for the given locale, nil if not found.
func Get(locale string) locales.Translator {
	locale = strings.ToLower(locale)

	mu.RLock()
	t, found := translators[locale]
	if found {
		mu.RUnlock()
		return t
	}

	fn, found := translatorFuncs[locale]
	mu.RUnlock()
	if !found {
		return nil
	}

	mu.Lock()
	t = fn()
	translators[locale] = t
	mu.Unlock()

	return t

}
