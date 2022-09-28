
## Branch information

### Main Branches
- Branch name for production releases = [main]
### Branch prefixes
- Feature branches = [feature/]
- Bugfix branches = [bugfix/]
- Release branches = [release/]
- Hotfix branches = [hotfix/]
- Support branches = [support/]
- documentation branches = [doc/]
- Version tag prefix = [v]

### releases
For local building from source, run `mage build` 
For Release generation, run `mage release` 

## Roadmap
### Goals
- Allow users to delete associations 
- Allow users to change priority of a goal
- Allow users to change dueDate of a goal
- Allow users to change the name of a goal
- Allow users to delete a goal
- Allow users to change the status of a goal
- Allow users to change teh priority of a goal
- Allow users to retrieve past due goals

### Log entries
- Allow users to migrate an entry to a new notebook
- Allow users to update the log entry 
- Allow users to add tags
- Allow users to remove tags
- Allow users to delete a log entry 
   
### Filtering
- Allow users to filter items that are returned under the list command
  - --filter [logentry/goal].[element_to_filter](>,<,=)[value]

### Import functionality
Allow users to import journal entries to migrate to the platform
Allow users to import goals
Allow users to import associations
Allow users to import export of the journal -- this can double as a backup method instead of backing up the sqlite db

