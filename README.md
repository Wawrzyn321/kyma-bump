#Mr BUMP
> Oh, poopity poop! Where is docs deployment?

###Commands:

####auto
Diffs current state of Kyma and Console and bumps related images.

Usage:
```
bump auto -c <console tag> -k <kyma tag> [--no-verify, -f]

bumo auto -c PR-1892 -f
bump auto -k ef2291 -c a3b23171
```
You can then inspect introduced changes in your local repositories.
####img
Updates tags of images.

Usage:

```
bump img <tag1> <...images> <tag2> <...images> [--no-verify, -f]

bump img tets addons --no-verify
bump img PR-1958 core core-ui PR-9337 cbs
```
####help
Displays help.

Usage:

```
bump help
bump -h
```
####list
Displays list of available images, along with their aliases.

Usage:

```
bump list
bump -l
```
###verify
Verifies changed images in repo. You can pass branch to diff changes by. Defaults to HEAD.

Usage:

```
bump verify [<revision>]
bump verify
bump verify issue-1925
```
####check-files
Checks if YAML configuration files exist and their tag variable paths match. Use for debugging.

Usage:

```
bump check-files
```
####egg
Haunts your nightmares.

```
bump egg
```

###Info
* `--no-verify (-f)` skips image check.
* Commit form of tag requires at least 8 characters. You can also use full commit tag, it will be trimmed automatically.

###Installation

* Run `go install mrbump.go`.
* Fill your Kyma and Console paths in `.bumprc` file and add it to your `.bashrc`, `.zshrc` etc.