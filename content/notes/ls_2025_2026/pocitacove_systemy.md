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

