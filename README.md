# Add annotation

Available at https://github.com/nbossard/taskwarrior-hook-addannotation.git

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

## Usage

Sample usage, as usual with taskwarrior:

```bash
task add "Fixing MR222"

🪄HookAddAnnotation: Found prefix "MR"
🪄HookAddAnnotation: ✅ Added annotation "https://taiga.tech.orange/project/thommil-mahali-poc/merge-request/222"

Created task 73.

task 73

Name          Value
ID            73
Description   Fixing MR222
                2023-03-25 00:42:32 https://gitlab.com/nbossard/taskwarrior/-/merge_requests/222
Status        Pending
Entered       2023-03-24 23:42:32 (28s)
Last modified 2023-03-24 23:42:32 (28s)
Virtual tags  ANNOTATED LATEST PENDING READY TAGGED UNBLOCKED
UUID          bfed0cee-9be1-4c8b-8df5-8eccb1d62bab
Urgency        3.4

    annotations      0.8 *    1 =    0.8
    tags             0.8 *    1 =    0.8
    UDA priority.      1 *  1.8 =    1.8
                                  ------
                                     3.4
```

<!-- vim: set conceallevel=0: -->
