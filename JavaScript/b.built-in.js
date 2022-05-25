/*
    ## JS obj 구분 (ES5기준)
        - 빌트인 오브젝트
        - 네이티브 오브젝트
        - 호스트 오브제게트
    
    ### 빌트인 오브젝트
        - 사전에 만들어 놓은 오브젝트
        - ex) 빌트인 Number object, 빌트인 String object
    ### 네이티브 오브젝트
        - JS 스펙에서 정의한 오브젝트
        - 빌트인 오브젝트도 포함
        - JS 코드를 실행할 때 만드는 오브젝트
        - 빌트인obj가 네이티브obj에 속하는 개념
    ### 호스트 오브젝트
        - 빌트인, 네이티브 obj를 제외한 오브젝트
        - ex) window, DOM object
        - JS는 호스트 환경에서
            - 브라우저의 모든 요소 기술은
            - 연결하고 융합하며 이를 제어

    #### 빌트인 오브젝트 구조
        - 오브젝트 이름 (Object, String, Number)
        - 오브젝트.prototype
            - 인스턴스 생성 가능 여부 기준
            - 프로퍼티를 연결하는 오브젝트
        - 오브젝트.prototype.constructor
            - 오브젝트의 생성자
        - 오브젝트.prototype.method
            - 메소드 이름과 함수 작성
        - 오브젝트 구조
*/


/*
    ### Built-in이란?
        - 값 타입, 연산자, 오브젝트를 사전에 만들어 놓은 것
        - js코드를 처리하는 영역에
        - 장점
            - 사전 처리를 하지 않고 즉시 사용
            - 자바스크립트 특징

    ### .Built-in Object
        - Number obj
        - String obj
        - Boolean obj
        - Object obj
        - Array obj
        - Function obj
        - Math obj
        - Date obj
        - JSON obj
        - RegExp obj
        - Global obj


*/ 

/*
    ### 함수와 메소드 연결
    #### 함수
        - 오브젝트에 연결
        - ex) Object.create()
    #### 메소드
        - 오브젝트의 prototype에 연결
        - Object.prototype.toString()
    
    ### 함수, 메소드 호출
    #### 함수 호출 방법
        - Object.create();
    #### 메소드 호출 방법
        - Object.prototype.toString();
        - 또는 인스턴스를 생성하여 호출
    #### 함수와 메소드를 구분해야 하는 이유
        - js 코드 작성 방법이 다르기 때문
        - 함수는 파라미터 값을 작성하고 메소드는 메소드 앞에 값을 작성
    
    ### 프로퍼티 처리 메소드
    #### hasOwnProperty()
        - 인스턴스에 파라미터 이름이
            - 존재하면 true 반환
            - 존재하지 않으면 false 반환
        - 자신이 만든 것이 아니라
            - 상속받은 프로퍼티면 false 반환
    #### propertyIsEnumerable()
        - 오브젝트에서 프로퍼티를
            - 열거할 수 있으면 true 반환
            - 열거할 수 없으면 false 반환
    
    ### 빌트인 Object 특징
        - 인스턴스를 만들 수 있는 옴든 빌트인 오브젝트의 __proto__에
            Object.prototype의 6개 메소드가 설정됨
        - 따라서 빌트인 오브젝트로 만든 인스턴스에도 설정됨
    #### isPrototypeOf()
        -파라미터에 작성한 오브젝트에
            - object 위치에 작성한 prototype이
            - 존재하면 true 반환
            - 존재하지 않으면 false 반환
*/
console.log(Number)