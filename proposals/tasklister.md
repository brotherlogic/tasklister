# Tasklister

I like to make task lists in markdown and then copy paste them into my
old tasklist manager to have them turn into github issues. I find this
works better for me than using any of the github provided issue or
project trackers.

The problem lies in that this is inefficient and it doesn't make
clear what is going on with the tasks, nor does it link in the subsequent
issues into the original tasklist. Hence the need for tasklister.

## Requirements

1. Tasklister can search all the repos to find any markdown files
1. Tasklister can take a section marked as Tasks and turn it into
   some issues in github
1. Tasklister can track these issues and update the tasklist file
1. Tasklister handles sub-issues, adding these to the list

## Building Process

Since a lot of tasklister is async, we have a processing queue.
So tasklister receives an item from the worker queue process it
and then acks the queue to remove the item. Rinse and Repeat. The queue
is push and sends the work item to the tasklister

We initially receive these type of work itmes:

1. Resync
1. Process Markdown
1. Process List
1. Recieve Issue close
1. Recieve repo update

### Resync

Resync is a periodic process that goes through each repo in github
and looks for Markdown files. So we get a repo, list the files in
the repo and see if there is any markdown in there. We keep a count
of repos -> hash and we only do a full search if the hash has changed.

New markdown files are stored in the markdown repo

user/repo/path -> (hash, last_checked)

If we see that the existing hash matches the read repo, we don't run an update.

### Process Markdown

The goal of the process piece is to evaluate if the markdown file contains
a tasklist. If it does we store it in the tasklist repo; If we already
have an entry for this file we skip.

user/repo/path -> Tasklist (parsed and synced)

### Process List

This is the key hub that syncs the tasklist with github. Process is straightforward,
we create an issue for each line in the taskslist and file a bug for the pending task
on the list. We also adjust the tasklist to include a ref to the issue directly in markdown.

A few things could then happen:

1. Issue is closed. We mark the task as complete, and strikeout the ref to the issue in markdown.
1. Issue has new subchildren. We add these as subtasks of the main taks and insert refs into the markdown.
   Once an issue has subchildren, then if those subchildren are all closed, then the task is considered complete
   and is closed automatically.

Once an issue is closed the next available is ready for write which will happen on the next pass. We run a full
pass of tasklists every hour, and on every issue close/open within github.

### Receive Issue Close

Issue closing can trigger a process list if we can identify the closed issue in a given list. We just refresh that list

### Receive Repo Update

Repo update (i.e. code push), may indicate new tasks, we run a resync on that repo.

### Receive Issue Open

Issue opens can be subchildren. Check for this and process a relevant list if we find a match.

## Queue

The queue is simple, and held in memory since we don't expect a large backlog of tasks. We fix the size at 100 (?) elements. We try
to avoid fanout where possible, relying on github caching to ensure we don't run into throttling issues.

## Tasks

1. Write process that checks out a github repo into local space - test on our repo
1. Have the process edit a file in the repo (e.g. replace (## Tasks > ## Tasks (handled by tasklister)))
1. Have the process write that change back to the repo
1. Run this process in the cluster
1. Write the queue handler
1. Handle repo sync
1. Handle process markdown
1. Handle process lists
