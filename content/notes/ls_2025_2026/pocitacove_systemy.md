# Intro to C/C++

- static type system (všechny proměnné musí mít předem určený typ)
- compiler to machine instructions -> dobré pro systémové programování, HPC (výpočetní cluster)
- case sensitive

```C++ Hello World
#include <stdio.h>

int main(int argc, char **argv) { // ** means pointer to a pointer
    printf("Hello, World!\n")
    return 0;
}
```

## Application Entry-Point

- pomocí funkce `main`, return main je exit code, pokud není explicitně return, tak implicitně je `return 0` exit code
- jednoduché: `int main() { }`
- složitější: `int main(int argc, char *argv[]) { }`

## Literály

### Integer

- decimální = `56` `-2` 
- binární = `0b00001101`
- hexadecimální = `0x0c1f`
- octal = `023` `07`

### Floating

`15.8` `-9.85e-4`

### Boolean

`true` `false`

### String

`"a string"`

### Char

`'x'`

### Escape sequence

`\n` = line feed
`\r` = carriage return
`\t` = tab
`\\` = back slash
`\'`
`\"`

## Data typy

### Integer

**základní**
- `char` (1 Byte)
- `int` (velikost se liší podle architektury, arduino má například 16-bit)
**modifikátory** (píšou se před `int`)
- `short` (16-bit)
- `long` (32-bit)
- `long long` (64-bit)
- `signed` (implicitní)
- `unsigned` (explicitní)
**další**
- `size_t` (hlavně pro indexování)

### Floating

- `float` (32-bit)
- `double` (64-bit)

### Ostatní typy

- `void` (prázdnota)
- `bool`

### Implicitní převody
- není třeba castit, funguje automaticky
- `int` $ \rightarrow $ `long` (nic se neztratí)
- `double` $ \rightarrow $ `float` (ztratí se přesnost, většinou warning v IDE)

### Explicitní převody (casting)
- odstraní warningy, ale nedoporučuje se

```C++
double d = 5.6;
float f = (float)d; 
```

## Proměnné

- musí být deklarované typem
- typ `auto` automaticky identifikuje typ, ale musí se deklaraci přiřadit hodnota

### Dosažitelnost

**lokální**
- existují v rámci jednoho bloku `{}`, po ukončení se smažou z paměti

**globální**
- existují v rámci více bloků

**statické lokální proměnné**
- jsou dostupné v rámci jednoho bloku/funkce, ale přežijí volání funkce, tedy nesmažou se z paměti

## Příkazy (Statements)

- (Compound) Složený příkaz (blok) = `{}`
- (Expression) Jednoduchý příkaz = `příkaz ;`
- (if) podmíněný příkaz = `if (podmínka) něco else něco_jiného`
- (return) vrácení = `return hodnotu ;`

### Switch

- nepoužívat, protože se může přelít z jedné case do jiné

```C++
switch(expr)
{
    case 0:
        // něco
        break;
    case 1:
    case 2:
    case 3:
        // něco jiného
        break;
    default:
        // něco
        break;
}
```

### Podmíněné příkazy (cykly)

**while cyklus**
- `while (expression) something`

**while cyklus s alespoň jedním průchodem**
- `do something while (expression);`

**for cyklus**
- `for (deklarace; test; end) { }`
- `for (int i=0;i<N;i++) { }`

**control flow cyklu**
- `break;` = zruší cyklus
- `continue;` = přeskočí konkrétní iteraci cyklu a pokračuje další

**bezpodmíněčné control flow (přesměrování)**

```
if(something) goto end_of_loop; // this is inside a loop or somewhere

end_of_loop: // something continues here
```

- často se nepoužívá, spíše nepěkné

## Operátory

**Aritmetické**
- základní: `+` `-` `*` `/` `%`
- inkrementace: `++` `--`

**Porovnávání**
- základní `<` `<=` `>` `>=` `==` `!=`

**Bitové operace**
- základní `~` `&` `|` `^` `<<` `>>`

**Logické operace**
- základní `&&` `||` `!`

**Pointers**
- základní `&` `*`

**Přiřazovací**
- základní `=` `+=` `-=` `*=` `/=` `%=` `&=` `|=` `^=`

**Velikost v bytech**
- `sizeof`

**Ternary expression**
- `expression ? something1 : something2`

**Komentáře**
- jednořádkový `// komentář`
- víceřádkový `/* komentář */`

## Arrays (Pole)

- sekvence objektů stejného typu

```C++
int a1[3][2] = {{3,1},{4,5},{6,5}};

int a2[3]; // nahodne hodnoty
```

**Strings**
- string je pole o velikosti vstupu + 1

```C++
char text[] = "text"; // pole velikosti 5, poslední je znak '\0' (zero) NUL character

char chars[] = {'t', 'e', 'x', 't'} // pole velikosti 4
```

## Aligned/Misaligned slova

- prostě proměnné o velikosti X musí být uloženy na adrese, která je dělitelná počtem bytů X
- 64-bit long long int musí být uložen na adrese dělitné 8

## Struktury (structs)

- jsou data, která jsou nějak spojená
- co se týče aligned/misaligned, tak jsou uspořádaná v paměti podle proměnné s největším slovem (outer alignment) a uvnitř jsou pak uspořádaná podle velikosti slov proměnných (inner alignment)

```C++
struct bod {
    int x,y;
};

bod novy;
novy.x = 1;
novy.y = 2;
```

## Konstanty

**regulární** (klasické read-only)
- `const int konstanta = 5;`

**compile-time (neexistují v paměti)**
- `constexpr int konstanta = 5;`

**enumerátory**
- jsou efektivní, protože se převádí na `int`

```C++
enum class skupenstvi {KAPALNE, PLYNNE, PEVNE}

skupenstvi skup = skupenstvi::KAPALNE
```

## Preprocessor C++

- upravuje textový procesor ještě před kompilací, tedy odstraňuje komentáře, spojí řádky a zpracuje příkazy začínající `#`

- typy direktiv:

```C++
#include "muj_soubor.h" // vloží obsah jiného souboru
#include <stdio.h> //

#define X 152 // raději se vyvarovat, nahradí identifikátor v kódu zadaným řetězcem, místo toho použít constexpr

#ifdef, #ifndef, #if, #else, #endif // umožňuje vynechat nebo nahradit některé části kódu na základě podmínek
```

## Pointers