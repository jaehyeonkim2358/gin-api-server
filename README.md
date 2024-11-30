curl http://localhost:8888/api/boards \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "게시글 제목","content": "게시글 내용"}'

curl http://localhost:8888/api/boards \
    --include \
    --header "Content-Type: application/json" \
    --request "GET"

curl http://localhost:8888/api/boards/3 \
    --include \
    --header "Content-Type: application/json" \
    --request "GET"

curl http://localhost:8888/api/boards/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "PATCH" \
    --data '{"title": "게시글 제목","content": "게시글 내용"}'
