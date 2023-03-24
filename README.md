# Add annotation

This is a program to add as a taskwarrior hook to automatically add annotations 
to tasks at time of creation.
It adds annotation when a special pattern is found in created task description.
For example if description contains "MR222" it will add annotation "https://gitlab.tech.orange/mahali/mahali-backend/-/merge_requests/222"
You can add as many patterns as you want by adding lines in hookaddannotation.go

## configuration

Change method main in hookaddannotation

## Installation

```bash
go build hookaddannotation.go
cp hookaddannotation ~/.task/hooks/on-add-addissueannotation
```

<!-- vim: set conceallevel=0: -->
