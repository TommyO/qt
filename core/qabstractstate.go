package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAbstractState struct {
	QObject
}

type QAbstractState_ITF interface {
	QObject_ITF
	QAbstractState_PTR() *QAbstractState
}

func PointerFromQAbstractState(ptr QAbstractState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractState_PTR().Pointer()
	}
	return nil
}

func NewQAbstractStateFromPointer(ptr unsafe.Pointer) *QAbstractState {
	var n = new(QAbstractState)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractState_") {
		n.SetObjectName("QAbstractState_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractState) QAbstractState_PTR() *QAbstractState {
	return ptr
}

func (ptr *QAbstractState) Active() bool {
	defer qt.Recovering("QAbstractState::active")

	if ptr.Pointer() != nil {
		return C.QAbstractState_Active(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractState) ConnectActiveChanged(f func(active bool)) {
	defer qt.Recovering("connect QAbstractState::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectActiveChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeChanged", f)
	}
}

func (ptr *QAbstractState) DisconnectActiveChanged() {
	defer qt.Recovering("disconnect QAbstractState::activeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectActiveChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeChanged")
	}
}

//export callbackQAbstractStateActiveChanged
func callbackQAbstractStateActiveChanged(ptrName *C.char, active C.int) {
	defer qt.Recovering("callback QAbstractState::activeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeChanged"); signal != nil {
		signal.(func(bool))(int(active) != 0)
	}

}

func (ptr *QAbstractState) ConnectEntered(f func()) {
	defer qt.Recovering("connect QAbstractState::entered")

	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "entered", f)
	}
}

func (ptr *QAbstractState) DisconnectEntered() {
	defer qt.Recovering("disconnect QAbstractState::entered")

	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "entered")
	}
}

//export callbackQAbstractStateEntered
func callbackQAbstractStateEntered(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractState::entered")

	if signal := qt.GetSignal(C.GoString(ptrName), "entered"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractState) ConnectExited(f func()) {
	defer qt.Recovering("connect QAbstractState::exited")

	if ptr.Pointer() != nil {
		C.QAbstractState_ConnectExited(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "exited", f)
	}
}

func (ptr *QAbstractState) DisconnectExited() {
	defer qt.Recovering("disconnect QAbstractState::exited")

	if ptr.Pointer() != nil {
		C.QAbstractState_DisconnectExited(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "exited")
	}
}

//export callbackQAbstractStateExited
func callbackQAbstractStateExited(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractState::exited")

	if signal := qt.GetSignal(C.GoString(ptrName), "exited"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractState) Machine() *QStateMachine {
	defer qt.Recovering("QAbstractState::machine")

	if ptr.Pointer() != nil {
		return NewQStateMachineFromPointer(C.QAbstractState_Machine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractState) ParentState() *QState {
	defer qt.Recovering("QAbstractState::parentState")

	if ptr.Pointer() != nil {
		return NewQStateFromPointer(C.QAbstractState_ParentState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractState) DestroyQAbstractState() {
	defer qt.Recovering("QAbstractState::~QAbstractState")

	if ptr.Pointer() != nil {
		C.QAbstractState_DestroyQAbstractState(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractState) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QAbstractState::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractState) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractState::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractStateTimerEvent
func callbackQAbstractStateTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractState::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractState) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QAbstractState::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractState) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractState::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractStateChildEvent
func callbackQAbstractStateChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractState::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractState) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QAbstractState::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractState) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractState::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractStateCustomEvent
func callbackQAbstractStateCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractState::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
