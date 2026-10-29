# and-plugin-pembis — Pembiş Tema

[English](README.md) · **Türkçe**

[AND](https://github.com/AND-community/AND) arayüzünü anında pembe tonlara boyayan (ya da
varsayılana geri döndüren) küçük bir tema eklentisi.

---

## Ne yapar

`POST /api/v1/theme` uç noktasını kullanarak AND'ın ana TUI'sinin vurgu/seçim/isim/rozet renklerini
değiştirir. Değişiklik **anında** uygulanır — AND'ı yeniden başlatmak gerekmez; DeltaX gibi diğer
eklentiler de açılışta aynı temayı `GET /api/v1/theme` ile okuyup kendi renklerine uygular
(bkz. AND'daki `internal/ui`).

Pembiş, ağdan/başka kullanıcılardan gelen hiçbir veriyi işlemez veya ekrana yazmaz — yalnızca
kendi sabit menüsünü gösterir. İlk açılışta tek bir izin ister: `theme`.

## Kullanım

Ana menüden **Pembiş Tema**'yı seç.

| Tuş | İşlev |
|-----|-------|
| `↑` / `↓` ya da `j` / `k` | Menüde gez |
| `enter` | Seçili işlemi uygula |
| `q` / `ctrl+c` | Çık |

Menü seçenekleri:

1. **Pembe Temayı Uygula** — vurgu/seçim/isim/rozet renklerini pembe tonlara ayarlar.
2. **Varsayılan Temaya Dön** — AND'ın öntanımlı Klasik temasına anında geri döner.
3. **Çıkış**.

Bir işlem uygulandıktan sonra sonuç ekranında `enter` ya da `q` ile eklentiden çıkılır.

Eklenti AND'ın arayüz dilinde açılır (Türkçe, English, Deutsch, Español, Français, Русский).

## Kurulum

Pembiş ayrı bir depodur; bir [AND](https://github.com/AND-community/AND) kopyasının `Eklentiler/`
klasörüne klonlanıp AND'ın kök dizininden derlenir:

```bash
git clone https://github.com/AND-community/Pembis.git Eklentiler/Pembis
go build -o Eklentiler/Pembis/and-plugin-pembis ./Eklentiler/Pembis

# Windows
go build -o Eklentiler/Pembis/and-plugin-pembis.exe ./Eklentiler/Pembis
```

AND açılışta `Eklentiler/Pembis/plugin.json`'ı ve yanındaki binary'yi (`name` alanına göre
`and-plugin-pembis`) bulur. Depoda Windows için derlenmiş `and-plugin-pembis.exe` de bulunur.

## Manifest

```json
{
  "name":        "pembis",
  "label":       "Pembiş Tema",
  "version":     "1.2.0",
  "description": "AND arayüzünü pembe tonlarda kız temasına boyar",
  "author":      "AND",
  "permissions": ["theme"],
  "category":    "theme"
}
```
