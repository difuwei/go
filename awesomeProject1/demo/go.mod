module demo

replace fuweidi.com/foo => ../foo
replace fuweidi.com/foo2 => ../foo2


go 1.17

require (
	fuweidi.com/foo v0.0.0-00010101000000-000000000000
	fuweidi.com/foo2 v0.0.0-00010101000000-000000000000
)

