#!/bin/bash

# 이 스크립트는 포스트그레스 SQL이 설치되어 있지않다면 직접 설치와 더불어 gin-study 데이터베이스를 생성합니다

if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "이 스크립트는 macOS에서 실행 중입니다."
else
    echo "이 스크립트는 macOS에서만 실행 가능합니다."
    exit 1
fi

brew install postgresql

brew services start postgresql

if ! psql -lqt | cut -d \| -f 1 | grep -qw gin-study; then
    createdb gin-study
    echo "gin-study 데이터베이스가 생성되었습니다."
else
    echo "gin-study 데이터베이스가 이미 존재합니다."
    exit 0
fi

brew services list | grep postgresql

# 유저를 생성하고 비밀번호를 설정합니다. (선택 사항)
echo "유저를 생성하시겠습니까? 비밀번호 123456, 유저네임 gin-user로 생성됩니다. (y/n)"

if read -r response && [[ "$response" == [Yy] ]]; then
    psql -c "CREATE USER \"gin-user\" WITH PASSWORD '123456';"
    psql -c "GRANT ALL PRIVILEGES ON DATABASE \"gin-study\" TO \"gin-user\";"
    echo "유저 gin-user가 생성되었고, gin-study 데이터베이스에 대한 모든 권한이 부여되었습니다."
else
    echo "유저 생성이 취소되었습니다."
fi
#DB의 URL은 이와같습니다 postgres://gin-user:123456@localhost:5432/gin-study