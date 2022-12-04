~~It's regex day! 🎉~~ [1]

I'm really liking [golangdocs](https://golanddocs.com)
Good example set of recipes

Generic Map/Reduce/Filter - #want
These are such _powerful_ hof I just can't live w/o them :'(

Hmm... []byte, string... this is painful

Maybe the point here is to make the obvious so painful,
we do the most performant since it's the same amount of work?
I'm tempted to just parse this explicitly, byte by byte

Maybe in a refactor

At some point I'm going to want to spend time figuring out
the go debugger here in nvim

Think I hit some bug in the renamer... could _not_ rename end
to End w/o a lua error

## Subnote 1
After starting down regex, it (as usual) was overpowered

I started looking through idiomatic scanning apis (text/scanner)
and realized that DOH fmt.Sscanf does all I need!
