package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QButtonGroup struct {
	core.QObject
}

type QButtonGroup_ITF interface {
	core.QObject_ITF
	QButtonGroup_PTR() *QButtonGroup
}

func PointerFromQButtonGroup(ptr QButtonGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QButtonGroup_PTR().Pointer()
	}
	return nil
}

func NewQButtonGroupFromPointer(ptr unsafe.Pointer) *QButtonGroup {
	var n = new(QButtonGroup)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QButtonGroup_") {
		n.SetObjectName("QButtonGroup_" + qt.Identifier())
	}
	return n
}

func (ptr *QButtonGroup) QButtonGroup_PTR() *QButtonGroup {
	return ptr
}

func NewQButtonGroup(parent core.QObject_ITF) *QButtonGroup {
	defer qt.Recovering("QButtonGroup::QButtonGroup")

	return NewQButtonGroupFromPointer(C.QButtonGroup_NewQButtonGroup(core.PointerFromQObject(parent)))
}

func (ptr *QButtonGroup) AddButton(button QAbstractButton_ITF, id int) {
	defer qt.Recovering("QButtonGroup::addButton")

	if ptr.Pointer() != nil {
		C.QButtonGroup_AddButton(ptr.Pointer(), PointerFromQAbstractButton(button), C.int(id))
	}
}

func (ptr *QButtonGroup) Button(id int) *QAbstractButton {
	defer qt.Recovering("QButtonGroup::button")

	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QButtonGroup_Button(ptr.Pointer(), C.int(id)))
	}
	return nil
}

func (ptr *QButtonGroup) CheckedButton() *QAbstractButton {
	defer qt.Recovering("QButtonGroup::checkedButton")

	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QButtonGroup_CheckedButton(ptr.Pointer()))
	}
	return nil
}

