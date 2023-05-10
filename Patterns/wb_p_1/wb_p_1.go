package main

import (
	"fmt"
	"strings"
)

// структуры
type bass struct {
	note string
}
type drum struct {
	voice string
}
type piano struct {
	keyboard string
}

// создание
func newBass(n string) *bass {
	return &bass{n}
}
func newDrum(v string) *drum {
	return &drum{v}
}
func newPiano(k string) *piano {
	return &piano{k}
}

// методы
func (b *bass) playBass(n string) string {
	return "bass: " + n + " "
}
func (b *drum) playDrum(v string) string {
	return "drum: " + v + " "
}
func (b *piano) playPiano(k string) string {
	return "piano: " + k + " "
}

// фасад
type musician struct {
	bassInstr  *bass
	drumInstr  *drum
	pianoInstr *piano
}

func newMusician() *musician {
	return &musician{newBass(""), newDrum(""), newPiano("")}
}

func strInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.ToLower(b) == strings.ToLower(a) {
			return true
		}
	}
	return false
}

func (m *musician) play(str string, song string) (bool, string) {
	s := strings.Split(str, " ")
	if len(s) > 3 {
		return false, "wrong format"
	}
	res := ""
	if strInSlice("bass", s) {
		res += m.bassInstr.playBass(song)
	}
	if strInSlice("drum", s) {
		res += m.drumInstr.playDrum(song)
	}
	if strInSlice("piano", s) {
		res += m.pianoInstr.playPiano(song)
	}
	return true, res
}
func main() {

	mus := newMusician()
	_, res := mus.play("bass piAno", "dum dum")
	fmt.Println(res)
}
