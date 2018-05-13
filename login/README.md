Requires api key for login to datasektionen to be stored in env. variable LOGIN\_API\_KEY

Used by sending request to `{host}:8021/{endpoint}`

Login is based on being redirected to a kth page.

## Endpoints

###### /login

Redirects to login2.datasektionen.se for login, returns token

Gives 401 if login failed

###### /logout

Takes token as json:
`{"token": token}`

###### /isLoggedin

Takes token as json:
`{"token": token}`

Returns 200 if token is logged in, 401 if not

###### /getUser

Takes token as json:
`{"token": token}`

if token is logged in, returns username as json: `{"user": kthid}`, returns 401 if not

###### /addDummyData

adds dummy users and tokens to database for testing
