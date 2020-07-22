package main

import (
	"context"
	"github.com/aws/aws-xray-sdk-go/xray"
)

func init() {
	xray.Configure(xray.Config{
		ServiceVersion:   "1.2.3",
	})
}

func main() {
	// Start a segment
	ctx, seg := xray.BeginSegment(context.Background(), "service-name")
	// Start a subsegment
	_, subSeg := xray.BeginSubsegment(ctx, "subsegment-name")
	// ...
	// Add metadata or annotation here if necessary
	// ...
	subSeg.Close(nil)
	// Close the segment
	seg.Close(nil)
}