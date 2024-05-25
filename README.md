# Scratch
Scratch is a cli tool for easily creating and navigating your programming projects


## Commands
`scratch create <template> <name>` - creates a new project
`scratch pad <template>` - creates a new scratchpad
`scratch template <name>` - creates a new template
`scratch dir` - takes you to your source directory
`scratch projects` - lists all projects
`scratch pads` - lists all scratchpads by name
`scratch elevate <name>` - elevates the current scratchpad to a project
`scratch archive <name>` - a


## Config File
Example config:
```yaml
projectsDirectory: "~/source" 
scratchesDir: "~/scratchpad"  # the directory to store temporary scratchpads
autoGit: true  # whether scratch will automatically initialize git


```
