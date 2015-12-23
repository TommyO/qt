package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QMediaContainerControl struct {
	QMediaControl
}

type QMediaContainerControl_ITF interface {
	QMediaControl_ITF
	QMediaContainerControl_PTR() *QMediaContainerControl
}

func PointerFromQMediaContainerControl(ptr QMediaContainerControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaContainerControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaContainerControlFromPointer(ptr unsafe.Pointer) *QMediaContainerControl {
	var n = new(QMediaContainerControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaContainerControl_") {
		n.SetObjectName("QMediaContainerControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaContainerControl) QMediaContainerControl_PTR() *QMediaContainerControl {
	return ptr
}

func (ptr *QMediaContainerControl) ContainerDescription(format string) string {
	defer qt.Recovering("QMediaContainerControl::containerDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContainerControl_ContainerDescription(ptr.Pointer(), C.CString(format)))
	}
	return ""
}

func (ptr *QMediaContainerControl) ContainerFormat() string {
	defer qt.Recovering("QMediaContainerControl::containerFormat")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaContainerControl_ContainerFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaContainerControl) SetContainerFormat(format string) {
	defer qt.Recovering("QMediaContainerControl::setContainerFormat")

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_SetContainerFormat(ptr.Pointer(), C.CString(format))
	}
}

func (ptr *QMediaContainerControl) SupportedContainers() []string {
	defer qt.Recovering("QMediaContainerControl::supportedContainers")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMediaContainerControl_SupportedContainers(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QMediaContainerControl) DestroyQMediaContainerControl() {
	defer qt.Recovering("QMediaContainerControl::~QMediaContainerControl")

	if ptr.Pointer() != nil {
		C.QMediaContainerControl_DestroyQMediaContainerControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaContainerControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaContainerControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaContainerControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaContainerControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaContainerControlTimerEvent
func callbackQMediaContainerControlTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaContainerControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMediaContainerControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaContainerControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaContainerControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaContainerControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaContainerControlChildEvent
func callbackQMediaContainerControlChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaContainerControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMediaContainerControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaContainerControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaContainerControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaContainerControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaContainerControlCustomEvent
func callbackQMediaContainerControlCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaContainerControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
