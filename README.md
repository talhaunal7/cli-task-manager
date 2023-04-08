# cli-task-manager

a CLI tool that can be used to manage your TODOs in the terminal. 

---

<p align="">
  <img src="https://github.com/talhaunal7/cli-task-manager/blob/main/task.gif" alt="your_gif_alt_text">
</p>

```
$ task
Task is a CLI tool that helps you add, delete and list tasks easily from your terminal.

	-task add (add a new task)
	-task list (list all tasks)
	-task do (remove a task)

Usage:
  task [command]

Available Commands:
  add         Add new task
  do          Remove a task with given task number
  empty       Deletes all tasks.
  help        Help about any command
  list        List the tasks
  version     Print the version number of CLI Task Manager

Flags:
  -h, --help     help for task
```
### Example Usage

```
$ task add finish project

$ task add clean the dishes

$ task list
1- finish project
2- clean the dishes

$ task do 1

$ task list
1- clean the dishes
```
*Note: Lines prefixed with `$` are lines where we type into the terminal, and other lines are output from our program.*

### Tech stack & Open-source libraries

* 	[Cobra](https://github.com/spf13/cobra) - Cobra is a library for creating powerful modern CLI applications.
* 	[BoltDB](https://github.com/boltdb/bolt) - BoltDB is an embedded ACID key/value database written in Go.

### Roadmap

- [x] Organize Project Layout
- [ ] Switch to a relational database
- [ ] Add task tags
- [ ] Add task priority
- [ ] Add task notes

### Installation & Usage

1. Installation 
```
git clone https://github.com/talhaunal7/cli-task-manager.git
cd cli-task-manager
```
2. Build and configuration
```
go install 
// Once you've done that the binary file will be in your $GOPATH/bin directory
```
After that, to reach to command from any directory open zshrc file `open ~/.zshrc` and add bin directory to there
```
export PATH=$PATH:$GOPATH/Users/Example/go/bin
```
