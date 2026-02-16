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

A znaménko: $\operatorname{sgn}(p) = (-1)^{\#inverzí}$

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

protože $|S_n| = 2$, nebo-li $S_2 = \{(1,2),(2,1)\}$

1. identita

pro $p=(1,2)$ platí $\operatorname{sgn}p=+1$ $\prod_{i=1}^{n} a_{i,p(i)} = a_{11}a_{22}$

2. #inverzí = 1

pro $p=(2,1)$ platí $\operatorname{sgn}p=-1$ $\prod_{i=1}^{n} a_{i,p(i)} = - a_{12}a_{21}$

Exaktní odvození lze získat z vyjádření neznámé obecně metodou sčítací, dosazovací. Pro dvě rovnice: $x_1$ a $x_2$. Pro tři rovnice pak $x_1$, $x_2$ a $x_3$.

Přesněji:
$$
\begin{aligned}
a_{11}x_1 + a_{12}x_2 &= b_1 \\\\
a_{21}x_1 + a_{22}x_2 &= b_2
\end{aligned}
$$

$$
\
x_1=\frac{b_1a_{22}-b_2a_{12}}{a_{11}a_{22}-a_{21}a_{12}},
\qquad
x_2=\frac{a_{11}b_2-a_{21}b_1}{a_{11}a_{22}-a_{21}a_{12}}
\
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
