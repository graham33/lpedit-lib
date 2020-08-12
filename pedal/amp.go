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

package pedal

import "fmt"

const (
    ampDisabled uint32 = 524287
)

type Amp struct {
    id      uint32
    atype   uint32
    active  bool
    name    string
    pos     uint16
    posType uint8
    params  []Parameter
    pb      *PedalBoard
}

var amps = []Amp {
    Amp{atype: ampDisabled, active: true, name: "Amp Disabled"},
    Amp{atype: 458752, active: true, name: "phD Motorway",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458753, active: true, name: "Tweed B-Man Normal",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458754, active: true, name: "Tweed B-Man Bright",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458755, active: true, name: "Blackface ‘Lux Normal",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458756, active: true, name: "Blackface ‘Lux Vibrato",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458757, active: true, name: "Blackface Double Normal",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458758, active: true, name: "Blackface Double Vibrato",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458760, active: true, name: "Hiway 100",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458761, active: true, name: "Brit J-45 Normal",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458762, active: true, name: "Brit J-45 Bright",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458763, active: true, name: "Treadplate",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458764, active: true, name: "Brit P-75 Normal",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458765, active: true, name: "Brit P-75 Bright",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458766, active: true, name: "Super O",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458768, active: true, name: "Class A-15",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458769, active: true, name: "Class A-30 TB",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458770, active: true, name: "Divide 9/15",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458772, active: true, name: "Gibtone 185",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458773, active: true, name: "Brit J-800",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458775, active: true, name: "Bomber Uber",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458777, active: true, name: "Angel F-Ball",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458779, active: true, name: "phD Motorway Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458780, active: true, name: "Tweed B-Man Normal Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458781, active: true, name: "Tweed B-Man Bright Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458782, active: true, name: "Blackface ‘Lux Normal Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458783, active: true, name: "Blackface ‘Lux Vibrato Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458784, active: true, name: "Blackface Double Normal Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458785, active: true, name: "Blackface Double Vibrato Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458787, active: true, name: "Hiway 100 Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458788, active: true, name: "Brit J-45 Normal Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458789, active: true, name: "Brit J-45 Bright Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458790, active: true, name: "Treadplate Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458791, active: true, name: "Brit P-75 Normal Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458792, active: true, name: "Brit P-75 Bright Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458793, active: true, name: "Super O Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458795, active: true, name: "Class A-15 Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458796, active: true, name: "Class A-30 TB Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458797, active: true, name: "Divide 9/15 Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458799, active: true, name: "Gibtone 185 Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458800, active: true, name: "Brit J-800 Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458802, active: true, name: "Bomber Uber Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458804, active: true, name: "Angel F-Ball Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458834, active: true, name: "Line 6 Elektrik",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458835, active: true, name: "Line 6 Elektrik Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458837, active: true, name: "Plexi Lead 100 Normal",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458838, active: true, name: "Plexi Lead 100 Normal Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458840, active: true, name: "Plexi Lead 100 Bright",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458841, active: true, name: "Plexi Lead 100 Bright Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458843, active: true, name: "Flip Top",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458844, active: true, name: "Flip Top Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458846, active: true, name: "Solo 100 Clean",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458847, active: true, name: "Solo 100 Clean Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458849, active: true, name: "Solo 100 Crunch",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458850, active: true, name: "Solo 100 Crunch Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458852, active: true, name: "Solo 100 OD",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458853, active: true, name: "Solo 100 OD Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458855, active: true, name: "Line 6 Doom",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458856, active: true, name: "Line 6 Doom Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
    Amp{atype: 458858, active: true, name: "Line 6 Epic",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000B, name: "Master"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100008, name: "SAG"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100007, name: "HUM"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100009, name: "Bias"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F10000A, name: "Bias X"}},
            }},
    Amp{atype: 458859, active: true, name: "Line 6 Epic Preamp",
        params: []Parameter{
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100003, name: "Drive"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100000, name: "Bass"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100001, name: "Mid"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100002, name: "Treble"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100004, name: "Presence"}},
            &PerCentParam{GenericParameter: GenericParameter{id: 0x3F100005, name: "Volume"}},
            }},
}

func newDisAmp(id uint32, pos uint16, posType uint8, pb *PedalBoard) *Amp {
    a := newAmp(id, pos, posType, pb, ampDisabled)
    return a
}

func newAmp(id uint32, pos uint16, posType uint8, pb *PedalBoard, atype uint32) *Amp {
    for _, newAmp := range amps {
        if newAmp.atype == atype {
            newAmp.id = id
            newAmp.pos = pos
            newAmp.posType = posType
            newAmp.pb = pb
            for i := range newAmp.params {
                newAmp.params[i] = newAmp.params[i].Copy()
                newAmp.params[i].SetParent(&newAmp)
            }
            return &newAmp
        }
    }
    return nil
}

func (a *Amp) GetActive() bool {
    return a.active
}

func (a *Amp) GetActive2() uint32 {
    if a.active {
        return 1
    }
    return 0
}

func GetAmpType() []string {
    m := make([]string, len(amps))
    for i, a := range amps {
        m[i] = a.GetName()
    }
    return m
}

func (a *Amp) GetID() uint32 {
    return a.id
}

func (a *Amp) GetName() string {
    return a.name
}

func (a *Amp) GetParam(id uint32) Parameter {
    for _, param := range a.params {
        if param.GetID() == id {
            return param
        }
    }
    return nil
}

func (a *Amp) GetParams() []Parameter {
    return a.params
}

func (a *Amp) GetParamLen() uint16 {
    return uint16(len(a.params))
}

func (a *Amp) GetPos() (uint16, uint8) {
    return a.pos, a.posType
}

func (a *Amp) GetType() uint32 {
    return a.atype
}

func (a *Amp) LockData() { a.pb.LockData() }

func (a *Amp) SetActive(active bool){
    a.active = active
}

func (a *Amp) SetPos(pos uint16, posType uint8) {
    if posType == PedalPosAStart || posType == PedalPosBStart ||
     posType == PedalPosAEnd || posType == PedalPosBEnd { return }
    a.pos = pos
    a.posType = posType
}

func (a *Amp) SetType(atype uint32) error{
    _a := newAmp(a.id, a.pos, a.posType, a.pb, atype)
    if _a == nil {
        return fmt.Errorf("Amp type not found, code: %d", atype)
    }
    *a = *_a
    for i := range a.params {
        a.params[i].SetParent(a)
    }
    return nil
}

func (a *Amp) SetType2(name string, none string) {
    for _, _a := range amps {
        if name == _a.name {
            a.SetType(_a.atype)
            break
        }
    }
}

func (a *Amp) UnlockData() { a.pb.UnlockData() }
