# reverse-shell
Simple demo of reverse shell

![Demo](https://github.com/jclegras/reverse-shell/blob/master/media/demo.gif)

# Overview

This is a simple demo of reverse shell.
I'm using a Golang payload for the purpose of this example got from [here](https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/Methodology%20and%20Resources/Reverse%20Shell%20Cheatsheet.md#golang)

# Template

We have three Docker containers popped up from a docker-compose file.
They share the same network and the browser app is exposed via the 5800 port.

# How it works

Basically we go to the form page and we inject a payload in order to create a reverse shell.
So the server container send its stdout/stderr to my given local host (via e.g netcat) while we send it each command to execute.

In a nutshell, this example is deliberately simple and awkward but it allows me to remember the principle.
