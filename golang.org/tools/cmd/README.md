```$xslt
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file holds the go generate command to run yacc on the grammar in expr.y.
// To build expr:
//	% go generate
//	% go build

//go:generate goyacc -o goyacc_example.go -p "Calc" goyacc_example.y

// Expr is a simple expression evaluator that serves as a working example of
// how to use Go's yacc implementation.
package main
```


Run
```
goyacc -o goyacc_example.go -p "Calc" goyacc_example.y
```

<code>Calc</code> 대신에 다른 변수를 넣으면 해당 변수를 prefix 로 갖는 변수명들이 쭉 생성된다.
즉 현재 코드 <code>goyacc_example.go</code> 에 있는 <code>Calc</code> 대신에 지정한 변수가 들어가게 된다.
  
기본적으로 yacc 는 token 을 single character 로 쓴다. 그 의미는 정의할 때 토큰을 문자열로 지정할 수 없다는 의미이고 <code>goyacc_exmaple.y</code> 에서 
token을 literal (string) 로 주면 __character token too long__ 이 뜬다. 해당 코드 [여기](https://github.com/golang/tools/blob/master/cmd/goyacc/yacc.go#L806) 서 확인 가능하다. 

single quote(') 로 감싸면 당연하게도 __invalid token__ 에러가 뜨고, 어떠한 quote도 주지 않으면 __syntax error__ 가 뜬다. 

결과적으로 literal token 을 주고 싶으면 조합해서 사용하면 된다. __3 -->__ 를 입력하면 두배를 곱한 결과값 6이 나오는 정의이다.
해당 파일에서는 다음과 같이 입력하였다.

<pre><code>|   expr '-' '-' '>'
    { $$ = $1 *  }
</code></pre>


결과로 생성된 <code>goyacc_example.go</code> 를 실행해서 다음과 같이 입력하면 원하는 결과 값이 나올 것이다.

<pre><code>equation: 3 -->
6
</code></pre>

goyacc 명령어 없이 직접 파일을 수정할 수도 있는데, 그럴 경우 최상단의 __DO NOT EDIT__ 을 지워주고 하자.

##### Ref Links
- [https://godoc.org/golang.org/x/tools/cmd/goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc)
- [https://github.com/golang-samples/yacc](https://github.com/golang-samples/yacc)
