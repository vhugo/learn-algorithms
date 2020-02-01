# Dynamic connectivity problem

Given a set of N objects, a union command to connect two objects at a time and a
find/connected query to know if there is a path connecting the two objects.

Assuming "is connected to" is an equivalence relation:
- reflexive: *p* is connected to *p*
- symmetric: if *p* is connected to *q*, then *q* is connected to *p*
- transitive: if *p* is connected to *q* and *q* is connected to *r*, then *p*
  is connected to *r*

## Quick-Find (eager approach)

Maintains the invariant that *p* and *q* are connected if and only if `id[p]`
is equal to `id[q]`. In other words, in a single connected component all sites
must have the same value in `id[]`.

**Data structure**:
- Integer array `id[]` of length N.
- Interpretation: *p* and *q* are connected iff they have the same id.

**Find method**: check if *p* and *q* have the same `id`.

**Union method**: merge components containing *p* and *q*, change all entries
whose id equals `id[p]` to `id[q]`.

## Quick-Union (lazy approach)

Based on the same data structure—the site-indexed `id[]` array as Quick-Find,
but it uses a different interpretation of the values that leads to more
complicated structures. Specifically, the `id[]` entry for each site will be
the name of another site in the same component (possibly itself).

To implement find() we start at the given site, follow its link to another site,
follow that sites link to yet another site, and so forth, following links until
reaching a root, a site that has a link to itself. Two sites are in the same
component if and only if this process leads them to the same root.

To validate this process, we need union() to maintain this invariant, which is
easily arranged: we follow links to find the roots associated with each of the
given sites, then rename one of the components by linking one of these roots to
the other.

**Data structure**:
- Integer array `id[]` of length N.
- Interpretation: `id[i]` is parent of *i*.
- Root of *i* is `id[id[id[...id[i]...]]]`.

**Find method**: check if *p* and *q* have the same root.

**Union method**: merge components containing *p* and *q*, set the id of *p*'s
root to the id of *q*'s root.

## Weighted Quick-Union (improved quick-union)

Rather than arbitrarily connecting the second tree to the first for union() in 
the quick-union algorithm, we keep track of the size of each tree and always 
connect the smaller tree to the larger. 

**Data structure**:
- Integer array `id[]` of length N.
- Integer array `size[]` to count number of objects in the tree rooted at *i*.
- Interpretation: `id[i]` is parent of *i*.
- Root of *i* is `id[id[id[...id[i]...]]]`.

**Find method**: check if *p* and *q* have the same root. Identical to 
quick-union.

**Union method**: merge components containing *p* and *q*, link root of smaller 
tree to root of larger tree. Update the `size[]` array.

## Analysis

**[Quick-Find](#quick-find-eager-approach)** is too slow for huge problems. *Union is too expensive*.
It takes N<sup>2</sup> (quadratic) array accesses to process a sequence of
N union commands on N objects. Trees are flat, but too expensive to keep them flat.

**[Quick-Union](#quick-union-lazy-approach)** is also too slow. *Find too expensive*,
could be N array accesses, but it has a linear growth. Trees can be too tall, so 
expensive to search roots. 

**[Weighted Quick-Union](#weighted-quick-union-improved-quick-union)** find now 
takes time proportional to the depth of *p* and *q*. Union takes constant time, 
given roots. It is guaranteed that the depth of any node in the tree is at most 
the logarithm to the base two of N. Trees are much lower, saving time to 
search roots.

### Cost Model

Number of array accesses (for read or write)

| algorithm   | initialize | union                | find    |
| ----------- | ---------- | -------------------- | ------- |
| quick-find  |     N      |   N                  |   1     |
| quick-union |     N      |   N<sup>†</sup>      |   N     |
| weighted QU |     N      |   lg N<sup>†</sup>   |   lg N  |

<sup>†</sup>  includes cost of finding roots
lg = base-2 logarithm

Ref.:
- [1.5 Case Study: Union-Find](https://algs4.cs.princeton.edu/15uf/)
- [Lecture slide](http://bit.ly/1_5UnionFind)
