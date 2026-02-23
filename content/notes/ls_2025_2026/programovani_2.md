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
```C#
int i = 5;
int x = i++; // x = 5, i = 6
```

- **preinkrement** - `++i;`
    - nejdříve se i zvýší o 1, pak se načte hodnota i.
```C#
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
- relační operátory - `==` `!=` `<` `>` `<=` `>=`
- logické spojky
    - zkrácené vyhodnocování = vyhodnotí jen potřebné množství výrazů
        - `&&` and
        - `||` or
    - úplné vyhodnocování = vyhodnotí všechny výrazy, může se hodit pokud je potřeba zavolat nějakou funkci a pracovat s jejím výsledkem
        - `&` and
        - `|` or
        - `!` not
        - `^` xor

### Switch

- je povinnost ukončit každou sekci příkazem `break;`
    - kromě výjimky např. u case 2 a case 3 v příkladu, kdy pro dvě podmínky se vykoná stejný příkaz
- V C/C++/Java/PHP - může se propadat mezi sekcemi, proto se nedoporučuje používat, V C# opraveno

```C#
int i = 5;
switch(i)
{
    case 1:
        i++;
        break;
    case 2:
    case 3:
        i--;
        break;
    case 4:
    case 5:
        break;
    default:
        i = 5; break;
}
```

### Cykly

#### For cyklus

```C#
int N = 12
int[] a = new int[N];

// provede se 12x, (inicializace; podmínka; příkaz iterace)
for (int i=0; i<N; i++){
    a[i] = i
}
```

#### While a Do-While

`while (podmínka) { příkazy }`

`do { příkazy } while (podmínka)`

#### Ukončení cyklu

- `continue` - pokračování do další iterace
- `break` - ukončení cyklu

### Funkce

- funkce patří třídě nebo objektu
- musí vracet výsledek nebo musí vrátit `void` a k ukončení funkce a vrácení hodnoty se napíše `return;`
- při deklaraci píšeme `()` i pro funkce bez parametrů
- Od C#7 lze definovat lokální funkci ve funkci

#### Předávání parametrů
1. hodnotou (výchozí pro C#)

```C#
using System;
class App {
    static public void Main() {
        string isVerified = false;
        isVerified = ChangeValue(isVerified);
        Console.WriteLine($"This user is verified: {isVerified ? "yes" : "no"}")
    }

    static bool ChangeValue(bool currentValue) {
        currentValue = !currentValue;
        return currentValue;
    }
}
```

2. odkazem - pomocí specifikátoru `ref`

```C#
using System;
class App {
    static public void Main(string[] args) {
        string name = "Peter";
        SetValue(ref name);
        Console.WriteLine(name); // it will print "John"
    }

    static void SetValue(ref string name) {
        name = "John";
    }
}
```

3. výstupním parametrem - pomocí specifikátoru `out`, ten nemusí mít vstupní hodnotu

```C#
using System;
class App {
    static public void Main(string[] args) {
        int num;
        Sum(out num);
        Console.WriteLine($"The sum is {num}")
    }

    public static void Sum(out int num) {
        num = 80;
        num += num;
    }
}
```

#### Výchozí metoda Main()

- vytvaří se pro ni samostatní třída
- určuje začátek a konce výpočtu
- většinou je jediná, jinak se musí určit, která se má spustit

### Vstup a výstup

- `Console.Read();` - vrací int, kód znaku ze vstupu
- `Console.ReadLine();` - vrací string, jeden řádek ze vstupu
- `Console.Write(výraz);` - vypíše kód znaku z parametru
- `Console.WriteLine(výraz);` - vypíše hodnotu výrazu, stringu

**formátovaný**
- `Console.WriteLine("x0 = {0}, x1 = {1}", x0, x1)`
- `Console.WriteLine($"x0 = {x0}, x1 = {x1}")`

## Soubory

- `.cs` - zdrojový kód, stačí překlad pomocí csc.exe

**Další z Visual Studia:**
- `.csproj` - projekt
- `.sln` - solution

**Top-level statementy** - zaobalí se do tříd
**global usingy** - co se používá, např. `using System;`, `using System.Collections.Generic;`
**file-scoped namespaces** - `namespace Aplikace {}`

## Dynamicky alokované proměnné

- `new + konstruktor objektu` - vrací vytvořenou instanci (ukazatel na ni)
- string, pole, třídy = referenční typy, potřeba vytvořit pomocí new
- `null` (jako None v pythonu)
- pro zkopírování existují připravené metody `CopyTo`, namísto `pole1 = pole2`, protože takhle se kopíruje reference

### Pole

- metody připravené:
    - `CopyTo`
    - `Sort`
    - `Reverse`
    - `BinarySearch`
    - `Array.Reverse(aaa);`

- vícerozměrné pole
    - obdélníkové `[,]`
    - nepravidelné `[][]`

```C# pole polí
int[][] pole = new int[3][]
pole[0] = new int[6]
pole[1] = new int[8]
pole[2] = new int[2]
```

### Znakový řetězec (string)

- deklarace `string str;`, string je referenční typ, alias třídy System.String
- indexování znaků od 0
- délka = `str.Length()`
- pro změnu stringu je třída `StirngBuilder`
- všechny objekty mají metodu `ToString()`

## Struktura (struct)

- zjednodušená třída, hodnotový typ (takže se nemusí alokovat)
- může mít i konstruktor
- omezení oproti třídám - nemůže dědit

```C#
struct Souradnice
{
    public int x,y;
    public Souradnice(int x, int y)
    {
        this.x = x;
        this.y = y;
    }
}
```

## Tuple - n-tice

- `System.Tuple`
    - indexují se vlastnostmi od jedničky `t.Item1` `t.Item2`

```C#
System.Tuple t = new System.Tuple<string, int, int, int, int, int>("Praha", 520, 55, 666, 856, 954);
```

nebo

```C#
System.Tuple t = System.Tuple.Create("Praha", 4561, 846, 52, 111);
```

- `System.ValueTuple`
    - novější doporučená verze

```C#
(double, int) t = (4.5, 3);

t.Item1
t.Item2
```

```C#
(double a, int b) t = (4.5,30);

t.a
t.b
```

# 2. Hodina