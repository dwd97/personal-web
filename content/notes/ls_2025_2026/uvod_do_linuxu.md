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
    - `sort -n -k 2 -r` - tohle řadí podle čísel ve druhém sloupci v opačném pořadí
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

- `git` je verzovací nástroj

- výběr editoru. Přidáním `export EDITOR=mcedit` nebo jiného editoru (např. vim, nano ...) na konec souboru `~/.bashrc`
    - pak lze spustit (po restartu terminálu) `$EDITOR ~/.bashrc` a tím se spustí daný editor se souborem `~/.bashrc`
    - takhle se nastavila proměnná prostředí EDITOR (používají ji ostatní programy k funkci)
    - TUI editory jsou superior od GUI, protože např. při psaní commit zprávy se správně zpracuje instance a nepošle se prázdný commit

## Příkaz git

- píše se vždy ve formě `git podprikaz --help`. Takže `git`, pak podpříkaz (např. `config`, `commit`) a pak parametry a nakonec data
    - podpříkazů je hodně (cca 100), ale většinou stačí jen asi méně než 20

### Nastavení Gitu

- každý commit má autora, proto je potřeba nastavit prostředí
    - přepínač `--global` znamená, že nastavení platí pro všechny gitové projekty

```sh
git config --global user.name "My Real Name"
git config --global user.email "My E-mail"
```
- lze použít i bez přepínače `--global`, nastaví identifikaci jen lokálně, např. pro odlišení identit

### Git init

- `git init` - vytvoří prázdný lokální repozitář gitu, lze později propojit se vzdáleným

### Git clone

- klon projektu je tzv. **Pracovní kopie** nebo-li **Working Copy**, je to 1:1 kopie a lze projekt na GitLabu obnovit z této kopie
    - nekopíruje se wiki, issues, ty nejsou verzované gitem

- `git clone` - příkaz na zkopírování a vytvoření **working copy**

- nastavení přihlášení se dá obnovit pomocí:
```sh
export GIT_ASKPASS=""
export SSH_ASKPASS=""
git clone ...
```

### Git status

- `git status` - vypíše aktuální změny lokálně oproti projektu v GitLabu
    - **Changes not staged for commit** = vidím změny v souborech
    - **Untracked files** = nové soubory (nevypisují se všechny soubory adresářů)

### Git diff

- `git diff` - zobrazí změny v souborech
    - `@@ -3,3 +3,5 @@` - změny v řádcích
        - `-3,3` - stará verze je od řádku 3 do řádku 3
        - `+3,5` - nová verze je od řádku 3 do řádku 5
    - `+ nějaký text` - přidaná řádka
    - `- nějaký text` - odebraná řádka

### Git add

- přidání do stagingu (tedy zahrnutí do další verze commitu)
- `git add`
    - `git add .` - přidá všechny soubory
    - `git add adresar/soubor.txt` - přidání určitého souboru

- v `git status` se to projeví tím, že se dané změny přesunou do sekce **Changes to be commited**

### Git commit

- `git commit` - otevře příslušný editor, kam se napíše commit zpráva a tím se vyčistí změny a může se pracovat na další změně
    - `git commit -m "příslušný commit název"` - rovnou nastaví zprávu commitu, `-m` znamená message

- `git status` v tomhle případě napíše **Your branch is ahead of 'origin/master' by 1 commit**, proto je potřeba změny poslat na server

### Git push

- `git push` - nahraje všechny změny na server

### Git log

- `git log` - vypíše všechny commity
- `git log --oneline` - vypíše všechny commity kompaktně
- `git log --max-count=20 --oneline` - jen 20 posledních kompaktně
    - lze použít `git ls` po nakonfigurování `git config --global alias.ls 'log --max-count=20 --oneline'`
    - další `ll = log --format='tformat:%C(yellow)%h%Creset %an (%cr) %C(yellow)%s%Creset' --max-count=20 --first-parent`

### Git pull

- `git pull` - stáhne změny ze serveru, dokáže začlenit změny z GitLabu a na lokálním repozitáři

- práci začít `git pull` a skončit `git push`

### Git blame

- pro získání overview, kdo modifikoval, kterou řádku kódů a v jakém commitu

## Použití jiných interpretrů

- do shebangů lze napsat jiné soubory a ty se spustí, např. `#!/bin/cat` nebo i vlastní soubor v pythonu, ale absolutní cestou
    - lze vlastně vypustit příponu souboru

- např. my-cat má obsah jen `#!/bin/cat` a my-echo jen `#!/bin/echo`, pak `./my-cat my-echo` spustí vlastně složení: `/bin/cat my-cat my-echo`

- pro úkoly (testy) lze spustit: `./bin/run_tests.sh` či jeho podmnožina `./bin/run_tests.sh 04` ...

