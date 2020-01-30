# Dynamic connectivity problem

Given a set of N objects, a union command to connect two objects at a time and a
find/connected query to know if there is a path connecting the two objects.

Assuming "is connected to" is an equivalence relation:
- reflexive: *p* is connected to *p*
- symmetric: if *p* is connected to *q*, then *q* is connected to *p*
- transitive: if *p* is connected to *q* and *q* is connected to *r*, then *p*
  is connected to *r*

## Quick-find (eager approach)

Maintains the invariant that *p* and *q* are connected if and only if `id[p]`
is equal to `id[q]`. In other words, in a single connected component all sites
must have the same value in `id[]`.

**Data structure**:
- Integer array `id[]` of length N.
- Interpretation: *p* and *q* are connected iff they have the same id.

**Find method**: check if *p* and *q* have the same `id`.

**Union method**: merge components containing *p* and *q*, change all entries
whose id equals `id[p]` to `id[q]`.

## Quick-union (lazy approach)

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

## Analysis

**[Quick-find](#quick-find-eager-approach)** is too slow for huge problems. *Union is too expensive*.
It takes N<sup>2</sup> (quadratic) array accesses to process a sequence of
N union commands on N objects. Trees are flat, but too expensive to keep them flat.

**[Quick-union](#quick-union-lazy-approach)** is also too slow. *Find too expensive*,
could be N array accesses. However it has a linear growth  

### Cost Model

Number of array accesses (for read or write)

| algorithm   | initialize | union             | find |
| ----------- | ---------- | ----------------- | ---- |
| quick-find  |     N      |   N               |   1  |
| quick-union |     N      |   N<sup>†</sup>   |   N  |

<sup>†</sup>  includes cost of finding roots

Ref.:
- [1.5 Case Study: Union-Find](https://algs4.cs.princeton.edu/15uf/)
- [Lecture slide](http://bit.ly/1_5UnionFind)
