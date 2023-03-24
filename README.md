# Add annotation

This is a taskwarrior hook script to automatically add annotations to tasks.
For example if description contains "MR222" will add annotation "https://gitlab.tech.orange/mahali/mahali-backend/-/merge_requests/222"

## configuration

Change method main in addissueannotation

## Installation

```bash
go build addissueannotation.go
cp addissueannotation ~/.task/hooks/on-add-addissueannotation
```

<!-- vim: set conceallevel=0: -->
