# xstatus

### References

+ xtesting*

### Functions

#### DbStatus

+ `type DbStatus int8`
+ `DbSuccess`
+ `DbNotFound`
+ `DbExisted`
+ `DbFailed`
+ `DbTagA`
+ `DbTagB`
+ `DbTagC`
+ `DbTagD`
+ `DbTagE`
+ `(d DbStatus) String() string`

#### FsmStatus

+ `type FsmStatus int8`
+ `FsmNone`
+ `FsmInState`
+ `FsmFinal`
+ `(f FsmStatus) String() string`

#### JwtStatus

+ `type JwtStatus int8`
+ `JwtSuccess`
+ `JwtExpired`
+ `JwtNotValid`
+ `JwtIssuer`
+ `JwtSubject`
+ `JwtInvalid`
+ `JwtUserErr`
+ `JwtFailed`
+ `JwtTagA`
+ `JwtTagB`
+ `JwtTagC`
+ `(j JwtStatus) String() string`