## Přesměrování vstupu a výstupu

### Standardní chybový výstup

- v případě, že například soubor neexistuje, tak místo toho, aby se chybová hláška dostala do `stdin` programu, tak se pošle pomocí `stderr` (standard error [output])
    - nepřesměrovává se spolu s `stdout` pomocí příkazu `>`

## Deskriptory souborů

- je to číslo otevřeného souboru, proto se musí soubory otevírat před tím, než se do nich zapisuje nebo se něco vypisuje
- je to identifikátor, který si operační systém udržuje v tabulce otevřených souborů a volají se pomocí nich syscall
- 0,1,2 je popořadě systdin, systdou, systderr v deskriptorech a syscallech

- lze přesměrovat stderr pomocí 2
    - např. neexistuje adresář nonexistent.txt `cat one.txt nonexistent.txt two.txt > merged.txt 2>err.txt`

## Obecné přesměrování

- `>&2` - přesměruje stdout na stderr
    - např. pokud nechceme vypsat všechny detaily při použití `wget`

- `>output.txt 2>&1` nebo `&>output.txt` - přesměrování stdout a stderr do jednoho souboru, `2>&1` říká tam, kam jde stdout
    - `>output.txt 2>output.txt` nefunguje, protože shell musí otevřít soubor dvakrát

### Speciální soubory

- `/dev/null`
    - `cat one.txt > /dev/null` - `/dev/null` je speciální typ souboru, kam lze něco přesměrovat, pokud nás to nezajímá

- `/dev/full`
    - simuluje disk, co je plný a vždy vrátí chybu při zápisu

- `/dev/stdin`
    - reprezentuje file descriptor 0, tedy stdin, v programech se dá zaměňovat často za `-`
    - `./funkce.py /dev/stdin one.txt < two.txt` - přesměruje složený one.txt a two.txt na stdin pro python program, two.txt je složený přes OS

- `/dev/stdout`
    - pokud program chce zapisovat do souboru, ale ty mu hodíš `/dev/stdout`, což přesměruje na výstup

### Návratová hodnota programu (exit code)

- `0` signalizuje úspěch, jakékoliv jiné číslo neúspěch

## Nastavení selhávání

- začít příkaz pomocí těchto zkratek:
    - `set -o pipefail` - chyba kterékoliv části pipeline způsobí ukončení celé pipeline
    - `set -e` - ukončení programu v případě chyby
    - `set -u` - ukončení skriptu v případě neinicializované proměnné
    - společně: `set -ueo pipefail` - dobrá praxe začínat každý skript tímhl příkazem

- `set -ueo pipefail cat /dev/urandom | head -n 1 | hexdump` - cat ukončí s chybou kvůli `set -o pipefail`, protože cat vypisuje, ale head ukončí příkaz jakmile vypíše jeho počet řádků a další nečtě, ale cat dál vypisuje, takže to, že to nikdo nečte se naštve a ukončí program, pomocí `||` tomu lze zamezit před head

## Přizpůsobení shellu

### Aliasy

- je to definování příkazů bez vytváření skriptů
- `alias l='ls -l -h'` - definuje nový příkaz l
    - ale jen na dobu běhu shellu, pak se přepíšou

- dají se nadefinovat v konfiguračním souboru shellu, např. `~/.bashrc`

### Nastavení promptu

- `PS1=''` - nastaví prompt na prázdný
- `PS1='Prikazy:'`
- `PS1='\u: \w '` - jméno uživatele a aktuální adresář

- lze měnit i barva, přidat datum pomocí `$( date )`

# 5. Hodina

- asymetrické šifrování, viz. Úvod do počítačových sítí ze ZS
- doporučuje se používat na více počítačích více párů veřejných/soukromých klíčů
- v Linuxu jsou veřejné/soukromé klíče uloženy v textových souborech

## SSH (Secure shell)

- port 22, ale může ho správce serveru přepnout

```sh
ssh LOGIN@NAZEV_VZDALENEHO_STROJE
```

- `ssh user@localhost` - připojení na loopback adresu (127.0.0.1), pokud je port otevřen k připojení

- `ssh` - většinou klient ssh, protože se připojuje k serveru SSH
    - `ssh -v` - vytiskne log a detaily připojení k SSH

- po připojení se zobrazí otisk serveru, např. ECDSA, RSA, ED25519
    - pokud se otisk změnil - může to být man-in-the-middle útok

- nastavení $PS1
    - `export PS1='\u@\h \w\$'` - zobrazí uživatelské jméno, jméno hostitele a pracovní adresář bez barvy

- `uname -a` a `hostname -f` - pro oveření, že jsem na jiném počítači
- `free -h` - zobrazí kolik paměti je volné
- `uptime` - zobrazí údaje o době běhu

