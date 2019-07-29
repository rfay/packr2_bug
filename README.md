This repo exists only to demonstrate a bug in packr2:

* If you use a directory name (perhaps "assets") in your code where the bundled files are supposed to be. 
* But then you run the built executable in a directory that happens to have a directory of that name.
* The files in "assets" in the *current directory* are listed, instead of the expected files that were bundled into the executable.

Instructions:

1. Build this project: `go install packr2_bug.go`
2. Run `~/go/bin/packr2_bug`, and you'll see the expected file, "junk-from-packr2.txt"
3. Now `mkdir -p /tmp/junk/assets && cd /tmp/junk && touch assets/this-was-in-pwd-assets.txt`. You've created a "junk" directory with a subdirectory named "assets", which has a file named "this-was-in-pwd-assets.txt" in it.
4. Run `~/bin/packr2_bug`. The output will be the pwd's assets subdirectory ("this-was-in-pwd-assets.txt") instead of the bundled file. 
