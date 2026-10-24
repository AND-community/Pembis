package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/ansi"

	"github.com/lucian95511/and/dil/diltest"
	"github.com/lucian95511/and/internal/pluginapi"
)

// TestEkranlar_HerDilde: Pembiş'in menüsü ve sonuç ekranları her dilde
// çevrilmemiş metin içermez ("PEMBİŞ" afişi ve eklentinin adı özel addır).
func TestEkranlar_HerDilde(t *testing.T) {
	denetci := diltest.YeniDenetci(t, []string{"P E M B İ Ş", "Pembiş"}, diltest.Oku(t, "ceviri.json", ceviriVerisi))
	diltest.HerDilde(t, func(kod string) {
		menu := newModel(&pluginapi.Client{})
		next, _ := menu.Update(tea.WindowSizeMsg{Width: 100, Height: 30})
		menu = next.(model)
		pembe, varsayilan := menu, menu
		pembe.ekran, pembe.uygulanan, pembe.mesaj = ekOnay, temaPembe, "ok"
		varsayilan.ekran, varsayilan.uygulanan, varsayilan.mesaj = ekOnay, temaVarsayilan, "ok"
		for ad, m := range map[string]model{"menu": menu, "pembe": pembe, "varsayilan": varsayilan} {
			for _, k := range denetci.Kalintilar(kod, ansi.Strip(m.View())) {
				t.Errorf("%s %s: %s", kod, ad, k)
			}
		}
	})
}
