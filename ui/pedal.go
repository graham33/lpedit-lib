/*
 * Copyright (C) 2020 Nicolas SCHWARTZ
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301, USA
 */

package ui

/*import "github.com/StarAurryon/qt/core"
import "github.com/StarAurryon/qt/gui"
import "github.com/StarAurryon/qt/svg"*/
import "github.com/StarAurryon/qt/widgets"

import "fmt"
import "sort"

import "github.com/StarAurryon/lpedit/model/pod"
import "github.com/StarAurryon/lpedit/qtctrl"

type Pedal struct {
    *PedalUI
    id         int
    pedalType  map[string][]string
    ctrl       *qtctrl.Controller
    parameters [5]Parameter
    parent     *LPEdit
}

func NewPedal(parent *LPEdit, w widgets.QWidget_ITF, c *qtctrl.Controller,
        pt map[string][]string, id int) *Pedal {
    p := &Pedal{PedalUI: NewPedalUI(w), ctrl: c, pedalType: pt, id: id}
    p.parameters[0] = Parameter{label: p.Param0Lbl, mid: p.Param0Mid,
        value: p.Param0Value, vfunc: p.parameter0Changed}
    p.parameters[1] = Parameter{label: p.Param1Lbl, mid: p.Param1Mid,
        value: p.Param1Value, vfunc: p.parameter1Changed}
    p.parameters[2] = Parameter{label: p.Param2Lbl, mid: p.Param2Mid,
        value: p.Param2Value, vfunc: p.parameter2Changed}
    p.parameters[3] = Parameter{label: p.Param3Lbl, mid: p.Param3Mid,
        value: p.Param3Value, vfunc: p.parameter3Changed}
    p.parameters[4] = Parameter{label: p.Param4Lbl, mid: p.Param4Mid,
        value: p.Param4Value, vfunc: p.parameter4Changed}
    p.parent = parent
    p.init()
    p.initUI()
    return p
}

func (p *Pedal) connectSignal() {
    keys := make([]string, 0, len(p.pedalType))
    for k := range p.pedalType {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    p.FxType.AddItems(keys)

    p.OnStatus.ConnectClicked(p.onStatusChanged)
    p.FxModel.ConnectActivated2(p.fxModelUserChanged)
    p.FxType.ConnectActivated2(p.fxTypeUserChanged)
    p.FxType.ConnectCurrentTextChanged(p.fxTypeChanged)
    for _, param := range p.parameters {
        param.value.ConnectActivated2(param.vfunc)
    }
}

func (p *Pedal) disconnectSignal() {
    p.OnStatus.DisconnectClicked()
    p.FxModel.DisconnectActivated2()
    p.FxType.DisconnectActivated2()
    p.FxType.DisconnectCurrentTextChanged()
    for i, param := range p.parameters {
        param.show()
        param.label.SetText(fmt.Sprintf("Parameter %d", i))
        param.value.DisconnectActivated2()
        param.value.Clear()
        param.value.SetEditable(false)
    }

    p.FxType.Clear()
    p.FxModel.Clear()
    p.OnStatus.SetChecked(false)
}

func (p *Pedal) initUI() {
    //Setting up knob
    /*svg := svg.NewQSvgRenderer2("ui/knob.svg", p)
    pix := gui.NewQPixmap2(p.Param1Knob.SizeHint())
    paint := gui.NewQPainter()
    paint.Scale(1, 1)

    pix.Fill(gui.NewQColor2(core.Qt__transparent))
    paint.Begin(pix)
    svg.Render(paint)
    paint.End()

    pal := gui.NewQPalette()
    pal.SetBrush(gui.QPalette__Background, gui.NewQBrush7(pix))
    p.Param1Knob.SetPalette(pal)*/

    //Setting up pod.type
}

func (p *Pedal) fxModelUserChanged(fxModel string) {
    p.ctrl.SetPedalType(uint32(p.id), p.FxType.CurrentText(), fxModel)
}

func (p *Pedal) fxTypeUserChanged(fxType string) {
    p.ctrl.SetPedalType(uint32(p.id), fxType, p.FxModel.CurrentText())
}

func (p *Pedal) fxTypeChanged(fxType string) {
    p.FxModel.Clear()
    p.FxModel.AddItems(p.pedalType[fxType])
}

func (p *Pedal) getParameter(id uint32) *Parameter{
    for i, param := range p.parameters {
        if param.id == id {
            return &p.parameters[i]
        }
    }
    return nil
}

func (p *Pedal) parameter0Changed(v string){ p.parameterChanged(&p.parameters[0], v) }
func (p *Pedal) parameter1Changed(v string){ p.parameterChanged(&p.parameters[1], v) }
func (p *Pedal) parameter2Changed(v string){ p.parameterChanged(&p.parameters[2], v) }
func (p *Pedal) parameter3Changed(v string){ p.parameterChanged(&p.parameters[3], v) }
func (p *Pedal) parameter4Changed(v string){ p.parameterChanged(&p.parameters[4], v) }

func (p *Pedal) parameterChanged(paramUI *Parameter, v string) {
    fmt.Println(paramUI.id)
    err := p.ctrl.SetPedalParameterValue(uint32(p.id), paramUI.id, v)
    if err != nil {
        mb := widgets.NewQMessageBox(p)
        mb.Critical(p, "An error occured", err.Error(), widgets.QMessageBox__Ok, 0)
    }
    param := p.ctrl.GetPedalBoard().GetPedal2(p.id).GetParam(paramUI.id)
    param.LockData()
    p.updateParam(param)
    param.UnlockData()
}

func (p *Pedal) onStatusChanged(checked bool) {
    p.ctrl.SetPedalActive(uint32(p.id), checked)
}

func (p *Pedal) setActive(status bool) {
    p.OnStatus.SetChecked(status)
}

func (pUI *Pedal) updatePedal(p *pod.Pedal) {
    pUI.setActive(p.GetActive())
    pUI.FxType.SetCurrentText(p.GetSType())
    pUI.FxModel.SetCurrentText(p.GetName())
    for i := range pUI.parameters {
        pUI.parameters[i].id = 0
        pUI.parameters[i].hide()
    }
    for i, param := range p.GetParams() {
        pUI.parameters[i].id = param.GetID()
        pUI.updateParam(param)
    }
}

func (pUI * Pedal) updateParam(p pod.Parameter) {
    param := pUI.getParameter(p.GetID())
    if param == nil { return }
    values := p.GetAllowedValues()

    valueIn := false
    for _, v := range values {
        if v == p.GetValueCurrent() {
            valueIn = true
            break
        }
    }

    if !valueIn {
        values = append([]string{p.GetValueCurrent()}, values...)
    }

    param.setValueList(values)
    if p.IsAllowingOtherValues() {
        param.setValueEditable(true)
    } else {
        param.setValueEditable(false)
    }
    param.setValue(p.GetValueCurrent())
    param.setLabel(p.GetName())
    param.show()
}
