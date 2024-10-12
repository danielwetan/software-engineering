```
rails g model Product name:string price:float stock:integer 

rails g model Transaction amount:integer status:string

rails g model Payment method:string amount:integer status:string 
```

Then manually add reference 
transaction has one product
transaction has many payments
