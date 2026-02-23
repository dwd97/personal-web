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
