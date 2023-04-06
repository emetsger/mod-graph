module github.com/emetsger/mod-graph/module-a

go 1.20

replace github.com/emetsger/mod-graph/module-b v0.0.0-unpublished => ../module-b

require github.com/emetsger/mod-graph/module-b v0.0.0-unpublished
