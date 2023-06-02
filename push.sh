#!/bin/bash

# Read the repository URL from user input
read -p "Enter the GitHub repository URL: " repository_url

# Initialize the local repository as a Git repository
git init

# Add the remote repository URL
git remote add origin "$repository_url"

# Commit the changes
git add .
git commit -m "Initial commit"

# Push the changes to GitHub
git push -u origin master

