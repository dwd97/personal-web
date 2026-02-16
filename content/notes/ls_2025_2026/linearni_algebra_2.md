## Teorie - bez důkazů

### 1. Determinanty

#### 1.1 Definice determinantu

Buď $A \in T^{nxn}$ čtvercová matice. Pak determinant matice $A$ je:

$$
det(A) = \sum_{p \in S_n} \operatorname{sgn}(p) \prod_{i=1}^{n} a_{i, p(i)}
$$

Značí se $det(A)$ nebo $|A|$

Pokud je $det(A) \neq 0$, pak $Ax = b$ má **jednoznačné** řešení, kde $A \in T^{nxn}$

Zjednodušene to znamená, že dostanu sčítance přes všechny permutace z $p \in S_n$. Tedy z matice se vybere n prvků tak, že se z každého řádku (index i) a z každého sloupce (index p(i)) se vybere právě jeden. Tyto prvky se pak vzájemně vynásobí.

Kde $S_n$ je grupa všech možných permutací na množině indexů $\{1,...,n\}$. Suma tedy iteruje přes všechny možné permutace.

Lze zapsat také:

$$
det(A) = \sum_{p \in S_n} \operatorname{sgn}(p) a_{1,p(1)} \cdot ... \cdot a_{n, p(n)}
$$

pro determinant matice řádu 2 $rank(A)=2$ platí:

$$
\\det\\begin{pmatrix}
a_{11} & a_{12} \\\\
a_{21} & a_{22}
\\end{pmatrix} = a_{11}a_{22} - a_{21}a_{12}
$$

Platí také:

$$
\det A = \det (A^T)
$$

## Důkazy

## Příklady

### Determinanty

- počítání determinantu matice nižšího řádu podle definice
- počítání determinantu matice vyššího řádu

## Aplikace
