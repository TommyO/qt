package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QQmlFileSelector struct {
	core.QObject
}

type QQmlFileSelector_ITF interface {
	core.QObject_ITF
	QQmlFileSelector_PTR() *QQmlFileSelector
}

func PointerFromQQmlFileSelector(ptr QQmlFileSelector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlFileSelector_PTR().Pointer()
	}
	return nil
}

func NewQQmlFileSelectorFromPointer(ptr unsafe.Pointer) *QQmlFileSelector {
	var n = new(QQmlFileSelector)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQmlFileSelector_") {
		n.SetObjectName("QQmlFileSelector_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlFileSelector) QQmlFileSelector_PTR() *QQmlFileSelector {
	return ptr
}

func NewQQmlFileSelector(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlFileSelector {
	defer qt.Recovering("QQmlFileSelector::QQmlFileSelector")

	return NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_NewQQmlFileSelector(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
}

func QQmlFileSelector_Get(engine QQmlEngine_ITF) *QQmlFileSelector {
	defer qt.Recovering("QQmlFileSelector::get")

	return NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_QQmlFileSelector_Get(PointerFromQQmlEngine(engine)))
}

func (ptr *QQmlFileSelector) SetExtraSelectors(strin []string) {
	defer qt.Recovering("QQmlFileSelector::setExtraSelectors")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetExtraSelectors(ptr.Pointer(), C.CString(strings.Join(strin, ",,,")))
	}
}

func (ptr *QQmlFileSelector) SetExtraSelectors2(strin []string) {
	defer qt.Recovering("QQmlFileSelector::setExtraSelectors")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetExtraSelectors2(ptr.Pointer(), C.CString(strings.Join(strin, ",,,")))
	}
}

func (ptr *QQmlFileSelector) SetSelector(selector core.QFileSelector_ITF) {
	defer qt.Recovering("QQmlFileSelector::setSelector")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetSelector(ptr.Pointer(), core.PointerFromQFileSelector(selector))
	}
}

func (ptr *QQmlFileSelector) DestroyQQmlFileSelector() {
	defer qt.Recovering("QQmlFileSelector::~QQmlFileSelector")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DestroyQQmlFileSelector(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQmlFileSelector) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQmlFileSelector::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQmlFileSelector::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQmlFileSelectorTimerEvent
func callbackQQmlFileSelectorTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlFileSelector::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQmlFileSelector) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQmlFileSelector::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQmlFileSelector::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQmlFileSelectorChildEvent
func callbackQQmlFileSelectorChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlFileSelector::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQmlFileSelector) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQmlFileSelector::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQmlFileSelector) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQmlFileSelector::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQmlFileSelectorCustomEvent
func callbackQQmlFileSelectorCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQmlFileSelector::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
