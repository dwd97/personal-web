## Teorie - bez důkazů

### 1. Determinanty

#### 1.1 Definice determinantu

Buď $A \in T^{n \times n}$ čtvercová matice. Pak determinant matice $A$ je:

$$
\det(A) = \sum_{p \in S_n} \operatorname{sgn}(p) \prod_{i=1}^{n} a_{i, p(i)}
$$

Značí se $\det(A)$ nebo $|A|$.

Pokud je $\det(A) \neq 0$, pak $Ax = b$ má **jednoznačné** řešení, kde $A \in T^{n \times n}$.

Zjednodušeně to znamená, že dostanu sčítance přes všechny permutace $p \in S_n$. Z matice se vybere $n$ prvků tak, že se z každého řádku (index $i$) a z každého sloupce (index $p(i)$) vybere právě jeden. Tyto prvky se pak vzájemně vynásobí.

Kde $S_n$ je grupa všech možných permutací na množině indexů $\{1,\dots,n\}$. Suma tedy iteruje přes všechny možné permutace.

A znaménko:
$$
\operatorname{sgn}(p) = (-1)^{\text{počet inverzí}}
$$

Lze zapsat také:

$$
\det(A) = \sum_{p \in S_n} \operatorname{sgn}(p)\, a_{1,p(1)} \cdot \dots \cdot a_{n,p(n)}
$$

Pro determinant matice řádu 2 ($\operatorname{rank}(A)=2$) platí:

$$
\det
\begin{pmatrix}
a_{11} & a_{12} \\
a_{21} & a_{22}
\end{pmatrix}
=
a_{11}a_{22} - a_{21}a_{12}
$$

Protože $|S_2| = 2$, tj. $S_2 = \{(1,2),(2,1)\}$:

1. identita  

pro $p=(1,2)$ platí $\operatorname{sgn}(p)=+1$ a
$$
\prod_{i=1}^{n} a_{i,p(i)} = a_{11}a_{22}
$$

2. jedna inverze  

pro $p=(2,1)$ platí $\operatorname{sgn}(p)=-1$ a
$$
\prod_{i=1}^{n} a_{i,p(i)} = a_{12}a_{21}
$$

Tedy determinant:

$$
\det A =
\begin{vmatrix}
a_{11} & a_{12} \\
a_{21} & a_{22}
\end{vmatrix}
=
(+1)\, a_{11}a_{22} + (-1)\, a_{12}a_{21}
$$

Exaktní odvození lze získat z vyjádření neznámých obecně metodou sčítací či dosazovací. Pro dvě rovnice: $x_1$ a $x_2$. Pro tři rovnice pak $x_1$, $x_2$ a $x_3$.

Přesněji:

$$
\begin{aligned}
a_{11}x_1 + a_{12}x_2 &= b_1 \\
a_{21}x_1 + a_{22}x_2 &= b_2
\end{aligned}
$$

$$
x_1=\frac{b_1a_{22}-b_2a_{12}}{a_{11}a_{22}-a_{21}a_{12}}, \qquad
x_2=\frac{a_{11}b_2-a_{21}b_1}{a_{11}a_{22}-a_{21}a_{12}}
$$

#### 1.2 Sarrusovo pravidlo

Je to mnemotechnická pomůcka na výpočet determinantu matice řádu 3, tedy $A \in T^{3 \times 3}$.

![sarrusovo_pravidlo](/images/notes/linearni_algebra_2/sarrusovo_pravidlo.png)

Protože
$$
S_3 = \{ (1,2,3),(1,3,2),(2,1,3),(2,3,1),(3,1,2),(3,2,1) \}
$$
má 6 prvků, kde permutace $(1,2,3), (2,3,1), (3,1,2)$ mají $\operatorname{sgn}(p)=+1$ a permutace $(1,3,2),(2,1,3),(3,2,1)$ mají $\operatorname{sgn}(p)=-1$.

#### 1.3 Vlastnosti determinantu

Má-li $A$ nulový řádek, pak $\det A = 0$.

## Důkazy

### 1. Determinanty

#### Determinant matice je roven determinantu transpozonované matice

**Pozorování:** $ \det A = \det(A^T) $

**Důkaz:**

#### Singulární matice mají determinant roven nule

**Pozorování:** Má-li $A$ nulový řádek, pak $\det A = 0$.

#### Pro trojúhelníkové matice platí, že determinant je roven součinu prvků diagonály.

**Pozorování:** Pro trojúhelníkové (i diagonální) matice $A \in T^{n \times n}$, kde $a_{ij} = 0$ pro $i>j$, platí:

$$
\left|
\begin{array}{cccc}
a_{11} & a_{12} & \cdots & a_{1n} \\
0      & a_{22} & \cdots & a_{2n} \\
\vdots & \ddots & \ddots & \vdots \\
0      & \cdots & 0      & a_{nn}
\end{array}
\right|
=
a_{11}a_{22}\cdots a_{nn}
$$

**Důkaz:**

## Příklady

### Determinanty

- počítání determinantu matice nižšího řádu podle definice
- počítání determinantu matice vyššího řádu

## Aplikace
