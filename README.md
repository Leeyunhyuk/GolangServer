# 📍 Go언어 Server
### 📌 목적
* API 웹 서버의 구조를 익힌다.
* Golang에서 사용되는 웹 프레임워크(__🥃Gin__)를 익힌다.
### 📌 서버 구조
-   **[config](https://github.com/Leeyunhyuk/GolangServer/tree/main/config)**  
    설정 파일과 설정 관련된 코드가 위치한 디렉터리입니다.
    -   **config.go**  
        설정 관련된 Go 코드가 포함되어 있습니다.

----------

-   **[init](https://github.com/Leeyunhyuk/GolangServer/tree/main/init)**  
    프로젝트 초기화 및 실행과 관련된 파일이 위치한 디렉터리입니다.
    -   **cmd**  
        명령어와 관련된 코드가 위치한 디렉터리입니다.
        -   **cmd.go**  
            명령어 관련 Go 코드가 포함되어 있습니다.
    -   **config.toml**  
        설정 파일입니다.
    -   **main.go**  
        메인 실행 파일입니다.

----------

-   **[network](https://github.com/Leeyunhyuk/GolangServer/tree/main/network)**  
    네트워크 관련 코드가 위치한 디렉터리입니다.
    -   **root.go**  
        네트워크 관련 Go 코드가 포함되어 있습니다.
    -   **user.go**  
        사용자 관련 네트워크 코드가 포함되어 있습니다.

----------

-   **[repository](https://github.com/Leeyunhyuk/GolangServer/tree/main/repository)**  
    데이터베이스나 외부 저장소와 상호작용하는 코드가 위치한 디렉터리입니다.
    -   **root.go**  
        저장소 관련 Go 코드가 포함되어 있습니다.

----------

-   **[service](https://github.com/Leeyunhyuk/GolangServer/tree/main/service)**  
    비즈니스 로직과 관련된 코드가 위치한 디렉터리입니다.
    -   **root.go**  
        서비스 관련 Go 코드가 포함되어 있습니다.

----------

-   **[types](https://github.com/Leeyunhyuk/GolangServer/tree/main/types)**  
    유틸리티 및 에러 유형 관련 코드가 위치한 디렉터리입니다.
    -   **errors**  
        에러 관련 코드가 위치한 디렉터리입니다.
       -   **erors.go**  
            에러 관련 Go 코드가 포함되어 있습니다.
    -   **utils.go**  
        유틸리티 함수 및 코드가 포함되어 있습니다.