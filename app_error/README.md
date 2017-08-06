DEPRECATED
==========

Take a look at https://github.com/creasty/apperrors.

<details>

app_error
=========

Why
---

Introducing a concept of "app error" which is "handled" errors -- like "yeah, I know it could happen and I'm taking good care of."  
App error can be exposed to clients so that they can tell what's happening inside a request, and what exactly causes the request to fail.


What
----

- Must not use `c.AbortWithError` for non-fatal errors
  - Use `c.Error` instead
- Must create custom error (i.e. `errors.New`) for "handled" errors
  - Create a const of an error string
  - Error string must follow the pattern `{status}.{package}.{struct/domain}...`

```go
r.Use(app_error.Wrap())

// or

r.Use(app_error.WrapWithCallback(func (c *gin.Context, body []byte, err error) {
	// notify the error to somewhere...
}))
```

``` go
package api

import ...

const ERROR_NAME_EMPTY = "422.api.sample.name.empty"  // This is the "app error"

func Sample(c *gin.Context) {
	// ...

	if somethingWentWrong {
		err := errors.New(ERROR_NAME_EMPTY)  // Create custom error
		c.Error(err)                         // Must not use c.AbortWithError
		return
	}

	// ...

	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
```

</details>
