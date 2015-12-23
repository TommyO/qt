package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QPushButton struct {
	QAbstractButton
}

type QPushButton_ITF interface {
	QAbstractButton_ITF
	QPushButton_PTR() *QPushButton
}

func PointerFromQPushButton(ptr QPushButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPushButton_PTR().Pointer()
	}
	return nil
}

func NewQPushButtonFromPointer(ptr unsafe.Pointer) *QPushButton {
	var n = new(QPushButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPushButton_") {
		n.SetObjectName("QPushButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QPushButton) QPushButton_PTR() *QPushButton {
	return ptr
}

func (ptr *QPushButton) AutoDefault() bool {
	defer qt.Recovering("QPushButton::autoDefault")

	if ptr.Pointer() != nil {
		return C.QPushButton_AutoDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPushButton) IsDefault() bool {
	defer qt.Recovering("QPushButton::isDefault")

	if ptr.Pointer() != nil {
		return C.QPushButton_IsDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPushButton) IsFlat() bool {
	defer qt.Recovering("QPushButton::isFlat")

	if ptr.Pointer() != nil {
		return C.QPushButton_IsFlat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPushButton) SetAutoDefault(v bool) {
	defer qt.Recovering("QPushButton::setAutoDefault")

	if ptr.Pointer() != nil {
		C.QPushButton_SetAutoDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QPushButton) SetDefault(v bool) {
	defer qt.Recovering("QPushButton::setDefault")

	if ptr.Pointer() != nil {
		C.QPushButton_SetDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QPushButton) SetFlat(v bool) {
	defer qt.Recovering("QPushButton::setFlat")

	if ptr.Pointer() != nil {
		C.QPushButton_SetFlat(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQPushButton(parent QWidget_ITF) *QPushButton {
	defer qt.Recovering("QPushButton::QPushButton")

	return NewQPushButtonFromPointer(C.QPushButton_NewQPushButton(PointerFromQWidget(parent)))
}

func NewQPushButton3(icon gui.QIcon_ITF, text string, parent QWidget_ITF) *QPushButton {
	defer qt.Recovering("QPushButton::QPushButton")

	return NewQPushButtonFromPointer(C.QPushButton_NewQPushButton3(gui.PointerFromQIcon(icon), C.CString(text), PointerFromQWidget(parent)))
}

func NewQPushButton2(text string, parent QWidget_ITF) *QPushButton {
	defer qt.Recovering("QPushButton::QPushButton")

	return NewQPushButtonFromPointer(C.QPushButton_NewQPushButton2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QPushButton) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPushButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QPushButton) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QPushButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQPushButtonFocusInEvent
func callbackQPushButtonFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QPushButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QPushButton) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QPushButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQPushButtonFocusOutEvent
func callbackQPushButtonFocusOutEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPushButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QPushButton) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QPushButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQPushButtonKeyPressEvent
func callbackQPushButtonKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) Menu() *QMenu {
	defer qt.Recovering("QPushButton::menu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QPushButton_Menu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPushButton) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QPushButton::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QPushButton_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPushButton) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QPushButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QPushButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QPushButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQPushButtonPaintEvent
func callbackQPushButtonPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QPushButton) SetMenu(menu QMenu_ITF) {
	defer qt.Recovering("QPushButton::setMenu")

	if ptr.Pointer() != nil {
		C.QPushButton_SetMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QPushButton) ShowMenu() {
	defer qt.Recovering("QPushButton::showMenu")

	if ptr.Pointer() != nil {
		C.QPushButton_ShowMenu(ptr.Pointer())
	}
}

func (ptr *QPushButton) SizeHint() *core.QSize {
	defer qt.Recovering("QPushButton::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QPushButton_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPushButton) DestroyQPushButton() {
	defer qt.Recovering("QPushButton::~QPushButton")

	if ptr.Pointer() != nil {
		C.QPushButton_DestroyQPushButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPushButton) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QPushButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QPushButton) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QPushButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQPushButtonChangeEvent
func callbackQPushButtonChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectCheckStateSet(f func()) {
	defer qt.Recovering("connect QPushButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "checkStateSet", f)
	}
}

func (ptr *QPushButton) DisconnectCheckStateSet() {
	defer qt.Recovering("disconnect QPushButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "checkStateSet")
	}
}

//export callbackQPushButtonCheckStateSet
func callbackQPushButtonCheckStateSet(ptrName *C.char) bool {
	defer qt.Recovering("callback QPushButton::checkStateSet")

	if signal := qt.GetSignal(C.GoString(ptrName), "checkStateSet"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QPushButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QPushButton) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QPushButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQPushButtonKeyReleaseEvent
func callbackQPushButtonKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPushButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QPushButton) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QPushButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQPushButtonMouseMoveEvent
func callbackQPushButtonMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPushButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QPushButton) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QPushButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQPushButtonMousePressEvent
func callbackQPushButtonMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPushButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QPushButton) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QPushButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQPushButtonMouseReleaseEvent
func callbackQPushButtonMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QPushButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QPushButton) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QPushButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQPushButtonNextCheckState
func callbackQPushButtonNextCheckState(ptrName *C.char) bool {
	defer qt.Recovering("callback QPushButton::nextCheckState")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextCheckState"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QPushButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPushButton) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPushButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPushButtonTimerEvent
func callbackQPushButtonTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QPushButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QPushButton) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QPushButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQPushButtonActionEvent
func callbackQPushButtonActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QPushButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QPushButton) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QPushButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQPushButtonDragEnterEvent
func callbackQPushButtonDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QPushButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QPushButton) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QPushButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQPushButtonDragLeaveEvent
func callbackQPushButtonDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QPushButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QPushButton) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QPushButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQPushButtonDragMoveEvent
func callbackQPushButtonDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QPushButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QPushButton) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QPushButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQPushButtonDropEvent
func callbackQPushButtonDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPushButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QPushButton) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QPushButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQPushButtonEnterEvent
func callbackQPushButtonEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QPushButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QPushButton) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QPushButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQPushButtonHideEvent
func callbackQPushButtonHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPushButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QPushButton) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QPushButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQPushButtonLeaveEvent
func callbackQPushButtonLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QPushButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QPushButton) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QPushButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQPushButtonMoveEvent
func callbackQPushButtonMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QPushButton::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QPushButton) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QPushButton::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQPushButtonSetVisible
func callbackQPushButtonSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QPushButton::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QPushButton::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QPushButton) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QPushButton::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQPushButtonShowEvent
func callbackQPushButtonShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QPushButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QPushButton) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QPushButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQPushButtonCloseEvent
func callbackQPushButtonCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QPushButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QPushButton) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QPushButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQPushButtonContextMenuEvent
func callbackQPushButtonContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QPushButton::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QPushButton) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QPushButton::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQPushButtonInitPainter
func callbackQPushButtonInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QPushButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QPushButton) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QPushButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQPushButtonInputMethodEvent
func callbackQPushButtonInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QPushButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QPushButton) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QPushButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQPushButtonMouseDoubleClickEvent
func callbackQPushButtonMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QPushButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QPushButton) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QPushButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQPushButtonResizeEvent
func callbackQPushButtonResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QPushButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QPushButton) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QPushButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQPushButtonTabletEvent
func callbackQPushButtonTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QPushButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QPushButton) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QPushButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQPushButtonWheelEvent
func callbackQPushButtonWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPushButton::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPushButton) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPushButton::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPushButtonChildEvent
func callbackQPushButtonChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QPushButton) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPushButton::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPushButton) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPushButton::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPushButtonCustomEvent
func callbackQPushButtonCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPushButton::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
