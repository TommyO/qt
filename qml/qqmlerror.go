package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlError struct {
	ptr unsafe.Pointer
}

type QQmlError_ITF interface {
	QQmlError_PTR() *QQmlError
}

func (p *QQmlError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlError(ptr QQmlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlError_PTR().Pointer()
	}
	return nil
}

func NewQQmlErrorFromPointer(ptr unsafe.Pointer) *QQmlError {
	var n = new(QQmlError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlError) QQmlError_PTR() *QQmlError {
	return ptr
}

func NewQQmlError() *QQmlError {
	defer qt.Recovering("QQmlError::QQmlError")

	return NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError())
}

func NewQQmlError2(other QQmlError_ITF) *QQmlError {
	defer qt.Recovering("QQmlError::QQmlError")

	return NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError2(PointerFromQQmlError(other)))
}

func (ptr *QQmlError) Column() int {
	defer qt.Recovering("QQmlError::column")

	if ptr.Pointer() != nil {
		return int(C.QQmlError_Column(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlError) Description() string {
	defer qt.Recovering("QQmlError::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) IsValid() bool {
	defer qt.Recovering("QQmlError::isValid")

	if ptr.Pointer() != nil {
		return C.QQmlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlError) Line() int {
	defer qt.Recovering("QQmlError::line")

	if ptr.Pointer() != nil {
		return int(C.QQmlError_Line(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlError) Object() *core.QObject {
	defer qt.Recovering("QQmlError::object")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlError_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlError) SetColumn(column int) {
	defer qt.Recovering("QQmlError::setColumn")

	if ptr.Pointer() != nil {
		C.QQmlError_SetColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QQmlError) SetDescription(description string) {
	defer qt.Recovering("QQmlError::setDescription")

	if ptr.Pointer() != nil {
		C.QQmlError_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QQmlError) SetLine(line int) {
	defer qt.Recovering("QQmlError::setLine")

	if ptr.Pointer() != nil {
		C.QQmlError_SetLine(ptr.Pointer(), C.int(line))
	}
}

func (ptr *QQmlError) SetObject(object core.QObject_ITF) {
	defer qt.Recovering("QQmlError::setObject")

	if ptr.Pointer() != nil {
		C.QQmlError_SetObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QQmlError) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlError::setUrl")

	if ptr.Pointer() != nil {
		C.QQmlError_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlError) ToString() string {
	defer qt.Recovering("QQmlError::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlError_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQmlError) Url() *core.QUrl {
	defer qt.Recovering("QQmlError::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQmlError_Url(ptr.Pointer()))
	}
	return nil
}