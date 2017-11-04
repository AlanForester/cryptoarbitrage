### Install GLIDE

`The easiest way to install the latest release on Mac or Linux is with the following script:

`curl https://glide.sh/get | sh`

On Mac OS X you can also install the latest release via Homebrew:

`$ brew install glide`

On Ubuntu Precise (12.04), Trusty (14.04), Wily (15.10) or Xenial (16.04) you can install from our PPA:

`sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update`
`sudo apt-get install glide`

On Ubuntu Zesty (17.04) the package is called golang-glide.

Binary packages are available for Mac, Linux and Windows.

For a development version it is also possible to go get github.com/Masterminds/glide.

To build from source you can:

Clone this repository into $GOPATH/src/github.com/Masterminds/glide and change directory into it
If you are using Go 1.5 ensure the environment variable GO15VENDOREXPERIMENT is set, for example by running export GO15VENDOREXPERIMENT=1. In Go 1.6 it is enabled by default and in Go 1.7 it is always enabled without the ability to turn it off.
Run make build
This will leave you with ./glide, which you can put in your $PATH if you'd like. (You can also take a look at make install to install for you.)

The Glide repo has now been configured to use glide to manage itself, too.`


https://github.com/sevlyar/go-daemon