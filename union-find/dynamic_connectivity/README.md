# Dynamic connectivity problem

Given a set of N objects, a union command to connect two objects at a time and a
find/connected query to know if there is a path connecting the two objects.

Assuming "is connected to" is an equivalence relation:
- reflexive: *p* is connected to *p*
- symmetric: if *p* is connected to *q*, then *q* is connected to *p*
- transitive: if *p* is connected to *q* and *q* is connected to *r*, then *p*
  is connected to *r*


## Quick-find (eager approach)

**Data structure**:
- Integer array id[] of length N.
- Interpretation: *p* and *q* are connected iff they have the same id.

**Find method**: check if *p* and *q* have the same id.

**Union method**: merge components containing *p* and *q*, change all entries
whose id equals `id[p]` to `id[q]`.
