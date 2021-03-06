// Package gestures maps hotkey events to key & gesture events.
package gestures

import (
	"time"

	"github.com/echocrow/Mouser/pkg/hotkeys/gestures/swipes"
	"github.com/echocrow/Mouser/pkg/hotkeys/hotkey"
	"github.com/echocrow/Mouser/pkg/hotkeys/monitor"
)

// Gesture identifies the type of gesture.
type Gesture string

// Gesture types -- Nil event.
const (
	NoGesture Gesture = ""
)

// Gesture types -- Basic key events.
const (
	KeyDown Gesture = "key_down"
	KeyUp   Gesture = "key_up"
)

// Gesture types -- Key gestures.
const (
	PressShort Gesture = "tap"
	PressLong  Gesture = "hold"
)

// Gesture types -- Mouse gestures.
const (
	SwipeUp    Gesture = "swipe_up"
	SwipeDown  Gesture = "swipe_down"
	SwipeLeft  Gesture = "swipe_left"
	SwipeRight Gesture = "swipe_right"
)

// Config defines gestures settings.
type Config struct {
	ShortPressTTL time.Duration
	GestureTTL    time.Duration
	Cap           int
}

// Event represents a key/mouse gestures event.
type Event struct {
	HkID  hotkey.ID
	Gests []Gesture
	T     time.Time
}

// Default settings.
const (
	defaultShortPressTTL = time.Millisecond * 500
	defaultGestureTTL    = time.Millisecond * 500
	defaultGesturesCap   = int(8)
)

// FromHotkeys maps hotkey events to key & mouse gestures.
func FromHotkeys(
	hkEvs <-chan monitor.HotkeyEvent,
) <-chan Event {
	config := Config{
		ShortPressTTL: defaultShortPressTTL,
		GestureTTL:    defaultGestureTTL,
		Cap:           defaultGesturesCap,
	}
	swpMon := swipes.NewDefaultMonitor()
	return FromHotkeysCustom(hkEvs, config, swpMon)
}

// FromHotkeysCustom maps hotkey events to key & mouse gestures with custom
// options.
//
// When hkEvs is depleted, swpMon still be stopped and the returned channel
// closed.
//
func FromHotkeysCustom(
	hkEvs <-chan monitor.HotkeyEvent,
	config Config,
	swpMon swipes.Monitor,
) <-chan Event {
	ch := make(chan Event)
	go mapHkEvs(hkEvs, swpMon, config, ch)
	return ch
}

// Match checks whether gesture sequences a & b hold the same gestures.
func Match(
	a []Gesture,
	b []Gesture,
) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// MatchSingle checks whether gesture sequences holds only g.
func MatchSingle(
	a []Gesture,
	g Gesture,
) bool {
	return len(a) == 1 && a[0] == g
}

// EndsIn checks whether gesture sequence a ends with gesture g.
func EndsIn(
	a []Gesture,
	g Gesture,
) bool {
	l := len(a)
	return l > 0 && a[l-1] == g
}

func mapHkEvs(
	hkEvs <-chan monitor.HotkeyEvent,
	swpMon swipes.Monitor,
	config Config,
	ch chan<- Event,
) {
	defer close(ch)
	var (
		prvT  time.Time
		hk    hotkey.ID
		prvHk hotkey.ID
		swpC  <-chan swipes.Event
		swpd  bool
		gests []Gesture
	)
	if swpMon != nil {
		defer swpMon.Stop()
		swpC = swpMon.Init()
	}
	handleSwpEv := func(hk hotkey.ID, swpEv swipes.Event) {
		if swpEv.IsSwipe() {
			swpd = true
			gests = appendGest(gests, config.Cap, swipeGesture(swpEv.Dir))
			ch <- Event{hk, gests, swpEv.T}
		}
	}
	for {
		select {

		case hkEv, ok := <-hkEvs:
			if !ok {
				return
			}
			ch <- Event{hkEv.HkID, []Gesture{keyGesture(hkEv)}, hkEv.T}

			t := hkEv.T
			dt := t.Sub(prvT)
			prvT = t

			if hkEv.IsOn {
				swpd = false
				if hkEv.HkID != prvHk || dt > config.GestureTTL {
					gests = nil
				}
				hk = hkEv.HkID
				prvHk = 0
				if swpMon != nil {
					swpMon.Restart()
				}
			} else if hkEv.HkID == hk {
				hk = 0
				prvHk = hkEv.HkID
				if swpMon != nil {
					swpEv := swpMon.Pause(hkEv.T)
					handleSwpEv(prvHk, swpEv)
				}
				if !swpd {
					if dt <= config.ShortPressTTL {
						gests = appendGest(gests, config.Cap, PressShort)
					} else {
						gests = appendGest(gests, config.Cap, PressLong)
					}
					ch <- Event{hkEv.HkID, gests, t}
				}
			}

		case swpEv, ok := <-swpC:
			if ok && hk != 0 {
				handleSwpEv(hk, swpEv)
			}
		}
	}
}

func appendGest(gests []Gesture, gestsCap int, gest Gesture) []Gesture {
	if gestsCap > 0 && len(gests)+1 > gestsCap {
		d := len(gests) + 1 - gestsCap
		gests = gests[d:]
	}
	gests = append(gests, gest)
	return gests
}

func keyGesture(hkEv monitor.HotkeyEvent) Gesture {
	if hkEv.IsOn {
		return KeyDown
	}
	return KeyUp
}

func swipeGesture(dir swipes.Dir) Gesture {
	switch dir {
	case swipes.SwipeUp:
		return SwipeUp
	case swipes.SwipeDown:
		return SwipeDown
	case swipes.SwipeLeft:
		return SwipeLeft
	case swipes.SwipeRight:
		return SwipeRight
	}
	panic("Invalid swipe direction")
}
