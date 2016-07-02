galaxygen
================================================================================

Package `galaxygen` generates realistic-looking star galaxies.

This algorithm was originally made in c++, by:

    Bryon Mau (www.cyberbrinedreams.com)

And was then modified (in c++) by:

    Patrick Down (www.codemoon.com)

This was done years ago and both sites now seems to be down.

This go port was made from the salvaged parts from my shitty python port
(done ages ago too), that was based on the original c++ source. It has
now been cleaned up and made easier to read.


Install
--------------------------------------------------------------------------------

    go get github.com/lmas/galaxygen

Usage
--------------------------------------------------------------------------------

    import (
        "time"
        "math/rand"
        "github.com/lmas/galaxygen"
    )

    // Set a new seed to get random galaxies each time.
    rand.Seed(time.Now().UnixNano())

    // Make a new galaxy with at most 5000 stars, in 5 arms spread out in a radius
    // of 100 points.
    starmap := galaxygen.New(5000, 5, 100.0)

See `examples/` for more usage examples.

License
--------------------------------------------------------------------------------

I have no idea if and what license the original code had, but I've decided to
release my implementation under the MIT license anyway.

If you are the original author and don't like it, please contact me and we'll
sort it out. :)
