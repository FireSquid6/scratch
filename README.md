# Scratch
Scratch allows you to instantly create projects from your own defined templates using a custom rofi menu.

# Requirements
**Scratch only works on linux.**  
Make sure the following programs are on your system:
- Python 3.10
- [rofi](https://github.com/davatorium/rofi)
- [python rofi menu](https://github.com/miphreal/python-rofi-menu)
  - Just run `pip install rofi-menu`

# Setup
1. Clone the repo somewhere you can easily access
2. Edit `config.py` with the config variables you want 
3. Create your templates. A template is just a directory inside the templates folder you can put anything inside. This template will then be copied into your new project.
4. Run the following command with the path filled in:  
  `rofi -modi mymenu:<path-to-main.py> -show mymenu`  
  In my case, path to main.py is `/home/firesquid/scratch/main.py`
5. Put that command in a shell script, or bind it to a key in your window manager.

If you're having trouble installing due to my error or you're just new to linux, feel free to create an issue.

# Usage
1. Open the menu
2. Create a new scratch based on a template you have created
3. Edit the property `name` in `.scratch.json` to easily find the project later

# Configuration
Scratch uses the templates folder