func (ptr *QButtonGroup) CheckedId() int {
	defer qt.Recovering("QButtonGroup::checkedId")

	if ptr.Pointer() != nil {
		return int(C.QButtonGroup_CheckedId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QButtonGroup) Exclusive() bool {
	defer qt.Recovering("QButtonGroup::exclusive")

	if ptr.Pointer() != nil {
		return C.QButtonGroup_Exclusive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QButtonGroup) Id(button QAbstractButton_ITF) int {
	defer qt.Recovering("QButtonGroup::id")

	if ptr.Pointer() != nil {
		return int(C.QButtonGroup_Id(ptr.Pointer(), PointerFromQAbstractButton(button)))
	}
	return 0
}

func (ptr *QButtonGroup) RemoveButton(button QAbstractButton_ITF) {
	defer qt.Recovering("QButtonGroup::removeButton")

	if ptr.Pointer() != nil {
		C.QButtonGroup_RemoveButton(ptr.Pointer(), PointerFromQAbstractButton(button))
	}
}

func (ptr *QButtonGroup) SetExclusive(v bool) {
	defer qt.Recovering("QButtonGroup::setExclusive")

	if ptr.Pointer() != nil {
		C.QButtonGroup_SetExclusive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QButtonGroup) SetId(button QAbstractButton_ITF, id int) {
	defer qt.Recovering("QButtonGroup::setId")

	if ptr.Pointer() != nil {
		C.QButtonGroup_SetId(ptr.Pointer(), PointerFromQAbstractButton(button), C.int(id))
	}
}

func (ptr *QButtonGroup) DestroyQButtonGroup() {
	defer qt.Recovering("QButtonGroup::~QButtonGroup")

	if ptr.Pointer() != nil {
		C.QButtonGroup_DestroyQButtonGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QButtonGroup) ConnectButtonClicked(f func(button *QAbstractButton)) {
	defer qt.Recovering("connect QButtonGroup::buttonClicked")

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonClicked", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonClicked() {
	defer qt.Recovering("disconnect QButtonGroup::buttonClicked")

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonClicked")
	}
}

//export callbackQButtonGroupButtonClicked
func callbackQButtonGroupButtonClicked(ptrName *C.char, button unsafe.Pointer) {
	defer qt.Recovering("callback QButtonGroup::buttonClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "buttonClicked"); signal != nil {
		signal.(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
	}

}

func (ptr *QButtonGroup) ConnectButtonClicked2(f func(id int)) {
	defer qt.Recovering("connect QButtonGroup::buttonClicked")

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonClicked2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonClicked2", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonClicked2() {
	defer qt.Recovering("disconnect QButtonGroup::buttonClicked")

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonClicked2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonClicked2")
	}
}

//export callbackQButtonGroupButtonClicked2
func callbackQButtonGroupButtonClicked2(ptrName *C.char, id C.int) {
	defer qt.Recovering("callback QButtonGroup::buttonClicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "buttonClicked2"); signal != nil {
		signal.(func(int))(int(id))
	}

}

func (ptr *QButtonGroup) ConnectButtonPressed(f func(button *QAbstractButton)) {
	defer qt.Recovering("connect QButtonGroup::buttonPressed")

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonPressed", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonPressed() {
	defer qt.Recovering("disconnect QButtonGroup::buttonPressed")

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonPressed")
	}
}

//export callbackQButtonGroupButtonPressed
func callbackQButtonGroupButtonPressed(ptrName *C.char, button unsafe.Pointer) {
	defer qt.Recovering("callback QButtonGroup::buttonPressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "buttonPressed"); signal != nil {
		signal.(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
	}

}

func (ptr *QButtonGroup) ConnectButtonPressed2(f func(id int)) {
	defer qt.Recovering("connect QButtonGroup::buttonPressed")

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonPressed2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonPressed2", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonPressed2() {
	defer qt.Recovering("disconnect QButtonGroup::buttonPressed")

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonPressed2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonPressed2")
	}
}

//export callbackQButtonGroupButtonPressed2
func callbackQButtonGroupButtonPressed2(ptrName *C.char, id C.int) {
	defer qt.Recovering("callback QButtonGroup::buttonPressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "buttonPressed2"); signal != nil {
		signal.(func(int))(int(id))
	}

}

func (ptr *QButtonGroup) ConnectButtonReleased(f func(button *QAbstractButton)) {
	defer qt.Recovering("connect QButtonGroup::buttonReleased")

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonReleased(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonReleased", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonReleased() {
	defer qt.Recovering("disconnect QButtonGroup::buttonReleased")

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonReleased")
	}
}

//export callbackQButtonGroupButtonReleased
func callbackQButtonGroupButtonReleased(ptrName *C.char, button unsafe.Pointer) {
	defer qt.Recovering("callback QButtonGroup::buttonReleased")

	if signal := qt.GetSignal(C.GoString(ptrName), "buttonReleased"); signal != nil {
		signal.(func(*QAbstractButton))(NewQAbstractButtonFromPointer(button))
	}

}

func (ptr *QButtonGroup) ConnectButtonReleased2(f func(id int)) {
	defer qt.Recovering("connect QButtonGroup::buttonReleased")

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonReleased2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonReleased2", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonReleased2() {
	defer qt.Recovering("disconnect QButtonGroup::buttonReleased")

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonReleased2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonReleased2")
	}
}

//export callbackQButtonGroupButtonReleased2
func callbackQButtonGroupButtonReleased2(ptrName *C.char, id C.int) {
	defer qt.Recovering("callback QButtonGroup::buttonReleased")

	if signal := qt.GetSignal(C.GoString(ptrName), "buttonReleased2"); signal != nil {
		signal.(func(int))(int(id))
	}

}

func (ptr *QButtonGroup) ConnectButtonToggled(f func(button *QAbstractButton, checked bool)) {
	defer qt.Recovering("connect QButtonGroup::buttonToggled")

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonToggled", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonToggled() {
	defer qt.Recovering("disconnect QButtonGroup::buttonToggled")

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonToggled")
	}
}

//export callbackQButtonGroupButtonToggled
func callbackQButtonGroupButtonToggled(ptrName *C.char, button unsafe.Pointer, checked C.int) {
	defer qt.Recovering("callback QButtonGroup::buttonToggled")

	if signal := qt.GetSignal(C.GoString(ptrName), "buttonToggled"); signal != nil {
		signal.(func(*QAbstractButton, bool))(NewQAbstractButtonFromPointer(button), int(checked) != 0)
	}

}

func (ptr *QButtonGroup) ConnectButtonToggled2(f func(id int, checked bool)) {
	defer qt.Recovering("connect QButtonGroup::buttonToggled")

	if ptr.Pointer() != nil {
		C.QButtonGroup_ConnectButtonToggled2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "buttonToggled2", f)
	}
}

func (ptr *QButtonGroup) DisconnectButtonToggled2() {
	defer qt.Recovering("disconnect QButtonGroup::buttonToggled")

	if ptr.Pointer() != nil {
		C.QButtonGroup_DisconnectButtonToggled2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "buttonToggled2")
	}
}

//export callbackQButtonGroupButtonToggled2
func callbackQButtonGroupButtonToggled2(ptrName *C.char, id C.int, checked C.int) {
	defer qt.Recovering("callback QButtonGroup::buttonToggled")

	if signal := qt.GetSignal(C.GoString(ptrName), "buttonToggled2"); signal != nil {
		signal.(func(int, bool))(int(id), int(checked) != 0)
	}

}

func (ptr *QButtonGroup) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QButtonGroup::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QButtonGroup) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QButtonGroup::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQButtonGroupTimerEvent
func callbackQButtonGroupTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QButtonGroup::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QButtonGroup) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QButtonGroup::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QButtonGroup) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QButtonGroup::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQButtonGroupChildEvent
func callbackQButtonGroupChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QButtonGroup::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QButtonGroup) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QButtonGroup::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QButtonGroup) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QButtonGroup::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQButtonGroupCustomEvent
func callbackQButtonGroupCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QButtonGroup::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
