import os

# ----------------------------------
#          CONFIGURATION
# ----------------------------------

# The executable to run to open your editor. I.e. "vim", "code", "nvim", etc.
EDITOR = "kitty"

# your rofi version. Find out with `rofi -v`
ROFI_VERSION = "1.7.1"

# The path to store your scratches in. I.e. "/home/username/scratch/scratches"
SCRATCHES_DIR = os.path.expanduser("~/scratch/scratches")

# The path to store your templates in. I.e. "/home/username/scratch/templates"
TEMPLATES_DIR = os.path.expanduser("~/scratch/templates")

# The path to the newscratch.py script. It's located right next to this one. I'm too lazy to write code to find it automatically, so it's your problem.
# If you do end up writing code to find it automatically, submit a pull request to be my best friend
NEWSCRATCH = os.path.expanduser("~/scratch/newscratch.py")
