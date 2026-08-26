#!/bin/bash

# Ensure the 'discogs' session exists
if ! tmux has-session -t discogs 2>/dev/null; then
  # Create a new session named 'discogs', detached
  cd /workspaces/discogs
  tmux new-session -d -s discogs
 fi
