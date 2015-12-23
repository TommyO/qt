package svg

//#include "svg.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QGraphicsSvgItem struct {
	widgets.QGraphicsObject
}

type QGraphicsSvgItem_ITF interface {
	widgets.QGraphicsObject_ITF
	QGraphicsSvgItem_PTR() *QGraphicsSvgItem
}

func PointerFromQGraphicsSvgItem(ptr QGraphicsSvgItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSvgItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSvgItemFromPointer(ptr unsafe.Pointer) *QGraphicsSvgItem {
	var n = new(QGraphicsSvgItem)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsSvgItem_") {
		n.SetObjectName("QGraphicsSvgItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsSvgItem) QGraphicsSvgItem_PTR() *QGraphicsSvgItem {
	return ptr
}

func (ptr *QGraphicsSvgItem) ElementId() string {
	defer qt.Recovering("QGraphicsSvgItem::elementId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsSvgItem_ElementId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsSvgItem) MaximumCacheSize() *core.QSize {
	defer qt.Recovering("QGraphicsSvgItem::maximumCacheSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QGraphicsSvgItem_MaximumCacheSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSvgItem) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget)) {
	defer qt.Recovering("connect QGraphicsSvgItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paint", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paint")
	}
}

//export callbackQGraphicsSvgItemPaint
func callbackQGraphicsSvgItemPaint(ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsSvgItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QGraphicsSvgItem) Renderer() *QSvgRenderer {
	defer qt.Recovering("QGraphicsSvgItem::renderer")

	if ptr.Pointer() != nil {
		return NewQSvgRendererFromPointer(C.QGraphicsSvgItem_Renderer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSvgItem) SetElementId(id string) {
	defer qt.Recovering("QGraphicsSvgItem::setElementId")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetElementId(ptr.Pointer(), C.CString(id))
	}
}

func (ptr *QGraphicsSvgItem) SetMaximumCacheSize(size core.QSize_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::setMaximumCacheSize")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetMaximumCacheSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QGraphicsSvgItem) SetSharedRenderer(renderer QSvgRenderer_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::setSharedRenderer")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetSharedRenderer(ptr.Pointer(), PointerFromQSvgRenderer(renderer))
	}
}

func (ptr *QGraphicsSvgItem) Type() int {
	defer qt.Recovering("QGraphicsSvgItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsSvgItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSvgItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsSvgItemTimerEvent
func callbackQGraphicsSvgItemTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsSvgItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsSvgItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsSvgItemChildEvent
func callbackQGraphicsSvgItemChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsSvgItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsSvgItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsSvgItemCustomEvent
func callbackQGraphicsSvgItemCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsSvgItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
