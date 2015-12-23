package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMouseEventTransition struct {
	core.QEventTransition
}

type QMouseEventTransition_ITF interface {
	core.QEventTransition_ITF
	QMouseEventTransition_PTR() *QMouseEventTransition
}

func PointerFromQMouseEventTransition(ptr QMouseEventTransition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMouseEventTransition_PTR().Pointer()
	}
	return nil
}

func NewQMouseEventTransitionFromPointer(ptr unsafe.Pointer) *QMouseEventTransition {
	var n = new(QMouseEventTransition)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMouseEventTransition_") {
		n.SetObjectName("QMouseEventTransition_" + qt.Identifier())
	}
	return n
}

func (ptr *QMouseEventTransition) QMouseEventTransition_PTR() *QMouseEventTransition {
	return ptr
}

func NewQMouseEventTransition2(object core.QObject_ITF, ty core.QEvent__Type, button core.Qt__MouseButton, sourceState core.QState_ITF) *QMouseEventTransition {
	defer qt.Recovering("QMouseEventTransition::QMouseEventTransition")

	return NewQMouseEventTransitionFromPointer(C.QMouseEventTransition_NewQMouseEventTransition2(core.PointerFromQObject(object), C.int(ty), C.int(button), core.PointerFromQState(sourceState)))
}

func NewQMouseEventTransition(sourceState core.QState_ITF) *QMouseEventTransition {
	defer qt.Recovering("QMouseEventTransition::QMouseEventTransition")

	return NewQMouseEventTransitionFromPointer(C.QMouseEventTransition_NewQMouseEventTransition(core.PointerFromQState(sourceState)))
}

func (ptr *QMouseEventTransition) Button() core.Qt__MouseButton {
	defer qt.Recovering("QMouseEventTransition::button")

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QMouseEventTransition_Button(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEventTransition) ModifierMask() core.Qt__KeyboardModifier {
	defer qt.Recovering("QMouseEventTransition::modifierMask")

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QMouseEventTransition_ModifierMask(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEventTransition) ConnectOnTransition(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMouseEventTransition::onTransition")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onTransition", f)
	}
}

func (ptr *QMouseEventTransition) DisconnectOnTransition() {
	defer qt.Recovering("disconnect QMouseEventTransition::onTransition")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onTransition")
	}
}

//export callbackQMouseEventTransitionOnTransition
func callbackQMouseEventTransitionOnTransition(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMouseEventTransition::onTransition")

	if signal := qt.GetSignal(C.GoString(ptrName), "onTransition"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMouseEventTransition) SetButton(button core.Qt__MouseButton) {
	defer qt.Recovering("QMouseEventTransition::setButton")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetButton(ptr.Pointer(), C.int(button))
	}
}

func (ptr *QMouseEventTransition) SetHitTestPath(path gui.QPainterPath_ITF) {
	defer qt.Recovering("QMouseEventTransition::setHitTestPath")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetHitTestPath(ptr.Pointer(), gui.PointerFromQPainterPath(path))
	}
}

func (ptr *QMouseEventTransition) SetModifierMask(modifierMask core.Qt__KeyboardModifier) {
	defer qt.Recovering("QMouseEventTransition::setModifierMask")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetModifierMask(ptr.Pointer(), C.int(modifierMask))
	}
}

func (ptr *QMouseEventTransition) DestroyQMouseEventTransition() {
	defer qt.Recovering("QMouseEventTransition::~QMouseEventTransition")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_DestroyQMouseEventTransition(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMouseEventTransition) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMouseEventTransition::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMouseEventTransition) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMouseEventTransition::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMouseEventTransitionTimerEvent
func callbackQMouseEventTransitionTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMouseEventTransition::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMouseEventTransition) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMouseEventTransition::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMouseEventTransition) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMouseEventTransition::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMouseEventTransitionChildEvent
func callbackQMouseEventTransitionChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMouseEventTransition::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMouseEventTransition) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMouseEventTransition::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMouseEventTransition) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMouseEventTransition::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMouseEventTransitionCustomEvent
func callbackQMouseEventTransitionCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMouseEventTransition::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
