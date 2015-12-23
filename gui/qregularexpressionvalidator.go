package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRegularExpressionValidator struct {
	QValidator
}

type QRegularExpressionValidator_ITF interface {
	QValidator_ITF
	QRegularExpressionValidator_PTR() *QRegularExpressionValidator
}

func PointerFromQRegularExpressionValidator(ptr QRegularExpressionValidator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegularExpressionValidator_PTR().Pointer()
	}
	return nil
}

func NewQRegularExpressionValidatorFromPointer(ptr unsafe.Pointer) *QRegularExpressionValidator {
	var n = new(QRegularExpressionValidator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRegularExpressionValidator_") {
		n.SetObjectName("QRegularExpressionValidator_" + qt.Identifier())
	}
	return n
}

func (ptr *QRegularExpressionValidator) QRegularExpressionValidator_PTR() *QRegularExpressionValidator {
	return ptr
}

func (ptr *QRegularExpressionValidator) RegularExpression() *core.QRegularExpression {
	defer qt.Recovering("QRegularExpressionValidator::regularExpression")

	if ptr.Pointer() != nil {
		return core.NewQRegularExpressionFromPointer(C.QRegularExpressionValidator_RegularExpression(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRegularExpressionValidator) SetRegularExpression(re core.QRegularExpression_ITF) {
	defer qt.Recovering("QRegularExpressionValidator::setRegularExpression")

	if ptr.Pointer() != nil {
		C.QRegularExpressionValidator_SetRegularExpression(ptr.Pointer(), core.PointerFromQRegularExpression(re))
	}
}

func NewQRegularExpressionValidator(parent core.QObject_ITF) *QRegularExpressionValidator {
	defer qt.Recovering("QRegularExpressionValidator::QRegularExpressionValidator")

	return NewQRegularExpressionValidatorFromPointer(C.QRegularExpressionValidator_NewQRegularExpressionValidator(core.PointerFromQObject(parent)))
}

func NewQRegularExpressionValidator2(re core.QRegularExpression_ITF, parent core.QObject_ITF) *QRegularExpressionValidator {
	defer qt.Recovering("QRegularExpressionValidator::QRegularExpressionValidator")

	return NewQRegularExpressionValidatorFromPointer(C.QRegularExpressionValidator_NewQRegularExpressionValidator2(core.PointerFromQRegularExpression(re), core.PointerFromQObject(parent)))
}

func (ptr *QRegularExpressionValidator) ConnectRegularExpressionChanged(f func(re *core.QRegularExpression)) {
	defer qt.Recovering("connect QRegularExpressionValidator::regularExpressionChanged")

	if ptr.Pointer() != nil {
		C.QRegularExpressionValidator_ConnectRegularExpressionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "regularExpressionChanged", f)
	}
}

func (ptr *QRegularExpressionValidator) DisconnectRegularExpressionChanged() {
	defer qt.Recovering("disconnect QRegularExpressionValidator::regularExpressionChanged")

	if ptr.Pointer() != nil {
		C.QRegularExpressionValidator_DisconnectRegularExpressionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "regularExpressionChanged")
	}
}

//export callbackQRegularExpressionValidatorRegularExpressionChanged
func callbackQRegularExpressionValidatorRegularExpressionChanged(ptrName *C.char, re unsafe.Pointer) {
	defer qt.Recovering("callback QRegularExpressionValidator::regularExpressionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "regularExpressionChanged"); signal != nil {
		signal.(func(*core.QRegularExpression))(core.NewQRegularExpressionFromPointer(re))
	}

}

func (ptr *QRegularExpressionValidator) DestroyQRegularExpressionValidator() {
	defer qt.Recovering("QRegularExpressionValidator::~QRegularExpressionValidator")

	if ptr.Pointer() != nil {
		C.QRegularExpressionValidator_DestroyQRegularExpressionValidator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRegularExpressionValidator) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QRegularExpressionValidator::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRegularExpressionValidator) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRegularExpressionValidator::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRegularExpressionValidatorTimerEvent
func callbackQRegularExpressionValidatorTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRegularExpressionValidator::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRegularExpressionValidator) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRegularExpressionValidator::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRegularExpressionValidator) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRegularExpressionValidator::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRegularExpressionValidatorChildEvent
func callbackQRegularExpressionValidatorChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRegularExpressionValidator::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRegularExpressionValidator) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRegularExpressionValidator::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRegularExpressionValidator) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRegularExpressionValidator::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRegularExpressionValidatorCustomEvent
func callbackQRegularExpressionValidatorCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRegularExpressionValidator::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
