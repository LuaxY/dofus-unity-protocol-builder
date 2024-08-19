go run ./cmd/builder/main.go -input_dir=./proto_descriptor -output_dir=./proto

for file in proto/connection/*.proto; do
  protoc -I=./ --go_out=. --proto_path=proto/connection "$file"
done

for file in proto/game/*.proto; do
  protoc -I=./ --go_out=. --proto_path=proto/game "$file"
done

rm -rf src/protocol/connection
rm -rf src/protocol/game
mv go-xp-dofus-unity-proto-builder/src/protocol/* src/protocol/
rm -rf go-xp-dofus-unity-proto-builder