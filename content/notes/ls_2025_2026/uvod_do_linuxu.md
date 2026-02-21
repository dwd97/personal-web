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
- **GUI (grafické rozhraní)** - většinou není potřeba
- **CLI (příkazová řádka)** - přesné a zautomatizovatelné
- **všechno je soubor** - proto stačí jen jedno API pro všechno
- **víceuživatelský** - více uživatelů může používat systém naráz, vzdáleně i lokálně
- **balíčky** - pro instalace aplikací, ty se pak lépe udržují a automaticky aktualizují
- **lze řídit všechno** - pomocí CLI

## Pracovní prostředí

- ovlivňují vzhled a způsob ovládání Linuxu.

mezi nejčastěji používáné patří:
- **GNOME, Plasma (KDE), Xfce, LXDE** - podobné prostředí ostatním OS
- **Openbox, i3** - ovládájí se většinou klávesnicí, steep learning curve

## Odbočka: Git (správa verzí kódů)

- sledování změn a zapamatování předchozích stavů.

# 2. hodina

## Příkazová řádka

- Pomocí CLI lze systém ovládat přesně. Uložené soubory s příkazy se nazývají skripty. Namísto programovacích jazyků, kde se spouští funkce, tak se v CLI se spouští programy.

## Názvy souborů a cesty

- název souboru + adresář = cesta
- Oddělovačem v adresáři je `/` dopředné lomítko.
- Na Linuxu se v cestě rozlišují velká a malá písmena.

### Kořenový adresář

- root/kořenový adresář je jediné dopředné lomítko `/`

### Relativní a absolutní cesta

- **relativní** - Nezačíná lomítkem. `adresar/soubor.txt` Potřeba zkombinovat s aktuální absolutní cestou. 
- **absolutní** - Začíná lomítkem. `/adresar/soubor.txt` Konkrétní soubor nehledě na tom, ve kterém adresáři právě jste.

### Speciální adresáře

- `..` - nadřazený (rodičovský) adresář. Příklad: lze zkombinovat `/home/documents/films/../dopis.odt` zkráceně jako `/home/documents/dopis.odt`
- `.` - aktuální adresář. Takže (relativní cesta) lze napsat `./bin/knihy/harry_poter.epub` nebo `bin/knihy/harry_poter.epub`