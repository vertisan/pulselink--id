package id

import (
	"fmt"
	"github.com/sony/sonyflake"
	"github.com/sony/sonyflake/awsutil"
	"net/http"
	"time"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings

	fmt.Printf("Initializing ID Generator ...\n")

	if isOnEc2() {
		st.MachineID = awsutil.AmazonEC2MachineID
	}

	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("Cannot start sonyflake!")
	}
}

func isOnEc2() bool {
	fmt.Printf("Checking environment ...\n")

	c := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := c.Get("http://169.254.169.254/latest/meta-data")
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("- AWS EC2 is not detected ...\n")
		return false
	}

	fmt.Printf("- Detected AWS EC2 machine ...\n")
	return true
}

func GenerateId() uint64 {
	id, err := sf.NextID()
	if err != nil {
		panic(fmt.Sprintf("Cannot create the ID: %v", err))
	}

	return id
}
