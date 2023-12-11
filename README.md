# bigcal(1)

Kozmix Go, 21-JAN-2021

bigcal


<a name="synopsis"></a>

# Synopsis

```
bigcal [[month] year]
```


<a name="description"></a>

# Description

**bigcal**
produces a monthly calendar just like
**cal**,
only bigger.


<a name="see-also"></a>

# See Also

**cal**(1)


<a name="bugs"></a>

# Bugs

When running
**bigcal**
_year_,
it might be more useful if the first day of the next month were
printed right after the last day of the previous month, producing a
continuous calendar for the whole year. Currently, it prints all
months separately from each other.

The
**cal**
_9 1752_
trick doesn't work as expected with
**bigcal**
(go pkg/time design decision).


<a name="author"></a>

# Author

svm

