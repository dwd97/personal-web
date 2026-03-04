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
- `etc/pass` - adresář sloužící ke konfiguraci Linuxu

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
- `tac` - rozšíření cat a vypíše řádky souboru obráceně
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
- `echo` - vypsání řetězce na stdout (standard output)
- `cut` - na extrahování částí z vstupu nebo souboru (tiskne části řádků)
    - `cut -d : -f 1` - spustí vstup pomocí stdin, `-d :` specifikuje delimiter `:`, na které se má text rozdělit a `-f 1` vezme první část z rozdělených úseků
    - `-b` - ignore leading blanks
- I/O redirection - pomocí operátoru `>`, tedy `> soubor.txt` se text nevypíše do příkazové řádky, ale do souboru `soubor.txt`. Např.: `echo Hello World > hello.txt`
    - operátor `>>` appenduje text na konec specifikovaného souboru místo přepisování/tvorby nového
    - operátor `<` umožňuje číst ze souboru na místo zapisování
- `uniq` - vytiskne jen unikátní
    - s argumentem `-c` - vypíše i počet
- `sort` - seřadí
    - `-n` - podle čísel, např. po `uniq -c` zavolám a předám stdout pomocí `|`
    - `-t,` - field delimiter
    - `-k2,2` - select fields, so selected from field 2 to field 2
- `|` - roura, předává stdout na stdin, např.: `cat text/*.txt | sort` - seřadí obsahu textových souborů podle abecedy
- `head -n 3` - vypíše jen první 3 řádky
- `tail` - opak head, vypíše poslední řádky
    - `tail -n +2` - odstraní první řádek
    - `tail -n 2 /etc/passwd /etc/groups` - vypíše poslední 2 řádky z obou souborů
- `paste` - vypíše do stdout předané řádky textu z stdin
    - `-d argument` - delimiter, přidá mezi předanými vstupy
    - `-s` - serial, vypíše najednou, ne paralelně
- `cp` - kopíruje soubory a adresáře
- `bc` - kalkulátor
- `grep` - používá se k hledání určitých informací v textu
- `sed [path]` - na hledání a úpravů dat rychle
- `wc` - vypíše počet řádků, slov a znaků v datech
    - `wc -l` - vypíše počet řádků souboru v argumentu
- `join` - spojí setříděné soubory podle společných sloupečků
- `rev` - otočí pořadí znaků na každém řádku
- `rm` - odstraní soubory nebo adresáře
- `rmdir` - odstraňuje adresáře
- `scp` - bezpečně kopíruje soubory napříč stroji
- `tar` - archivační nástroj
- `test` - porovnává hodnoty a rozhoduje o typech souborů
- `tr` - přeloží (nahradí) nebo odstraní znaky (písmena)
- `who` - vypíše 2 přihlášené uživatele, druhý se právě přihlásil do GUI a první do CLI
- `whoami` - vypíše, kdo je přihlášen
- `export` - vypíše dočasné proměnné


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

### Argumenty příkazové řádky (Python)

- pomocí `import sys` a argumenty jsou ve formátu `'-d'`, `'argument2'` ...

```Python
#!/usr/bin/env python3

import sys

def main():
    for arg in sys.argv:
        print(arg)

if __name__ == '__main__':
    main()
```

### Stdout, Stdin

- jsou knihovny pro čtení a vypisování. V pythonu je k nim možné přistoupit přes `import sys.stdin` a nebo `import sys.stdout`
    - pak se používají pomocí `sys.stdin.readline()` a nebo pro čtení více řádek: `for line in sys.stdin`
    - nemusí se soubor uzavírat, protože se uzavře sám po dokončení přístupu
- např.: při použití `cut -d : -f 1` se spustí stdin a ukončí se `Ctrl+D`, což je něco jiného než `Ctrl+C`


#### I/O redirection

- low-level přesměrování výstupu pomocí argumentu `> soubor.txt`, ale tohle se nepředává v sys.argv, nemusí o tom program vědět
- soubor se přepíše a nelze obnovit, proto se doporučuje používat `Tab` pro zjištění zda daný souboru už existuje
- příklad: `cat 1.txt 2.txt > 12.txt`
- příklad č.2: `tac 1.txt 2.txt > 2.txt` (při přesměrování na úrovni shellu se souboru 2.txt smaže, takže se uloží jen opačné pořadí souboru 1.txt do 2.txt)

### Filtry

- jsou vlastně příkazy na úpravu standard inputu a předání, např.: `sort`, `cut`, `cat`, `tac`, `head`, `tail`, `uniq`, `wc`, `grep`, `sed`, `nl`
- `cut -d : -f 1 </etc/passwd` a `cut -d : -f 1 /etc/passwd` vypíšou to samé, ale první je předán text cutu pomocí shellu a musí to otevírat shell a kdyžtak nahlásit chybu, ten druhý případ se předá souboru příkazu cut a ten zpracuje soubor a řeší případné chyby

### Roury (pipes)

- skládání proudových dat
- místo toho, abych zapisoval do dočasného souboru a omlyme přepsal nějaký soubor, tak použiju `|`, což vlastně znamená, že `stdout` pošlu na `stdin`
- `cat logs/*.csv | cut -d , -f 5 | sort | uniq -c | sort -n -r | head -n 3` - seřadí csv soubory v určitém formátu podle počtu výskytů a vypíše první 3
- `cat logs/*.csv | cut -d , -f 5 | paste -s -d + | bc` - spočítá byty logů pomocí bc kalkulátoru a paste, který zformátuje do stdout předaný stdin

- stdin je dostupný i uvnitř skriptů pro tvoření pipes
```sh column.sh
#!/bin/sh

cut -d : -f 1 | sort
```
- tento skript se propojí s pipeline: `cat /etc/passwd | ./column.sh | tail -n 5`

další příklady:
- `cut -d : -f 1,3 /etc/group` - vypíše první a třetí sloupec souboru /etc/group
- `<skore.txt tr -s ' ' | cut -d ' ' -f 2- | tr ' ' '+' | bc | paste score.txt - | tr '\t' ' '` - pomocí rour, takže -s (squeeze zredukuje mezeru na jednu), 2- je od 2. až na konec, pak nahradí mezeru za +, pak spočítá, paste pak spojí 2 vstupy po sloupcích, kde `-` znamená vzít z stdin, `tr '\t' ' '` nahradí tabulátor mezerou
- `grep -F vendor_id /proc/cpuinfo | cut -d : -f 2 | cut -b 2- | sort | uniq` - vytiskne výrobce CPU
- `sort -t, -k2,2n file` - seřadí podle čísel v druhém sloupic odděleném čárkou

# 4. Hodina