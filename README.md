# Inspire

<img src="http://www.metamorph.gr/wp-content/uploads/2013/10/inspire_0.jpg" width="800" height="400" />

Feeling down? Want to feel like a boss? Want to feel on top of the world? And want to do it all from the command line??? Introducing Inspire: a command line tool that produces words of inspiration, written in Go Language!!

## Instructions
To use, simply clone the repo, run "make all", and then run the inspire file ("./inspire"). Note that, due to how the repo is set up, you must have the following settings:

	- Running on an *nix system (not tested on anything other than 64-bit Ubuntu, but should work equally well)
	- Downloaded and installed Go language
	- Installed the following Go packages:
		- "strings"
		- "math/rand"
		- "time"
		- "color"

## Features
When the application is run, you are presented with a quick menu for using the app:

	- Quick
	- Schedule

Each of these are quite straightforward, but can be used for customizing your inspirational experience!

## Quick
The standard go-to run of the application, where we just present you with a quote of inspiration. Once produced, the application will exit with the assumption you've been inspired enough!

## Regular/Random
Want to be surprised with your inspirations? After all, the best times to get inspired are the times you don't know you want to be inspired! That's probably not true, but we'll stick with it anyway. In here, you can adjust all time settings so that either on a regular or completely random schedule, you will receive some inspiration. 

## Favorites
Cycles through the quotes from your favorites! 

## About
This was mostly intendend as an intro projects for the Go programming language (very basic). The forking section of it turned out to be pretty interesting (for the regular/random quotes) but was pretty basic otherwise. Check out Go though -- it's pretty sick!! If you want additional settings or changes, please feel free to fork or drop a comment!