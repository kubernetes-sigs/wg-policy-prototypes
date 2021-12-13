package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	report "github.com/haardikdharma10/kubearmor-adapter/pkg/report"
	pb "github.com/kubearmor/KubeArmor/protobuf"
	"google.golang.org/grpc"
	"k8s.io/client-go/util/homedir"
)

func GetOSSigChannel() chan os.Signal {
	c := make(chan os.Signal, 1)

	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)

	return c
}

func main() {

	StopChan := make(chan struct{})

	gRPCPtr := flag.String("gRPC", "", "gRPC server information")

	flag.Parse()
	gRPC := ""

	if *gRPCPtr != "" {
		gRPC = *gRPCPtr
	} else {
		if val, ok := os.LookupEnv("KUBEARMOR_SERVICE"); ok {
			gRPC = val
		} else {
			gRPC = "localhost:32767"
		}
	}

	conn, err := grpc.Dial(gRPC, grpc.WithInsecure())

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	client := pb.NewLogServiceClient(conn)

	req := pb.RequestMessage{}

	//Stream Alerts
	go func() {
		defer conn.Close()
		if stream, err := client.WatchAlerts(context.Background(), &req); err == nil {
			for {
				select {
				case <-StopChan:
					return

				default:
					res, err := stream.Recv()
					if err != nil {
						fmt.Print("system alerts stream stopped: " + err.Error())
					}

					r, err := report.New(res, res.NamespaceName)

					if err != nil {
						fmt.Printf("failed to create policy reports: %v \n", err)
						os.Exit(-1)
					}
					r, err = report.Write(r, res.NamespaceName, filepath.Join(homedir.HomeDir(), ".kube", "config"))
					if err != nil {
						fmt.Printf("failed to create policy reports: %v \n", err)
						os.Exit(-1)
					}
					fmt.Printf("wrote policy report %s \n", r.Name)
				}
			}
		} else {
			fmt.Print("unable to stream systems alerts: " + err.Error())
		}
	}()
	sigChan := GetOSSigChannel()
	<-sigChan
	close(StopChan)
}
