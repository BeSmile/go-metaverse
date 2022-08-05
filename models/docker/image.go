package docker

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"go-metaverse/helpers/app"
	//"github.com/docker/docker/api/types"
	//"github.com/docker/docker/api/types/container"
	//"github.com/docker/docker/pkg/stdcopy"
)

//
//func main() {
//	ctx := context.Background()
//	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
//	if err != nil {
//		panic(err)
//	}
//
//	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
//	if err != nil {
//		panic(err)
//	}
//	io.Copy(os.Stdout, reader)
//
//	resp, err := cli.ContainerCreate(ctx, &container.Config{
//		Image: "alpine",
//		Cmd:   []string{"echo", "hello world"},
//	}, nil, nil, "")
//	if err != nil {
//		panic(err)
//	}
//
//	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
//		panic(err)
//	}
//
//	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
//	select {
//	case err := <-errCh:
//		if err != nil {
//			panic(err)
//		}
//	case <-statusCh:
//	}
//
//	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
//	if err != nil {
//		panic(err)
//	}
//
//	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
//}

func (cmd *CMD) ImageList(c *gin.Context) {
	fmt.Println("GetImages")
	fmt.Println(cmd)
	images, err := cmd.client.ImageList(cmd.ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	for _, image := range images {
		fmt.Printf("%s\n", image.RepoTags)
		//fmt.Printf("%v\n", image)
	}
	app.OK(c, images, "success")
}
