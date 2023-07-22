#!/bin/python3
import rofi_menu
import json
import os
from config import EDITOR, ROFI_VERSION, SCRATCHES_DIR, TEMPLATES_DIR, NEWSCRATCH

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

    for dir in os.listdir(TEMPLATES_DIR):
        template = os.path.join(TEMPLATES_DIR, dir)
        if os.path.isdir(template):
            shell_command = f"{NEWSCRATCH} {template}"
            items.append(rofi_menu.ShellItem(f"new: {dir}", shell_command))


if __name__ == "__main__":
    main()
