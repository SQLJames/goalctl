# Associate goals and logentries

## Use case
Need the ability to print out goals with all the log entries 



## commands neeeded
Associate entry to goals
goalctl.exe link \
--goalid 1 --goalid 2 --goalid 3 \
--entryid #if null assumes latest

goalctl.exe list \
goals 
--long #if null assumes short otherwise print all logentries

