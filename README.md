# ShoppingList

BE Side of the simple Shopping List app. 


##Endpoint list
> All calls except **signin** and **signup** should contain "auth" parameter in header. 
> "auth" value received in **signin** call response
### **POST** Sign in metod call example: 
>POST [YOUR_SERVER_ADDRESS:1881/signin]
```JSON
Request:
{
    "name": "AwesomeUser",
    "pass": "AwesomePass"
}
Response: 200 OK
{
    "Name": "AwesomeUser",
    "Token": "F2D052C796A264A666EE76B5350EB7BE"
}
```
### **POST** Sign up metod call example:
>POST [YOUR_SERVER_ADDRESS:1881/signup]
```JSON
Request:
{
    "name": "AwesomeUser",
    "pass": "AwesomePass"
}
Response: 200 OK
{
    "name": "AwesomeUser",
    "pass": "AwesomePass"
}

```

### **GET** Product List metod call example:
>GET [YOUR_SERVER_ADDRESS:1881/productList]
```JSON
Request:
{} //empty request
Response: 200 OK
[  
   {  
      "id":"B6889609-33DD-FA88-05A1-28670104AFC8",
      "name":"Milk",
      "isBought":false
   },
   {  
      "id":"C91272E0-34E0-F806-7792-A630E01F2876",
      "name":"Bread",
      "isBought":false
   },
   {  
      "id":"EB3FFCC6-D6C7-31F7-0A14-1D9F68D8A986",
      "name":"Meat",
      "isBought":false
   },
   {  
      "id":"EBA68C19-5B4B-B81B-559C-B67004661162",
      "name":"Water",
      "isBought":true
   },
   {  
      "id":"EE2A2F9D-EDE8-38E9-1D7B-5D7638DB8D80",
      "name":"Cereal",
      "isBought":false
   }
]
```

### **POST** Product metod call example:
>POST [YOUR_SERVER_ADDRESS:1881/product]
```JSON
Request:
{
    "name":"Milk"
}
Response: 200 OK
{
    "id":"FDD22A9C-BB7E-8ACE-76E2-17C0FCA80C17",
    "name":"Milk",
    "isBought":false
}
```

### **GET** Product by ID metod call example:
>GET [YOUR_SERVER_ADDRESS:1881/product/{id}]
```JSON
Request:
{}
Response: 200 OK
{
    "id":"FDD22A9C-BB7E-8ACE-76E2-17C0FCA80C17",
    "name":"Milk",
    "isBought":false
}
```
### **DELETE** Product by ID metod call example:
>DELETE [YOUR_SERVER_ADDRESS:1881/product/{id}]
```JSON
Request:
{}
Response: 200 OK
```
### **UPDATE[PUT]** Product by ID metod call example:
>PUT [YOUR_SERVER_ADDRESS:1881/product/{id}]
```JSON
Request:
{}
Response: 200 OK
{
    "id":"FDD22A9C-BB7E-8ACE-76E2-17C0FCA80C17",
    "name":"Milk",
    "isBought":true
}
```
