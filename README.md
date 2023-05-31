# Add annotation

Available at <https://github.com/nbossard/taskwarrior-hook-addannotation.git>

This is a program to add as a taskwarrior hook to automatically add annotations
when creating or editing a task.

It adds annotation when a special pattern is found in created task description or in annotations.
For example if description contains "MR222" it will add annotation
`"https://gitlab.tech.orange/mahali/mahali-backend/-/merge_requests/222"`

You can add as many patterns as you want by adding lines in .taskrc config file.

## configuration

Add rules in taskwarrior .taskrc config file with following format:

```text
hookaddannotation.rule.r1= {"prefix":"MR", "URL":"https://gitlab.com/nbossard/taskwarrior/-/merge_requests/"}
hookaddannotation.rule.r2= {"prefix":"ISS", "URL":"https://taiga.tech.orange/project/thommil-mahali-poc/issue/"}
hookaddannotation.rule.r3= {"prefix":"US", "URL":"https://taiga.tech.orange/project/thommil-mahali-poc/us/"}
hookaddannotation.rule.r4= {"prefix":"TSK", "URL":"https://taiga.tech.orange/project/thommil-mahali-poc/task/"}
...
```

## Installation

```bash
go build hookaddannotation.go
cp hookaddannotation ~/.task/hooks/on-add-hookaddannotation
cp hookaddannotation ~/.task/hooks/on-modify-hookaddannotation
```

## Usage

Sample usage, as usual with taskwarrior:

```bash
task add "Fixing MR222"

🪄HookAddAnnotation: Found prefix "MR"
🪄HookAddAnnotation: ✅ Added annotation "https://taiga.tech.orange/project/thommil-mahali-poc/merge-request/222"

Created task 73.

```

<!-- vim: set conceallevel=0: -->
