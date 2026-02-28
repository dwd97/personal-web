# 1. hodina

## Co je Linux

[Vývoj Unixu (wikipedie)](https://upload.wikimedia.org/wikipedia/commons/thumb/7/77/Unix_history-simple.svg/1920px-Unix_history-simple.svg.png)

### UNIX

- Ochranná známka pro operační systémy, co dodržují Single UNIX Specification. Definují API (systémové volání) pro komunikaci s hardwarem.

### Linux

- Jedná se o operační systémy založené na Linux kernelu, který je Unix-like a dodržuje POSIX (portable operating system interface).
- Správným názvem se myslí GNU/Linux, tedy se sem počítá jak Linux kernel, tak aplikace, utility a další software.
- Linuxová distribuce je balíček Linuxového jádra a uživatelských aplikací.

#### vlastnosti Linuxu
- **OSS (open source software)**
- **konfigurace v textových souborech**
- **GUI (grafické uživatelské rozhraní)** - většinou není potřeba
- **CLI (rozhraní příkazové řádky)** - přesné a zautomatizovatelné
- **všechno je soubor** - proto stačí jen jedno API pro všechno
- **víceuživatelský** - více uživatelů může používat systém naráz, vzdáleně i lokálně
- **balíčky** - pro instalace aplikací, ty se pak lépe udržují a automaticky aktualizují
- **lze řídit všechno** - pomocí CLI

## Pracovní prostředí

- ovlivňují vzhled a způsob ovládání Linuxu.

mezi nejčastěji používáné patří:
- **GNOME, Plasma (KDE), Xfce, LXDE** - podobné prostředí ostatním OS
- **Openbox, i3** - ovládájí se většinou klávesnicí, steep learning curve

### Zkratky v Linuxu

- `Super(Windows key) + L` - logs the user out
- `Ctrl + Q` - quits current application

## Odbočka: Git (správa verzí kódů)

- sledování změn a zapamatování předchozích stavů.

# 2. hodina

## Příkazová řádka

- Pomocí CLI lze systém ovládat přesně. Uložené soubory s příkazy se nazývají skripty. Namísto programovacích jazyků, kde se spouští funkce, tak se v CLI se spouští programy.

## Názvy souborů a cesty

- název souboru + adresář = cesta
- Oddělovačem v adresáři je `/` dopředné lomítko.
- Na Linuxu se v cestě rozlišují velká a malá písmena.
- adresář je speciální typ souboru. Vše je soubor.

### Kořenový adresář

- root/kořenový adresář je jediné dopředné lomítko `/`

### Relativní a absolutní cesta

- **relativní** - Nezačíná lomítkem. `adresar/soubor.txt` Potřeba zkombinovat s aktuální absolutní cestou. 
- **absolutní** - Začíná lomítkem. `/adresar/soubor.txt` Konkrétní soubor nehledě na tom, ve kterém adresáři právě jste.

### Speciální adresáře

- `..` - nadřazený (rodičovský) adresář. Příklad: lze zkombinovat `/home/documents/films/../dopis.odt` zkráceně jako `/home/documents/dopis.odt`
- `.` - aktuální adresář. Takže (relativní cesta) lze napsat `./bin/knihy/harry_poter.epub` nebo `bin/knihy/harry_poter.epub`

### Přípony souborů

- Linux nevyžaduje použití přípon u soborů a soubory mohou fungovat bez nich.
- Jiné soubory mohou mít přípon více, např. `tar.gz` (tape archive komprimován gzipem). Ty pak jsou vyvíjeny nezávisle a lze změnit gzip jiným komprimovacím algoritmem bez ohledu na archivační algoritmus.

### Skryté soubory

- Soubory s předponou `.` Jsou většinou skryté, protože se nepředpokládá, že by s nimi uživatel potřeboval přímo pracovat. Lze je odhalit v CLI nebo GUI.
- např.: `.gitignore` pro nastavení pravidel gitu nebo `.config` pro nastavení Linuxu

## Práce s terminálem

- `[intro@localhost] ~` je prompt. A `~` odkazuje na domovský adresář, ale ten může být delší. V domovském adresáři se ukládají všechny soubor uživatele včetně nastavení prostředí jako `.config`.
- **Shell** - zobrazuje výzvu a jedná se o plnohodnotný programovací jazyk. Shell je program běžící uvnitř terminálu (emulátoru terminálu).

### Zkratky/Používání

- `Ctrl + Left/Right Arrow` - přeskakování slov
- `Up/Down Arrow` - načítání předchozí zadané příkazy
- `Enter` - spuštění příkazů
- `Ctrl + D` - ukončení terminálu (nejčistší způsob)
- `Ctrl + C` - násilné ukončení programu
- `Ctrl + Shift + C` nebo `Select Text` - zkopírování obsahu do schránky
- `Middle Mouse Button` - vloží se text ze schránky

### Příkazy

- `exit` - ukončení terminálu
- `uptime` - vytiskne dobu běhu stroje
- `ls` - "list", zobrazí seznam souborů v aktuálním adresáři
    - `ls -l` - "list long", tzv. dlouhý režime, kde -l je přepínač
    - `ls -la` - vypíše i skryté soubory (v long list)
    - `ls -a` - vypíše i skryté soubory bez přepínače long
    - `ls -lh` - vypíše velikosti souborů čitelné pro normálního uživatele
    - `ls -lt` - vypíše seřazené podle času modifikace
    - `ls -l one.txt two.txt three.txt four.txt` - zobrazit si informace pouze k určitým souborům v adresáři
    - `ls -l *.txt` - zobrazí informace jen k souborům s příponou .txt (tzv. **wildcards** = zástupné znaky pro určení více souborů najednou, řadí abecedně)
    - `ls t*` - vytiskne všechny soubory začínající t
    - `ls [of]*.txt` - vytiskne soubory začínající na o nebo f s příponou .txt
    - `ls *[a-f].txt` - vytiskne všechny soubory s příponou .txt, co končí na písmena a až f
    - `ls -d D*` - vypíše soubory začínající na písmeno d, ale s přepínačem -d, který zabrání, aby byl vypsán i obsah těchto souborů (adresářů)
    - `ls x?.txt` - vypíše soubory, co mají 6 znaků, tedy x_.txt, kde _ je libovolný znak
    - `ls -a ~` nebo `ls -d ~/.*` - vypíše všechny skryté soubory bez obsahu v domovském adresáři (ne root)
    - `ls --color` zobrazí soubory barvami
    - `ls -F` spustitelné soubory označí * a adresáře /
- `cd` - "change directory", změní aktuální adresář na adresář daný argumentem. Např.:
    - `cd Documents` - do složky Documents
    - `cd ..` - do nadřazeného adresáře
    - `cd .` - změní na aktuální adresář
    - `cd` - bez argumentů změní na domovský adresář
- `pwd` - vypíše celou (absolutní) cestu k aktuálnímu adresáři
- `cat` - "concatenate" - vypíše obsah souboru, lze použít i ke spojování
    - `cat *` - vypíše obsah všech souborů
    - `cat [0-9][0-9][0-9].txt` - vypíše obsah všech souborů s názvy [000-999].txt
- `hexdump` - vypíše bajty souboru v hexadecimálním tvaru
    - `hexdump -C c/image.bmp` - přepínač -C slouží k vypsání ASCII znaků vedle hexdumpu
- `file` - slouží k zobrazení metadat k souboru jako typ souboru, u obrázku velikost atd.
- `man [argument/cmd]` - "manual", slouží k zobrazení dokumentace k určitému příkazu
    - `man cat` - zobrazí přepínače a způsob použití příkazu cat
    - `man man` - zobrazí úplný seznam sekcí
    - `/` - po otevření manuálu se dá pomocí lomítka vyhledávat
    - `q` - vypne manuál
    - `Up/Down Arrow` - navigace v manuálu
- `--help` - přepínač, funguje u většiny programů GNU, vypíšou nápovědu
- `--version` nebo `-v` nebo `-V` - přepínač, vypíše verzi a autorská práva programu
- `--verbose` nebo `--debug` (občas `-v` nebo `-d`) - spustí program v debug režimu, kdy program vypisuje podrobně, co dělá
- `--dry-run` (občas `-n`) - spustí program bez toho, aniž by provedl jakékoliv změny, např. vypsání souborů, které  by odstranil
- `--interactive` (občas `-i`) - program požádá o potvrzení při destruktivních akcích
- `--` - ukončí seznam přepínaču, hodí se pokud existuje v adresáři soubor začínající pomlčkou. Hodí se často používat pro jistotu.
    - `ls -l -h -- *.txt` - bezpečná konvence, příkaz, přepínače, ukončení přepínaču a wildcard nebo název souboru. Toto vypíše "lidsky" velikosti a názvy souborů s příponou .txt
- `\ ` - pokud použijeme soubor s mezerami v názvu, tak mezery jdou escapovat pomocí zpátečního lomítka `\`, tedy `ls tady\ je\ mezera.txt`
- `wget [URL]` - stáhnutí souboru do aktuálního adresáře

### Doplnění tabulátorem

- Např. při psaní argumentu `cd Doc`, tak pokud v daném adresáři existuje jediný možný soubor začínající na Doc, tedy Documents, tak se to doplní na `cd Documents`.
- V ostatních případech je třeba `Tab` stisknout vícekrát pro zobrazení možných příkazů nebo všech souborů, co odpovídá danému argumentu pokud nejjednoznačné.

## Midnight commander

- Spustí se příkazem `mc`
- čísla dole odkazují na fuknční klávesy

#### Ovládání

- `F1` - Help
- `F2` - Menu, otevře víc příkazů pro soubor
- **`F3`** - zobrazí obsah souboru
- **`F4`** - otevře textový editor (s syntax highlighting) pro editaci souboru
- `F5` - Copy, zkopíruje soubor do jiného adresáře
- `F6` - Move, změní adresář souboru
- `F7` - MKDIR, vytvoří nový soubor
- `F8` - Delete
- `F9` - Vybere nabídku MC a umožní interagovat pomocí klávesnice
- `F10` - Quit

- `Tab` - přepínání mezi panely vpravo a vlevo
- `Insert` - umožňuje vybrat soubory (pro provádění akcí jako smazání)
- `+` - umožní zadat masku pro výběr více souborů najednou
- `Ctrl + o` - skryje panely a dočasně přepne do shellu, ale nezamkne ho

## Ranger

- správce souborů inspirovaný Vimem

#### Ovládání - procházení

- `q` - ukončí ranger
- `j` - posun dolů
- `k` - posun nahoru
- `h` - přesun do nadřazeného adresáře
- `l` - otevření souboru nebo přesun do adresáře
- `gg` - přejít na začátek seznamu
- `G` - přejít na konec seznamu
- `gh` - přesune aktuální adresář do domovského adresáře, nebo-li `cd ~`
- `gm` - přesune do souboru /media, nebo-li `cd /media`
- `gr` - přesune aktuální adresář do `/`, nebo-li root, tedy `cd /`

#### Ovládání - práce se soubory

- `zh` - zobrazí skryté soubory
- `cw` - přejmenuje aktuální soubor
- `Spacebar` - vybere aktuální soubor
- `yy` - vyjmutí (kopírování) souboru nebo více souborů
- `dd` - označení vybraných souborů pro operaci vystřižení
- `pp` - vložit vyjmutý nebo vystřižený soubor/soubory
- `dD` - odstranění vybraných souborů

## Úpravy souborů

- mezi nejpoužívanější patří: Emacs, Joe, mcedit, nano a Vim
- otevírají se pomocí příkazu + argumentu, např. `mcedit hello.py`
- jedná se o TUI editory, takže jsou vesměs dostupné v terminálu bez grafického rozhraní a tedy lze s nimi pracovat vzdáleně.

# 3. Hodina

## Skript

- pomocí shellu, což je jazyk příkazové řádky
- Ukládají se do souborů s příponou `.sh`, napriklad `skript.sh`
- Následně se spouští pomocí příkazu `sh`, celý příkaz bude `sh skript.sh`. Který spustí interpretr jazyka shell.

- Je potřeba specifikovat interpretr a ten se specifikuje pomocí **shebangu/hasbangu**, nebo-li `#!`
    - Pro shellové skripty se používá `#!/bin/sh`
    - Pro pythonové skripty se používá `#!/usr/bin/env python3`
    - Další: `#!/bin/bash` - bash je rozšíření shellu

- Pro označení souboru jako spustitelný je třeba provést `chmod +x skript.sh`, pak lze soubor spustit pomocí: `./skript.sh`
    - Musí se použít `./`, protože se shell dívá první do `usr/bin` nebo-li `$PATH`, kde je většina spustitelných souborů a do aktuálního adresáře se nepodívá, proto je potřeba specifikovat relativní cestu

- Při spuštění skriptů má proces vlastní pracovní adresář, který nijak neovlivní ostatní procesy, tedy shell uživatele pokud budu chtít ve skriptu měnit adresu.
    - Z toho důvod je `cd` builtin funkce shellu a nejedná se o normální skript/program

