#!/bin/bash
# git.sh - A script to initialize a repository, commit changes, and push to GitHub using hub,
# while sending notifications using notify-send.

# Check if hub is installed.
if ! command -v hub &> /dev/null; then
    notify-send "Error" "hub is not installed. Please install hub and try again."
    exit 1
fi

# Check if notify-send is installed.
if ! command -v notify-send &> /dev/null; then
    echo "notify-send is not installed. Please install libnotify-bin (or equivalent) and try again."
    exit 1
fi

# If repository is not initialized, initialize it.
if [ ! -d ".git" ]; then
    notify-send "Initializing Repository" "Repository not found. Running git init..."
    git init
    if [ $? -ne 0 ]; then
        notify-send "Error" "Failed to initialize git repository."
        exit 1
    fi
fi

# Add all changes.
notify-send "Adding Files" "Running: hub add ."
hub add .
if [ $? -ne 0 ]; then
    notify-send "Error" "Failed to add files."
    exit 1
fi

# Commit the changes.
notify-send "Committing Changes" "Running: hub commit -m 'Piece De Revolution'"
hub commit -m "Piece De Revolution"
if [ $? -ne 0 ]; then
    notify-send "Error" "Failed to commit changes."
    exit 1
fi

# Create the GitHub repository.
notify-send "Creating Repository" "Running: hub create"
create_output=$(hub create 2>&1)
if [ $? -ne 0 ]; then
    notify-send "Error" "Failed to create repository on GitHub: $create_output"
    exit 1
else
    notify-send "Repository Created" "$create_output"
fi

# Rename the current branch to 'main'.
notify-send "Renaming Branch" "Running: hub branch -M main"
hub branch -M main
if [ $? -ne 0 ]; then
    notify-send "Error" "Failed to rename branch to main."
    exit 1
fi

# Push the changes to GitHub.
notify-send "Pushing Changes" "Running: hub push -u origin main"
push_output=$(hub push -u origin main 2>&1)
if [ $? -ne 0 ]; then
    notify-send "Error" "Failed to push to GitHub: $push_output"
    exit 1
else
    notify-send "Push Complete" "$push_output"
fi

notify-send "Done" "Repository setup complete."

