package localization

import (
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	localizerCache = make(map[string]*i18n.Localizer)
	cacheMutex     = &sync.RWMutex{}
)

func getLocalizerFromCache(lang string) (*i18n.Localizer, bool) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	localizer, ok := localizerCache[lang]
	return localizer, ok
}

func setLocalizerInCache(lang string, localizer *i18n.Localizer) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	localizerCache[lang] = localizer
}
