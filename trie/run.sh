go build -o trie trie.go && ./trie  < ~/proj/wiki/enwiki-20170820-pages-articles-multistream-index.txt && go tool pprof trie memprofile.out
