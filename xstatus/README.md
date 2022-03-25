# xstatus

## Dependencies

+ (xtesting)

## Document

### Types

+ `type DbStatus uint64`
+ `type FsmStatus uint64`
+ `type JwtStatus uint64`

### Variables

+ None

### Constants

+ `const DbUnknown DbStatus`
+ `const DbSuccess DbStatus`
+ `const DbNotFound DbStatus`
+ `const DbExisted DbStatus`
+ `const DbFailed DbStatus`
+ `const DbTagA DbStatus`
+ `const DbTagB DbStatus`
+ `const DbTagC DbStatus`
+ `const DbTagD DbStatus`
+ `const DbTagE DbStatus`
+ `const DbTagF DbStatus`
+ `const DbTagG DbStatus`
+ `const FsmNone FsmStatus`
+ `const FsmState FsmStatus`
+ `const FsmFinal FsmStatus`
+ `const FsmTagA FsmStatus`
+ `const FsmTagB FsmStatus`
+ `const FsmTagC FsmStatus`
+ `const FsmTagD FsmStatus`
+ `const FsmTagE FsmStatus`
+ `const FsmTagF FsmStatus`
+ `const FsmTagG FsmStatus`
+ `const JwtUnknown JwtStatus`
+ `const JwtSuccess JwtStatus`
+ `const JwtBlank JwtStatus`
+ `const JwtInvalid JwtStatus`
+ `const JwtTokenNotFound JwtStatus`
+ `const JwtUserNotFound JwtStatus`
+ `const JwtFailed JwtStatus`
+ `const JwtTagA JwtStatus`
+ `const JwtTagB JwtStatus`
+ `const JwtTagC JwtStatus`
+ `const JwtTagD JwtStatus`
+ `const JwtTagE JwtStatus`
+ `const JwtTagF JwtStatus`
+ `const JwtTagG JwtStatus`
+ `const JwtAudience JwtStatus`
+ `const JwtExpired JwtStatus`
+ `const JwtId JwtStatus`
+ `const JwtIssuedAt JwtStatus`
+ `const JwtIssuer JwtStatus`
+ `const JwtNotValidYet JwtStatus`
+ `const JwtSubject JwtStatus`
+ `const JwtClaimsInvalid JwtStatus`

### Functions

+ None

### Methods

+ `func (d DbStatus) String() string`
+ `func (f FsmStatus) String() string`
+ `func (j JwtStatus) String() string`
