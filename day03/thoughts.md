# Part 1
Geez.... Go likes iteration and mutation. I freaking hate it.
Rust?

Golang is made for interviewing, not for getting shit done. Change my mind.
Want to get the keys out of map? Just write another GD loop

Surely there's an easier way to get these stack traces readable?

Off by one... I've had too many beers for this.

# Part 2
These elves are always causing shit. Santa needs more reliable workers.

Do we have chunking loops?

## Using /x/exp
I was trying to use only the std library for the first week...but too basic for my tastes.

I am flexing the rules a bit, and pulling in the 1.18 experimental package for generic operations.

This makes golang less of a toy.

## Why do we not have Set?

Me: "Mom can we have Sets?"
Go: "We've got Sets at home!"
Sets at Home:

```go
	iAmSeT := map[rune]struct{}{}
	iAmSeT['A'] = struct{}{}
	_, has := iAmSeT['A']
```