### spuštění jednoho příkazu přes ssh

- `ssh login@host uname -r > local.txt` - spustí se na vzdáleném stroji a zapíše do lokálního souboru
- `ssh login@host "uname -r > remote.txt"` - celý příkaz se spustí na vzdáleném stroji

### přihlášení bez hesla

- vytvořit pár klíčů a zkopírovat veřejný na vzdálený počítač
    - generuje vždy na lokálním počítači (nebo na vzdálenem pokud se připojuje dál, např. github)

- `ssh-keygen -t ed25519 -C "name@email.com"`
    - uloží se do adresáře `~/.ssh` klíče `id_ed25519` soukromý a `id_ed25519.pub` veřejný
    - pro správu více klíčů se používá přepínař `-i` nebo-li `ssh-copy-id -i ~/.ssh/muj_specificky_klic uzivatel@vzdaleny_hostitel` - tohle nahraje klíč na vzdálený počítač, uloží se do `~/.ssh/authorized_keys`

### kopírování souborů

- na vzdálený počítač: `cat local.txt | ssh login@host "cat > remote.txt"`
- na lokální počítač: `ssh login@host cat remote.txt > local.txt`

- `scp` a `rsync` - pro kopírování více souborů přes SSH

- pomocí `mc` - v panelu nahoře right a tam lze zadat `login@host` pod shell link

### konfigurace SSH

- v `~/.ssh/config`, syntaxe je v `man 5 shh_config`

```sh
Host *
    IdentityFile ~/.ssh/id_ed25519

Host intro
    Hostname server.cz
    User LOGIN
```

- pak lze napsat jen `ssh intro` a připojí se pomocí klíče k SSH serveru

### Git

- lze se připojit k SSH git serveru

- pomocí ssh adresy `git clone git@server:adresar/repozitar.git`

## Síťové nástroje

- `ip` a `ip addr` - příkaz na správu sítě

### Network Manager

- CLI rozhraní se spustí pomocí `nmcli`

### Změna ip konfigurace

- pro nastavení ip adresy mimo DHCP, např. propojí počítačů při přenášení souborů

```sh
sudo nmcli connection add con-name wired-static-temporary ifname enp0s31f6 type ethernet ip 192.168.177.201/24
```
- nastaví spojení přes ethernet, kde `con-name` znamená název profilu, a musí se uvést adresa s maskou

- `sudo nmcli connection up/down/delete nazev-profilu` - na správu profilu a připojení, sudo je administrátor

```sh
nmcli connection add con-name wifi-temporary ifname wlp58s0 ssid "nazev site" wifi-sec.key-mgmt "wpa-psk" wifi-sec.psk "heslo" ip4 192.168.177.203/24
```
- tohle vytvoří profil pro wifi, kde `ssid` je název sítě, `wifi-sec.key-mgmt` je zabezpečení a `wifi-sec.psk` je heslo
    - ale u WiFi většinou je DHCP, takže tam to není potřeba

### Ping

- pro nalezení problému, pokud ztratím spojení se serverem.
- `ping 192.168.177.201/24` jako na ethernet profil nahoře

### Traceroute

- `traceroute 1.1.1.1` - cesta k serveru přes routery, posílají se ICMP package

### nmap

- často zakázáno správci sítě, protože se většinou používá jako příprava na kybernetický útok, kdy se skenují otevřené porty

- `nmap localhost` - najde všechny otevřené porty na loopback ip adrese
    - `sudo nmap -A localhost` - více informací
    - `sudo -p1-65535 localhost` - větší rozsah portů (všechny)

### nc (netcat)

- `nc --listen 8888 --keep-open --sh-exec 'cat transfer_file.txt'` - otevře port 8888 a lze se na něj připojit pomocí `nc localhost 8888`

### tmux

- pro ovládání více relací terminálu v jednom terminálu pomocí serveru běžícího ssh
- `tmux new -s nazev` nebo `tmux` spustí novou relaci
- `tmux ls` - vypíše všechny aktivní relace
- `tmux attach -t nazev` - připojení k relaci
- `tmux kill-session -t nazev` - ukončení relace

- `Ctrl + B +`
    - `d` - odpojení od relace
    - `c` - vytvoření okna
    - `n` - další okno
    - `p` - předchozí okno
    - `f` - najít okno
    - `,` - název okna
    - `&` - zavřít okno
    - `%` - vertikální rozdělení
    - `"` - horizontální rozdělení
    - `o` - prohození panelů
    - `q` - zobrazit čísla panelů
    - `x` - zavřít panel
    - šípky (přepnutí panelů)

- dá se upravit v `.tmux.conf`

### pdfpc

- na prezentování

