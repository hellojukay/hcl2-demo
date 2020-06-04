# hcl2-demo
parse hcl config demo

# usage
```bat
PS C:\Users\Administrator\Desktop\hcl2-demo> .\hcl2-demo.exe --help 
Usage of C:\Users\Administrator\Desktop\hcl2-demo\hcl2-demo.exe:
  -f string
        config file path (default "./taskfile")
PS C:\Users\Administrator\Desktop\hcl2-demo> .\hcl2-demo.exe -f .\taskfile
{"pipeline":{"env":{"PWD":"C:\\Users\\Administrator\\Desktop\\hcl2-demo"},"stages":{"stage":[{"Name":"first","Describe":"hellojukay"},{"Name":"second","Describe":""}]}}}
PS C:\Users\Administrator\Desktop\hcl2-demo>
```