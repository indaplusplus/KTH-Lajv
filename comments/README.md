Modulen lyssnar alltid efter JSON på :8080.    

/comment/post  
Takes parameters:  
user - string - unique user id  
token - string - session token  
text - string - content of comment  
response - int - id of the comment this comment is responding to, otherwise -1      


/comment/like  
Takes parameters:  
user - string - unique user id  
token - string - session token  
postid - id - id of post    

/comment/delete  
Takes parameters:  
user - string - unique user id  
token - string - session token  
postid - id - id of post    

/comment/get  
Takes parameters:  
user - string - unique user id  
token - string - session token  
video - string - id of video
