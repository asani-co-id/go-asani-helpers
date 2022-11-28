# asani-helper@2022
## How to use
- get to your repo project, below example:
```
    go get github.com/asani-co-id/go-asani-helpers
```
- import when use , 
```
import (
  asaniHelpers "github.com/asani-co-id/go-asani-helpers"
)
``` 
### Logging Helpers
- config your Logging , using imported helpers
```
//can be active or inactive like if in prod will be false
 asaniHelpers.LoggerConfig = true
 asaniHelpers.LogInfo("test info")
 asaniHelpers.LogWarning("test info")
 asaniHelpers.LogError(fmt.Errorf("test info"))
 asaniHelpers.LogDebug("test info")
 asaniHelpers.LogFatal(fmt.Errorf("test info"))
```

#### Asani Fiber Response Helpers
- if wanted to reponse of fiber in asani you can use this helpers so , all respone will be same structure in 1 call,base on status code
- example,
```
func AnyControllersFunction(c *fiber.Ctx) error { 

    return asaniHelpers.FiberReponses(c,fiber.StatusOK,"whatever status", "whatever message success", "whatever data wanted to return to reponse") 
}
 
```  

-here status code can be handle in return FiberReponses, 
```
    StatusOK = 200
    StatusBadRequest = 400
    StatusNotFound = 404
    StatusInternalServerError = 500
    StatusConflict = 409
    StatusUnauthorized = 401
    StatusForbidden = 403
    StatusTooManyRequests = 429
    StatusRequestEntityTooLarge = 413
    StatusNoConten = 204
```

- if you have any inputs just fork, PR , merge , and have fun .

