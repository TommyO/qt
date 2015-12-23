package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioDecoderControl struct {
	QMediaControl
}

type QAudioDecoderControl_ITF interface {
	QMediaControl_ITF
	QAudioDecoderControl_PTR() *QAudioDecoderControl
}

func PointerFromQAudioDecoderControl(ptr QAudioDecoderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioDecoderControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioDecoderControlFromPointer(ptr unsafe.Pointer) *QAudioDecoderControl {
	var n = new(QAudioDecoderControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioDecoderControl_") {
		n.SetObjectName("QAudioDecoderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioDecoderControl) QAudioDecoderControl_PTR() *QAudioDecoderControl {
	return ptr
}

func (ptr *QAudioDecoderControl) BufferAvailable() bool {
	defer qt.Recovering("QAudioDecoderControl::bufferAvailable")

	if ptr.Pointer() != nil {
		return C.QAudioDecoderControl_BufferAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioDecoderControl) ConnectBufferAvailableChanged(f func(available bool)) {
	defer qt.Recovering("connect QAudioDecoderControl::bufferAvailableChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectBufferAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferAvailableChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectBufferAvailableChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::bufferAvailableChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectBufferAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferAvailableChanged")
	}
}

//export callbackQAudioDecoderControlBufferAvailableChanged
func callbackQAudioDecoderControlBufferAvailableChanged(ptrName *C.char, available C.int) {
	defer qt.Recovering("callback QAudioDecoderControl::bufferAvailableChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferAvailableChanged"); signal != nil {
		signal.(func(bool))(int(available) != 0)
	}

}

func (ptr *QAudioDecoderControl) ConnectBufferReady(f func()) {
	defer qt.Recovering("connect QAudioDecoderControl::bufferReady")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectBufferReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferReady", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectBufferReady() {
	defer qt.Recovering("disconnect QAudioDecoderControl::bufferReady")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectBufferReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferReady")
	}
}

//export callbackQAudioDecoderControlBufferReady
func callbackQAudioDecoderControlBufferReady(ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::bufferReady")

	if signal := qt.GetSignal(C.GoString(ptrName), "bufferReady"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoderControl) Duration() int64 {
	defer qt.Recovering("QAudioDecoderControl::duration")

	if ptr.Pointer() != nil {
		return int64(C.QAudioDecoderControl_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioDecoderControl) ConnectDurationChanged(f func(duration int64)) {
	defer qt.Recovering("connect QAudioDecoderControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectDurationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "durationChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectDurationChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::durationChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectDurationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "durationChanged")
	}
}

//export callbackQAudioDecoderControlDurationChanged
func callbackQAudioDecoderControlDurationChanged(ptrName *C.char, duration C.longlong) {
	defer qt.Recovering("callback QAudioDecoderControl::durationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "durationChanged"); signal != nil {
		signal.(func(int64))(int64(duration))
	}

}

func (ptr *QAudioDecoderControl) ConnectError(f func(error int, errorString string)) {
	defer qt.Recovering("connect QAudioDecoderControl::error")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectError() {
	defer qt.Recovering("disconnect QAudioDecoderControl::error")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQAudioDecoderControlError
func callbackQAudioDecoderControlError(ptrName *C.char, error C.int, errorString *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(int, string))(int(error), C.GoString(errorString))
	}

}

func (ptr *QAudioDecoderControl) ConnectFinished(f func()) {
	defer qt.Recovering("connect QAudioDecoderControl::finished")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectFinished() {
	defer qt.Recovering("disconnect QAudioDecoderControl::finished")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAudioDecoderControlFinished
func callbackQAudioDecoderControlFinished(ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoderControl) Position() int64 {
	defer qt.Recovering("QAudioDecoderControl::position")

	if ptr.Pointer() != nil {
		return int64(C.QAudioDecoderControl_Position(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioDecoderControl) ConnectPositionChanged(f func(position int64)) {
	defer qt.Recovering("connect QAudioDecoderControl::positionChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectPositionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "positionChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectPositionChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::positionChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectPositionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "positionChanged")
	}
}

//export callbackQAudioDecoderControlPositionChanged
func callbackQAudioDecoderControlPositionChanged(ptrName *C.char, position C.longlong) {
	defer qt.Recovering("callback QAudioDecoderControl::positionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "positionChanged"); signal != nil {
		signal.(func(int64))(int64(position))
	}

}

func (ptr *QAudioDecoderControl) SetAudioFormat(format QAudioFormat_ITF) {
	defer qt.Recovering("QAudioDecoderControl::setAudioFormat")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetAudioFormat(ptr.Pointer(), PointerFromQAudioFormat(format))
	}
}

func (ptr *QAudioDecoderControl) SetSourceDevice(device core.QIODevice_ITF) {
	defer qt.Recovering("QAudioDecoderControl::setSourceDevice")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetSourceDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioDecoderControl) SetSourceFilename(fileName string) {
	defer qt.Recovering("QAudioDecoderControl::setSourceFilename")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_SetSourceFilename(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QAudioDecoderControl) ConnectSourceChanged(f func()) {
	defer qt.Recovering("connect QAudioDecoderControl::sourceChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectSourceChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::sourceChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQAudioDecoderControlSourceChanged
func callbackQAudioDecoderControlSourceChanged(ptrName *C.char) {
	defer qt.Recovering("callback QAudioDecoderControl::sourceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sourceChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioDecoderControl) SourceDevice() *core.QIODevice {
	defer qt.Recovering("QAudioDecoderControl::sourceDevice")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioDecoderControl_SourceDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDecoderControl) SourceFilename() string {
	defer qt.Recovering("QAudioDecoderControl::sourceFilename")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoderControl_SourceFilename(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioDecoderControl) Start() {
	defer qt.Recovering("QAudioDecoderControl::start")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Start(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) State() QAudioDecoder__State {
	defer qt.Recovering("QAudioDecoderControl::state")

	if ptr.Pointer() != nil {
		return QAudioDecoder__State(C.QAudioDecoderControl_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioDecoderControl) ConnectStateChanged(f func(state QAudioDecoder__State)) {
	defer qt.Recovering("connect QAudioDecoderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QAudioDecoderControl::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioDecoderControlStateChanged
func callbackQAudioDecoderControlStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QAudioDecoderControl::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QAudioDecoder__State))(QAudioDecoder__State(state))
	}

}

func (ptr *QAudioDecoderControl) Stop() {
	defer qt.Recovering("QAudioDecoderControl::stop")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioDecoderControl) DestroyQAudioDecoderControl() {
	defer qt.Recovering("QAudioDecoderControl::~QAudioDecoderControl")

	if ptr.Pointer() != nil {
		C.QAudioDecoderControl_DestroyQAudioDecoderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioDecoderControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioDecoderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioDecoderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioDecoderControlTimerEvent
func callbackQAudioDecoderControlTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAudioDecoderControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAudioDecoderControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioDecoderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioDecoderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioDecoderControlChildEvent
func callbackQAudioDecoderControlChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAudioDecoderControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAudioDecoderControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioDecoderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioDecoderControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioDecoderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioDecoderControlCustomEvent
func callbackQAudioDecoderControlCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAudioDecoderControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
