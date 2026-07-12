# todos

1. Apply:
- revert
- status
- conflicts
- overlook
- config


2. record
    1. we validate the path is not empty string 
    2. we validate the file is valid 
    3. we validate that the file has changes else an empty patch is rejected 
    4. we check if a patch exists for that file already, if yes defer/todo 
    5. else we create a patch