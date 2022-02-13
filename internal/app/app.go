package app

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	jsoniter "github.com/json-iterator/go"

	"github.com/shootingfans/redis-dashboard/internal/locales"
	"github.com/shootingfans/redis-dashboard/internal/logger"
	"github.com/shootingfans/redis-dashboard/pkg/event"
	"github.com/shootingfans/redis-dashboard/pkg/utils"
)

const appUniqueId = "com.shootingfans.redis-dashboard"
const appVersion = "v0.1.0"

const (
	preferenceKeyOfSettingLanguage  = "SettingLanguage"
	preferenceKeyOfTheme            = "SettingTheme"
	preferenceKeyOfRecordWindowSize = "SettingRecordWindowSize"
	preferenceKeyOfAppSize          = "ApplicationWindowSize"
	preferenceKeyOfRedisConfigList  = "RedisConfigList"
)

const (
	defaultMainAppWindowsWidth  = 800
	defaultMainAppWindowsHeight = 600
	defaultSettingDialogWidth   = 300
	defaultSettingDialogHeight  = 100
	defaultNewRedisDialogWidth  = 400
	defaultNewRedisDialogHeight = 200
)

const (
	eventNameOfRebootWindow      = "event.reboot.window"
	eventNameOfThemeChanged      = "event.theme.changed"
	eventNameOfToggleLeftMenu    = "event.toggle.left.menu"
	eventNameOfCloseWindow       = "event.close.window"
	eventNameOfRedisHostsChanged = "event.redis.hosts.changed"
)

var currentApp App
var redisStore *redisEntryStore

// App define application interface
type App interface {
	// Start is start the application
	Start() error

	// Stop is stop the application
	Stop()

	EventManager() event.Manager
}

type guiApp struct {
	eventMgr event.Manager
}

func (g *guiApp) EventManager() event.Manager {
	return g.eventMgr
}

func (g *guiApp) initialize() error {
	g.eventMgr = event.NewManager()
	p := app.NewWithID(appUniqueId)
	locales.SetLanguage(currentLanguage())
	p.Lifecycle().SetOnStarted(func() {
		logger.Info(locales.Get(locales.LOG_INFO_APPLICATION_STARTED))
	})
	p.Lifecycle().SetOnStopped(func() {
		logger.Info(locales.Get(locales.LOG_INFO_APPLICATION_STOPED))
	})
	setAppTheme()
	g.eventMgr.FocusOn(eventNameOfThemeChanged, func(_ string, _ interface{}) {
		setAppTheme()
	})
	g.renderMain()
	size := []int{defaultMainAppWindowsWidth, defaultMainAppWindowsHeight}
	if currentRecordWindowSize() {
		for k, v := range strings.Split(p.Preferences().String(preferenceKeyOfAppSize), ",") {
			if tmp, err := strconv.Atoi(v); err == nil {
				size[k] = tmp
			}
		}
	}
	p.Driver().AllWindows()[0].Resize(fyne.NewSize(float32(size[0]), float32(size[1])))
	g.eventMgr.FocusOn(eventNameOfCloseWindow, func(_ string, data interface{}) {
		if size, ok := data.(fyne.Size); ok {
			p.Preferences().SetString(preferenceKeyOfAppSize, fmt.Sprintf("%d,%d", int(size.Width), int(size.Height)))
		}
	})
	return nil
}

func (g *guiApp) renderMain() {
	main := makeMainWindows()
	main.Show()
	g.eventMgr.FocusOn(eventNameOfRebootWindow, func(_ string, _ interface{}) {
		rebootMainWindows()
	})
}

func (g *guiApp) Start() error {
	currentApp = g
	redisStore = new(redisEntryStore)
	if err := g.initialize(); err != nil {
		return err
	}
	fyne.CurrentApp().Run()
	return nil
}

func (g *guiApp) Stop() {
	fyne.CurrentApp().Quit()
}

// NewGUI return new gui application
func NewGUI() App {
	return new(guiApp)
}

func rebootMainWindows() {
	w := fyne.CurrentApp().Driver().AllWindows()[0]
	main := makeMainWindows()
	main.Resize(fyne.NewSize(w.Canvas().Size().Width, w.Canvas().Size().Height))
	main.Show()
	logger.Info(locales.Get(locales.LOG_INFO_RENEW_RENDER_WINDOWS))
	w.Close()
}

func currentLanguage() string {
	return fyne.CurrentApp().Preferences().StringWithFallback(preferenceKeyOfSettingLanguage, locales.CurrentLanguage().String())
}

func currentRecordWindowSize() bool {
	return fyne.CurrentApp().Preferences().BoolWithFallback(preferenceKeyOfRecordWindowSize, true)
}

type redisEntry struct {
	Name     string
	Hosts    []string
	Password string
}

func (r redisEntry) Equal(e redisEntry) bool {
	return r.Name == e.Name
}

type redisEntryStore struct {
	locker sync.RWMutex
}

func (r *redisEntryStore) List() ([]redisEntry, error) {
	r.locker.RLock()
	defer r.locker.RUnlock()
	return r.getList()
}

func (r *redisEntryStore) getList() ([]redisEntry, error) {
	s := fyne.CurrentApp().Preferences().String(preferenceKeyOfRedisConfigList)
	if len(s) == 0 {
		return nil, nil
	}
	var list []redisEntry
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(utils.NocopyStr2ByteSlice(s), &list); err != nil {
		return nil, errors.New(locales.Get(locales.ERROR_UNMARSHAL_FAILED))
	}
	return list, nil
}

func (r *redisEntryStore) saveList(list []redisEntry) error {
	b, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(list)
	if err != nil {
		return errors.New(locales.Get(locales.ERROR_MARSHAL_FAILED))
	}
	fyne.CurrentApp().Preferences().SetString(preferenceKeyOfRedisConfigList, utils.NocopyByteSlice2Str(b))
	return nil
}

func (r *redisEntryStore) Add(entry redisEntry) ([]redisEntry, error) {
	r.locker.Lock()
	defer r.locker.Unlock()
	list, err := r.getList()
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		if item.Equal(entry) {
			return list, errors.New(locales.Get(locales.ERROR_ALREADY_EXISTS))
		}
	}
	list = append(list, entry)
	if err := r.saveList(list); err != nil {
		return list, errors.New(locales.Get(locales.ERROR_START_APPLICATION_FAILED, err))
	}
	return list, nil
}

func (r *redisEntryStore) Del(entry redisEntry) ([]redisEntry, error) {
	r.locker.Lock()
	defer r.locker.Unlock()
	list, err := r.getList()
	if err != nil {
		return nil, err
	}
	for i, item := range list {
		if item.Equal(entry) {
			switch i {
			case 0:
				list = list[1:]
			case len(list) - 1:
				list = list[0:i]
			default:
				list = append(list[0:i], list[i+1:]...)
			}
			break
		}
	}
	if err := r.saveList(list); err != nil {
		return list, errors.New(locales.Get(locales.ERROR_REMOVE_FAILED, err))
	}
	return list, nil
}

func (r *redisEntryStore) Exists(name string) bool {
	if list, err := r.List(); err == nil {
		for _, entry := range list {
			if entry.Name == name {
				return true
			}
		}
		return false
	}
	return false
}
