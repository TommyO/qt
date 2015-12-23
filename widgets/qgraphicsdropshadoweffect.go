package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsDropShadowEffect struct {
	QGraphicsEffect
}

type QGraphicsDropShadowEffect_ITF interface {
	QGraphicsEffect_ITF
	QGraphicsDropShadowEffect_PTR() *QGraphicsDropShadowEffect
}

func PointerFromQGraphicsDropShadowEffect(ptr QGraphicsDropShadowEffect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsDropShadowEffect_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsDropShadowEffectFromPointer(ptr unsafe.Pointer) *QGraphicsDropShadowEffect {
	var n = new(QGraphicsDropShadowEffect)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsDropShadowEffect_") {
		n.SetObjectName("QGraphicsDropShadowEffect_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsDropShadowEffect) QGraphicsDropShadowEffect_PTR() *QGraphicsDropShadowEffect {
	return ptr
}

func (ptr *QGraphicsDropShadowEffect) BlurRadius() float64 {
	defer qt.Recovering("QGraphicsDropShadowEffect::blurRadius")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsDropShadowEffect_BlurRadius(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsDropShadowEffect) Color() *gui.QColor {
	defer qt.Recovering("QGraphicsDropShadowEffect::color")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QGraphicsDropShadowEffect_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsDropShadowEffect) SetBlurRadius(blurRadius float64) {
	defer qt.Recovering("QGraphicsDropShadowEffect::setBlurRadius")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_SetBlurRadius(ptr.Pointer(), C.double(blurRadius))
	}
}

func (ptr *QGraphicsDropShadowEffect) SetColor(color gui.QColor_ITF) {
	defer qt.Recovering("QGraphicsDropShadowEffect::setColor")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QGraphicsDropShadowEffect) SetOffset(ofs core.QPointF_ITF) {
	defer qt.Recovering("QGraphicsDropShadowEffect::setOffset")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_SetOffset(ptr.Pointer(), core.PointerFromQPointF(ofs))
	}
}

func NewQGraphicsDropShadowEffect(parent core.QObject_ITF) *QGraphicsDropShadowEffect {
	defer qt.Recovering("QGraphicsDropShadowEffect::QGraphicsDropShadowEffect")

	return NewQGraphicsDropShadowEffectFromPointer(C.QGraphicsDropShadowEffect_NewQGraphicsDropShadowEffect(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsDropShadowEffect) ConnectBlurRadiusChanged(f func(blurRadius float64)) {
	defer qt.Recovering("connect QGraphicsDropShadowEffect::blurRadiusChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_ConnectBlurRadiusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "blurRadiusChanged", f)
	}
}

func (ptr *QGraphicsDropShadowEffect) DisconnectBlurRadiusChanged() {
	defer qt.Recovering("disconnect QGraphicsDropShadowEffect::blurRadiusChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_DisconnectBlurRadiusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "blurRadiusChanged")
	}
}

//export callbackQGraphicsDropShadowEffectBlurRadiusChanged
func callbackQGraphicsDropShadowEffectBlurRadiusChanged(ptrName *C.char, blurRadius C.double) {
	defer qt.Recovering("callback QGraphicsDropShadowEffect::blurRadiusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "blurRadiusChanged"); signal != nil {
		signal.(func(float64))(float64(blurRadius))
	}

}

func (ptr *QGraphicsDropShadowEffect) ConnectColorChanged(f func(color *gui.QColor)) {
	defer qt.Recovering("connect QGraphicsDropShadowEffect::colorChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "colorChanged", f)
	}
}

func (ptr *QGraphicsDropShadowEffect) DisconnectColorChanged() {
	defer qt.Recovering("disconnect QGraphicsDropShadowEffect::colorChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "colorChanged")
	}
}

//export callbackQGraphicsDropShadowEffectColorChanged
func callbackQGraphicsDropShadowEffectColorChanged(ptrName *C.char, color unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsDropShadowEffect::colorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(color))
	}

}

func (ptr *QGraphicsDropShadowEffect) SetOffset3(d float64) {
	defer qt.Recovering("QGraphicsDropShadowEffect::setOffset")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_SetOffset3(ptr.Pointer(), C.double(d))
	}
}

func (ptr *QGraphicsDropShadowEffect) SetOffset2(dx float64, dy float64) {
	defer qt.Recovering("QGraphicsDropShadowEffect::setOffset")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_SetOffset2(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QGraphicsDropShadowEffect) SetXOffset(dx float64) {
	defer qt.Recovering("QGraphicsDropShadowEffect::setXOffset")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_SetXOffset(ptr.Pointer(), C.double(dx))
	}
}

func (ptr *QGraphicsDropShadowEffect) SetYOffset(dy float64) {
	defer qt.Recovering("QGraphicsDropShadowEffect::setYOffset")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_SetYOffset(ptr.Pointer(), C.double(dy))
	}
}

func (ptr *QGraphicsDropShadowEffect) XOffset() float64 {
	defer qt.Recovering("QGraphicsDropShadowEffect::xOffset")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsDropShadowEffect_XOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsDropShadowEffect) YOffset() float64 {
	defer qt.Recovering("QGraphicsDropShadowEffect::yOffset")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsDropShadowEffect_YOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsDropShadowEffect) DestroyQGraphicsDropShadowEffect() {
	defer qt.Recovering("QGraphicsDropShadowEffect::~QGraphicsDropShadowEffect")

	if ptr.Pointer() != nil {
		C.QGraphicsDropShadowEffect_DestroyQGraphicsDropShadowEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsDropShadowEffect) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsDropShadowEffect::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsDropShadowEffect) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsDropShadowEffect::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsDropShadowEffectTimerEvent
func callbackQGraphicsDropShadowEffectTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsDropShadowEffect::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsDropShadowEffect) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsDropShadowEffect::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsDropShadowEffect) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsDropShadowEffect::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsDropShadowEffectChildEvent
func callbackQGraphicsDropShadowEffectChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsDropShadowEffect::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsDropShadowEffect) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsDropShadowEffect::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsDropShadowEffect) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsDropShadowEffect::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsDropShadowEffectCustomEvent
func callbackQGraphicsDropShadowEffectCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsDropShadowEffect::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
