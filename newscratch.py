#!/bin/python3
import sys
import os
import random
import shutil
import json
from config import SCRATCHES_DIR, TEMPLATES_DIR, EDITOR


def main():
    name = f"{get_random_adjective()}-{get_random_noun()}-{random.randint(1000, 9999)}"
    template = os.path.join(TEMPLATES_DIR, sys.argv[1])
    path = os.path.join(SCRATCHES_DIR, name)

    os.mkdir(path)
    copy_files(template, path)

    with open(os.path.join(path, ".scratch.json"), "w") as f:
        json.dump({"name": name}, f)

    os.system(f"{EDITOR} {path}")


def copy_files(src_dir, dest_dir):
    for root, dirs, files in os.walk(src_dir):
        for file in files:
            src_file = os.path.join(root, file)
            dest_file = os.path.join(
                dest_dir, os.path.relpath(src_file, src_dir))
            os.makedirs(os.path.dirname(dest_file), exist_ok=True)
            shutil.copy2(src_file, dest_file)


def get_random_noun():
    nouns_file = os.path.join(os.path.dirname(
        os.path.realpath(__file__)), "nouns.txt")
    with open(nouns_file, "r") as f:
        nouns = f.read().split("\n")
        return random.choice(nouns)


def get_random_adjective():
    adjectives_file = os.path.join(os.path.dirname(
        os.path.realpath(__file__)), "adjectives.txt")
    with open(adjectives_file, "r") as f:
        adjectives = f.read().split("\n")
        return random.choice(adjectives)


if __name__ == "__main__":
    main()
