package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QMimeData struct {
	QObject
}

type QMimeData_ITF interface {
	QObject_ITF
	QMimeData_PTR() *QMimeData
}

func PointerFromQMimeData(ptr QMimeData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMimeData_PTR().Pointer()
	}
	return nil
}

func NewQMimeDataFromPointer(ptr unsafe.Pointer) *QMimeData {
	var n = new(QMimeData)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMimeData_") {
		n.SetObjectName("QMimeData_" + qt.Identifier())
	}
	return n
}

func (ptr *QMimeData) QMimeData_PTR() *QMimeData {
	return ptr
}

func NewQMimeData() *QMimeData {
	defer qt.Recovering("QMimeData::QMimeData")

	return NewQMimeDataFromPointer(C.QMimeData_NewQMimeData())
}

func (ptr *QMimeData) Clear() {
	defer qt.Recovering("QMimeData::clear")

	if ptr.Pointer() != nil {
		C.QMimeData_Clear(ptr.Pointer())
	}
}

func (ptr *QMimeData) ColorData() *QVariant {
	defer qt.Recovering("QMimeData::colorData")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QMimeData_ColorData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMimeData) Data(mimeType string) *QByteArray {
	defer qt.Recovering("QMimeData::data")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QMimeData_Data(ptr.Pointer(), C.CString(mimeType)))
	}
	return nil
}

func (ptr *QMimeData) Formats() []string {
	defer qt.Recovering("QMimeData::formats")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QMimeData_Formats(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QMimeData) HasColor() bool {
	defer qt.Recovering("QMimeData::hasColor")

	if ptr.Pointer() != nil {
		return C.QMimeData_HasColor(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeData) HasFormat(mimeType string) bool {
	defer qt.Recovering("QMimeData::hasFormat")

	if ptr.Pointer() != nil {
		return C.QMimeData_HasFormat(ptr.Pointer(), C.CString(mimeType)) != 0
	}
	return false
}

func (ptr *QMimeData) HasHtml() bool {
	defer qt.Recovering("QMimeData::hasHtml")

	if ptr.Pointer() != nil {
		return C.QMimeData_HasHtml(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeData) HasImage() bool {
	defer qt.Recovering("QMimeData::hasImage")

	if ptr.Pointer() != nil {
		return C.QMimeData_HasImage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeData) HasText() bool {
	defer qt.Recovering("QMimeData::hasText")

	if ptr.Pointer() != nil {
		return C.QMimeData_HasText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeData) HasUrls() bool {
	defer qt.Recovering("QMimeData::hasUrls")

	if ptr.Pointer() != nil {
		return C.QMimeData_HasUrls(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMimeData) Html() string {
	defer qt.Recovering("QMimeData::html")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeData_Html(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMimeData) ImageData() *QVariant {
	defer qt.Recovering("QMimeData::imageData")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QMimeData_ImageData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMimeData) RemoveFormat(mimeType string) {
	defer qt.Recovering("QMimeData::removeFormat")

	if ptr.Pointer() != nil {
		C.QMimeData_RemoveFormat(ptr.Pointer(), C.CString(mimeType))
	}
}

func (ptr *QMimeData) SetColorData(color QVariant_ITF) {
	defer qt.Recovering("QMimeData::setColorData")

	if ptr.Pointer() != nil {
		C.QMimeData_SetColorData(ptr.Pointer(), PointerFromQVariant(color))
	}
}

func (ptr *QMimeData) SetData(mimeType string, data QByteArray_ITF) {
	defer qt.Recovering("QMimeData::setData")

	if ptr.Pointer() != nil {
		C.QMimeData_SetData(ptr.Pointer(), C.CString(mimeType), PointerFromQByteArray(data))
	}
}

func (ptr *QMimeData) SetHtml(html string) {
	defer qt.Recovering("QMimeData::setHtml")

	if ptr.Pointer() != nil {
		C.QMimeData_SetHtml(ptr.Pointer(), C.CString(html))
	}
}

func (ptr *QMimeData) SetImageData(image QVariant_ITF) {
	defer qt.Recovering("QMimeData::setImageData")

	if ptr.Pointer() != nil {
		C.QMimeData_SetImageData(ptr.Pointer(), PointerFromQVariant(image))
	}
}

func (ptr *QMimeData) SetText(text string) {
	defer qt.Recovering("QMimeData::setText")

	if ptr.Pointer() != nil {
		C.QMimeData_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QMimeData) Text() string {
	defer qt.Recovering("QMimeData::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeData_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMimeData) DestroyQMimeData() {
	defer qt.Recovering("QMimeData::~QMimeData")

	if ptr.Pointer() != nil {
		C.QMimeData_DestroyQMimeData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}