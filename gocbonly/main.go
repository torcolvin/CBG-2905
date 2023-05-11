package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/couchbase/gocb/v2"
)

const (
	cbsUserName = "Administrator"
	cbsPassword = "password"
	cbsAddr     = "localhost"
	bucketName  = "bucket1"
)

func main() {
	gocb.SetLogger(gocb.VerboseStdioLogger())
	log.Print("Start test program")
	cluster := getCluster()
	defer func() {
		err := cluster.Close(nil)
		if err != nil {
			log.Fatalf("Error closing cluster %+v", err)
		}
	}()
	buckets, err := cluster.Buckets().GetAllBuckets(nil)

	for _, bucket := range buckets {
		fmt.Println("bucket1:", cluster.Bucket(bucket.Name))
		fmt.Println("bucket2:", cluster.Bucket(bucket.Name))
	}
	time.Sleep(10 * time.Second)
	cmd := exec.Command("docker", "exec", "couchbase", "couchbase-cli", "bucket-delete", "-c", "localhost", "-u", "Administrator", "-p", "password", "--bucket", "bucket1")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error deleting bucket %+v", err)
	}
	time.Sleep(1 * time.Hour)
	log.Print("SUCCESS")
}

func getCluster() *gocb.Cluster {
	DefaultGocbV2OperationTimeout := 10 * time.Second

	clusterOptions := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: cbsUserName,
			Password: cbsPassword,
		},
		SecurityConfig: gocb.SecurityConfig{
			TLSSkipVerify: false,
		},
		TimeoutsConfig: gocb.TimeoutsConfig{
			ConnectTimeout:    DefaultGocbV2OperationTimeout,
			KVTimeout:         DefaultGocbV2OperationTimeout,
			ManagementTimeout: DefaultGocbV2OperationTimeout,
			QueryTimeout:      90 * time.Second,
			ViewTimeout:       90 * time.Second,
		},
	}
	connStr := fmt.Sprintf("couchbase://%s?idle_http_connection_timeout=90000&kv_pool_size=2&max_idle_http_connections=64000&max_perhost_idle_http_connections=256", cbsAddr)
	cluster, err := gocb.Connect(connStr, clusterOptions)
	if err != nil {
		log.Fatalf("Error connecticting to cluster %+v", err)
	}
	err = cluster.WaitUntilReady(15*time.Second, nil)
	if err != nil {
		log.Fatalf("Can't connect to cluster %+v", err)
	}

	err = cluster.WaitUntilReady(90*time.Second,
		&gocb.WaitUntilReadyOptions{ServiceTypes: []gocb.ServiceType{gocb.ServiceTypeQuery}},
	)
	if err != nil {
		log.Fatalf("Query service not online")
	}
	return cluster
}
