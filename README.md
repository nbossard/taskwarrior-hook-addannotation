# Add annotation

This is a program to add as a taskwarrior hook to automatically add annotations
to tasks at time of creation.
It adds annotation when a special pattern is found in created task description.
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
cp hookaddannotation ~/.task/hooks/on-add-addissueannotation
```

<!-- vim: set conceallevel=0: -->
