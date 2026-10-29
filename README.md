# and-plugin-pembis — Pembiş Theme

**English** · [Türkçe](README.tr.md)

A small theme plugin for [AND](https://github.com/AND-community/AND) that instantly paints the AND
interface in shades of pink (or switches it back to the default).

---

## What it does

Using the `POST /api/v1/theme` endpoint, it changes the accent/selection/name/badge colours of AND's
main TUI. The change applies **instantly** — there is no need to restart AND. Other plugins such as
DeltaX read the same theme with `GET /api/v1/theme` when they start and apply it to their own
colours (see `internal/ui` in AND).

Pembiş doesn't process or display any data that comes from the network or from other users — it
only shows its own fixed menu. It asks for one permission, `theme`, the first time it is opened.

## Usage

Choose **Pembiş Theme** in AND's main menu.

| Key | Action |
|-----|--------|
| `↑` / `↓` or `j` / `k` | Move through the menu |
| `enter` | Apply the selected action |
| `q` / `ctrl+c` | Quit |

Menu options:

1. **Apply the Pink Theme** — sets the accent/selection/name/badge colours to shades of pink.
2. **Back to the Default Theme** — instantly returns to AND's default Classic theme.
3. **Quit**.

After an action is applied, press `enter` or `q` on the result screen to leave the plugin.

The plugin opens in AND's interface language (Turkish, English, German, Spanish, French or Russian).

## Installation

Pembiş lives in its own repository; it is cloned into the `Eklentiler/` folder of an
[AND](https://github.com/AND-community/AND) checkout and built from AND's root directory:

```bash
git clone https://github.com/AND-community/Pembis.git Eklentiler/Pembis
go build -o Eklentiler/Pembis/and-plugin-pembis ./Eklentiler/Pembis

# Windows
go build -o Eklentiler/Pembis/and-plugin-pembis.exe ./Eklentiler/Pembis
```

AND finds `Eklentiler/Pembis/plugin.json` and the binary next to it (named after the `name` field:
`and-plugin-pembis`) when it starts. The repository also contains a prebuilt Windows binary,
`and-plugin-pembis.exe`.

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

The label and description are in Turkish in the manifest; AND shows them translated in the other
interface languages.
