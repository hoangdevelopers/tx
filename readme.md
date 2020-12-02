GET("/send/:data", sendTo)
GET("/verify/:txHash/:data", verify)
eg:
http://localhost:1323/send/data-sample
http://localhost:1323/verify/0xfe800c3d8c0f43feb51cd6f566f461805e4af848328668290e42f5a2e9bc475f/data-sample