package main

import (
	"log"
)

// A CommentFilter will only pass TSPackets
// which have a non-empty Comment field.
// Concretely, len(Comment) > 0
type CommentFilter struct {
	TsNode
}

//register with global AvailableNodes map
func init() {
	AvailableNodes["CommentFilter"] = NewCommentFilter
}

func NewCommentFilter(fname string) (*CommentFilter, error) {
	// try to open file

	node := &CommentFilter{}
	node.input = make(chan TsPacket, CHAN_BUF_SIZE)
	node.outputs = make([]chan<- TsPacket, 0)

	go node.process()
	return node, nil
}

func (node *CommentFilter) process() {
	defer node.closeDown()
	for pkt := range node.input {
		node.pktsIn++
		if len(pkt.Comment) > 0 {
			for _, output := range node.outputs {
				node.pktsOut++
				output <- pkt
			}
		}
	}
}

func (node *CommentFilter) closeDown() {
	log.Print("closing down CommentFilter")
	for _, output := range node.outputs {
		close(output)
	}
}
