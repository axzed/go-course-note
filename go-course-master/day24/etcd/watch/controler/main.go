package main

import "gitee.com/infraboard/go-course/day24/etcd/watch"

func main() {
	watch.WatchConfig("/registry/configs")
}
