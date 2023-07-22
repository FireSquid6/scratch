#!/bin/python3
import rofi_menu
import json
import os

# ----------------------------------
#          CONFIGURATION
# ----------------------------------

# The executable to run to open your editor. I.e. "vim", "code", "nvim", etc.
EDITOR = "code"

# your rofi version. Find out with `rofi -v`
ROFI_VERSION = "1.7.1"

# The path to store your scratches in. I.e. "/home/username/scratch/scratches"
SCRATCHES_DIR = os.path.expanduser("~/scratch/scratches")

# The path to store your templates in. I.e. "/home/username/scratch/templates"
TEMPLATES_DIR = os.path.expanduser("~/scratch/templates")

# ----------------------------------
#               CODE
# -----------------------------------


def main():
    rofi_menu.run(Menu(), rofi_version=ROFI_VERSION)


class Menu(rofi_menu.Menu):
    prompt = "Scratch"
    items = []

    for dir in os.listdir(SCRATCHES_DIR):
        project = os.path.join(SCRATCHES_DIR, dir)
        scratch_json = os.path.join(SCRATCHES_DIR, dir, ".scratch.json")
        if os.path.isfile(scratch_json):
            shell_command = ""
            description = ""

            with open(scratch_json, "r") as f:
                scratch = json.load(f)
                shell_command = f"{EDITOR} {project}"
                description = f"open: {scratch['name']}"

            items.append(rofi_menu.ShellItem(description, shell_command))


if __name__ == "__main__":
    main()
