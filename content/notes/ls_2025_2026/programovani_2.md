# 1. Hodina

## Přehled

**Java**
- z roku 1995
- běží nad Java Virtual Machine (JVM)
- kompiluje se do bytecode, ten je přenositelný mezi platformami

**C#**
- z roku 2002, součást platformy .NET (skupiny jazyků)
- kompiluje se do intermediate language (IL), který běží nad Common Language Runtime (CLR)
- ECMA a ISO/IEC (standardizovaný)

- Visual Studio je plnohodnotné IDE na .NET projektry oproti VS Code

## Rozdíly mezi Python a C#

- tzv. whitespaces nic neznamenají (mezery, tabulátory, konce řádek), ale doporučuje se indentace
- před použitím proměnně se musí deklarovat její typ
- proměnna představuje místo v paměti (s deklarovaným typem, který se musí dodržovat (static-typing))

## Konvence

- používá se angličtina, i když se dají použít mezinárodní abecedy

**malá_písmena**
- klíčová slova - `class`, `return`, `if`, `struct`, `namespace`, ...

**VELKÁ_PÍSMENA**
- konstanty - `const int MAX_SIZE = 100;`

**PascalCase**
- JménaProstorů - `namespace Project1`
- Třídy - `class Person {}`
- Metody - `void GetTotal(){}`
- VeřejnéČleny (cokoliv označené public) - `public int Add(int a, int b) { return a+b; }`
- Vlastnosti - `public string FirstName {get; set;}`

**camelCase**
- lokální proměnné
- parametry metod
- soukromé metody

## Struktura programu

- celý program se skládá ze tříd. Je spuštěn statickou metodou main - `static void main(){}`
    - volný kód = top-level statements (ten se ale obalí v IDE zase do tříd)
- položky deklarované ve třídě: metody (členské funkce), datové složky (členské proměnné)
- lokálně deklarované proměnné jsou dostupné jen v blocích, kde byly deklarované

## Konstanty

- syntaxe stejná, co inicializované proměnné
- `const int NUM = 15;`

## Typy

### Hodnotové
- uložené rovnou v paměti

- **celé číslo** `int` = `System.Int32`, tedy 32 bitů
    - `byte`
    - `sbyte`
    - `short`
    - `ushort`
    - `uint`
    - `long`
    - `ulong`
- **desetinné číslo** `double` = `System.Double`, tedy 64 bitů
    - `float`
    - `decimal`
- **logická hodnota** `bool`
- **znak** `char` = `System.Char`, 16 bitů Unicode
- **výčtový** `enum`
- **struktura** `struct`

### Referenční
- hodnota odkazuje jen na místo v paměti

- **pole** `[]` = `System.Array`
- **znakový řetězec** `string` = `System.String`
- **objekt třídy** `class`
    - `ArrayList`, `StringBuilder`, `List<>`

## Aritmetické výrazy

- `+` `-` `*` `/` `%`
    - `/` může představovat desetinné i celočíselné dělení, záleží na zvoleném typu
- `checked` a `unchecked` - určuje, zda se má kontrolovat aritmetické přetečení, zpravidla unchecked

## Další syntaxe

- středník `;` - ukončuje příkaz
- čárka `,` - odděluje index v poli, parametry v funkci, deklarace více proměnných najednou

### Komentáře

- jednořádkové `// komentář`
- víceřádkové `/* komentář */`
- dokumentace (pro automatickou generaci) - `///`

### Blok

- `{}` - spojení více příkazů dohromady

### Dosazovací příkaz

- proměnná = výraz

### Modifikace hodnoty

- **postinkrement** - `i++`
    - nejdříve načte hodnotu i, pak inkrementuje o 1
```C# postinkrement
int i = 5;
int x = i++; // x = 5, i = 6
```

- **preinkrement** - `++i;`
    - nejdříve se i zvýší o 1, pak se načte hodnota i.
```C# preinkrement
int i = 5;
int x = ++i; // i = 6, x = 6
```

- **postdekrement** - `i--;`, **predekrement** - `--i;`
- `i += 10;`
- `i -= 10;`
- `i *= 10;`
- `i /= 10;`
- `i %= 10;`

### Podmíněný příkaz

- `if (a == 5) b = 17;`