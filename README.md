# What is this repository?
Dump the sample implementation if I start a new project

# Concept of each directory
| directory | explanation |
| --- | --- |
| app | application logis |
| cmd | starting point of any operations |
| config | struct of setting variables |
| environment | put variables that depends on environment. Use env var for credentials |
| infrastructure | pkg to interact with external things such as DB. Use interface on the side of app so that we avoid dependencies in test |
| testsupport | pkg to support test |

# Setup
```
make up
make mod
make migrate
```
