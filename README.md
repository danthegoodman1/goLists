# GO Linked lists testing

Turns out mine is about 2x-2.1x faster than the `container/list` library

See it in action here: https://repl.it/@danthegoodman/Go-Linked-Lists

## Performance

#### Data Structures and Why we used Skip Lists

After extensively testing golang's built-in linked lists library, `container/list`, raw linked lists (smaller feature set), and `github.com/MauriceGit/skiplist` (skip lists), we found skip lists to have an impressive advantage over other data structures.

![speed comparison](/assets/unknown-1.png)
_Fig. 1 - built-in linked lists vs. raw linked lists vs. skip list operations_

As seen in _Fig. 1_, the speed in which a skip list can search for an element was over 25,000x faster than the built in `container/list` library, and almost 12,000x faster than the raw linked lists.

This is the major advantage that skip lists present, as a significantly large majority of operations will be search, as opposed to data modification operations.

By sacrificing some time on the speed of modification, we by far make that time up in search performance.
