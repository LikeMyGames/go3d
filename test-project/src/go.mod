module game

go 1.23.3

// temporary

require go3d/engine v0.0.0

replace go3d/engine => ../../engine

// replace with
// replace go3d/engine => github.com/likemygames/go3d/engine

// require ../../engine v0.0.0

require (
	github.com/eiannone/keyboard v0.0.0-20220611211555-0d226195f203 // indirect
	golang.org/x/sys v0.34.0 // indirect
)
