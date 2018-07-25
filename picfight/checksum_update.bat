cd ..

cd pfcec\secp256k1

del secp256k1.go

go run -tags gensecp256k1 genprecomps.go

cd ..

cd ..
