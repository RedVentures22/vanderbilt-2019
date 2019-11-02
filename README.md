## If you have any problems just let us know (or if you're remote, create an issue) and we will be glad to help you

# Install git

First check if you have git installed

In your command line run:

`git --version`

If you see a version show up (something like `git version 2.23.0` then go to the next step)

# Install Golang (aka Go)

Check if you have Go installed on your machine by running this in your command line: 

`go`

If you see a big ol list of shit pop up then you're good

Otherwise, install Go:

[For Mac](https://www.callicoder.com/golang-installation-setup-gopath-workspace/#mac-os-x)

[For Windows](https://www.callicoder.com/golang-installation-setup-gopath-workspace/#windows)

[For Linux](https://www.callicoder.com/golang-installation-setup-gopath-workspace/#linux)

# Set your Go path

Now let's set up your Go path. Again, let's check if you have this already by running: 

`echo $GOPATH`

If you see something like:

`/Users/sstanton/go`

Then you are good

Other wise follow these instructions:

[For Mac/Linux](https://www.callicoder.com/golang-installation-setup-gopath-workspace/#unix-systems-linux-and-macos)

[For Windows](https://www.callicoder.com/golang-installation-setup-gopath-workspace/#windows-system)

# Hello World

Let's make sure you didn't fuck up. Follow these instructions:

[Build Hello World App](https://www.callicoder.com/golang-installation-setup-gopath-workspace/#testing-your-go-installation-with-the-hello-world-program)

Once you run `go run hello.go`, then you can stop here.

# Install Dependencies

We need one dependency for this demo. Don't worry, we will explain what this does:

In your command line run:

`go get -u github.com/gorilla/mux`

Don't worry, it's safe. You can trust us...