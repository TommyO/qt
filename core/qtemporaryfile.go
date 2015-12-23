package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTemporaryFile struct {
	QFile
}

type QTemporaryFile_ITF interface {
	QFile_ITF
	QTemporaryFile_PTR() *QTemporaryFile
}

func PointerFromQTemporaryFile(ptr QTemporaryFile_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTemporaryFile_PTR().Pointer()
	}
	return nil
}

func NewQTemporaryFileFromPointer(ptr unsafe.Pointer) *QTemporaryFile {
	var n = new(QTemporaryFile)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTemporaryFile_") {
		n.SetObjectName("QTemporaryFile_" + qt.Identifier())
	}
	return n
}

func (ptr *QTemporaryFile) QTemporaryFile_PTR() *QTemporaryFile {
	return ptr
}

func NewQTemporaryFile() *QTemporaryFile {
	defer qt.Recovering("QTemporaryFile::QTemporaryFile")

	return NewQTemporaryFileFromPointer(C.QTemporaryFile_NewQTemporaryFile())
}

func NewQTemporaryFile3(parent QObject_ITF) *QTemporaryFile {
	defer qt.Recovering("QTemporaryFile::QTemporaryFile")

	return NewQTemporaryFileFromPointer(C.QTemporaryFile_NewQTemporaryFile3(PointerFromQObject(parent)))
}

func NewQTemporaryFile2(templateName string) *QTemporaryFile {
	defer qt.Recovering("QTemporaryFile::QTemporaryFile")

	return NewQTemporaryFileFromPointer(C.QTemporaryFile_NewQTemporaryFile2(C.CString(templateName)))
}

func NewQTemporaryFile4(templateName string, parent QObject_ITF) *QTemporaryFile {
	defer qt.Recovering("QTemporaryFile::QTemporaryFile")

	return NewQTemporaryFileFromPointer(C.QTemporaryFile_NewQTemporaryFile4(C.CString(templateName), PointerFromQObject(parent)))
}

func (ptr *QTemporaryFile) AutoRemove() bool {
	defer qt.Recovering("QTemporaryFile::autoRemove")

	if ptr.Pointer() != nil {
		return C.QTemporaryFile_AutoRemove(ptr.Pointer()) != 0
	}
	return false
}

func QTemporaryFile_CreateNativeFile(file QFile_ITF) *QTemporaryFile {
	defer qt.Recovering("QTemporaryFile::createNativeFile")

	return NewQTemporaryFileFromPointer(C.QTemporaryFile_QTemporaryFile_CreateNativeFile(PointerFromQFile(file)))
}

func QTemporaryFile_CreateNativeFile2(fileName string) *QTemporaryFile {
	defer qt.Recovering("QTemporaryFile::createNativeFile")

	return NewQTemporaryFileFromPointer(C.QTemporaryFile_QTemporaryFile_CreateNativeFile2(C.CString(fileName)))
}

func (ptr *QTemporaryFile) FileName() string {
	defer qt.Recovering("QTemporaryFile::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTemporaryFile_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTemporaryFile) FileTemplate() string {
	defer qt.Recovering("QTemporaryFile::fileTemplate")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTemporaryFile_FileTemplate(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTemporaryFile) Open() bool {
	defer qt.Recovering("QTemporaryFile::open")

	if ptr.Pointer() != nil {
		return C.QTemporaryFile_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTemporaryFile) SetAutoRemove(b bool) {
	defer qt.Recovering("QTemporaryFile::setAutoRemove")

	if ptr.Pointer() != nil {
		C.QTemporaryFile_SetAutoRemove(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTemporaryFile) SetFileTemplate(name string) {
	defer qt.Recovering("QTemporaryFile::setFileTemplate")

	if ptr.Pointer() != nil {
		C.QTemporaryFile_SetFileTemplate(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QTemporaryFile) DestroyQTemporaryFile() {
	defer qt.Recovering("QTemporaryFile::~QTemporaryFile")

	if ptr.Pointer() != nil {
		C.QTemporaryFile_DestroyQTemporaryFile(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTemporaryFile) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QTemporaryFile::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTemporaryFile) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTemporaryFile::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTemporaryFileTimerEvent
func callbackQTemporaryFileTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTemporaryFile::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTemporaryFile) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QTemporaryFile::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTemporaryFile) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTemporaryFile::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTemporaryFileChildEvent
func callbackQTemporaryFileChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTemporaryFile::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTemporaryFile) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QTemporaryFile::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTemporaryFile) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTemporaryFile::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTemporaryFileCustomEvent
func callbackQTemporaryFileCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTemporaryFile::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
