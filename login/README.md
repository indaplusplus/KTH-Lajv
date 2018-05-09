port 8021

Requires api key to login for datasektionen to be stored in LOGIN\_API\_KEY

/login

redirects to login2.datasektionen.se for login, returns token

gives 401 if login failed

/logout

takes token

/isLoggedin

Takes token, returns 200 if token is logged in, 401 if not

/getUser

takes token, returns username if logged in, 401 if not
