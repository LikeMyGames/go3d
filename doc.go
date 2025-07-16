/*
This package contains, once installed, contains many useful and powerful commands

The first, and arguably most important command, is the help command:

The command lists out all of the commands, except for itself,
along with a short description for it. This information is then printed to the console.

Structure:

	go3d help <?command-name ...>

The help command can be run by itself or in conjunction with the names
of other valid commands as arguments.

When run with the specificity of other command names, the help command will not print
the short description of a command, but rather, the long description,
which goes more in depth about what it does.

When the names of multiple commands are given, their long descriptions
are listed sequentially.

Example:

> Command:

>>	go3d help

> Result:

>>	new: ...

>>	build: ...

>>	run: ...

> Command:

>>	go3d help new

> Result:

>>	new:

>>	...

> Command:

>>	go3d help new build run

> Result:

>>	new:

>>	...

>>	build:

>>	...

>>	run:

>>	...

The next command is new:

The command uses a function inside the package to generate a file
hierarchy and files necessary to the development of a go-3D project.

Structure:

>	go3d new <project-name>

For this command, the <project-name> gets replaced with the desired name for the new project.
All other changes that a user may wish to make to the settings or otherwise of the project can
be made in the settings.json file (also created by this command).

Example:

>	go3d new test-project

>	go3d new "example project name"

This command creates the base for a go-3D project, which includes:

> 1. A settings.json file that holds all build settings used in [Build()].

> 2. A file and folder structure that is used by the builder for input and output.

> 3. A base [main.go] file in the [./src] directory that contains boilerplate code.

Example:

>	go3d new test-project

The following is the resulting file tree.

	> test-project
		> bin
		> public
			icon.png
			splash.png
			opening.mp4
		> src
			main.go
		settings.json
*/
package main
